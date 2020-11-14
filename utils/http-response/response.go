package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CafeLucuma/go-play/utils/logging"
)

type Response struct {
	StatusCode int
	Body       interface{}
	Headers    http.Header
}

type GenericBody struct {
	Code    int    `json:"code,ommitempty"`
	Message string `json:"message"`
}

func NewGenericBody(code int, message string) GenericBody {
	return GenericBody{
		Code:    code,
		Message: message,
	}
}

func New(s int, b interface{}, h http.Header) Response {
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
		log.Printf("Cant encode '%s' to JSON", r.Body)
		return err
	}

	r.Body = encodedBody

	return nil
}

func (r *Response) RespondJSON(w http.ResponseWriter) error {

	if err := r.encodeBodyToJson(); err != nil {
		logging.Error.Printf("Error encoding json in RespondJson: %s", err)
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
