package main

import (
	"encoding/json"
	"errors"
	"github.com/magiconair/properties"
	"os"
	"path"
	"testing"
)

func TestAppMainWithCorrectData(t *testing.T) {
	serverDir := path.Join(os.TempDir(), "omsms-server-init-test")
	defer os.RemoveAll(serverDir)
	ServerMountPath = serverDir

	filesInit := filesInit{
		CustomStartScript: "",
		ServerIconUrl:     "https://placehold.co/64",

		// server.properties stuff
		Motd:               `An \u00A7cOMSMS\u00A7r managed server, more info at \u00A79\u00A7nomsms.octsrv.org`,
		EnableCommandBlock: true,
		OnlineMode:         true,
		AllowFlight:        false,
		MaxTickTime:        int64((^uint64(0)) >> 1), // Max value for int64
		MaxPlayers:         60,
		SpawnProtection:    0,
		ViewDistance:       10,
		SimulationDistance: 9,
	}
	// Required because the library parses the unicode escape characters
	expectedMotd := `An §cOMSMS§r managed server, more info at §9§nomsms.octsrv.org`

	filesInitString, err := json.Marshal(filesInit)
	if err != nil {
		panic(err)
	}

	t.Setenv("OMSMS_SERVER_FILES_INIT", string(filesInitString))
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_TYPE", DeploymentTypeZip)
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_VALUE", "https://mediafilez.forgecdn.net/files/3822/691/ATM3-SERVER-FULL-6.1.1.zip")
	t.Setenv("OMSMS_SERVER_START_SCRIPT_NAME", "startserver.sh")

	main()

	if _, err := os.Stat(path.Join(serverDir, "mods")); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected mods folder to be in " + serverDir)
	}
	if _, err := os.Stat(path.Join(serverDir, "config")); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected config folder to be in " + serverDir)
	}
	if _, err := os.Stat(path.Join(serverDir, "scripts")); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected scripts folder to be in " + serverDir)
	}

	if _, err := os.Stat(path.Join(serverDir, "server-icon.png")); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected server-icon.png to be in " + serverDir)
	}

	p := properties.MustLoadFile(path.Join(serverDir, "server.properties"), properties.UTF8)

	motd := p.MustGetString("motd")
	if motd != expectedMotd {
		t.Fatalf(`Expected motd to be "%s" but got "%s"`, expectedMotd, motd)
	}
	EnableCommandBlock := p.MustGetBool("enable-command-block")
	if EnableCommandBlock != filesInit.EnableCommandBlock {
		t.Fatalf("Expected enable-command-block to be %v but got %v", filesInit.EnableCommandBlock, EnableCommandBlock)
	}
	OnlineMode := p.MustGetBool("online-mode")
	if OnlineMode != filesInit.OnlineMode {
		t.Fatalf("Expected online-mode to be %v but got %v", filesInit.OnlineMode, OnlineMode)
	}
	AllowFlight := p.MustGetBool("allow-flight")
	if AllowFlight != filesInit.AllowFlight {
		t.Fatalf("Expected allow-flight to be %v but got %v", filesInit.OnlineMode, OnlineMode)
	}
	MaxTickTime := p.MustGetInt64("max-tick-time")
	if MaxTickTime != filesInit.MaxTickTime {
		t.Fatalf("Expected max-tick-time to be %v but got %v", filesInit.MaxTickTime, MaxTickTime)
	}
	MaxPlayers := p.MustGetUint("max-players")
	if MaxPlayers != filesInit.MaxPlayers {
		t.Fatalf("Expected max-players to be %v but got %v", filesInit.MaxPlayers, MaxPlayers)
	}
	SpawnProtection := p.MustGetUint("spawn-protection")
	if SpawnProtection != filesInit.SpawnProtection {
		t.Fatalf("Expected spawn-protection to be %v but got %v", filesInit.SpawnProtection, SpawnProtection)
	}
	ViewDistance := p.MustGetUint("view-distance")
	if ViewDistance != filesInit.ViewDistance {
		t.Fatalf("Expected view-distance to be %v but got %v", filesInit.ViewDistance, ViewDistance)
	}
	SimulationDistance := p.MustGetUint("simulation-distance")
	if SimulationDistance != filesInit.SimulationDistance {
		t.Fatalf("Expected simulation-distance to be %v but got %v", filesInit.ViewDistance, ViewDistance)
	}
}

