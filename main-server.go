package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nightlyone/lockfile"
)

var port = flag.Int("port", 5000, "Port to run webserver on")
var logs = flag.String("logs", "/home/ec2-user/", "Directory to use for log storage")

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	logInit(*logs)

	// Quit if already running
	lock, err := lockfile.New("/tmp/main-server.lck")
	if err != nil {
		log.Fatal("Cannot init lock. reason: ", err)
	}
	err = lock.TryLock()

	// Error handling is essential, as we only try to get the lock.
	if err != nil {
		Error.Println("Cannot get application lock file initiatied: ", err)
		os.Exit(1)
	}

	defer lock.Unlock()

	Info.Println("Application log file: ", *logs+"app_logs")
	Info.Println("Server port numer: ", *port)

	dist := http.FileServer(http.Dir("/home/ec2-user/dist/"))

	http.Handle("/dist/", http.StripPrefix("/dist", dist))
	http.Handle("/css/", dist)
	http.Handle("/assets/", dist)
	http.Handle("/js/", dist)
	http.HandleFunc("/", homePageHandler)

	Error.Println(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
