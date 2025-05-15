package url

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestIdempotentSave(t *testing.T) {

	testStorage := make(map[string]string)

	var db = DB{
		storage: testStorage,
	}

	db.idempotentSave("test", "val1")
	db.idempotentSave("test", "val1")

	assert.Equal(t, 1, len(testStorage), "The save is idempotent")
}

func TestParseRequestTableDriven(t *testing.T) {

	tests := []struct {
		isError 		bool
		payload     []byte
		result      *RequestError
		description string
	}{
		{isError: false, payload: []byte(`{"url": "http://google.com"}`), result: nil, description: "No error for a correct URL" },
		{isError: true, payload: []byte(`{"url": "mmm"}`), result: &RequestError{Status: 400}, description: "Parsing error for invalid JSON"},
		{isError: true, payload: []byte(`{"url": "mmm"`), result: &RequestError{Status: 400}, description: "Unmarshaling error for invalid JSON"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			var r Request

			err := parseRequest(bytes.NewBuffer(tt.payload), &r)

			if tt.isError {
				var re *RequestError
				errors.As(err, &re)
	
				if re.Status != tt.result.Status {
					t.Errorf("got %d, want %d", re.Status, tt.result.Status)
				}
			}

			if !tt.isError {
				if tt.result != nil {
					t.Errorf("got %d, want no error", tt.result.Status)
				}
			}
		})
	}
}

type DBMock struct {
	storage map[string]string
}

func (db DBMock) idempotentSave(url string, shortUrl string) {
	db.storage[url] = shortUrl
}

func TestGenerateTableDriven(t *testing.T) {

	testStorage := make(map[string]string)

	dbMock := DBMock{
		storage: testStorage,
	}

	testUrl := "http://google.com/very-long-url"

	tests := []struct {
		payload     []byte 
		responseCode     int
		description string
	}{
		{payload: []byte(fmt.Sprintf(`{"url": "%s"}`, testUrl)), responseCode: 200, description: "No error for a correct URL" },
		{payload: []byte(`{"url": "mmm"}`), responseCode:400 , description: "Parsing error for invalid JSON"},
		{payload: []byte(`mmm`), responseCode:400, description: "Unmarshaling error for invalid JSON"},
	}

	for _, tt := range tests {
		request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(tt.payload))
		response := httptest.NewRecorder()
		generate(dbMock)(response, request)	
		assert.Equal(t, response.Code, tt.responseCode, "Response code should be %n", tt.responseCode)
	
		if tt.responseCode == 200 {
			expectedResponseStruct := Response{Url: testStorage[testUrl]}
			expectedResponse, _ := json.Marshal(expectedResponseStruct)

			assert.Equal(t, expectedResponse, response.Body.Bytes(), "")
		}
	}
}

