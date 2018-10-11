package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/gops/agent"
	ifsconfig "github.com/yalvinz/go-sample/config"
	"github.com/yalvinz/go-sample/hello"
	"github.com/yalvinz/go-sample/util/grace"
	"github.com/yalvinz/go-sample/util/logging"
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
