package main

import (
	"flag"
	"fmt"
	"github.com/hultan/manjaro-torrent/internal/download"
	"github.com/hultan/manjaro-torrent/internal/manjaro"
	notify_user "github.com/hultan/manjaro-torrent/internal/notifier"
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

	notify := notify_user.New()
	notify.NotifyUserIfNeeded(manjaroNew, manjaroOld)
}
