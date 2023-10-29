package llg_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLaunchMeta(launchdata LaunchBody) (response LaunchMeta, err error) {
	url := "https://api.lunarclientprod.com/launcher/launch"

	params := map[string]string{
		"hwid":             "0",
		"launch_type":      "OFFLINE",
		"branch":           "master",
		"launcher_version": "3.0.0",
		"os":               launchdata.OS,
		"arch":             launchdata.Arch,
		"version":          launchdata.Version,
		"module":           launchdata.Module,
	}

	jsonVal, _ := json.Marshal(params)
	var res LaunchMeta

	client := &http.Client{}

	if req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonVal)); err == nil {
		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("User-Agent", "Lunar Client Launcher v2.16.1")

		if resp, err := client.Do(req); err == nil {
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))

			if err := json.Unmarshal(body, &res); err != nil {
				fmt.Printf("Couldn't Unmarshal the response: %s\n", err)
			}
			return res, err
		}
	}
	return res, err
}
