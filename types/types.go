package types

import "time"

// Requesting svc creation on NodePort
type SvcRequest struct {
	Port      int32 `json:port`
	NodePort  int32 `json:nodeport`
	Label     PodLabel
	Namespace string `json:namespace`
	Time      time.Time
}

type PodLabel struct {
	Key   string
	Value string
}

type Namespacedata struct {
	Phase         string `json:"phase"`
	Name          string `json:"name"`
	Readypods     int    `json:"readypods"`
	Pendingpods   int    `json:"pendingpods"`
	Succeededpods int    `json:"succeededpods"`
	Failedpods    int    `json:"failedpods"`
	Unknownpods   int    `json:"unknownpods"`
}

type Svccache map[string]SvcRequest
