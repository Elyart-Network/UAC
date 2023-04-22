package config

var serverDef = server{
	ListenPort: "3000",
	DebugMode:  true,
	FileLogger: false,
}

var postgresDef = postgres{
	Host: "localhost:5432",
	Name: "uac",
}

var redisDef = redis{
	Hosts: []string{"localhost:6379"},
}
