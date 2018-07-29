package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
	"./versions"
	"flag"
	"fmt"
	"isoft/isoft_storage/cfg"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		sectionSearch := flag.String("sectionSearch", os.Args[1], "sectionSearch")
		cfg.InitConfig(*sectionSearch)
	} else {
		fmt.Println("os args length error...")
	}

	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)

	fmt.Println(fmt.Sprintf("Start ListenAndServe address %s", cfg.GetConfigValue(cfg.LISTEN_ADDRESS)))

	log.Fatal(http.ListenAndServe(cfg.GetConfigValue(cfg.LISTEN_ADDRESS), nil))
}
