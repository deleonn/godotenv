package godotenv

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
    Config(".env.example")

    host := os.Getenv("HOST")
    port := os.Getenv("PORT")

    if host != "localhost=123" {
         t.Errorf("got %q, wanted %q", host, "localhost=123")
    }

    if port != "8080" {
         t.Errorf("got %q, wanted %q", port, "8080")
    }
}
