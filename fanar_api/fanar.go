package fanar

import (
	"fmt"
	"net/http"
)

type FanarServer struct {
	Addr    string
	Server  *http.ServeMux
	Storage Storage
}

func NewFanarServer(addr string, storage Storage) *FanarServer {
	server := &http.ServeMux{}

	fs := &FanarServer{
		Addr:    addr,
		Server:  server,
		Storage: storage,
	}

	server.HandleFunc("GET /check", withServer(fs, checkHealth))

	server.HandleFunc("POST /register", withServer(fs, register))

	return fs
}

func (fs *FanarServer) Start() error {
	fmt.Println("server is now running at " + fs.Addr)
	return http.ListenAndServe("localhost"+fs.Addr, fs.Server)
}
