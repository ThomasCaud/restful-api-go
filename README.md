[![Go Report Card](https://goreportcard.com/badge/github.com/ThomasCaud/restful-api-go)](https://goreportcard.com/report/github.com/ThomasCaud/restful-api-go)

# Learning roadmap
Learnings Go concepts using the classical Book API: Dockerized and tested API, connected to database, with logs and authent management...

- ✓ Manage CRUD
- ✓ Dockerize the API
- ✓ Dockerize a database
- ✓ Use this database from API
- ✓ Use UUID
- ✓ Better handler and validator management using Gin & Tonic
- ✓ Setup CI
- ✓ Improve integrations tests using ovh/venom
- ✓ Use juju/errors
- ✓ Generate swagger
- Manage migrations neatly
- Keep HTTP calls logs
- Add broker message using
- Manage authentication
- Add channel using example

## Questions
- How to manage pagination?
- Use ORM?

# To make it work
```
#cp docker-compose.override.yml.dist docker-compose.override.yml
```
Prerecommit hook is setting up, using [pre-commit](https://pre-commit.com/)


# To launch tests
```
#cd tests/venom
#venom run
```

# Swagger
```
Go to /swagger.json
```

# Licence
MIT License

Copyright (c) 2020 Thomas Caudrelier

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.