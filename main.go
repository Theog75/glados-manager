package main

import (
	"glados-manager/config"
	"glados-manager/controllers"
	"glados-manager/svccache"
	"log"
	"net/http"
)

func main() {
	config.ReadEnv()
	svccache.CacheInit()
	go svccache.CachePruner(10, 30)
	http.Handle("/ping", controllers.Ping())
	http.Handle("/api/v1/getns", controllers.GetNS(config.KubeConfigPath))
	http.HandleFunc("/api/v1/createsvc", controllers.CreateSVC(config.KubeConfigPath))
	http.HandleFunc("/api/v1/deletesvc", controllers.DeleteSVC(config.KubeConfigPath))
	http.HandleFunc("/api/v1/getcache", controllers.GetCache())

	log.Print("Glados Manager starting to listen on port " + config.Port)
	http.ListenAndServe(":"+config.Port, nil)
}
