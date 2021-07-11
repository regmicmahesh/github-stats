package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/patrickmn/go-cache"
)

var languageCountMap map[string]int

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 128*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}
func walker(path string, d os.DirEntry, err error) error {

	if d.IsDir() && d.Name() == ".git" {
		return filepath.SkipDir
	}
	if err != nil {
		return nil
	}
	if d.Type().IsRegular() {
		fileInfo, _ := d.Info()
		ext := filepath.Ext(fileInfo.Name())
		if len(ext) != 0 {
			reader, _ := os.Open(path)
			count, _ := lineCounter(reader)
			languageCountMap[ext] += count
		}
		return nil
	}
	return nil
}

func processRepository(writer io.Writer, repository string) error {

	temp, _ := ioutil.TempDir("", "repository")
	defer os.RemoveAll(temp)
	git.PlainClone(temp, false, &git.CloneOptions{
		URL:   "https://github.com/" + repository,
		Depth: 1,
	})
	filepath.WalkDir(temp, walker)
	jsonBytes, _ := json.Marshal(languageCountMap)

	jsonString := string(jsonBytes)

	fmt.Println("k cha")
	c.Set(repository, jsonString, cache.DefaultExpiration)
	fmt.Fprintln(writer, jsonString)
	return nil
}
