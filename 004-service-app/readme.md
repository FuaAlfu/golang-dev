# service-app
..

## Requirements

This project was designed against Go 1.12. It should work for 1.11 but 1.12 is
recommended.

Supporting services like the database are hosted in Docker. If you cannot
install Docker on your machine you can still follow most of this material by
hosting a database elsewhere and modifying the connection information to your
needs.

## Setup

Clone your repository somewhere on your computer. The location does not
especially matter but if it is outside of your `$GOPATH` then the Go modules
features will work automatically.

In a separate folder make a directory where you will be building your API. We
recommend you initialize that folder as a Git repository to track your work.


```sh
mkdir ~/project-name
cd ~/project-name
git clone https://github.com/user/project-name.git
mkdir garagesale
cd garagesale
git init .
```

---

You must also use `go mod init` to set the import path for this project. Doing
this exactly as shown will allow you to copy and paste code without a need to
modify import paths.

```sh
go mod init github.com/ardanlabs/garagesale
git add go.mod
git commit -m "Initial commit"
```

## Postman API Client

For the class we will be building up a REST API. You may use any HTTP client
you prefer to make requests but we recommend [Postman](https://www.getpostman.com/).
For convenience you may use the import button in the top left to import the
included `postman_environment.json` and `postman_collection.json` files to get
a client up and running quickly. Be sure to select the "Garage Sale Service"
environment in the top right.

## Diffing Folders

Reviewing the differences between the successive steps helps to reinforce the
ideas each change is about. This is made easier by running the following
command to define a git alias called `dirdiff`:

```sh
git config --global alias.dirdiff 'diff -p --stat -w --no-index'
```

With that alias in place, run this command from the top level folder to see the
differences between the `01-startup` directory and the `02-shutdown` directory.

```sh
git dirdiff 01-startup 02-shutdown`
```

---
---
---