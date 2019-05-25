package middlewares

import (
	"encoding/json"
	"net/http"
	"nucleous/enhancer"
	"nucleous/payloads"
)

type NextServe struct {
	EResponse *enhancer.Responser
}

func (ser *NextServe) SendNext(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload payloads.ResendPacket

		defer r.Body.Close()

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			ser.EResponse.ResponseWithError(w, r, http.StatusBadRequest, map[string]string{
				"status":  "404",
				"context": "nucleous.Resender",
				"code":    "does not convert need service and route on them",
			},
				"application/json",
			)
		}

		next(w, r)
	})
}
