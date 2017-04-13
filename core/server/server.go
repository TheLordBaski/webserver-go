package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Info struct {
	Hostname	string `json:"Hostname"`	// Server name
	UseHTTP         bool   `json:"UseHTTP"`         // Listen on HTTP
	HTTPPort        int    `json:"HTTPPort"`        // HTTP port
	//HTTPS will be probably omitted because we can use reverse nginx proxy engine
}

// if we would use https we can specify httpsHandler http.Handler in parameter and rework run method
func Run(httpHandler http.Handler,  info Info){
	if info.UseHTTP {
		startHttp(httpHandler, info);
	} else {
		log.Println("Config file does not specify a listener to start")
	}
}
func startHttp(handlers http.Handler, info Info) {
	fmt.Println(time.Now().Format("31-12-2016 15:04:05"), "Running HTTP "+httpAddress(info))
	// Start the HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(info), handlers))
}
func httpAddress(info Info) string {
	return info.Hostname + ":" + fmt.Sprintf("%d", info.HTTPPort)
}