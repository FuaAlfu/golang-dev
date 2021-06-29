---
stack: GO
pack: Fiber
---

## creating restful api using Fiber
We will create arrestee API and go using Fiber package.
and we will define endpoints for the Restfull API.

## about Fiber
If you are coming from another language, such node-js and trying your hand at developing Go applications then Fiber is an incredibly easy framework to start working with.
It presents a familiar feel to Node.js developers who have previously built systems using Express.js. It’s also built on top of Fasthttp which is an incredibly performant and minimal HTTP engine built for Go.
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