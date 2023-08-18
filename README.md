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
export MONGODB_URL=mongodb://localhost:27017
export DB_TIMEOUT=5000

go run cmd/api/*
```

> The `HOST` is where the server servers, default value of `HOST` is `127.0.0.1`
> The `PORT` is where the server listens, default value of `PORT` is `9999`
> The `MONGODB_URL` is where the mongodb database is running, the default value is `mongodb://localhost:27017`
> The `DB_TIMEOUT` is the timeout time for each mongodb operation, the default value is `5000`

### Run using Docker

```bash
docker run \
   --name api \
   -d -p 9999:9999 \
   -e HOST=0.0.0.0 \
   -e PORT=9999 \
   -e MONGODB_URL=mongodb://localhost:27017 \
   -e DB_TIMEOUT=5000 \
   ghcr.io/kunalsin9h/api:latest
```

### API Endpoints

## 1. Generate Image for **Blog Post**

> This is for OpenGraph and twitter Card of SEO
> The background image and font are in `assets` folder

**GET** /v1/image/`:title`

Returns a image, whose `Content-Type` is `image/jpeg`

## 2. Get **Views** for the blog

> This will not update the view count

**GET** /v1/views/`:slug`

**Response**

```json
{
  "data": [
    "views": 0
  ],
  "success": true
}
```

## 3. Update **Views** for the blog

> This will update and return the updated view count

**POST** /v1/views/`:slug`

**Response**

```json
{
  "data": {
    "views": 0
  },
  "success": true
}
```
