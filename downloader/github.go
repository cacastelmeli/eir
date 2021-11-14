package downloader

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cacastelmeli/eir/util"
)

const (
	githubTarballUrl = "https://github.com/%s/tarball/master"
)

type GithubDownloader struct {
}

func NewGithubDownloader() *GithubDownloader {
	return &GithubDownloader{}
}

// Download downloads github's tarball from given `templatePath`, then decompress it
// into `CacheTemplateDir`
func (downloader *GithubDownloader) Download(templatePath string) error {
	targetTarUrl := fmt.Sprintf(githubTarballUrl, templatePath)

	if resp, err := http.Get(targetTarUrl); err != nil {
		return err
	} else if uncompressedStream, err := gzip.NewReader(resp.Body); err != nil {
		return err
	} else {
		tarReader := tar.NewReader(uncompressedStream)

		for {
			header, err := tarReader.Next()

			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			switch header.Typeflag {
			case tar.TypeDir:
				if err = os.Mkdir(util.PathCacheTemplate(header.Name), os.ModePerm); err != nil {
					return err
				}

			case tar.TypeReg:
				if file, err := os.Create(util.PathCacheTemplate(header.Name)); err != nil {
					return err
				} else if _, err := io.Copy(file, tarReader); err != nil {
					return err
				} else {
					file.Close()
				}
			}
		}
	}

	return nil
}
