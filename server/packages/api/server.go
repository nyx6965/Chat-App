package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Address string
}

func NewServer(Address string) *Server {
	return &Server{
		Address: Address,
	}
}

func (s *Server) StartServer() error {

	router := mux.NewRouter()
	
	// Auth routes
	router.Use(Logging())
	// router.Use(Authtication())
	router.HandleFunc("/register",HTTPHandler(s.RegisterUser)).Methods("POST")
	router.HandleFunc("/login",HTTPHandler(s.LoginUser)).Methods("POST")
	router.HandleFunc("/logout",HTTPHandler(s.Logout)).Methods("POST")
	
	// Channels Routes
	router.HandleFunc("/create-channel",HTTPHandler(s.CreateChannel)).Methods("POST")
	router.HandleFunc("/join-channel",HTTPHandler(s.JoinChannel)).Methods("POST")

	// Messages Routes
	// pool := NewPool()
    // go pool.Start()

	// router.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request) {
	// 	serveWs(pool, w, r)
	// })
	router.HandleFunc("/get-msg",HTTPHandler(s.GetMessages)).Methods("GET")

	log.Println("Server running on port:",s.Address)
	return http.ListenAndServe(s.Address,router)
}
