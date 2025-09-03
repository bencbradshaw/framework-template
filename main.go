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
		AuthGuard:                  middleware.AuthMiddleware,
		AutoRegisterTemplateRoutes: true,
	})

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
	mux.Handle("/api/user", middleware.LoggingMiddleware(middleware.AuthMiddleware(api.UserHandler())))

	print("Server started on http://localhost:2026\n")
	http.ListenAndServe(":2026", mux)
}
