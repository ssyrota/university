package main

import (
	"linear_system_lab/linear_system"
	"linear_system_lab/request"
	"log"
	"net/http"
)

type DetReqBody struct {
	linear_system.LinearSystem
}

type DetResBody struct {
	linear_system.Solution
}

func DetPost(w http.ResponseWriter, r *http.Request) {
	req := request.NewT(r, w)
	body, err := request.ParseJSONBody[SolveReqBody](req)
	if err != nil {
		req.SendErr(err, http.StatusBadRequest)
		return
	}
	if err := req.SendJSON(map[string]float64{"determinant": body.Determinant()}, http.StatusOK); err != nil {
		log.Println("Error sending response: ", err)
	}
}
