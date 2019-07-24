package middlewares

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"nucleous/enhancer"
	"nucleous/payloads"
)

type NextServe struct {
	EResponse *enhancer.Responser
}

const (
	nameServer = "[NUCLEOUS: RESENDER]: "
)

var (
	servicesEntries map[string]string = map[string]string{
		"nickel": "http://nickel:9999",
	}
)

/*SendNext - проксирование запроса по его внутреннему пакету*/
func (ser *NextServe) SendNext(w http.ResponseWriter, r *http.Request) {

	var payload payloads.ResendPacket

	log.Println(nameServer + " handle new send")

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		ser.EResponse.ResponseWithError(w, r, http.StatusBadRequest, map[string]string{
			"status":  "404",
			"context": "nucleous.Resender",
			"code":    "does not convert need service and route on them",
		},
			"application/json",
		)
		return
	}

	buffer := new(bytes.Buffer)
	jsonF, err := json.Marshal(payload.Body)

	if err != nil {
		ser.EResponse.ResponseWithError(w, r, http.StatusBadRequest, map[string]string{
			"status":  "404",
			"context": "nucleous.Resender",
			"code":    err.Error(),
		},
			"application/json",
		)
		return
	}

	buffer.Write(jsonF)

	httpClient := &http.Client{}

	req, _ := http.NewRequest(payload.ResolutionMethod, payload.ResolutionAddress+payload.ResolutionRoute, buffer)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := httpClient.Do(req)
	// log.Println(nameServer+"status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// Work / inspect body. You may even modify it!

	// And now set a new body, which will simulate the same data we read:
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	ser.EResponse.ResponseWithJSON(w, r, resp.StatusCode, map[string]interface{}{
		"status": resp.Status,
		"data":   string(body),
	}, "application/json")
	return
}
