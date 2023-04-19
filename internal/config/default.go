package config

var serverDef = server{
	ListenPort: "3000",
	DebugMode:  true,
	FileLogger: false,
}

var databaseDef = database{
	Type: "internal",
}

var queueDef = queue{
	Type: "internal",
}

var cacheDef = cache{
	Type: "internal",
}

var searchDef = search{
	Type: "internal",
}
