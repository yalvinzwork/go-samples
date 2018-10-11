package config

import (
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

func ReadModuleConfig(cfg interface{}, path string, module string) error {
	environ := os.Getenv("PENV")
	if environ == "" {
		environ = "development"
	}

	fname := path + "/" + module + "." + environ + ".ini"
	err := gcfg.ReadFileInto(cfg, fname)
	if err != nil {
		return err
	}

	return nil
}

// func MustReadModuleConfig(cfg interface{}, paths []string, module string) {
// 	res := false
// 	for _, path := range paths {
// 		res = ReadModuleConfig(cfg, path, module)
// 		if res == true {
// 			break
// 		}
// 	}

// 	if res == false {
// 		log.Fatalln("couldn't read config for ", os.Getenv("IFSENV"))
// 	}
// }
