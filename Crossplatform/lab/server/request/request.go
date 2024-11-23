package request

import (
	"encoding/json"
	"net/http"
)

func NewT(r *http.Request, w http.ResponseWriter) T {
	return T{r, w}
}

type T struct {
	r *http.Request
	w http.ResponseWriter
}

func (r T) SendErr(err error, status int) {
	http.Error(r.w, err.Error(), status)
}

func ParseJSONBody[R any](r T) (R, error) {
	var reqBody R
	err := json.NewDecoder(r.r.Body).Decode(&reqBody)
	return reqBody, err
}

func (r T) SendJSON(resBody any, status int) error {
	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(status)
	err := json.NewEncoder(r.w).Encode(resBody)
	if err != nil {
		http.Error(r.w, err.Error(), http.StatusInternalServerError)
	}
	return err
}
