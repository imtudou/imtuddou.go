package server

import (
	"net/http"
)

var App = &Server{}

type Server struct{}

func (*Server) Run(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
