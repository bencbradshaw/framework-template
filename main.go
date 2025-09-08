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

// build compiles frontend assets for production deployment.
// Uses ESBuild to bundle TypeScript, minify code, and output to static directory.
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
	// Handle build command for production asset compilation
	// This builds and bundles frontend assets into the /static directory
	if len(os.Args) > 1 && os.Args[1] == "build" {
		build()
		return
	}
	
	// Initialize the framework with authentication and auto-routing enabled
	// Auto-routing creates routes based on template filenames (e.g., about.html -> /about)
	mux := framework.Run(framework.InitParams{
		AuthGuard:                  middleware.AuthMiddleware, // Protect routes that need authentication
		AutoRegisterTemplateRoutes: true,                      // Enable automatic template-based routing
	})

	// Register custom route handlers for complex business logic
	// Shop handler demonstrates template rendering with dynamic data
	mux.Handle("/shop", middleware.LoggingMiddleware(shop.Handler()))

	// Authentication routes for user login/logout functionality
	// These handle both GET (show forms) and POST (process submissions)
	mux.Handle("/login", middleware.LoggingMiddleware(auth.LoginHandler()))
	mux.Handle("/signup", middleware.LoggingMiddleware(auth.SignupHandler()))
	mux.Handle("GET /logout", middleware.LoggingMiddleware(auth.LogoutHandler()))

	// Protected API endpoints that require authentication
	// AuthMiddleware ensures only logged-in users can access these routes
	mux.Handle("/api/user", middleware.LoggingMiddleware(middleware.AuthMiddleware(api.UserHandler())))

	print("Server started on http://localhost:2026\n")
	http.ListenAndServe(":2026", mux)
}
