// Copyright (c) 2021. Homeland Interactive Technology Ltd. All rights reserved.

// Package manager manage and print build info of server, also provide PProf HTTP implement.
package manager

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

// config
var (
	GITHASH   string // current git branch commit short hash
	GITBRANCH string // current git branch name
	BUILDTIME string // application build time
	GOVERSION string // version of golang at building application
)

// GetBuildInfo get current application build info
func GetBuildInfo() string {
	return fmt.Sprintf(
		"Build time: %s, Git branch: %s, Git commit hash: %s, Golang version: %s",
		BUILDTIME, GITBRANCH, GITHASH, GOVERSION)
}

// ServePProf starting net/http/pprof to debug your app
func ServePProf(addr string) error {
	router := http.NewServeMux()
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return http.ListenAndServe(addr, router)
}
