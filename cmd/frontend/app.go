package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/yalvinz/go-sample/util/grace"
	"github.com/yalvinz/go-sample/util/logging"
)

var tpl *template.Template

func main() {
	// Initialize logger
	flag.Parse()
	logging.LogInit()

	// Initialize debug logger
	debug := logging.Debug.Println
	debug("app started")

	http.HandleFunc("/comingsoon", idx)

	log.Fatal(grace.Serve(":9001", nil))
}

// please move below codes to module instead defining in app.go
func init() {
	tpl = template.Must(template.ParseGlob("/var/www/go-sample/bin/templates/*.gohtml"))
}

func idx(w http.ResponseWriter, req *http.Request) {
	debug := logging.Debug.Println

	args := make(map[string]interface{})
	args["AssetUrl"] = "http://dev.ifreight.id" // please load url asset from config, don't hardcode
	args["MetaTitle"] = "Coming Soon :: iFreight"
	args["Header"] = "iFreight is coming soon"
	args["Desc"] = "Stay tuned, iFreight is going to bring you a life changing experience for your freight needs."

	err := tpl.ExecuteTemplate(w, "comingsoon.gohtml", args)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	debug(req.URL.Path)
}
