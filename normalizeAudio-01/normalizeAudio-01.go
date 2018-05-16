package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	inputFile := "dafoe.wav"
	outputFile := "newOutput.wav"

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

	cmd2 := exec.Command("ffmpeg", "-i", inputFile, "-filter:a", normalizedVol, outputFile)
	err = cmd2.Start()
	if err != nil {
		panic(err)
	}
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
