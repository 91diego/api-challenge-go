# Code Challenge
## REST API

## Features

- Autheticate using JWT.
- Get claims from authenticated user.
- Web Scraping to obtain the links of a page.

## Tech

- [Golang] - GO!

## Installation

Rest Api requires [golang.org](https://golang.org/doc/install).

Install all dependencies and plugins.
```sh
go mod tidy
```

Update or create vendor folder.
```sh
go mod vendor
```

Create .env file.

```sh
cp .env.example .env
```

Run server.

```sh
run go *.go
```

Verify the deployment by navigating to your server address in
your preferred browser.

```sh
http://localhost:8080
```

You can run unit testing with

```sh
php artisan test
```

## License

MIT

**Free Software, Hell Yeah!**

