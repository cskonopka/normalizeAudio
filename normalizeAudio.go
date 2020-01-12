package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	getWavs, arraySize, dir := GetFiles()

	// Iterate
	for i := 0; i < arraySize; i++ {
		// Find the volume of the file
		findVol := FindVolume(getWavs[i])

		// Isolate filenames from directory
		splitFilename := strings.SplitAfter(getWavs[i], dir+"/")

		// Remove the file extension from the filename
		filename := splitFilename[1][:len(splitFilename[1])-4]

		// Normalized output filename
		normalizedOutput := filename + "-normalized.wav"

		/* Normalize the file
		files[i]: the current audio file
		normalizedOutput: the filename of the normalized file
		findVol: The dB peak of the current audio file
		*/
		NormalizeFile(getWavs[i], normalizedOutput, findVol)
	}
}

func GetFiles() ([]string, int, string) {
	// String array for files
	var files []string

	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Filewalk the directory
	err2 := filepath.Walk(string(dir), func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".wav" {
			// Append file paths to the file string array
			files = append(files, path)
		}
		return nil
	})
	if err2 != nil {
		panic(err2)
	}

	arraySize := len(files)

	// return a string array of files in the directory
	return files, arraySize, dir
}

func RemoveHeader(s string) string {
	if idx := strings.Index(s, "max_volume: "); idx != -1 {
		return s[idx+12:]
	}
	return s
}

func FindVolume(inputFile string) string {

	// Get the peak of the audio file using 'volumedetect'
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-af", "volumedetect", "-f", "null", "-y", "nul")

	// Output of the command
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start command
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Read the output of the command
	ffmpegCLIOutput, _ := ioutil.ReadAll(stderr)

	// Convert the command output from bytes to a string
	convertCLIOutput := string(ffmpegCLIOutput[:])

	// Index the command output and output the 'max_volume" value
	currentVol := RemoveHeader(convertCLIOutput)

	// Remove the 'dB' extension
	var removedDB string

	if strings.Contains(currentVol[:], "0.0 dB") {
		removedDB = currentVol[0:4]
	} else {
		removedDB = currentVol[1:5]
	}
	normalizedVol := "volume=" + removedDB

	return normalizedVol
}

func NormalizeFile(inputFile, outputFile, normalizedVol string) {
	cmd2 := exec.Command("ffmpeg", "-i", inputFile, "-filter:a", normalizedVol, outputFile)
	// err = cmd2.Start()
	if err := cmd2.Start(); err != nil {
		log.Fatal(err)
	}
}
