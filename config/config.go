package config

var ServerConfig ServerConfiguration

type ServerConfiguration struct {
	BaseURL           string
	EmailAuthName     string
	EmailAuthPassword string
	EmailServer       string
	EmailPort         int
	ErrorLogEmail     string
	Debug             bool
}
