# normalizeAudio

<p align="center">
  <img width="55%" height="55%" src="https://i.ibb.co/4FWHrbp/DBX-566-tube-comp.jpg"/>
</p>

A batch audio normalization tool built with Go and FFmpeg.

# Why?
I wanted to create a tool to analyze *.wav* files in a directory, find the peak dB, and export normalized audio files.

# Requirements
- FFmpeg

# How to use?
- Compile the tool.
```go
go build normalizeAudio.go
```

- Run the tool.
```go
./normalizeAudio location/of/directory
```

# Example
![gif](https://i.ibb.co/Y8HJNhn/normalize-Audio.gif)
