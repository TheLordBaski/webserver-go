package main

import(
	"log"
	"runtime"

	"./core/server"
	"./core/boot"
	"./core/router"
	"./lib/env"
)


// Setup runtime settings
// Runs before main
func init(){
	// Verbose logging with file name and number
	log.SetFlags(log.Lshortfile)

	// We can use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){
	// Load the configuration file
	config, err := env.LoadConfig( "/home/gabco/Projects/go/src/github.com/thelordbaski/test-web/env.json")
	if err != nil {
		log.Fatalln(err)
	}

	handler := boot.SetUpMiddleware(router.Instance())
	server.Run(
		handler,       // HTTP handler
		config.Server, // Server settings
	)
}

