package downloader

import "os"

type DFile struct {
	URL      string
	FilePath string
}

func (f *DFile) IsFileExists() bool {
	_, err := os.Stat(f.FilePath)
	return os.IsExist(err)
}
