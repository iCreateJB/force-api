package main

import (
  "fmt"
  "net/http"
	"net/http/httptest"
  "bufio"
  "bytes"
  // "fmt"
  "log"
  "os"
  "regexp"
  "testing"
  "encoding/json"
)

type Response struct { Message string `json:"message"` }

func TestIndexReturnsWithStatusOk(t *testing.T){
  request, _ := http.NewRequest("GET", "/", nil)
  response := httptest.NewRecorder()

  IndexHandler(response, request)

  if response.Code != http.StatusOK {
      t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
  }
}

func TestMoralsReturnsWithStatusOk(t *testing.T){
  request, _ := http.NewRequest("GET", "/morals", nil)
  response := httptest.NewRecorder()

  MoralHandler(response, request)

  if response.Code != http.StatusOK {
      t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
  }
}

func TestIndexHasNoJSON(t *testing.T){
  request, _ := http.NewRequest("GET", "/", nil)
  response := httptest.NewRecorder()

  IndexHandler(response,request)

  var resp Response
  json.NewDecoder(response.Body).Decode(&resp)

  if resp.Message != "" {
    t.Fatalf("Response body did contain a message")
  }
}

func TestMoralsJSONHasMessage(t *testing.T){
  request, _ := http.NewRequest("GET", "/morals", nil)
  response := httptest.NewRecorder()

  MoralHandler(response, request)

  var resp Response
  json.NewDecoder(response.Body).Decode(&resp)

  if resp.Message == "" {
    t.Fatalf("Response body did not contain a message")
  }
}

func TestCommonHeaders(t *testing.T){
  testHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}

  ts := httptest.NewServer(commonHeaders(testHandler))
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

	testHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}

	ts := httptest.NewServer(logHandler(testHandler))
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
