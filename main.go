package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"strings"

    "github.com/schollz/progressbar/v3"
)






func main() {
    url := "https://shdw-drive.genesysgo.net/snapshots/latest"
    incUrl := "https://shdw-drive.genesysgo.net/snapshots/latest-incremental"
    // cli flag to determine if incremental
    incFlag := flag.Bool("incremental", false, "Set the download to pull the latest incremental snapshot")

    flag.Parse()

    if *incFlag == true {
        DownloadSnapshot(incUrl)
    } else {
        DownloadSnapshot(url)
    }

}

func DownloadSnapshot (url string) error {

    // Make request, grab new URL, and return filename and extension
    req, err := http.NewRequest("GET", url, nil)
    resp, err := http.DefaultClient.Do(req)
    finalURL := resp.Request.URL.String()
    parts := strings.Split(finalURL, "/")
    filename := parts[len(parts)-1]
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer out.Close()

    // Create progress bar to track progress and filesize
    bar := progressbar.DefaultBytes(
        resp.ContentLength,
        "Downloading",
    )
    io.Copy(io.MultiWriter(out, bar), resp.Body)
    return err
}
