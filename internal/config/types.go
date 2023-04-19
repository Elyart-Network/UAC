package config

type config struct {
	Server   server   `yaml:"server"`
	Database database `yaml:"handlers"`
	Queue    queue    `yaml:"queue"`
	Cache    cache    `yaml:"cache"`
	Search   search   `yaml:"search"`
}

type server struct {
	ListenPort string `yaml:"listen_port"`
	DebugMode  bool   `yaml:"debug_mode"`
	FileLogger bool   `yaml:"file_logger"`
}

type database struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type queue struct {
	Type      string `yaml:"type"`
	Host      string `yaml:"host"`
	IndexName string `yaml:"index_name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

type cache struct {
	Type      string `yaml:"type"`
	Host      string `yaml:"host"`
	IndexName string `yaml:"index_name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

type search struct {
	Type      string `yaml:"type"`
	Host      string `yaml:"host"`
	IndexName string `yaml:"index_name"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}
