## Overview

This is a Go web server project designed for learning and experimentation with RESTful APIs, long-running service workers, RSS feed processing, and PostgreSQL database integration.

## Learning Objectives

By building this project, you'll gain experience with:

- **Go Web Development**: Creating RESTful APIs, handling HTTP requests and responses, and working with middleware.
- **RSS Feed Handling**: Fetching, parsing, and processing RSS feeds using Go libraries.
- **Database Integration**: Storing and retrieving data from a PostgreSQL database, including migrations and type-safe queries.
- **Concurrency**: Managing concurrent service workers for efficient feed processing.
- **Tooling**: Using tools like Air for hot reloading, Goose for database migrations, and SQLC for generating type-safe database code.

## Getting Started 

### System Requirements

- Go: 1.22 or higher
- Air: (for hot-reloading)
- PostgreSQL: 13 or higher
- Goose: 3.1.0 or higher

### Installation

1. Clone this repository: `git clone github.com/ggrangel/blog-aggregator`
2. Install dependencies: `go mod download`
3. Set environment variables in `.env`:
```
    DATABASE_URL=<database-url>
    PORT=<port-number>
```
4. Create migration files: `goose postgres "<database-driver>://<username>:<password>@localhost:5432/<database-name>" up`

5. Review the migration files and run the migrations against your database if everything is ok.

### Usage

1. Start the server: air -c .air.toml (by default, runs on port 8080)

2. Refer to the comments in main.go for basic API usage examples and available routes.

