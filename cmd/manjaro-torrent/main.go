package main

import (
	"flag"
	"fmt"
	"github.com/hultan/manjaro-torrent/internal/download"
	"github.com/hultan/manjaro-torrent/internal/manjaro"
	"github.com/keybase/go-notifier"
	"os"
)

func main() {
	pUrl := flag.String("url", "", "URL to be processed")
	flag.Parse()
	url := *pUrl
	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: empty URL!\n")
		return
	}

	d := download.New(url)
	err := d.DownloadURL(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	manjaroNew := manjaro.New()
	manjaroNew.ParseHtml(d.Path)

	manjaroOld := manjaro.New()
	manjaroOld.ParseHtml(d.OldPath)

	var toDownload = false
	for name, oldVersion := range manjaroOld.Distributions {
		toDownload = false
		newVersion, ok := manjaroNew.Distributions[name]
		if ok == false {
			toDownload = true
		}
		if newVersion.Version != oldVersion.Version {
			toDownload = true
		}

		if toDownload {
			n, err := notifier.NewNotifier()
			if err != nil {
				fmt.Fprintf(os.Stderr, fmt.Sprintf("Failed to create notifier! %s", err.Error()))
				break
			}
			n.DeliverNotification(notifier.Notification{
				Title:   fmt.Sprintf("Manjaro %s has been updated!", newVersion.Name),
				Message: fmt.Sprintf("Update torrent for %s.\n\nOld Version : %s\nNew version : %s", name, oldVersion.Version, newVersion.Version),
			})
		}
	}
}
