package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ljahier/reverse-proxy/pkg/args"
)

/*
 * TODO
 * Catch all incoming requests on the configured port
 * Read config file
 * In this priority order
 * - -c path/to/config/file
 * - $HOME/.reverse-proxy.json
 * - /etc/reverse-proxy.json
 * Redirecting incoming requests from a domain to its own configuration
 * Run goroutine(?) to read each X seconds the config file and compare it to previous config file hash
 */

type Config struct {
	Sites []Site
}

type Site struct {
	Host string
	Port string
}

func checkConfigFile() {

}

func ReadConfigFile() {

}

func RunHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got / request\n")
		fmt.Println("r.Host = ", r.Host)
		io.WriteString(w, "This is my website!\n")
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	config := Config{}
	// read config file
	args.CheckCliArgs(&config)
	// run http server in other thread
	RunHttpServer()
	// run configuration changes watcher
}
