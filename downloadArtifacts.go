package llgutils

import (
	"fmt"
	"log"
	"sync"
)

// Downloads Lunar's artifacts from API.
//
// Specify where to download artifacts with `path`.
func (data LaunchMeta) DownloadArtifacts(path string) (err error) {
	if !data.Success {
		return fmt.Errorf("[API] Success: False")
	}

	var wg sync.WaitGroup
	counter := 0

	for _, val := range data.LaunchTypeData.Artifacts {
		wg.Add(1)

		go func(artifact Artifacts) {
			defer wg.Done()

			fp := fmt.Sprintf("%s/%s", path, artifact.Name)

			if IfExists(fp) && CheckHash(fp, artifact.Sha1) {
				log.Printf("[INFO] Artifact already up to date: %s\n", artifact.Name)
				return
			}

			if err := DownloadFile(fp, artifact.Url); err != nil {
				log.Printf("[WARN] Error downloading artifact: %s (%s)\n", artifact.Name, err)
				counter++
				return
			}
			log.Printf("[INFO] Downloaded artifact: %s\n", artifact.Name)
		}(val)
	}

	wg.Wait()
	log.Printf("[INFO] Downloaded artifacts with %d failures", counter)

	return nil
}
