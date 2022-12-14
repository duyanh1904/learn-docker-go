package main

import (
	"flag"
	"fmt"
	_ "github.com/santosh/gingo/docs"
	"os"

	"github.com/duyanh1904/learn-docker-go/config"
	"github.com/duyanh1904/learn-docker-go/db"
	"github.com/duyanh1904/learn-docker-go/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
