package http

import "github.com/gorilla/mux"

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/getData", s.User.UserHandler).Methods("GET")
	r.HandleFunc("/insertData", s.User.UserHandler).Methods("POST")
	r.HandleFunc("/updateData", s.User.UserHandler).Methods("PUT")

	return r
}
