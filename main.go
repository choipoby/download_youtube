package main

import (
	"fmt"
	"net/url"
	"os"

	. "./api"
)

const VIDEO_EXTENSION = ".mp4"
const VIDEO_TYPE = "video/mp4"
const VIDEO_QUALITY = "highres"

func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	if len(os.Args) < 3 {
		fmt.Println("usage: main URL (nickname) (path)")
		fmt.Println("if path is not specified, it will be under current directory")
		return
	}
	videoURL := os.Args[1]
	var nickname string
	var filePath string = "./"
	if len(os.Args) > 2 {
		nickname = os.Args[2]
		if len(os.Args) > 3 {
			filePath = os.Args[3]
		}
	}

	fmt.Printf("video ID: %s will be downloaded to %s\n", videoURL, filePath)
	outputPath, err := DownloadYoutubeVideo(videoURL, filePath, nickname)
	if nil != err {
		fmt.Printf("video download failed %s", err)
	}
	fmt.Printf("video downloaded to %s\n", filePath+outputPath)
}

// DownloadYoutubeVideo is download Video
func DownloadYoutubeVideo(u string, outpath string, nickname string) (string, error) {
	videoID, err := getYoutubeVideoIDFromURL(u)
	if nil != err {
		return "", fmt.Errorf("Can't not found video id in url : %s", err)
	}

	videoFileName := videoID + VIDEO_EXTENSION
	if nickname != "" {
		videoFileName = nickname + VIDEO_EXTENSION
	}

	err = downloadYoutubeVideo(videoID, videoFileName, outpath, VIDEO_QUALITY, VIDEO_TYPE)
	if nil != err {
		return "", err
	}
	return videoFileName, nil
}

func downloadYoutubeVideo(videoID, filename, outpath, quality, videoType string) error {
	vl, err := GetVideoListFromId(videoID)
	if err != nil {
		return err
	}

	err = vl.Download(outpath, filename, "", videoType)
	return err
}
func getYoutubeVideoIDFromURL(u string) (string, error) {
	parsedURL, err := url.Parse(u)
	if nil != err {
		return "", err
	}
	m, err := url.ParseQuery(parsedURL.RawQuery)
	if nil != err {
		return "", err
	}
	id := m.Get("v")
	return id, nil
}
