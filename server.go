package main

import(
	"log"
	"runtime"
)


// Setup runtime settings
func init(){
	// Verbose logging with file name and number
	log.SetFlags(log.Lshortfile)

	// We can use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){

}

