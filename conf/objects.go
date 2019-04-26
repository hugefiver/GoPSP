package conf

type LogObject struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

type UserObject struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

type DownloaderObject struct {
	UA        string `yaml:"ua"`
	Threads   int    `yaml:"threads"`
	SleepTime int    `yaml:"sleeptime"`
}

type ProxyObject struct {
	Method  string `yaml:"method"`
	Address string `yaml:"address"`
}

//PixivPathObject : path for Pixiv saving pictures
type PixivPathObject struct {
	BasePath string `yaml:"base_path"`
	Daily    string `yaml:"daily"`
	Month    string `yaml:"month"`
	Member   string `yaml:"member"`
}
