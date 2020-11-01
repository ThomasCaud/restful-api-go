[![Go Report Card](https://goreportcard.com/badge/github.com/ThomasCaud/restful-api-go)](https://goreportcard.com/report/github.com/ThomasCaud/restful-api-go)
![Go](https://github.com/ThomasCaud/restful-api-go/workflows/Go/badge.svg)
[![Known Vulnerabilities](https://snyk.io/test/github/ThomasCaud/restful-api-go/badge.svg?targetFile=Gopkg.lock)](https://snyk.io/test/github/ThomasCaud/restful-api-go?targetFile=Gopkg.lock)

# Learning roadmap
Learnings Go concepts using the classical Book API: Dockerized and tested API, connected to database, with logs and authent management...

- ✓ Manage CRUD
- ✓ Dockerize the API
- ✓ Dockerize a database
- ✓ Use this database from API
- ✓ Use UUID instead of integer
- ✓ Better handler and validator management, using gin-gonic/gin & loopfz/gadgeto/tonic
- ✓ Setup CI
- ✓ Improve integrations tests, using ovh/venom
- ✓ Use juju/errors
- ✓ Generate swagger
- ✓ Manage migrations neatly, using rubenv/sql-migrate
- Keep HTTP calls logs
- Add broker message using
- Manage authentication
- Add channel using example

## Questions
- How to manage pagination?
- Use ORM?

# To make it work
```
$ cp docker-compose.override.yml.dist docker-compose.override.yml
```
Prerecommit hook is setting up, using [pre-commit](https://pre-commit.com/)


# To launch tests
```
$ cd tests/venom
$ venom run
```

# Swagger
```
Go to /swagger.json
```