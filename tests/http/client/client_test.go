package client

import (
	"testing"

	"go-command/pkg/utils/http"

	"github.com/valyala/fasthttp"
)

func TestClient(t *testing.T) {

	client := http.NewClient(&fasthttp.Client{})
	client.SetTimeout(1)
	_, status := client.Get("https://www.google.com")

	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

}
