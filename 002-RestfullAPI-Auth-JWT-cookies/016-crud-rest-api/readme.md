---
stack: GO
pack: mux, GORM
---

## creating restful api using gorilla-mux
We will create arrestee API and go using guerrilla's mux package.
and we will define endpoints for the restful API.

## about orm
It stands for Object-Relational-Mapper and actually makes it easier to write code. right?! ^^
---

## create modules
```
go mod init www.github.com/username/repo
```

## get orm
```
go get gorm.io/gorm
```

## get mysql
```
go get -u gorm.io/driver/mysql
```

## get gorilla mux
```
go get -u github.com/gorilla/mux
```

## to install latest version of any pkg
```
go mod tidy
```

## if you unset your go root and set your GOPATH:
```
export GOPATH=$GOROOT
unset GOROOT
```