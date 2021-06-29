---
stack: GO
pack: Fiber
---

## creating restful api using Fiber
We will create arrestee API and go using guerrilla's mux package.
and we will define endpoints for the restful API.

## about orm
It stands for Object-Relational-Mapper and actually makes it easier to write code. right?! ^^
---

## create modules
```
go mod init www.github.com/username/repo
```

## get Fiber
```
go get github.com/gofiber/fiber
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