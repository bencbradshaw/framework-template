# Framework Template

A complete starter template for web applications built with the [bencbradshaw/framework](https://github.com/bencbradshaw/framework).

## Quick Start

### Prerequisites

- Go 1.21+
- Node.js 18+ (required only for `npm install`)

### Installation

1. **Use this template**

Create a new repo from this template. A github action will create a PR to rename various references throughout your new repo.

2. **Clone your repo**

```bash
git clone your-new-repo
```

3. **Install dependencies**

```bash
make install
```

4. **Start development server**

```bash
make
```

4. **Open browser** to [http://localhost:2026](http://localhost:2026)

## Developer Scenario

Picture this: You're a developer tasked with building a complete web presence for a growing company. They need a lightning-fast marketing website that Google can easily index, a secure user application with rich interactive features, and an admin panel for managing everything behind the scenes. Traditionally, this would mean juggling multiple technologies: a static site generator for marketing pages, a separate Node.js server for your JavaScript applications, complex build processes, and headaches managing different deployment pipelines. But what if you could build all of this with a single Go binary that handles everything?

Enter Framework+ Framework Template. A clean approach that lets you serve SEO-optimized HTML pages alongside dynamic JavaScript applications, all from one streamlined Go process. When users land on your marketing site, Framework delivers blazing-fast static HTML that search engines love. When they're ready to sign up, the same server seamlessly handles authentication and serves your interactive React, Vue, or vanilla JavaScript applications. No need for a separate Node.js process, no complex orchestration between services, and no deployment nightmares. Framework's built-in TypeScript compilation and hot-reload development server means you can build modern, type-safe frontend applications while keeping the simplicity of a single Go binary.

The magic happens during development and deployment—Framework compiles your TypeScript, bundles your assets, and serves everything through intelligent routing that knows when to deliver static HTML for SEO and when to serve your JavaScript applications for authenticated users. Deploy this single binary to Google Cloud Run's free tier, and you have a complete web application that can handle thousands of users without breaking the bank. No container orchestration, no microservice complexity, just one process that scales beautifully and costs pennies to run. Framework gives you the developer experience of modern JavaScript tooling with the operational simplicity and performance of Go, letting you focus on building features instead of managing infrastructure.
Company Web Application Architecture

## Your Website Architecture

```txt
                                    ┌─────────────────┐
                                    │   Google Bot    │
                                    │   (Indexing)    │
                                    └─────────┬───────┘
                                              │
                                              ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                          PUBLIC MARKETING WEBSITE                          │
│                               (Plain HTML)                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐            │
│  │  Landing Page   │  │   About Page    │  │  Features Page  │            │
│  │   (index.html)  │  │  (about.html)   │  │ (features.html) │   ...more  │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘            │
│                                                                             │
│  Features:                                                                  │
│  • SEO Optimized                                                           │
│  • Fast Loading                                                            │
│  • Google Indexable                                                        │
│  • No Authentication Required                                              │
└─────────────────────────┬───────────────────────────────────────────────────┘
                          │
                          │ User clicks "Get Started" / "Login"
                          ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                        AUTHENTICATION GATEWAY                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────┐                    ┌─────────────────┐                 │
│  │  Login Page     │◄──────────────────►│  Signup Page    │                 │
│  │                 │                    │                 │                 │
│  └─────────────────┘                    └─────────────────┘                 │
│                                                                             │
│  Features:                                                                  │
│  • User Authentication                                                      │
│  • Session Management                                                       │
│  • Role-based Access Control                                                │
└─────────────────────────┬───────────────────────────────────────────────────┘
                          │
                          │ Authentication Success
                          ▼
                    ┌─────────────┐
                    │   Router    │
                    │ (Role-based)│
                    └─────┬───────┘
                          │
                ┌─────────┴─────────┐
                ▼                   ▼
┌─────────────────────────┐    ┌─────────────────────────┐
│     USER APPLICATION    │    │    ADMIN APPLICATION    │
│    (JavaScript SPA)     │    │    (JavaScript SPA)     │
├─────────────────────────┤    ├─────────────────────────┤
│                         │    │                         │
│ ┌─────────────────────┐ │    │ ┌─────────────────────┐ │
│ │    Dashboard        │ │    │ │   User Management   │ │
│ └─────────────────────┘ │    │ └─────────────────────┘ │
│                         │    │                         │
│ ┌─────────────────────┐ │    │ ┌─────────────────────┐ │
│ │    Profile          │ │    │ │   Analytics         │ │
│ └─────────────────────┘ │    │ └─────────────────────┘ │
│                         │    │                         │
│ ┌─────────────────────┐ │    │ ┌─────────────────────┐ │
│ │    Features         │ │    │ │   System Settings   │ │
│ └─────────────────────┘ │    │ └─────────────────────┘ │
│                         │    │                         │
│ Features:               │    │ Features:               │
│ • Full App Experience   │    │ • Limited Access        │
│ • Protected Routes      │    │ • Admin Only Features   │
│ • User-specific Data    │    │ • System Management     │
└─────────────────────────┘    └─────────────────────────┘
```

## Development

## Features

- 🚀 **Modern Go Web Framework** - Built on the bencbradshaw/framework v0.0.14
- 🏗️ **Auto-routing** - Automatic route registration from templates
- 🎯 **SPA Support** - Subroute templates for single-page application sections
- 🔐 **Protected Routes** - API endpoints with authentication middleware
- 🎨 **Frontend Integration** - TypeScript with go-web-framework package, ESBuild bundling, hot reload
- 🔒 **Authentication System** - Sample login/signup flow with cookie-based sessions
- 📊 **Request Logging** - Sample HTTP request/response logging middleware
- � **Shop Demo** - Example shop implementation with product display
- 🐳 **Docker Support** - Multi-stage Dockerfile and docker-compose for development

### Available Commands

```bash
make dev          # Start development server with hot reload
make build        # Build frontend assets for production
make install      # Install Go and frontend dependencies
make clean        # Clean build artifacts
make build-binary # Build production binary
```

### Project Structure

```
framework-template/
├── main.go              # Application entry point with route setup
├── go.mod               # Go dependencies (framework v0.0.14, esbuild)
├── .env.example         # Environment configuration template
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Development environment setup
├── makefile             # Build and development commands
├── middleware/          # HTTP middleware components
│   ├── auth.go          # Authentication middleware with cookie validation
│   └── logging.go       # Request logging with timing information
├── auth/                # Authentication system
│   ├── handlers.go      # Login/signup form handling and processing
│   └── utils.go         # Authentication utilities
├── api/                 # RESTful API endpoints
│   └── handlers.go      # Protected user API with JSON responses
├── shop/                # Shop functionality demo
│   └── handlers.go      # Product listing and shop page rendering
├── templates/           # HTML templates (auto-routed)
│   ├── base.html        # Base layout template
│   ├── index.html       # Home page
│   ├── about.html       # About page
│   ├── login.custom.html # Custom login form
│   ├── signup.custom.html # Custom signup form
│   ├── shop.custom.html # Shop page with product display
│   ├── account.subroute.auth.html # Protected account SPA
│   ├── entry.html       # Entry page template
│   └── error.html       # Error page template
├── frontend/            # Frontend source code
│   ├── src/
│   │   ├── index.ts     # TypeScript entry with route-based code splitting
│   │   ├── index.css    # Global styles
│   │   └── account/     # Account page components
│   ├── package.json     # Node.js dependencies (go-web-framework)
│   └── tsconfig.json    # TypeScript configuration
├── static/              # Built assets (auto-generated by ESBuild)
└── .github/
    └── workflows/
        └── template-init.yml # Auto-replace template name on first use
```

### Development Workflow

1. **Edit Go files** - Server does not automatically restart. You can use the restart strategy of your choice.
2. **Edit templates** - Changes reflected on next browser refresh or page load
3. **Edit frontend** - any change to a file in `frontend/src` will autorebuild. if you are listening to the 'sse' event, the browser will reload automatically.
4. **View logs** - All requests logged with time and method + status + time

### Adding Routes

The template demonstrates multiple routing approaches:

**Option 1: Template-based (automatic)**

- Create `templates/newpage.html` → Available at `/newpage`
- Create `templates/admin.subroute.auth.html` → Available at `/admin/*` (protected SPA)
- Create `templates/my-spa.subroute.html` -> available at `/my-spa/*` (unprotected SPA)
- Templates with `.custom.html` extension require custom handlers (like login/signup), and will not be automatically registered

**Option 2: Programmatic with handlers**

```go
// Simple route with logging
mux.Handle("/shop", middleware.LoggingMiddleware(shop.Handler()))

// Protected route with authentication
mux.Handle("/api/user", middleware.LoggingMiddleware(middleware.AuthMiddleware(api.UserHandler())))

// Authentication routes with method handling
mux.Handle("/login", middleware.LoggingMiddleware(auth.LoginHandler()))
mux.Handle("GET /logout", middleware.LoggingMiddleware(auth.LogoutHandler()))
```

**Implemented Routes:**

- `/` - Home page (auto-routed from `index.html`)
- `/about` - About page (auto-routed from `about.html`)
- `/shop` - Product showcase with custom handler
- `/login` - Login form (GET) and submission (POST)
- `/signup` - Signup form (GET) and submission (POST)
- `/logout` - Logout handler (GET only)
- `/account/*` - Protected SPA section (requires authentication)
- `/api/user` - Protected user API endpoint

## Deployment

### Docker Deployment

1. **Build and run with Docker Compose**

```bash
docker-compose up --build
```

### Manual Deployment

1. **Build for production**

```bash
make build
make build-binary
```

2. **Deploy binary** with `static/` folder to your server

### Environment Variables

Create a `.env` file based on `.env.example`:

```bash
# Core application settings
APP_ENV=development    # or 'production' for production mode
PORT=2026             # Server port (default: 2026)

# Add your configuration as needed:
# DATABASE_URL=
# JWT_SECRET=
# EXTERNAL_API_KEY=
```

## Customization

### Authentication

The template includes a **sample authentication system**:

- **Cookie-based sessions** - Simple session management via HTTP cookies
- **Login/Signup forms** - Ready-to-use authentication pages
- **Protected routes** - Middleware automatically redirects unauthenticated users
- **User context** - Authenticated user ID available in request context

**Current Implementation:**

- Cookie value stores user ID (email) directly
- No password hashing (demo purposes only)
- Session stored in browser cookie named "framework"

**For production, enhance with:**

- JWT tokens with proper signing/validation
- Database-backed user storage with password hashing
- OAuth providers (Google, GitHub, etc.)
- Secure session management with database storage

### Frontend Architecture

The template uses **route-based code splitting**:

```typescript
// index.ts - Entry point
switch (path) {
  case '/login':
    // No additional JS needed
    break;
  case '/account':
    await import('./account/app-root.js'); // Load account-specific code
    break;
}
```

**Technologies:**

- TypeScript with `go-web-framework` package
- ESBuild for bundling and hot reload - managed through go, not node
  - CSS imports processed automatically
  - Code splitting for optimized loading

### Styling

- Edit `frontend/src/index.css` for global styles
- Component-specific CSS in respective directories
- Templates use Go template syntax with automatic escaping
- Built assets output to `static/` directory

## Learn More

- [Framework Documentation](https://github.com/bencbradshaw/framework)
- [Go Templates Guide](https://pkg.go.dev/text/template)
- [ESBuild Documentation](https://esbuild.github.io/)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)

## License

MIT License - see LICENSE file for details.
