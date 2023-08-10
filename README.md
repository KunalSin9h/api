# api

Api backend for my personal website (kunalsin9h.com)

## Docs

### Setup

1. Clone and cd into the repo

```bash
git clone https://github.com/KunalSin9h/api
cd api
```

2. Run the local server

```bash

export HOST=127.0.0.1
export PORT=9999

go run cmd/api/*.go
```

> The `HOST` is where the server servers, default value of `HOST` is `127.0.0.1`
> The `PORT` is where the server listens, default value of `PORT` is `9999`

### Run using Docker

```bash
docker run --name api -d -p 9999:9999 ghcr.io/kunalsin9h/api:latest
```

### API Endpoints

1. Generate Image for **Blog Post**

> This is for OpenGraph and twitter Card of SEO
> The background image and font are in `assets` folder

**GET** /v1/image/`:title`

Returns a image, whose `Content-Type` is `image/jpeg`
