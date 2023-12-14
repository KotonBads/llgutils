# llgutils

Pretty much a library for creating a launcher for Lunar Client in Go. 


# Installation
run
```sh
go get github.com/KotonBads/llgutils@latest
```
then
```sh
go mod tidy
```


# Example
```go
package main

import (
	utils "github.com/KotonBads/llgutils"
)

func main() {
	launchbody := utils.LaunchBody{
		OS:      "linux",
		Arch:    "x64",
		Version: "1.8.9",
		Module:  "forge",
	}
	launchmeta, _ := launchbody.FetchLaunchMeta()
	launchmeta.DownloadArtifacts("temp")
	launchmeta.DownloadCosmetics("temp/textures")
}
```