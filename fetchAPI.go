package llgutils

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
