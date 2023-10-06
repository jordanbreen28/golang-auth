# Go auth

Simple authentication API which utilises JWT and Cookies for session authentication.
Technologies used:

- Docker
- Golang
- Postgres

## Setup

This project has been designed to run in two docker containers, one for the application and one for a postgres database.
You can set it up yourself by running:

```bash
docker-compose build
docker-compose up
```

The application will then start listening to `0.0.0.0:8000` on your local machine. A comprehensive list of all available routes can be found [here](https://github.com/jordanbreen28/golang-auth/blob/main/main.go) in the main.go file. All routes are prefixed with `/api/$api_version`, with a current version of `v1`.

### Example request

```bash
$ curl -vX POST http://0.0.0.0:8000/api/v1/users -d @user.json \
--header "Content-Type: application/json"
```

Where `user.json` is a valid JSON object containing your form data.
