package respond

import (
	"encoding/json"
	"net/http"
)

func With(w http.ResponseWriter, r *http.Request, httpStatus int, bodyData any) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(bodyData)
}
