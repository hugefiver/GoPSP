package downloader

type PicDownloader struct {
	Files []DFile
	Limit int
}

func (d *PicDownloader) StartDownload() {
	var threads chan struct{}

	if d.Limit > 0 {
		threads = make(chan struct{}, d.Limit)
	}

	// TODO downloader
}
