package main


import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetNotesHandler(t *testing.T) {

	notes = []Note{
		{"suspension note"},
	}

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getNoteHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Note{"suspension note"}
	b := []Note{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
func TestCreateNotesHandler(t *testing.T) {

	notes = []Note{
		{"suspend note"},
	}

	form := newCreateNoteForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createNoteHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Note{"suspension note"}

	if err != nil {
		t.Fatal(err)
	}

	actual := notes[1]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateNoteForm() *url.Values {
	form := url.Values{}
	form.Set("content", "suspension note")
	return &form
}