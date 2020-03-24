# KintoHub Backend Api Template

> This repo provide a template for starting a new golang API project @Kintohub.

The template itself is a simple ping/pong api.

## Requirements

* Install Go version 1.12 or higher

## How to use this template

1) Click on `use this template` at the top right of the [repo](https://github.com/kintohub/backend-api-template)
2) `git clone {the_newly_generated_repo}` in your workspace directory of choice
3) `./setup_git` in the root of the project

## Development Setup

Duplicate the `.env.example` file into a `.env` file.  
Modify the variables if needed.

```shell script
$ go run cmd/main.go
{"level":"info","msg":"Successfully started server listening to port 8080","time":"2020-03-24T10:41:59+08:00"}
{"level":"debug","msg":"Processing ping request {Hey it is me}","time":"2020-03-24T10:42:02+08:00"}
{"level":"debug","msg":"Successful ping response \u0026{Coop Was Here}","time":"2020-03-24T10:42:02+08:00"}
```

## Send a ping

```shell script
$ curl -d '{"message": "Who was here?"}' localhost:8080/ping
```

## Meta

KintoHub Goons - [@KintoHub](https://twitter.com/kintohub)

https://www.kintohub.com

## Contributing

1. Fork it (<https://github.com/kintohub/backend-api-template/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request