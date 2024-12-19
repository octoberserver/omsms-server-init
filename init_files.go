package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
)

type filesInit struct {
	CustomStartScript string
	ServerIconUrl     string

	// server.properties stuff
	Motd               string
	EnableCommandBlock bool
	OnlineMode         bool
	AllowFlight        bool
	MaxTickTime        int64
	MaxPlayers         uint
	SpawnProtection    uint
	ViewDistance       uint
	SimulationDistance uint
}

var filesInitDefault = filesInit{
	Motd:               `An \u00A7cOMSMS\u00A7r managed server, more info at \u00A79\u00A7nomsms.octsrv.org`,
	EnableCommandBlock: true,
	OnlineMode:         true,
	AllowFlight:        false,
	MaxTickTime:        -1,
	MaxPlayers:         60,
	SpawnProtection:    0,
	ViewDistance:       10,
	SimulationDistance: 9,
}

func initServerFiles(filesInit filesInit, startScriptName string, serverFolderPath string) {
	slog.Info("Initialising server files...")

	eulaTxt, err := os.OpenFile(path.Join(serverFolderPath, "eula.txt"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic("Failed to open eula.txt: " + err.Error())
	}
	defer eulaTxt.Close()

	_, err = eulaTxt.WriteAt([]byte("eula=true"), 0)
	if err != nil {
		panic("Failed to write eula.txt: " + err.Error())
	}
	slog.Info("Successfully written to eula.txt")

	if filesInit.CustomStartScript != "" {
		startScriptPath := path.Join(serverFolderPath, startScriptName)
		slog.Info("Found custom script in config, writing custom start script " + startScriptPath + " with content: \n" + filesInit.CustomStartScript)

		startScriptFile, err := os.OpenFile(startScriptPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic("Failed to open destination file: " + startScriptPath + ", error: " + err.Error())
		}
		defer startScriptFile.Close()

		_, err = startScriptFile.WriteAt([]byte(filesInit.CustomStartScript), 0)
		if err != nil {
			panic("Failed to write script file: " + startScriptPath + ", error: " + err.Error())
		}
		slog.Info("Successfully written custom start script in server folder")
	}

	if filesInit.ServerIconUrl != "" {
		slog.Info("Found server icon url in config, downloading server icon: " + filesInit.ServerIconUrl)
		resp, err := http.Get(filesInit.ServerIconUrl)
		if err != nil {
			panic("Failed to download server icon from " + filesInit.ServerIconUrl + ", error: " + err.Error())
		}
		defer resp.Body.Close()
		slog.Info("Successfully downloaded server icon from " + filesInit.ServerIconUrl + ", saving to server folder")

		iconPath := path.Join(serverFolderPath, "server-icon.png")
		iconFile, err := os.OpenFile(iconPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic("Failed to open destination file: " + iconPath + ", error: " + err.Error())
		}
		defer iconFile.Close()

		_, err = io.Copy(iconFile, resp.Body)
		if err != nil {
			panic("Failed to save icon file: " + iconPath + ", error: " + err.Error())
		}
		slog.Info("Successfully downloaded and saved server icon")
	}

	slog.Info("Initialising server.properties")
	properties := map[string]any{
		"motd":                 filesInit.Motd,
		"enable-command-block": filesInit.EnableCommandBlock,
		"online-mode":          filesInit.OnlineMode,
		"allow-flight":         filesInit.AllowFlight,
		"max-tick-time":        filesInit.MaxTickTime,
		"max-players":          filesInit.MaxPlayers,
		"spawn-protection":     filesInit.SpawnProtection,
		"view-distance":        filesInit.ViewDistance,
		"simulation-distance":  filesInit.SimulationDistance,
	}

	file, err := os.Create(path.Join(serverFolderPath, "server.properties"))
	if err != nil {
		panic("Error creating server.properties: " + err.Error())
	}
	defer file.Close()

	slog.Info("Created server.properties, writing content")

	_, err = file.Seek(0, 0)
	if err != nil {
		panic("Failed to reset server properties file write pointer to start: " + err.Error())
	}
	for key, value := range properties {
		_, err := file.WriteString(fmt.Sprintf("%s=%v\n", key, value))
		if err != nil {
			fmt.Println("Error writing to server.properties:", err)
			return
		}
	}
	slog.Info("Successfully written content to server.properties")
}
