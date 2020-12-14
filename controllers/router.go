package controllers

import (
	"encoding/json"
	"glados-manager/config"
	"glados-manager/k8sclient"
	"glados-manager/svccache"
	"glados-manager/types"
	"io/ioutil"
	"log"
	"net/http"
)

//GetCache get service cache from memory
func GetCache() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := svccache.GetCachData()

		// fmt.Println("Activated Ping test")
		responseforping := res
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

//GetNS test
func GetNS(KubeConfigPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		k8sclient.GetListOfNamespaces(config.KubeConfigPath)

		// fmt.Println("Activated Ping test")
		responseforping := Pong{"pong", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

//CreateSVC test
func CreateSVC(KubeConfigPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqd types.SvcRequest
		b, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(b, &reqd)
		if err != nil {
			log.Print(err)
		} else {
			// log.Print(reqd)
		}

		go k8sclient.ServiceCreate(config.KubeConfigPath, reqd)

		// fmt.Println("Activated Ping test")
		responseforping := Pong{"ServiceCreateCommand", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

//DeleteSVC test
func DeleteSVC(KubeConfigPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqd types.SvcRequest
		b, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(b, &reqd)
		if err != nil {
			log.Print(err)
		} else {
			log.Print(reqd)
		}

		k8sclient.ServiceDelete(config.KubeConfigPath, reqd)

		// fmt.Println("Activated Ping test")
		responseforping := Pong{"ServiceDeleteCommand", "ok"}
		js, err := json.Marshal(responseforping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
