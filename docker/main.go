package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

var languageCountMap = make(map[string]int)

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

func main() {

	hostname, _ := os.Hostname()

	outFile, _ := os.Create(hostname +".csv")
	defer outFile.Close()

	args := os.Args

	if len(args) == 1 {
		fmt.Println("Please provide a repository name.")
		os.Exit(1)
	}

	repository := args[1]

	temp, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(temp)
	git.PlainClone(temp, false, &git.CloneOptions{
		URL:   "https://github.com/" + repository,
		Depth: 1,
	})
	filepath.WalkDir(temp, walker)
	for k, v := range languageCountMap {
		fmt.Fprintf(outFile, "%v,%v\n", k, v)
	}

}
