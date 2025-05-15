package url

import (
	"fmt"
	"io"

	"net/http"
	"net/url"

	"crypto/rand"
	"encoding/base64"

	"encoding/json"
)

const urlBase string = "http://s.dev/"

type RequestError struct {
	Msg    string
	Status int
}

func (re *RequestError) Error() string {
	return re.Msg
}

type Request struct {
	Url string
}

type Response struct {
	Url string `json:"short_url"`
}

type Storage interface {
	idempotentSave(string, string)
}

type DB struct {
	storage map[string]string
}

func (db DB) idempotentSave(url string, shortUrl string) {
	db.storage[url] = shortUrl
}

func generateShortURL() string {
	b := make([]byte, 6)
	rand.Read(b)
	return urlBase + base64.URLEncoding.EncodeToString(b)[:6]
}

func parseRequest(reqBody io.Reader, r *Request) *RequestError {

	err := json.NewDecoder(reqBody).Decode(r)

	if err != nil {
		return &RequestError{Status: http.StatusBadRequest, Msg: err.Error()}
	}

	_, err = url.ParseRequestURI(r.Url)

	if err != nil {
		return &RequestError{Status: http.StatusBadRequest, Msg: err.Error()}
	}

	return nil
}

func generate(db Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var r Request

		// the basic mux doesn't handle VERBS
		if req.Method != http.MethodPost {
			http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

		res.Header().Set("Content-Type", "application/json")

		parsingErr := parseRequest(req.Body, &r)

		if parsingErr != nil {

			jsonResp, e := json.Marshal(parsingErr)

			if e != nil {
				http.Error(res, "server error", http.StatusInternalServerError)
			}

			res.WriteHeader(parsingErr.Status)
			res.Write(jsonResp)
			return

		}

		rr := Response{Url: generateShortURL()}

		jsonResp, err := json.Marshal(rr)
		if err != nil {
			http.Error(res, "server error", http.StatusInternalServerError)
		}

		db.idempotentSave(r.Url, rr.Url)
		res.Write(jsonResp)

	}
}

func Run(out io.Writer) int {
	fmt.Fprintln(out, "Hello, World!")

	// The job of mux or multiplexer is to route request to the registered handler
	// based upon API signature (and also request Method sometimes).
	// If the signature and its handler is not registered with the mux, it raises a 404
	mux := http.NewServeMux()

	db := &DB{
		storage: make(map[string]string),
	}

	mux.HandleFunc("/", generate(*db))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()

	return 0
}
