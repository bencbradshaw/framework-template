package main

import (
	"net/http"

	"github.com/bencbradshaw/framework"
)

func main() {
	mux := framework.Run(nil)

	mux.Handle("/shop", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		framework.RenderWithHtmlResponse(w, "shop.custom.html", map[string]any{
			"Title": "Shop",
			"Body":  "Welcome to the shop!",
			"Items": []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"},
		})
	}))
	print("Server started on http://localhost:2026\n")
	http.ListenAndServe(":2026", mux)
}
