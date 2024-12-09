package main

import (
	"archive/zip"
	"crypto/tls"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func downloadAndExtractZip(url string, distPath string) {
	// Get the file form http
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(url)
	if err != nil {
		panic("Failed to download server files from " + url + ", error: " + err.Error())
	}
	defer resp.Body.Close()

	// Create a temporary file to store the downloaded ZIP
	tmpFile, err := os.CreateTemp("", "zip-*.zip")
	if err != nil {
		panic("Failed to create temporary file: " + err.Error())
	}
	defer os.Remove(tmpFile.Name())

	// Copy the downloaded content to the temporary file
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		panic("Failed to write server zip fil from " + url + " to " + tmpFile.Name() + ", error: " + err.Error())
	}

	// Open the ZIP file for extraction
	reader, err := zip.OpenReader(tmpFile.Name())
	if err != nil {
		panic("Failed to open zip: " + err.Error())
	}
	defer reader.Close()

	zipMembers := reader.File
	// Check if all files share a common top-level directory
	var topLevelDir string
	foundTopLevelDir := false
	for _, f := range reader.File {
		dir, _ := filepath.Split(f.Name)

		if !foundTopLevelDir {
			topLevelDir = dir
			foundTopLevelDir = true
		} else if dir != topLevelDir {
			slog.Info("Files have different top-level directories, extracting as normal")
			extractZipFile(zipMembers, distPath)
			return
		}
	}

	slog.Warn("All files share a common top-level directory, extracting while omitting the top level directory")
	for i, f := range reader.File {
		zipMembers[i].Name = strings.Join(strings.Split(f.Name, "/")[1:], "/")
	}
	extractZipFile(zipMembers, distPath)
}

func extractZipFile(zipMembers []*zip.File, path string) {
	for _, f := range zipMembers {
		filePath := filepath.Join(path, f.Name)
		slog.Debug("Extracting file: " + filePath)

		// Create an empty dir in the destination if the zip file member is an empty dir
		if f.FileInfo().IsDir() {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				panic("Failed to create directory: " + err.Error())
			}
			continue
		}

		// Create the dir for destination file
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic("Failed to create directory: " + err.Error())
		}
		// Create the destination file
		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic("Failed to open destination file: " + filePath + ", error: " + err.Error())
		}
		defer dstFile.Close()

		// Open the zip member file
		fileInArchive, err := f.Open()
		if err != nil {
			panic("Failed to open soure file: " + filePath + ", error: " + err.Error())
		}
		defer fileInArchive.Close()

		// Copy contents from the file in zip to the destination file
		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic("Failed to write contents to destination file: " + filePath + ", error: " + err.Error())
		}
	}
}