func TestAppMainWithCorrectDataAndCustomScript(t *testing.T) {
	serverDir := path.Join(os.TempDir(), "omsms-server-init-test")
	defer os.RemoveAll(serverDir)
	ServerMountPath = serverDir

	filesInit := filesInit{
		CustomStartScript: "echo \"ducky is cool\"\njava -Xmx8G -Xms8G server.jar",
		ServerIconUrl:     "https://placehold.co/64",

		// server.properties stuff
		Motd:               `An \u00A7cOMSMS\u00A7r managed server, more info at \u00A79\u00A7nomsms.octsrv.org`,
		EnableCommandBlock: true,
		OnlineMode:         true,
		AllowFlight:        false,
		MaxTickTime:        int64((^uint64(0)) >> 1), // Max value for int64
		MaxPlayers:         60,
		SpawnProtection:    0,
		ViewDistance:       10,
		SimulationDistance: 9,
	}
	// Required because the library parses the unicode escape characters
	expectedMotd := `An §cOMSMS§r managed server, more info at §9§nomsms.octsrv.org`

	filesInitString, err := json.Marshal(filesInit)
	if err != nil {
		panic(err)
	}

	t.Setenv("OMSMS_SERVER_FILES_INIT", string(filesInitString))
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_TYPE", DeploymentTypeZip)
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_VALUE", "https://mediafilez.forgecdn.net/files/3822/691/ATM3-SERVER-FULL-6.1.1.zip")
	t.Setenv("OMSMS_SERVER_START_SCRIPT_NAME", "startserver.sh")

	main()

	if _, err := os.Stat(path.Join(serverDir, "mods")); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected mods folder to be in " + serverDir)
	}
	if _, err := os.Stat(path.Join(serverDir, "config")); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected config folder to be in " + serverDir)
	}
	if _, err := os.Stat(path.Join(serverDir, "scripts")); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected scripts folder to be in " + serverDir)
	}

	if _, err := os.Stat(path.Join(serverDir, "server-icon.png")); errors.Is(err, os.ErrNotExist) {
		t.Fatal("Expected server-icon.png to be in " + serverDir)
	}

	p := properties.MustLoadFile(path.Join(serverDir, "server.properties"), properties.UTF8)

	motd := p.MustGetString("motd")
	if motd != expectedMotd {
		t.Fatalf(`Expected motd to be "%s" but got "%s"`, expectedMotd, motd)
	}
	EnableCommandBlock := p.MustGetBool("enable-command-block")
	if EnableCommandBlock != filesInit.EnableCommandBlock {
		t.Fatalf("Expected enable-command-block to be %v but got %v", filesInit.EnableCommandBlock, EnableCommandBlock)
	}
	OnlineMode := p.MustGetBool("online-mode")
	if OnlineMode != filesInit.OnlineMode {
		t.Fatalf("Expected online-mode to be %v but got %v", filesInit.OnlineMode, OnlineMode)
	}
	AllowFlight := p.MustGetBool("allow-flight")
	if AllowFlight != filesInit.AllowFlight {
		t.Fatalf("Expected allow-flight to be %v but got %v", filesInit.OnlineMode, OnlineMode)
	}
	MaxTickTime := p.MustGetInt64("max-tick-time")
	if MaxTickTime != filesInit.MaxTickTime {
		t.Fatalf("Expected max-tick-time to be %v but got %v", filesInit.MaxTickTime, MaxTickTime)
	}
	MaxPlayers := p.MustGetUint("max-players")
	if MaxPlayers != filesInit.MaxPlayers {
		t.Fatalf("Expected max-players to be %v but got %v", filesInit.MaxPlayers, MaxPlayers)
	}
	SpawnProtection := p.MustGetUint("spawn-protection")
	if SpawnProtection != filesInit.SpawnProtection {
		t.Fatalf("Expected spawn-protection to be %v but got %v", filesInit.SpawnProtection, SpawnProtection)
	}
	ViewDistance := p.MustGetUint("view-distance")
	if ViewDistance != filesInit.ViewDistance {
		t.Fatalf("Expected view-distance to be %v but got %v", filesInit.ViewDistance, ViewDistance)
	}
	SimulationDistance := p.MustGetUint("simulation-distance")
	if SimulationDistance != filesInit.SimulationDistance {
		t.Fatalf("Expected simulation-distance to be %v but got %v", filesInit.ViewDistance, ViewDistance)
	}

	// Validate custom script
	startScript, err := os.ReadFile(path.Join(serverDir, "startserver.sh"))
	if string(startScript) != filesInit.CustomStartScript {
		t.Fatalf("Expected startserver.sh to be %v but got %v", filesInit.CustomStartScript, startScript)
	}
}

