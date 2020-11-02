package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type Download struct {
	Url string
	Path string
	OldPath string

	Error error
}

func New(url string) *Download {
	download := new(Download)
	download.Url = url
	download.Path = suggestedPath(url)
	download.OldPath = suggestedPath(url) + ".old"

	if _, err := os.Stat(download.Path); err==nil {
		if _, err := os.Stat(download.OldPath); err==nil {
			os.Remove(download.OldPath)
		}
		os.Rename(download.Path, download.OldPath)
	}
	return download
}

func (d *Download) DownloadURL(url string) (err error) {
	fmt.Println("Downloading ", url, " to ", d.Path)

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(d.Path)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}

func suggestedPath(url string) string {
	return "/var/tmp/" + path.Base(url)
}