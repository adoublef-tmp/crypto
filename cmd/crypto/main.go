package main

import (
	"github.com/temp/crypto/internal/alert"
	"github.com/temp/crypto/internal/net/http"
)

func main() {
	app := http.Handler(alert.Alerter{})
	app.Listen(":3000")
}
