package server

import "github.com/duyanh1904/learn-docker-go/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
