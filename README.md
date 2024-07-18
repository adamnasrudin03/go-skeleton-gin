# go-skeleton-gin
 Skeleton project golang with framework <a href="https://gin-gonic.com/" target="_blank"> Gin / Gin-gonic </a>

## Development Guide

### Collection Using Postman
- ./go-skeleton.postman_collection.json
  
### Installation
- Clone this repo
    ```sh
    git clone https://github.com/adamnasrudin03/go-skeleton-gin.git
    ```

- Copy `.env.example` to `.env`

    ```sh
    cp .env.example .env
    ```
- Setup local database
- Start service API
    ```sh
    go run main.go
    ```

### Coverage Unit test
```sh
  make cover
```

## Structure Response RESTfull API 
- Error
```json
{
  "status": "status error",
  "code": 10, // code custom error
  "message": {
    "id": "message error language Indonesian",
    "en": "message error language English"
  }
}
```

- Success Single Data
```json
{
  "status": "Created",
  "data": {}
}
```

- Success Multiple Data
```json
{
  "status": "Success",
  "meta": {
    "page": 1,
    "limit": 10,
    "total_records": 3,
    "total_pages": 1
  },
  "data": []
}
```

- Success response message
```json
{
  "status": "Created",
  "message": "data created"
}
```

- Success response Multiple message
```json
{
  "status": "Created",
  "message": {
    "id": "Data berhasil dibuat",
    "en": "Data created successfully"
  }
}
```


<br clear="both">
<h2 align="left">Connect with me:</h2>
<div align="left">
  <a href="https://www.linkedin.com/in/adam-nasrudin/" target="_blank">
    <img src="https://img.shields.io/static/v1?message=LinkedIn&logo=linkedin&label=&color=0077B5&logoColor=white&labelColor=&style=for-the-badge" height="35" alt="linkedin logo"  />
  </a>
  <a href="https://adamnasrudin.vercel.app/blog" target="_blank">
    <img 
        src="https://img.shields.io/static/v1?message=My%20Blog&logo=blogger&label=&color=blue&logoColor=white&labelColor=&style=for-the-badge" height="35" alt="blog"  />
  </a>
</div>