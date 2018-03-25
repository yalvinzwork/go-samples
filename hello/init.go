package hello

import (
	//"context"
	"fmt"
	//"log"
	"net/http"

	ifsconfig "github.com/yalvinz/ipreit/config"
	//"github.com/yalvinz/ipreit/util/router"
)

type HelloWorldModule struct {
	Cfg       *ifsconfig.HelloCfg
	Something string
}

func NewHelloWorldModule(cfg *ifsconfig.MainConfig) *HelloWorldModule {

	return &HelloWorldModule{
		Cfg:       &cfg.Hello,
		Something: "John Doe",
	}
}

func (hlm *HelloWorldModule) SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	// span, ctx := opentracing.StartSpanFromContext(r.Context(), r.URL.Path)
	// defer span.Finish()

	var word string
	for i := 0; i < hlm.Cfg.HelloCount; i++ {
		word = word + fmt.Sprintf("Hello %s \n", hlm.Something)
	}

	// resp := make(map[string]string)
	// resp["text"] = word

	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(word))
}
