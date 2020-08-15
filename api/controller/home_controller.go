package controller

import (
	"net/http"

	responses "github.com/prajwalnayak7/Blog/api/response"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "[SERVER] Bam!! This is working")

}
