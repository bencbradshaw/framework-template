package main

import (
	"fmt"
	"net/http"
	"os"

	"framework-template/api"
	"framework-template/auth"
	"framework-template/middleware"
	"framework-template/shop"

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
	// this will bundle the js into the /static dir in preparation for running on a server
	if len(os.Args) > 1 && os.Args[1] == "build" {
		build()
		return
	}
	// render simple html pages. the url is based on the template file name
	// e.g., my-blog.html will be served from /my-blog
	// use the returned mux to handle other routes later on
	mux := framework.Run(framework.InitParams{
		AuthGuard:                  middleware.AuthMiddleware,
		AutoRegisterTemplateRoutes: true,
	})

	// a sample of handling a route that needs more logic than the autoregistered templates
	mux.Handle("/shop", middleware.LoggingMiddleware(shop.Handler()))

	// Authentication routes
	mux.Handle("/login", middleware.LoggingMiddleware(auth.LoginHandler()))
	mux.Handle("/signup", middleware.LoggingMiddleware(auth.SignupHandler()))
	mux.Handle("GET /logout", middleware.LoggingMiddleware(auth.LogoutHandler()))

	// Example API endpoints (protected by auth middleware)
	mux.Handle("/api/user", middleware.LoggingMiddleware(middleware.AuthMiddleware(api.UserHandler())))

	print("Server started on http://localhost:2026\n")
	http.ListenAndServe(":2026", mux)
}
