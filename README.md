# Blox Engineering Service
<!-- Table of contents -->
A simple Go API for managing construction related work

## REQUIREMENTS

- Go
- PostgreSQL

## RUN LOCALLY

1. Run an instance of postgreSQL on your machine
2. Fill out your [config.toml](./config/config.toml) with your database creds
3. Run the service:

  ```bash
    git clone git@github.com:blox-eng/service.git
    cd service
    air
  ```

4. Open <http://localhost:7000/api/docs/index.html> to open the Swagger UI

## References

- Main technology is [Go Chi](https://github.com/go-chi/chi)
- Project layout as per <https://github.com/golang-standards/project-layout/tree/master>
- Logging via [Logrus](https://github.com/sirupsen/logrus)
