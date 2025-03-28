package main

import (
	"net/http"

	"github.com/bencbradshaw/framework"
)

func main() {
	http.ListenAndServe(":2026", framework.Run(nil))
}
