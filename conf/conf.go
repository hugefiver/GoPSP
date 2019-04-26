package conf

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const configpath = "./config.yml"

var (
	isInited   = false
	Log        LogObject
	User       UserObject
	Downloader DownloaderObject
	Proxy      ProxyObject
	PixivPath  PixivPathObject
)

type ConfigObject struct {
	Log          LogObject        `yaml:"log"`
	User         UserObject       `yaml:"user"`
	Downloader   DownloaderObject `yaml:"downloader"`
	Proxy        ProxyObject      `yaml:"proxy"`
	DownloadPath struct {
		Pixiv PixivPathObject `yaml:"pixiv"`
	} `yaml:"download_path"`
}

func getConfig() []byte {
	basePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	config := filepath.Join(basePath, configpath)

	file, err := os.Open(config)
	defer file.Close()

	if err != nil {
		log.Fatal("[Error] Read config error:", err)
		return []byte{}
	} else {
		b, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal("[Error] Read config error:", err)
		}
		return b
	}
}

func init() {
	if isInited {
		return
	}

	config := getConfig()
	c := ConfigObject{}

	if err := yaml.Unmarshal(config, &c); err != nil {
		log.Fatal("[Error] Unmarshal config failed:", err)
		return
	} else {
		// Log
		if c.Log.Path == "" {
			Log.Path = "./gopsp.log"
		} else {
			Log.Path = c.Log.Path
		}
		if c.Log.Level == "" {
			Log.Level = "warning"
		} else {
			Log.Level = c.Log.Level
		}

		// User
		User = c.User

		// Downloader
		Downloader.UA = c.Downloader.UA
		Downloader.SleepTime = c.Downloader.SleepTime
		if c.Downloader.Threads == 0 {
			Downloader.Threads = 1
		} else {
			Downloader.Threads = c.Downloader.Threads
		}

		// Proxy
		Proxy = c.Proxy

		// Download Path
		// Pixiv
		p := c.DownloadPath.Pixiv
		if p.BasePath != "" {
			PixivPath.BasePath = p.BasePath

			if strings.HasPrefix(p.Daily, "/") {
				PixivPath.Daily = p.Daily
			} else {
				PixivPath.Daily = filepath.Join(p.BasePath, p.Daily)
			}

			if strings.HasPrefix(p.Month, "/") {
				PixivPath.Month = p.Month
			} else {
				PixivPath.Month = filepath.Join(p.BasePath, p.Month)
			}

			if strings.HasPrefix(p.Member, "/") {
				PixivPath.Member = p.Member
			} else {
				PixivPath.Member = filepath.Join(p.BasePath, p.Member)
			}
		} else {
			PixivPath = p
		}

		isInited = true
	}
}
