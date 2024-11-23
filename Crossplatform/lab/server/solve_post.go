package main

import (
	"linear_system_lab/linear_system"
	"linear_system_lab/request"
	"log"
	"net/http"
)

type SolveReqBody struct {
	linear_system.LinearSystem
}

type SolveResBody struct {
	linear_system.Solution
}

func SolvePost(w http.ResponseWriter, r *http.Request) {
	req := request.NewT(r, w)
	body, err := request.ParseJSONBody[SolveReqBody](req)
	if err != nil {
		req.SendErr(err, http.StatusBadRequest)
		return
	}
	if err := req.SendJSON(body.SolveMatrix(), http.StatusOK); err != nil {
		log.Println("Error sending response: ", err)
	}
}
