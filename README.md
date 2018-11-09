# Starish-Wars-Backend
[![CircleCI](https://circleci.com/gh/cbusby/Starish-Wars-Backend.svg?style=svg)](https://circleci.com/gh/cbusby/Starish-Wars-Backend)

**Go quirks we've discovered**

- The Go build tools prefer all Go projects to be in a single directory. That directory should be pointed at by an environment variable `GOPATH`. Make a src directory within it and put all your projects there.
- Projects are named by where they can be downloaded by CVS, e.g., github.com/cbusby/Starish-Wars-Backend. To clone this project into an appropriate place, navigate to `$GOPATH/src` and do `git clone <repo> github.com/cbusby/Starish-Wars-Backend`. This will create the correct subdirectory and place the code in the right place.
- If `GOPATH` is set appropriately, you can build your code from anywhere on your machine: `env GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/main github.com/cbusby/Starish-Wars-Backend/cmd/main`.
- AWS Lambda requires a zip file with the executable: `zip -j main.zip main`.
- The team has settled on Ginkgo and Gomega as our testing framework. This provides an expressive test syntax similar to Jest in JavaScript. Ginkgo provides the basic testing functionality, while Gomega provides a DSL that allows fluent testing and test grouping patterns.
  - https://onsi.github.io/ginkgo/
  - http://onsi.github.io/gomega/
- Run all tests in your project: from the root of your project do `go test ./...`.

**AWS quirks we've discovered**

- The Handler in the Lambda setup page should be the last component of the package that contains main.go, e.g., `main`. It should also be the name of the executable you create in the build step above, but not necessarily the name of the zip file you create.
- We can't currently test the Lambda in the Lambda Services section, becuase it does not get the HTTP method that the router function requires. Right now we're testing through the API Gateway on AWS or Postman.
- Postman has functionality for doing Amazon authorization. You can enter a user's Access Key and Secret Key in the Authorization section of the request specification. That's currently the most robust way to test the Lambda from the outside.

**Other tips**
- To paste an animated GIF in a GitHub comment, put an exclamation point before the description: `![description](http://the-link.gif)`.
