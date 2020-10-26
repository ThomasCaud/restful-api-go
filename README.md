# Learning roadmap
Learnings Go concepts using the classical Book API: Dockerized and tested API, connected to database, with logs and authent management...

- ✓ Manage CRUD
- ✓ Dockerize the API
- ✓ Dockerize a database
- ✓ Use this database from API
- Generate swagger
- e2e testing
- Keep HTTP calls logs
- Use juju/errors
- Add validator?
- Manage authent
- Context

# To make it work
```
#cp docker-compose.override.yml.dist docker-compose.override.yml
```

Change id to uuid
Validators
Fix interface (BooksDatabase is useless ; main should use it)