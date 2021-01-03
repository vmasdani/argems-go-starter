# ARGEMS Go Starter

This is a starter Golang ARGEMS code which has REST API and database interaction.



## Development

- Install prerequisites
    1. [Go](https://golang.org/dl/)
    2. [Elm](https://guide.elm-lang.org/install/elm.html)

- Go
```
go get
go build
./argems-go-starter
```

- Elm
```
cd frontend
./run.sh
```

# Build
To build, you need Python 3, and also [Node.JS](https://nodejs.org/en/download/) for minification.  
Edit your app URL in `release.json`.
```
sudo npm i -g uglify-js
./release.py
```