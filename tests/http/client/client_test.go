package client

import (
	"testing"

	"go-command/pkg/utils/http"
)

func TestClient(t *testing.T) {

	client := http.Client{}
	client.Init("https://www.google.com")
	client.SetTimeout(1)
	_, status := client.Get()

	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

}
