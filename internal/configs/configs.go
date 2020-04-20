package configs

import (
	"flag"
	"go-initializr/pkg/application"
	"os"
)

func Init() application.Application {
	appName := flag.String("appName", "", "[Mandatory] Application to generate.")
	flag.Parse()
	if *appName == "" {
		flag.Usage()
		os.Exit(1)
	}
	return application.NewApplication(*appName)
}
