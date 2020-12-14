package svccache

import (
	"fmt"
	"glados-manager/types"
	"log"
	"strconv"
	"time"
)

type svccache map[string]types.SvcRequest

var cachedb types.Svccache

//CachePruner prune items in the service cache
func CachePruner(interval time.Duration, pruneafter time.Duration) {
	for {
		time.Sleep(interval * time.Second)
		log.Print("Starting Pruning process")
		for _, v := range cachedb {
			if time.Since(v.Time) > pruneafter*time.Second {
				fmt.Print(time.Since(v.Time))
				//TODO add pruning of services over time
				// k8sclient.ServiceDelete(config.KubeConfigPath, v)
				// RemoveFromCache(v)
			}
		}
	}
}

//CacheInit - Initialize the cache map
func CacheInit() {
	cachedb = make(map[string]types.SvcRequest)

}

//GetCachData get service cache to rest call
func GetCachData() types.Svccache {
	return cachedb
}

//AddToCache - add a nodeport service which is created to the cache
func AddToCache(svctoadd types.SvcRequest) bool {
	now := time.Now()

	portname := svctoadd.Label.Value + "-" + strconv.Itoa(int(svctoadd.NodePort))
	svctoadd.Time = now

	cachedb[portname] = svctoadd
	return true
}

//RemoveFromCache remove service from cache (usually after service deletion)
func RemoveFromCache(svctoadd types.SvcRequest) bool {

	portname := svctoadd.Label.Value + "-" + strconv.Itoa(int(svctoadd.NodePort))

	delete(cachedb, portname)
	return true
}

//CheckSvcToCache - add a nodeport service which is created to the cache
func CheckSvcToCache(svctoadd types.SvcRequest) bool {
	portname := svctoadd.Label.Value + "-" + strconv.Itoa(int(svctoadd.NodePort))
	for k, v := range cachedb {
		// fmt.Println(k)
		if portname == k {
			log.Print("Node port on port " + strconv.Itoa(int(svctoadd.NodePort)) + " already in use by " + v.Label.Value)
		}
		return false
	}

	return true
}
