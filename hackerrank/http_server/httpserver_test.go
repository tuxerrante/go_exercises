package main

//
// --- NOT WORKING
//
// Would require a major refactoring of main module
// like a dedicated server function to invoke from testing functions
// 	 https://github.com/quii/learn-go-with-tests/blob/162b4b9c3e00e2f89ec4d2929278f904cf008db5/http-server/v5/server.go#L20
// Leaving this here for reference.
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type HttpPayload struct {
	Type    string `json:"type"`
	Payload Lake   `json:"payload"`
}

var postAction Action = Action{
	Type:    "post",
	Payload: "{\"id\":\"id0000\", \"name\":\"TestLake\", \"area\":100}\"}",
}

func TestPostGet(t *testing.T) {

	t.Run("Post a payload without any object", func(t *testing.T) {
		http.NewRequest(http.MethodPost, "/post", nil)
		response := httptest.NewRecorder()
		log.Printf("response.Code=%d", response.Code)
		assertInt(t, response.Code, http.StatusOK)
	})

	t.Run("Post a payload with one Lake object", func(t *testing.T) {
		// json_payload, err := json.Marshal(postPayload)
		// if err != nil {
		// 	t.Errorf("Can't convert payload to JSON. %v", err)
		// }
		// log.Printf("Payload=%v", postPayload)
		respPost, err := http.Post(Address+"/post", "application/json", strings.NewReader(postAction.Payload))
		if err != nil {
			t.Errorf("Can't read response from Post call: %s", err)
		}
		defer respPost.Body.Close()
		log.Println(respPost.Body)
		log.Println(respPost.StatusCode)

		http.NewRequest(http.MethodGet, "/get?id=id0000", nil)
		respGet := httptest.NewRecorder()

		gotBytes, _ := io.ReadAll(respGet.Body)
		var got Lake
		json.Unmarshal(gotBytes, &got)
		log.Println(got)
		want := "id0000"
		assertString(t, got.Id, want)
	})

	t.Run("Parallel posts with same id", func(t *testing.T) {
		assertBool(t, true, false)
	})
}

// Assert function for simple data types comparable through ==
func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got a different result: %q != %q", got, want)
	}
}
func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got a different result: %d != %d", got, want)
	}
}
func assertBool(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("Got a different result: %v != %v", got, want)
	}
}
