package main

import (
	"fmt"

	utils "github.com/KotonBads/llgutils"
)

func main() {
	launchbody := utils.LaunchBody{
		OS:      "linux",
		Arch:    "x64",
		Version: "1.8.9",
		Module:  "forge",
	}
	if launchmeta, err := utils.FetchLaunchMeta(launchbody); err == nil {
		fmt.Println(launchmeta)
		launchmeta.DownloadArtifacts("temp")
	} else {
		fmt.Println(err)
	}
}
