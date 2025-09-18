# Framework Template

A complete starter template for web applications built with the [bencbradshaw/framework](https://github.com/bencbradshaw/framework).

## Features

- 🚀 **Modern Go Web Framework** - Built on the bencbradshaw/framework
- 🎨 **Frontend Integration** - TypeScript, ESBuild bundling, hot reload
- 🔒 **Authentication Ready** - Middleware patterns for auth implementation
- 📊 **Request Logging** - Built-in HTTP request/response logging
- 🏗️ **Auto-routing** - Automatic route registration from templates
- 📱 **Responsive Design** - Mobile-first CSS framework
- 🐳 **Docker Support** - Ready for containerized deployment

## Quick Start

### Prerequisites

- Go 1.21+
- Node.js 18+ (for frontend dependencies)

### Installation

1. **Clone this template** (or use as GitHub template)

```bash
git clone <your-repo-url>
cd framework-template
```

2. **Install dependencies**

```bash
make install
```

3. **Start development server**

```bash
make dev
```

4. **Open browser** to [http://localhost:2026](http://localhost:2026)

## Development

### Available Commands

```bash
make dev          # Start development server with hot reload
make run          # Start production server
make build        # Build frontend assets for production
make install      # Install Go and Node.js dependencies
make clean        # Clean build artifacts
make build-binary # Build production binary
```

### Project Structure

```
framework-template/
├── main.go              # Application entry point
├── middleware/          # HTTP middleware
│   └── logging.go       # Request logging middleware
├── templates/           # HTML templates (auto-routed)
│   ├── base.html        # Base layout
│   ├── index.html       # Home page
│   ├── about.html       # About page
│   └── *.html           # Other pages
├── frontend/            # Frontend source code
│   ├── src/
│   │   └── index.ts     # TypeScript entry point
│   └── package.json     # Node.js dependencies
├── static/              # Built assets (auto-generated)
└── docker-compose.yml   # Local development with Docker
```

### Development Workflow

1. **Edit Go files** - Server restarts automatically in dev mode
2. **Edit templates** - Changes reflected immediately
3. **Edit frontend** - Assets rebuild automatically
4. **View logs** - All requests logged with timing information

### Adding Routes

**Option 1: Template-based (automatic)**

- Create `templates/newpage.html` → Available at `/newpage`
- Create `templates/admin.subroute.html` → Available at `/admin/*` (SPA)

**Option 2: Programmatic**

```go
mux.Handle("/api/users", middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Your handler logic
})))
```

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

- `APP_ENV=production` - Sets production mode
- `PORT=8080` - Server port (default: 2026)

## Customization

### Authentication

The template includes a basic auth middleware example. Replace with your preferred auth system:

- JWT tokens
- OAuth providers
- Database sessions
- External auth services

### Styling

- Edit `frontend/src/index.css` for global styles
- Use the built-in CSS framework or replace with your preferred solution
- Templates use Go template syntax with automatic escaping

### Database

Add your preferred database:

```go
// Example: Add to main.go
db := setupDatabase() // Your database setup
mux.Handle("/api/users", middleware.LoggingMiddleware(userHandler(db)))
```

## Learn More

- [Framework Documentation](https://github.com/bencbradshaw/framework)
- [Go Templates Guide](https://pkg.go.dev/text/template)
- [ESBuild Documentation](https://esbuild.github.io/)

## License

MIT License - see LICENSE file for details.
