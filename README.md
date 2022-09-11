# sekareco_srv
[![codecov](https://codecov.io/gh/ebiy0rom0/sekareco_srv/branch/develop/graph/badge.svg?token=KV6DKG67DF)](https://codecov.io/gh/ebiy0rom0/sekareco_srv)
[![Generate-API-Doc](https://github.com/ebiy0rom0/sekareco_srv/actions/workflows/genarate_api_doc.yml/badge.svg?branch=develop)](https://github.com/ebiy0rom0/sekareco_srv/actions/workflows/genarate_api_doc.yml)
[![Unit-Test](https://github.com/ebiy0rom0/sekareco_srv/actions/workflows/unit_testing.yml/badge.svg?branch=develop)](https://github.com/ebiy0rom0/sekareco_srv/actions/workflows/unit_testing.yml)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/ebiy0rom0/sekareco_srv/tree/develop.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/ebiy0rom0/sekareco_srv/tree/develop)
## Introduction
This project is server-side application for [sekareco](https://github.com/ebiy0rom0/sekareco).  


## Support `make` command
In this project, some operations in development supported by Makefile.  
For more information on possible operations,  
see help displayed by execute `make` or `make help` or refer directly to the Makefile.

## Launch with Docker
This project is supports launch production and development mode with `Docker`.  
command for example.
```
# if you want to launch production mode, set ENV="prod" to the build arguments
$ docker build --tag sekareco_srv:latest --build-arg ENV="dev" .
$ docker run -p 8000:8000 --name sekareco_srv sekareco_srv:latest
```
and access the health check endpoint using the `curl` command.
```
$ curl -v http://localhost:8000/health

*   Trying 127.0.0.1:8000...
* Connected to localhost (127.0.0.1) port 8000 (#0)
> GET /health HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.83.1
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json
< Vary: Origin
< Date: Sun, 11 Sep 2022 07:41:22 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

## API documentation
API documentation is published on GitHub Pages.  
- [https://ebiy0rom0.github.io/sekareco_srv/page/index.html](https://ebiy0rom0.github.io/sekareco_srv/page/index.html)
