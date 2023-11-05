package llgutils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

func (data LaunchMeta) DownloadCosmetics(path string) (err error) {
	if !data.Success {
		return fmt.Errorf("[API] Success: False")
	}

	var wg sync.WaitGroup
	var failed []map[string]string
	counter := 0

	// get main index
	res, err := http.Get(data.Textures.IndexURL)

	if err != nil {
		return err
	}

	// read body
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	index := strings.Split(string(body), "\n")

	for _, val := range index {
		wg.Add(1)
		go func(cosmetic string) {
			defer wg.Done()

			l := strings.Split(cosmetic, " ")
			fp := fmt.Sprintf("%s/%s", path, l[0])
			hash := l[1]

			if IfExists(fp) && CheckHash(fp, hash) {
				log.Printf("[INFO] Asset already up to date: %s\n", fp)
				return
			}

			if err := DownloadFile(fp, data.Textures.BaseURL+hash); err != nil {
				log.Printf("[WARN] Error downloading asset: %s (%s)\n", fp, err)
				failed = append(failed, map[string]string{
					"path": fp,
					"url":  data.Textures.BaseURL + hash,
				})
				counter++
				return
			}
			log.Printf("[INFO] Downloaded asset: %s\n", fp)
		}(val)
	}

	wg.Wait()
	if counter != 0 {
		log.Printf("[INFO] Downloaded assets with %d failures, retrying...", counter)

		for i := 0; 0 < counter; i++ {
			if err := DownloadFile(failed[i]["path"], failed[i]["url"]); err != nil {
				log.Printf(
					"[WARN] Error downloading asset: %s (%s)\n",
					failed[i]["path"],
					failed[i]["url"],
				)
				failed = append(failed, map[string]string{
					"path": failed[i]["path"],
					"url":  failed[i]["url"],
				})
				counter++
				continue
			}
			log.Printf("[INFO] Downloaded asset: %s\n", failed[i]["path"])
			counter--
		}
	}

	return nil
}
