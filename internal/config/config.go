package config

type Config struct {
	global Global
	connections []Connection
}

type Global struct {
	ssh_key string
}

type Connection struct {
	name string
}
