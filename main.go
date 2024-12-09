package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"log/slog"
	"os"
)

const (
	DeploymentTypeZip = "ZIP"
	DeploymentTypeGit = "GIT"
)

var (
	DeploymentTypes = []string{DeploymentTypeZip, DeploymentTypeGit}
	ServerMountPath = "/minecraft-server"
)

type envs struct {
	filesInit       filesInit
	deploymentType  string
	deploymentValue string
	startScriptName string
}

func main() {
	envs := getEnvs()

	if envs.deploymentType == DeploymentTypeZip {
		slog.Info("Deploying server from zip file...")
		downloadAndExtractZip(envs.deploymentValue, ServerMountPath)
		initServerFiles(envs.filesInit, envs.startScriptName, ServerMountPath)
	}

	if envs.deploymentType == DeploymentTypeGit {
		slog.Info("Deploying server from git repository...")
		_, err := git.PlainClone(ServerMountPath, false, &git.CloneOptions{
			URL:      envs.deploymentValue,
			Progress: os.Stdout,
		})
		if err != nil {
			panic("Failed to clone repository:" + err.Error())
		}
	}
}

func getEnvs() envs {
	// Get and validate fileInitString
	fileInitString := os.Getenv("OMSMS_SERVER_FILES_INIT")
	if fileInitString == "" {
		panic("OMSMS_SERVER_FILES_INIT environment variable not set")
	}
	filesInit := filesInitDefault
	err := json.Unmarshal([]byte(fileInitString), &filesInit)
	if err != nil {
		panic("Failed to decode config" + err.Error())
	}

	// Get and validate deploymentType
	deploymentType := os.Getenv("OMSMS_SERVER_DEPLOYMENT_TYPE")
	if deploymentType == "" {
		panic("OMSMS_SERVER_DEPLOYMENT_TYPE environment variable not set")
	}
	if !checkStringMatches(deploymentType, DeploymentTypes) {
		panic("Invalid Deployment Type: " + deploymentType)
	}

	// Get and validate deploymentValue
	deploymentValue := os.Getenv("OMSMS_SERVER_DEPLOYMENT_VALUE")
	if deploymentValue == "" {
		panic("OMSMS_SERVER_DEPLOYMENT_VALUE environment variable not set")
	}
	if !isURL(deploymentValue) {
		panic("Invalid deployment url: " + deploymentValue + " for type: " + deploymentType)
	}

	startScriptName := os.Getenv("OMSMS_SERVER_START_SCRIPT_NAME")
	if startScriptName == "" {
		panic("OMSMS_SERVER_DEPLOYMENT_VALUE environment variable not set")
	}
	if startScriptName[0] == '/' {
		startScriptName = startScriptName[1:]
	}

	slog.Info(fmt.Sprintf(`Successfully read environmental variables:
OMSMS_SERVER_DEPLOYMENT_TYPE: %s
OMSMS_SERVER_DEPLOYMENT_Value: %s
OMSMS_SERVER_START_SCRIPT_NAME: %s
OMSMS_SERVER_FILES_INIT: %s
`, deploymentType, deploymentValue, startScriptName, fileInitString))

	return envs{
		filesInit:       filesInit,
		deploymentType:  deploymentType,
		deploymentValue: deploymentValue,
		startScriptName: startScriptName,
	}
}
