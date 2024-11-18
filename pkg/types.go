package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func DefaultResponse() []byte {
	res := &Response{
		Message: "Operation Successful",
		Status:  http.StatusOK,
	}

	result, err := json.Marshal(&res)
	if err != nil {
		log.Panic(err)
	}
	return result
}
