package main

import (
	"fmt"
	"net/http"
	"os"

	"framework-template/api"
	"framework-template/auth"
	"framework-template/middleware"

	"github.com/bencbradshaw/framework"
	esbuild "github.com/evanw/esbuild/pkg/api"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("AUTH CHECK: [%s %s]\n", r.Method, r.URL.Path)
		_, err := r.Cookie("framework")
		if err != nil {
			// Redirect to login page instead of showing 401 error
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func build() {
	buildParams := framework.InitParams{
		EsbuildOpts: esbuild.BuildOptions{
			EntryPoints:       []string{"./frontend/src/index.ts"},
			MinifyWhitespace:  true,
			MinifyIdentifiers: true,
			MinifySyntax:      true,
			Sourcemap:         esbuild.SourceMapNone,
			Outdir:            "static",
		},
	}
	framework.Build(buildParams)
	fmt.Println("Build complete")
}

func main() {
	// Handle build command
	if len(os.Args) > 1 && os.Args[1] == "build" {
		build()
		return
	}

	mux := framework.Run(framework.InitParams{
		AuthGuard:                  AuthMiddleware,
		IsDevMode:                  true,
		AutoRegisterTemplateRoutes: false, // Disable auto-registration to prevent conflicts
	})

	// Manually register all routes with proper auth control
	mux.Handle("/", middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		framework.RenderWithHtmlResponse(w, "index.html", map[string]any{
			"title": "Framework Template",
		})
	})))

	mux.Handle("/about", middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		framework.RenderWithHtmlResponse(w, "about.html", map[string]any{
			"title": "About",
		})
	})))

	// Protected account route that redirects to login
	mux.Handle("/account/", middleware.LoggingMiddleware(AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		framework.RenderWithHtmlResponse(w, "account.subroute.auth.html", map[string]any{
			"title": "Account",
		})
	}))))

	mux.Handle("/shop", middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		framework.RenderWithHtmlResponse(w, "shop.custom.html", map[string]any{
			"Title": "Shop",
			"Body":  "Welcome to the shop!",
			"Items": []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"},
		})
	})))

	// Authentication routes
	mux.Handle("/login", middleware.LoggingMiddleware(auth.LoginHandler()))
	mux.Handle("/signup", middleware.LoggingMiddleware(auth.SignupHandler()))
	mux.Handle("GET /logout", middleware.LoggingMiddleware(auth.LogoutHandler()))

	// Example API endpoints (protected by auth middleware)
	mux.Handle("/api/user", middleware.LoggingMiddleware(AuthMiddleware(api.UserHandler())))

	print("Server started on http://localhost:2026\n")
	http.ListenAndServe(":2026", mux)
}
