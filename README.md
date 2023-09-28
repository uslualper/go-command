# GO Command

### Project Structure
https://github.com/golang-standards/project-layout

## Makefile

make {command}

- ### Docker

  - docker-build
  - docker-up
  - docker-restart
  - docker-exec

## Commands

-  ``` go run main.go warmup https://{your_domain}/sitemap.xml ```
-  ``` go run main.go warmup https://{your_domain}/sitemap.xml --max-worker=1 --max-request-in-time=10 --max-run-time=50 ```


## Test

```
go run cmd/cli/main.go test
```