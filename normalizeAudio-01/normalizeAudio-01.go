package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	inputFile := "RayWise.wav"
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
	// fmt.Println(currentVol)
	removedDB := currentVol[0:5]
	// fmt.Println(removedDB)
	// s := "0.0 d"
	var cleanedVol string
	t := strings.Replace(removedDB, "d", "", -1)
	if t != "0.0" {
		cleanedVol = t
		fmt.Println("j5itmf : ", cleanedVol)
	}

	normalizedVol := "volume=" + cleanedVol

	cmd2 := exec.Command("ffmpeg", "-i", inputFile, "-filter:a", normalizedVol, outputFile)
	err = cmd2.Start()
	if err != nil {
		panic(err)
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func RemoveHeader2(s string) string {
	if idx := strings.Index(s, "max_volume:"); idx != -1 {
		newSlice := s[idx+12:]
		return newSlice
	}
	return s
}

func RemoveHeader(s string) string {
	if idx := strings.Index(s, "max_volume:"); idx != -1 {
		return s[idx+12:]
	}
	return s
}
