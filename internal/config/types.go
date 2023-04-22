package config

type config struct {
	Server   server   `yaml:"server"`
	Postgres postgres `yaml:"postgres"`
	Redis    redis    `yaml:"redis"`
}

type server struct {
	ListenPort string `yaml:"listen_port"`
	DebugMode  bool   `yaml:"debug_mode"`
	FileLogger bool   `yaml:"file_logger"`
}

type postgres struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type redis struct {
	Hosts    []string `yaml:"hosts"`
	Master   string   `yaml:"master"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}
