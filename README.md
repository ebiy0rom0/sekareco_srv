# sekareco_srv
[![codecov](https://codecov.io/gh/ebiy0rom0/sekareco_srv/branch/develop/graph/badge.svg?token=KV6DKG67DF)](https://codecov.io/gh/ebiy0rom0/sekareco_srv)
[![Generate-API-Doc](https://github.com/ebiy0rom0/sekareco_srv/actions/workflows/genarate_api_doc.yml/badge.svg?branch=develop)](https://github.com/ebiy0rom0/sekareco_srv/actions/workflows/genarate_api_doc.yml)
[![Unit-Test](https://github.com/ebiy0rom0/sekareco_srv/actions/workflows/unit_testing.yml/badge.svg?branch=develop)](https://github.com/ebiy0rom0/sekareco_srv/actions/workflows/unit_testing.yml)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/ebiy0rom0/sekareco_srv/tree/develop.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/ebiy0rom0/sekareco_srv/tree/develop)
## Introduction
This project is server-side application for [sekareco](https://github.com/ebiy0rom0/sekareco).  


## Launch with Docker
This project is supports launch production and development mode with `Docker`.  
command for example.
```
# if you want to launch production mode, set ENV="prod" to the build arguments
$ docker build --tag sekareco_srv:v1.0 --build-args ENV="dev" .
$ docker run -p 8000:8000 --name sekareco sekareco_srv:v1.0
```
and access to `localhost:8000` on your browser.

## API documentation
API documentation is published on GitHub Pages.  
- [https://ebiy0rom0.github.io/sekareco_srv/page/index.html](https://ebiy0rom0.github.io/sekareco_srv/page/index.html)
