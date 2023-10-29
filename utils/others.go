package utils

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"
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
