package config

type Config struct {
	Server Server
}

type Server struct {
	Host string
	Port string
}
