# goapp-echo-mysql

- [echo](https://github.com/labstack/echo)を利用
- [database/sql](https://golang.org/pkg/database/sql)を利用
- [validator](https://github.com/go-playground/validator)を利用


## env
Docker version 18.06.1-ce

## build
```
make
```

## initialize db
```
make init-db
```

## run
```
make run
```

## request
```
curl localhost/api/users
```
