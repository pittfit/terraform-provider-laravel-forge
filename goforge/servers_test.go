package goforge

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

// testServersList
func TestServersList(t *testing.T) {
	client, ts, err := createTestClient(ServersListSuccessfulResponse(t))

	if err != nil {
		t.Errorf("Unable to create test client: %v", err)
	}

	defer ts.Close()

	result, err := client.ServersList()
	if err != nil {
		t.Errorf("Error getting credentials: %v", err)
	}

	expected := []Server{
		Server{ID: 1, CredentialID: 1, Name: "test-via-api", PHPVersion: "php71", IsReady: true},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ServersList returned %+v, expected %+v", result, expected)
	}
}

func ServersListSuccessfulResponse(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got ‘%s’", r.Method)
		}

		if r.URL.EscapedPath() != "/api/v1/servers" {
			t.Errorf("Expected request to ‘/api/v1/servers, got ‘%s’", r.URL.EscapedPath())
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		io.WriteString(w, `{
			"servers":[
				{
					"id": 1,
					"credential_id": 1,
					"name": "test-via-api",
					"php_version": "php71",
					"revoked": false,
					"is_ready": true
				}
			]
		}`)
	})
}