func TestEnvParserWithIncorrectEnvs(t *testing.T) {
	filesInit := filesInit{
		CustomStartScript: "echo \"ducky is cool\"\njava -Xmx8G -Xms8G server.jar",
		ServerIconUrl:     "https://placehold.co/64",

		// server.properties stuff
		Motd:               `An \u00A7cOMSMS\u00A7r managed server, more info at \u00A79\u00A7nomsms.octsrv.org`,
		EnableCommandBlock: true,
		OnlineMode:         true,
		AllowFlight:        false,
		MaxTickTime:        int64((^uint64(0)) >> 1), // Max value for int64
		MaxPlayers:         60,
		SpawnProtection:    0,
		ViewDistance:       10,
		SimulationDistance: 9,
	}
	filesInitString, err := json.Marshal(filesInit)
	if err != nil {
		panic(err)
	}

	t.Setenv("OMSMS_SERVER_FILES_INIT", string(filesInitString))
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_TYPE", "Ducky")
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_VALUE", "https://mediafilez.forgecdn.net/files/3822/691/ATM3-SERVER-FULL-6.1.1.zip")
	t.Setenv("OMSMS_SERVER_START_SCRIPT_NAME", "startserver.sh")

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic recovered: %v", r)
		} else {
			t.Fatal("Expected panic, but none occurred")
		}
	}()

	getEnvs()
}

func TestEnvParserWithIncorrectEnvs2(t *testing.T) {
	filesInit := filesInit{
		CustomStartScript: "echo \"ducky is cool\"\njava -Xmx8G -Xms8G server.jar",
		ServerIconUrl:     "https://placehold.co/64",

		// server.properties stuff
		Motd:               `An \u00A7cOMSMS\u00A7r managed server, more info at \u00A79\u00A7nomsms.octsrv.org`,
		EnableCommandBlock: true,
		OnlineMode:         true,
		AllowFlight:        false,
		MaxTickTime:        int64((^uint64(0)) >> 1), // Max value for int64
		MaxPlayers:         60,
		SpawnProtection:    0,
		ViewDistance:       10,
		SimulationDistance: 9,
	}
	filesInitString, err := json.Marshal(filesInit)
	if err != nil {
		panic(err)
	}

	t.Setenv("OMSMS_SERVER_FILES_INIT", string(filesInitString))
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_TYPE", DeploymentTypeZip)
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_VALUE", "ducky://mediafilez.forgecdn.net/files/3822/691/ATM3-SERVER-FULL-6.1.1.zip")
	t.Setenv("OMSMS_SERVER_START_SCRIPT_NAME", "startserver.sh")

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic recovered: %v", r)
		} else {
			t.Fatal("Expected panic, but none occurred")
		}
	}()

	getEnvs()
}

func TestEnvParserWithIncorrectEnvs3(t *testing.T) {
	filesInit := filesInit{
		CustomStartScript: "echo \"ducky is cool\"\njava -Xmx8G -Xms8G server.jar",
		ServerIconUrl:     "https://placehold.co/64",

		// server.properties stuff
		Motd:               `An \u00A7cOMSMS\u00A7r managed server, more info at \u00A79\u00A7nomsms.octsrv.org`,
		EnableCommandBlock: true,
		OnlineMode:         true,
		AllowFlight:        false,
		MaxTickTime:        int64((^uint64(0)) >> 1), // Max value for int64
		MaxPlayers:         60,
		SpawnProtection:    0,
		ViewDistance:       10,
		SimulationDistance: 9,
	}
	filesInitString, err := json.Marshal(filesInit)
	if err != nil {
		panic(err)
	}

	t.Setenv("OMSMS_SERVER_FILES_INIT", string(filesInitString))
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_TYPE", DeploymentTypeZip)
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_VALUE", "https://mediafilez.forgecdn.net/files/3822/691/ATM3-SERVER-FULL-6.1.1.zip")

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic recovered: %v", r)
		} else {
			t.Fatal("Expected panic, but none occurred")
		}
	}()

	getEnvs()
}

func TestEnvParserWithIncorrectEnvs4(t *testing.T) {
	t.Setenv("OMSMS_SERVER_FILES_INIT", `{"ducky": 0"}`)
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_TYPE", DeploymentTypeZip)
	t.Setenv("OMSMS_SERVER_DEPLOYMENT_VALUE", "https://mediafilez.forgecdn.net/files/3822/691/ATM3-SERVER-FULL-6.1.1.zip")
	t.Setenv("OMSMS_SERVER_START_SCRIPT_NAME", "startserver.sh")

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic recovered: %v", r)
		} else {
			t.Fatal("Expected panic, but none occurred")
		}
	}()

	getEnvs()
}
