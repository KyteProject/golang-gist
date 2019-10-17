# Snippetbox
> A fullstack snippet posting app build with Go and SQL.
![GoReport A+ rank](https://goreportcard.com/badge/github.com/KyteProject/golang-gist)

This is an app made for educational purposes, it is **not intended for production** and **is not secure!** It uses locally generates TLS certificates to listen and serve on HTTPS. There are some tests which you can see below.

Check out Alex Edwards' book "Let's Go" to see where this project came from.

![Image of snippetbox website](https://ukemi.ninja/1efc86b72fe4ebea0367b036dde2325f.png)

## Setup

1) Clone the repo:

```sh
$ git clone https://github.com/KyteProject/golang-gist.git
$ cd golang-gist
```

2) Run docker compose:

```sh
$ docker-compose build
$ docker-compose up
```

3) Check the website in your browser:

```sh
https://localhost:8080
```

## Architecture
```
app
├── cmd
│   ├── migrate
│   │   └── main.go
│   └── web
│       ├── handlers_test.go
│       ├── handlers.go
│       ├── helpers.go
│       ├── main.go
│       ├── middleware_test.go
│       ├── middleware.go
│       ├── routes.go
│       ├── templates_test.go
│       ├── templates.go
│       └── testutils_test.go
├── docker
│   ├── app
│   │   ├── bin
│   │   │   ├── init.sh
│   │   │   └── wait-for-mysql.sh
│   │   ├── .env
│   │   └── Dockerfile
│   └── mariadb
│       └── Dockerfile
├── migrations
│   ├── 20170506082420_create_users_table.sql
│   └── 20191017132470_create_snippets_table.sql
├── pkg
│   ├── forms
│   │   ├── errors.go
│   │   └── form.go
│   └── models
│       ├── mock
│       │   ├── snippets.go
│       │   └── users.go
│       ├── mysql
│       │   ├── testdata
│       │   │   ├── setup.sql
│       │   │   └── teardown.sql
│       │   ├── snippets.go
│       │   ├── testutils_test.go
│       │   ├── users_test.go
│       │   └── users.go
│       └── models.go
├── tls
│   ├── cert.pem
│   └── key.pem
├── ui
│   ├── html
│   │   ├── base.layout.tmpl
│   │   ├── create.page.tmpl
│   │   ├── footer.partial.tmpl
│   │   ├── home.page.tmpl
│   │   ├── login.page.tmpl
│   │   ├── show.page.tmpl
│   │   └── signup.page.tmpl
│   └── static
│       ├── css
│       │   └── main.css
│       ├── img
│       │   ├── favicon.ico
│       │   └── logo.png
│       └── js
│           └── main.js
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## API

| ENDPOINT        | HTTP Method | Parameters |
| --------------- | ----------- | ---------- |
| /               | GET         |            |
| /snippet/create | GET         |            |
| /snippet/create | POST        |            |
| /snippet/:id    | GET         | ?id=[int]  |
| /user/signup    | GET         |            |
| /user/signup    | POST        |            |
| /user/login     | GET         |            |
| /user/login     | POST        |            |
| /user/logout    | POST        |            |
| /user/ping      | GET         |            |

## Tests

```
$ go test -v -short ./...
?       github.com/kyteproject/golang-gist/cmd/migrate  [no test files]
=== RUN   TestPing
--- PASS: TestPing (0.01s)
=== RUN   TestShowSnippet
=== RUN   TestShowSnippet/Valid_ID
=== RUN   TestShowSnippet/Non-existent_ID
=== RUN   TestShowSnippet/Negative_ID
=== RUN   TestShowSnippet/Decimal_ID
=== RUN   TestShowSnippet/String_ID
=== RUN   TestShowSnippet/Empty_ID
=== RUN   TestShowSnippet/Trailing_slash
--- PASS: TestShowSnippet (0.01s)
    --- PASS: TestShowSnippet/Valid_ID (0.00s)
    --- PASS: TestShowSnippet/Non-existent_ID (0.00s)
    --- PASS: TestShowSnippet/Negative_ID (0.00s)
    --- PASS: TestShowSnippet/Decimal_ID (0.00s)
    --- PASS: TestShowSnippet/String_ID (0.00s)
    --- PASS: TestShowSnippet/Empty_ID (0.00s)
    --- PASS: TestShowSnippet/Trailing_slash (0.00s)
=== RUN   TestSignupUser
=== RUN   TestSignupUser/Valid_submission
=== RUN   TestSignupUser/Empty_name
=== RUN   TestSignupUser/Empty_email
=== RUN   TestSignupUser/Empty_password
=== RUN   TestSignupUser/Invalid_email_(incomplete_domain)
=== RUN   TestSignupUser/Invalid_email_(missing_@)
=== RUN   TestSignupUser/Invalid_email_(missing_local_part)
=== RUN   TestSignupUser/Short_password
=== RUN   TestSignupUser/Duplicate_email
=== RUN   TestSignupUser/Invalid_CSRF_Token
--- PASS: TestSignupUser (0.01s)
    --- PASS: TestSignupUser/Valid_submission (0.00s)
    --- PASS: TestSignupUser/Empty_name (0.00s)
    --- PASS: TestSignupUser/Empty_email (0.00s)
    --- PASS: TestSignupUser/Empty_password (0.00s)
    --- PASS: TestSignupUser/Invalid_email_(incomplete_domain) (0.00s)
    --- PASS: TestSignupUser/Invalid_email_(missing_@) (0.00s)
    --- PASS: TestSignupUser/Invalid_email_(missing_local_part) (0.00s)
    --- PASS: TestSignupUser/Short_password (0.00s)
    --- PASS: TestSignupUser/Duplicate_email (0.00s)
    --- PASS: TestSignupUser/Invalid_CSRF_Token (0.00s)
=== RUN   TestSecureHeaders
--- PASS: TestSecureHeaders (0.00s)
=== RUN   TestHumanDate
=== RUN   TestHumanDate/UTC
=== RUN   TestHumanDate/Empty
=== RUN   TestHumanDate/CET
--- PASS: TestHumanDate (0.00s)
    --- PASS: TestHumanDate/UTC (0.00s)
    --- PASS: TestHumanDate/Empty (0.00s)
    --- PASS: TestHumanDate/CET (0.00s)
PASS
ok      github.com/kyteproject/golang-gist/cmd/web      (cached)
?       github.com/kyteproject/golang-gist/pkg/forms    [no test files]
?       github.com/kyteproject/golang-gist/pkg/models   [no test files]
?       github.com/kyteproject/golang-gist/pkg/models/mock      [no test files]
=== RUN   TestUserModelGet
--- SKIP: TestUserModelGet (0.00s)
    users_test.go:13: mysql: skipping integration test
PASS
ok      github.com/kyteproject/golang-gist/pkg/models/mysql     (cached)```
