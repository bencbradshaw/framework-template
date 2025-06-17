package main

import (
	"fmt"
	"net/http"

	"github.com/bencbradshaw/framework"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("AUTH CHECK: [%s %s]\n", r.Method, r.URL.Path)
		_, err := r.Cookie("framework")
		if err != nil {
			framework.RenderWithHtmlResponse(w, "error.html", map[string]any{
				"title": "Unauthorized",
				"code":  401,
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := framework.Run(framework.InitParams{
		AuthGuard:                  AuthMiddleware,
		IsDevMode:                  true,
		AutoRegisterTemplateRoutes: true,
	})

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
