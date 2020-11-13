package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       interface{}
	Headers    http.Header
}

func NewResponse(s int, b interface{}, h http.Header) Response {
	return Response{
		StatusCode: s,
		Body:       b,
		Headers:    h,
	}
}

func (r *Response) encodeBodyToJson() error {

	encodedBody := new(bytes.Buffer)
	encoder := json.NewEncoder(encodedBody)
	if err := encoder.Encode(r.Body); err != nil {
		log.Printf("Cant encode '%+v' to JSON", r.Body)
		return err
	}

	r.Body = encodedBody

	return nil
}

func (r *Response) RespondJSON(w http.ResponseWriter) error {

	if err := r.encodeBodyToJson(); err != nil {
		log.Printf("Error encoding json in RespondJson: %s", err.Error())
		return err
	}

	for key, values := range r.Headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	fmt.Fprintf(w, "%s", r.Body)
	return nil
}
