package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/text", s.handlesendtext()).Methods("POST")

}
