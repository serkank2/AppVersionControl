package loader

import (
	"appVersionControl/api/config"
	"log"
)

func Loader() {

	envControl := config.Env()
	if !envControl {
		log.Fatal("Error Env load")
	}

}
