# echo-mvc

CLI does 2 things:

1. clones [template repository](https://github.com/sk1t0n/echo-mvc-template),
2. installs [CLI](https://github.com/sk1t0n/echo-mvc-generator) to generate CRUD or separate controller/model/view.

## Install

```sh
go install github.com/sk1t0n/echo-mvc@latest
```

## Run CLI

```sh
# you need to set the GOBIN environment variable
echo-mvc new directory
```

## Run tests

```sh
make run_tests
```

## Show code coverage

```sh
make run_cover_tests
```
