package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"slices"
	"strings"
	"sync"

	"emperror.dev/emperror"
	"go.yaml.in/yaml/v4"
)

func downloadSpec() {
	emperror.Panic(os.MkdirAll("spec", 0o755))

	os.RemoveAll("spec")

	downloader := &specDownloader{
		baseURL: "https://wa-cdn.nyc3.cdn.digitaloceanspaces.com/assets/prod/boromir-documentation/swagger/",
		baseDir: "spec",
	}

	downloader.wg.Add(1)
	go downloader.downloadRecursively("openapi.yml")
	downloader.wg.Wait()
}

type specDownloader struct {
	wg         sync.WaitGroup
	mu         sync.Mutex
	baseURL    string
	baseDir    string
	downloaded []string
}

func (d *specDownloader) downloadRecursively(specFile string) {
	defer d.wg.Done()

	fullURL := d.baseURL + specFile
	res := must(http.Get(fullURL))
	defer res.Body.Close()

	filePath := path.Join(d.baseDir, specFile)
	parentDir := path.Dir(filePath)

	d.mu.Lock()
	if slices.Contains(d.downloaded, specFile) {
		d.mu.Unlock()
		return
	}
	d.downloaded = append(d.downloaded, specFile)
	emperror.Panic(os.MkdirAll(parentDir, 0o755))
	d.mu.Unlock()

	slog.Info("Downloading spec file", "file", specFile)

	f := must(os.Create(filePath))
	defer f.Close()

	buf := new(bytes.Buffer)
	must(io.Copy(buf, io.TeeReader(res.Body, f)))

	var doc map[string]any
	if err := yaml.Unmarshal(buf.Bytes(), &doc); err != nil {
		slog.Error("Failed to parse spec file", "file", specFile, "url", fullURL, "error", err)
		return
	}

	objStack := []any{doc}

	for len(objStack) > 0 {
		var obj any
		obj, objStack = objStack[len(objStack)-1], objStack[:len(objStack)-1]

		switch v := obj.(type) {
		case map[string]any:
			if ref, ok := v["$ref"].(string); ok {
				if len(ref) > 0 && ref[0] != '#' {
					ref, _, _ = strings.Cut(ref, "#")

					// Compute spec file path relative to the current spec file
					ref = path.Join(path.Dir(specFile), ref)

					d.wg.Add(1)
					go d.downloadRecursively(ref)
				}
			}
			for _, value := range v {
				objStack = append(objStack, value)
			}
		case []any:
			objStack = append(objStack, v...)
		}
	}
}
