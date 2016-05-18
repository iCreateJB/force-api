package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
)

type Response struct {
	Id       int
	Quote    string  `json:"quote"`
	Category string  `json:"category"`
	Key      string  `json:"request_id"`
	Errors   []Error `json:"errors"`
}

func TestIndexReturnsWithStatusOk(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	IndexHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}
}

// func TestIndexReturnsEmptyMoral(t *testing.T) {
// 	var resp Response
// 	request, _ := http.NewRequest("GET", "/", nil)
// 	response := httptest.NewRecorder()
//
// 	IndexHandler(response, request)
//
// 	json.NewDecoder(response.Body).Decode(&resp)
//
// 	if response.Code != http.StatusOK {
// 		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
// 	}
// }

func TestMoralsReturnsWithStatusOk(t *testing.T) {
	request, _ := http.NewRequest("GET", "/morals", nil)
	response := httptest.NewRecorder()

	MoralHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}
}

func TestNewMoralHandler(t *testing.T) {
	var resp Response
	moral := Response{Quote: "Simple Test", Category: "art"}
	req, _ := json.Marshal(moral)
	request, _ := http.NewRequest("POST", "/morals", bytes.NewReader(req))
	response := httptest.NewRecorder()

	NewMoralHandler(response, request)

	json.NewDecoder(response.Body).Decode(&resp)

	if response.Code != http.StatusCreated {
		t.Fatalf("Request did not get created.")
	}

	if resp.Quote != moral.Quote {
		t.Fatalf("Request did not return the generated message.")
	}
}

func TestNewMoralHandlerInvalidMessage(t *testing.T) {
	var resp = new(Response)
	moral := Response{Category: "music"}
	req, _ := json.Marshal(moral)
	request, _ := http.NewRequest("POST", "/morals", bytes.NewReader(req))
	response := httptest.NewRecorder()

	NewMoralHandler(response, request)

	json.NewDecoder(response.Body).Decode(&resp)

	if response.Code != http.StatusBadRequest {
		t.Fatal("Request not recorded correctly")
	}

	if len(resp.Errors) == 0 {
		t.Fatal("Error was not recorded on Message")
	}
}

func TestIndexHasNoJSON(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	IndexHandler(response, request)

	var resp Response
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Quote != "" {
		t.Fatalf("Response body did contain a message")
	}
}

func TestMoralsJSONHasMessage(t *testing.T) {
	request, _ := http.NewRequest("GET", "/morals", nil)
	response := httptest.NewRecorder()

	MoralHandler(response, request)

	var resp Response
	json.NewDecoder(response.Body).Decode(&resp)

	if resp.Quote == "" {
		t.Fatalf("Response body did not contain a message")
	}
}

func TestCommonHeaders(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake twitter json string"}`)
	})

	ts := httptest.NewServer(commonHeaders(handler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Header.Get("Accept") != "application/vnd.morals."+appVersionStr+"+json" {
		t.Errorf("commonHeader should specify version via accept header")
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("commonHeader should specific Content-Type as application/json")
	}
}

func TestLogHandler(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	log.SetOutput(w)
	defer log.SetOutput(os.Stderr)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake twitter json string"}`)
	})

	ts := httptest.NewServer(logHandler(handler))
	defer ts.Close()

	_, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	w.Flush()

	re := regexp.MustCompile(`^(\d\d\d\d/\d\d/\d\d) (\d{1,2}:\d\d:\d\d) (\[\w+\]) ("/")`)
	if !re.Match(b.Bytes()) {
		t.Errorf("logHandler wrote %q", b.String())
	}
}

func TestPortNumber(t *testing.T) {
	expected := "5000"
	result := portNumber()
	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}
