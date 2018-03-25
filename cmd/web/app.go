package main

import (
	"flag"
	"github.com/google/gops/agent"
	"log"
	//"net/http"

	//"../../api"
	"../../logging"
)

func main() {

	// Initialize logger
	flag.Parse()
	logging.LogInit()

	// Initialize debug logger
	debug := logging.Debug.Println
	debug("app started")

	// Initialize Gops agent
	if err := agent.Listen(agent.Options{ShutdownCleanup: true}); err != nil {
		log.Fatal(err)
	}
	debug("gops started")

	// Initialize config

	// hwm := hello.NewHelloWorldModule()

	// http.Handle("/metrics", promhttp.Handler())

	// http.HandleFunc("/hello", hwm.SayHelloWorld)
	// go logging.StatsLog()

	// tracer.Init(&tracer.Config{Port: 8700, Enabled: true})

	// log.Fatal(grace.Serve(":9000", nil))
}
