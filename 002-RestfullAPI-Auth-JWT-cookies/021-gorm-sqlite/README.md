---
stack: GO, sqlite, GORM, gorilla/mux
lang: all
---

## get gorilla mux
```
go get -u github.com/gorilla/mux
```

## get orm
```
go get gorm.io/gorm
```

## to install latest version of any pkg
```
go mod tidy
```

## Got a problem with lunch? GOPATH should be set to
```
export GOPATH=$GOROOT
unset GOROOT
```

##  GO111MODULE set "on" or "off"
```
go env -w GO111MODULE=off
```