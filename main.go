package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	DioExportServer = "http://localhost:8000"
	ConvertType     = "png"
)

func main() {
	rootDir := "input"
	destDir := "output"
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		rel, _ := filepath.Rel(rootDir, path)
		if info.IsDir() {
			os.Mkdir(filepath.Join(destDir, rel), os.ModePerm)
		} else {
			// TODO: check ext (.dio, ???, ...)
			if err := convert(rootDir, destDir, rel, ConvertType); err != nil {
				log.Printf("ERR: %v", err)
			}
		}

		return nil
	})
}

func convert(origin, dest, rel, ext string) error {
	originFilePath := filepath.Join(origin, rel)
	new, err := getNewFilename(rel, ext)
	if err != nil {
		return err
	}
	destFilePath := filepath.Join(dest, new)

	out, err := os.Create(destFilePath)
	if err != nil {
		return fmt.Errorf("failed to create dest file: %w", err)
	}

	// Set up request
	form := url.Values{}
	form.Add("format", ext)
	xml, err := ioutil.ReadFile(originFilePath)
	if len(xml) == 0 {
		return fmt.Errorf("input file is empty: %s", originFilePath)
	}
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	form.Add("xml", string(xml))

	body := strings.NewReader(form.Encode())
	req, err := http.NewRequest(http.MethodPost, DioExportServer, body)
	if err != nil {
		return fmt.Errorf("failed to create a request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("encountered unexpected error: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("request is failed with status code: %s", res.Status)
	}

	// Write results
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return fmt.Errorf("failed to copy response body to file: %w", err)
	}
	return nil
}

func getNewFilename(rel, ext string) (string, error) {
	idx := strings.LastIndex(rel, ".")
	if idx < 0 {
		return "", fmt.Errorf("file doen't have any ext: %s", rel)
	}
	new := rel[:idx]

	return fmt.Sprintf("%s.%s", new, ext), nil
}
