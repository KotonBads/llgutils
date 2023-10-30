package main

import (
	utils "github.com/KotonBads/llgutils"
)

func main() {
	launchbody := utils.LaunchBody{
		OS: "linux",
		Arch: "x64",
		Version: "1.8.9",
		Module: "forge",
	}
	if launchmeta, err := utils.FetchLaunchMeta(launchbody); err == nil {
		launchmeta.DownloadArtifacts("temp")
	}
}