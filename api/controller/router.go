package controller

import middlewares "github.com/prajwalnayak7/Blog/api/middleware"

func (s *Server) initializeRoutes() {

	// Default Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

}
