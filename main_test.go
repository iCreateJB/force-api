package main

import (
  "fmt"
  "net/http"
	"net/http/httptest"
  "bufio"
  "bytes"
  "log"
  "os"
  "regexp"
  "testing"
  "encoding/json"
)

type Response struct {
  Message string `json:"message"`
  Key     string `json:"key"`
}

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

func TestNewMoralHandler(t *testing.T){
  type Moral struct {
    Message string `json:"message"`
    Key     int `json:"moral_id"`
  }

  moral := Moral{ Message: "Simple Test", Key: 2 }
  req, _:= json.Marshal(moral)
  request,_ := http.NewRequest("POST", "/morals", bytes.NewReader(req))
  response  := httptest.NewRecorder()

  NewMoralHandler(response, request)

  var resp Moral
  json.NewDecoder(response.Body).Decode(&resp)

  if response.Code != http.StatusCreated {
    t.Fatalf("Request did not get created.")
  }

  if resp.Message != moral.Message {
    t.Fatalf("Request did not return the generated message.")
  }

  if resp.Key != moral.Key {
    t.Fatalf("Request did not return the user supplied key.")
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

func TestPortNumber(t *testing.T){
  var portNo string
  portNo = portNumber()
  if portNo != "5000" {
    t.Errorf("Default port number was not set")
  }
}

func TestPortNumberOS(t *testing.T){
  var portNo string
  os.Setenv("PORT","12345")
  portNo = portNumber()
  if portNo != "12345" {
    t.Errorf("Port Number was not set")
  }
}
