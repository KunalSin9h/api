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
export MEILI_HOST=http://localhost:7700
export MEILI_MASTER_KEY="-meilisearch-master-key-"

go run cmd/api/*
```

> [!NOTE]
> The `HOST` is where the server servers, default value of `HOST` is `127.0.0.1`. 
> The `PORT` is where the server listens, default value of `PORT` is `9999`. 
> The `MONGODB_URL` is where the mongodb database is running, the default value is `mongodb://localhost:27017`. 
> The `DB_TIMEOUT` is the timeout time for each mongodb operation, the default value is `5000`. 
> The `MEILI_HOST` is the host for meilisearch server. 
> The `MEILI_MASTER_KEY` is the master key (auth) for meilisearch server.

### Run using Docker

```bash
docker run \
   --name api \
   -d -p 9999:9999 \
   -e HOST=0.0.0.0 \
   -e PORT=9999 \
   -e MONGODB_URL=mongodb://localhost:27017 \
   -e DB_TIMEOUT=5000 \
   -e MEILI_HOST=http://localhost:7700 \
   -e MEILI_MASTER_KEY="-meilisearch-master-key-" \
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

## 4. Add Document in the `Meilisearch` 

> This will and or update the document in meilisearch at some index

**POST** /v1/search/`:index`

**Request Payload**
```json
{
    "data": {
        // json data
    }
}
```

**Response** : 200 if ok

## 5. Get Document from the `Meilisearch` after query

**GET** /v1/search/`:index`?text="how to search"

**Response Payload**

```json
[
    {// data},
    {// data},
    {// data},
]
```
