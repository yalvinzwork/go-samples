package main

import (
	"flag"
	"fmt"
	"github.com/google/gops/agent"
	"log"
	"net/http"

	ifsconfig "github.com/yalvinz/ipreit/config"
	"github.com/yalvinz/ipreit/hello"
	"github.com/yalvinz/ipreit/util/grace"
	"github.com/yalvinz/ipreit/util/logging"
)

func main() {
	var err error

	// Initialize logger
	flag.Parse()
	logging.LogInit()

	// Initialize debug logger
	debug := logging.Debug.Println
	debug("app started")

	// Initialize Gops agent
	if err = agent.Listen(agent.Options{ShutdownCleanup: true}); err != nil {
		log.Fatal(err)
	}
	debug("gops started")

	// Initialize config
	moduleName := "main"
	mainCfg := &ifsconfig.MainConfig{}
	if err = ifsconfig.InitConfig(mainCfg, moduleName, "files/etc/ipreit"); err != nil {
		log.Fatal(err)
	}
	debug("successfully read config")
	fmt.Printf("loaded %s module config is: %+v\n", moduleName, mainCfg)

	// Initialize modules
	hModule := hello.NewHelloWorldModule(mainCfg)

	// Initialize handlers
	http.HandleFunc("/api/hello", hModule.SayHelloWorld)

	// tracer.Init(&tracer.Config{Port: 8700, Enabled: true})

	log.Fatal(grace.Serve(":9000", nil))
}
