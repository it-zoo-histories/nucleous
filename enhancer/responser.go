package enhancer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	nameServer = "[NUCLEOUS: RESPONSER]: "
)

type Responser struct {
}

/*ResponseWithError - ответ с ошибкой*/
func (resp *Responser) ResponseWithError(w http.ResponseWriter, r *http.Request, httpStatus int, payload interface{}, contentType string) {
	resp.ResponseWithJSON(w, r, httpStatus, payload, contentType)
}

/*ResponseWithJSON - ответ в формате json*/
func (resp *Responser) ResponseWithJSON(w http.ResponseWriter, r *http.Request, httpStatus int, payload interface{}, contentType string) {
	fmt.Println(nameServer+"request header type: ", r.Header.Get("Content-Type"))
	if !strings.Contains(r.Header.Get("Content-Type"), contentType) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)

		log.Println("[NUCLEOUS]: error request from: ", r.RemoteAddr)

		response, _ := json.Marshal(map[string]string{"error": "you packet in non " + contentType + " format!"})
		w.Write(response)

	} else {

		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")

		log.Println("[NUCLEOUS]: success request from: ", r.RemoteAddr)

		w.WriteHeader(httpStatus)
		w.Write(response)
	}
}
