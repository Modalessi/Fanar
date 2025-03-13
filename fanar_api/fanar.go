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

	// middlwares
	fanarHandler := NewFanarHandler(fs)
	adminOnly := NewAdminOnlyMiddlware(adminEmail)
	authorized := NewAuthorizedMiddleware(jwtSecret)

	server.HandleFunc("GET /check", fanarHandler(checkHealth))

	server.HandleFunc("POST /register", fanarHandler(register))
	server.HandleFunc("POST /login", fanarHandler(login))

	server.HandleFunc("POST /course", fanarHandler(authorized(adminOnly(createCourse))))
	server.HandleFunc("DELETE /course", fanarHandler(authorized(adminOnly(deleteCourse))))
	server.HandleFunc("PUT /course", fanarHandler(authorized(adminOnly(editCourse))))
	server.HandleFunc("GET /course", fanarHandler(getCourse))

	server.HandleFunc("POST /resource", fanarHandler(authorized(addResource)))
	server.HandleFunc("GET /resource/link", fanarHandler(getResourceLink))

	server.HandleFunc("GET /protected", fanarHandler(authorized(protected)))

	return fs
}

func (fs *FanarServer) Start() error {
	fmt.Println("server is now running at " + fs.Addr)
	return http.ListenAndServe("localhost"+fs.Addr, fs.Server)
}
