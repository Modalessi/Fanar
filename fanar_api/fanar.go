package fanar

import (
	"fmt"
	"net/http"
)

type FanarServer struct {
	Addr       string
	JWTSecret  string
	AdminEmail string
	Server     *http.ServeMux
	Storage    Storage
}

func NewFanarServer(addr string, jwtSecret string, adminEmail string, storage Storage) *FanarServer {
	server := &http.ServeMux{}

	fs := &FanarServer{
		Addr:       addr,
		JWTSecret:  jwtSecret,
		AdminEmail: adminEmail,
		Server:     server,
		Storage:    storage,
	}

	server.HandleFunc("GET /check", withServer(fs, checkHealth))

	server.HandleFunc("POST /register", withServer(fs, register))
	server.HandleFunc("POST /login", withServer(fs, login))

	server.HandleFunc("GET /protected", withServer(fs, authorized(protected, fs.JWTSecret)))

	return fs
}

func (fs *FanarServer) Start() error {
	fmt.Println("server is now running at " + fs.Addr)
	return http.ListenAndServe("localhost"+fs.Addr, fs.Server)
}
