package llgutils

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func IfExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func CheckHash(path string, hash string) bool {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	if fmt.Sprintf("%x", h.Sum(nil)) == hash {
		return true
	}
	return false
}

func DownloadFile(path string, url string) (err error) {
	// create file
	out, err := os.Create(path)

	if err != nil {
		return err
	}
	defer out.Close()

	// send request
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// write to file
	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}
