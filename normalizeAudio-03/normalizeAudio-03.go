package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	getWavs := GetFiles()

	numofFiles := len(getWavs)
	fmt.Println(numofFiles)

	for i := 0; i < numofFiles; i++ {
		fmt.Println("File Location : ", getWavs[i])
		// outputFile := getWavs[i]
		findVol := FindVolume(getWavs[i])
		NormalizeFile(getWavs[i], "outputDafoe.wav", findVol)
		fmt.Println("Max Volume : ", findVol)
	}
}

func FindVolume(inputFile string) string {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-af", "volumedetect", "-f", "null", "-y", "nul")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slug, _ := ioutil.ReadAll(stderr)

	convertSlug := BytesToString(slug)
	currentVol := RemoveHeader(convertSlug)

	removedDB := currentVol[1:5]
	normalizedVol := "volume=" + removedDB
	// fmt.Println(normalizedVol)
	return normalizedVol
}

func NormalizeFile(inputFile, outputFile, normalizedVol string) {
	cmd2 := exec.Command("ffmpeg", "-i", inputFile, "-filter:a", normalizedVol, outputFile)
	// err = cmd2.Start()
	if err := cmd2.Start(); err != nil {
		log.Fatal(err)
	}

}

func GetFiles() []string {

	var files []string

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dir)

	root := string(dir)
	err2 := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".wav" {
			files = append(files, path)
		}
		// files = append(files, path)
		return nil
	})
	if err2 != nil {
		panic(err2)
	}
	// for _, file := range files {
	// 	fmt.Println(file)
	// }
	return files
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func RemoveHeader(s string) string {
	if idx := strings.Index(s, "max_volume: "); idx != -1 {
		return s[idx+12:]
	}
	return s
}
