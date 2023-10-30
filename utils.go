package llgutils

import (
	"archive/zip"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
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
	// create all folders
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	// create file
	out, err := os.Create(path)

	if err != nil {
		return err
	}
	defer out.Close()

	// send request
	client := http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Get(url)

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

func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// Create destination directory if it doesn't exist
	os.MkdirAll(dest, os.ModePerm)

	// Extract files
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
		} else {
			// Create file
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			// Write file contents
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CreateLog(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, err
	}

	return os.Create(path)
}
