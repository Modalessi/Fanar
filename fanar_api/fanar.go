package fanar

import (
	"fmt"
	"net/http"
)

type FanarServer struct {
	Addr   string
	Server *http.ServeMux
}

func NewFanarServer(addr string) *FanarServer {
	server := &http.ServeMux{}

	fs := &FanarServer{
		Addr:   addr,
		Server: server,
	}

	server.HandleFunc("GET /check", withServer(fs, checkHealth))

	return fs
}

func (fs *FanarServer) Start() error {
	fmt.Println("server is now running at " + fs.Addr)
	return http.ListenAndServe("localhost"+fs.Addr, fs.Server)
}
