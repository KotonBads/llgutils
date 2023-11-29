package llgutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Fetches Lunar's API.
//
// Takes in a `LaunchBody` type. 
//
// Returns a `LaunchMeta` type.
//
// TODO: make this into LaunchBody.FetchLaunchMeta
func FetchLaunchMeta(launchdata LaunchBody) (response LaunchMeta, err error) {
	url := "https://api.lunarclientprod.com/launcher/launch"

	params := map[string]string{
		"hwid":             "0",
		"launch_type":      "OFFLINE",
		"branch":           "master",
		"launcher_version": "3.0.0",
		"os_release":       "6.5.10-200.fc38.x86_64",
		"installation_id":  "f1e23d4c-5a67-8a9c-0123-45678ac9e10a",
		"os":               launchdata.OS,
		"arch":             launchdata.Arch,
		"version":          launchdata.Version,
		"module":           launchdata.Module,
	}

	jsonVal, _ := json.Marshal(params)
	var res LaunchMeta

	client := &http.Client{}

	// create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonVal))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return res, err
	}

	// send request
	resp, err := client.Do(req)

	if err != nil {
		return res, err
	}

	// unmarshal response into res
	body, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &res); err != nil {
		return res, fmt.Errorf("couldn't Unmarshal response: %s", err)
	}

	return res, fmt.Errorf("%s (%s)", err, body)
}
