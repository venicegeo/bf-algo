# bf-algo

"bf-algo" is intended to grow significantly.  Currently, however, it provides only a dummy service that vaguely pretends to be like one of the algorithm services for beachfront.

## Installing and running

Make sure you have Go installed on you machine, and an appropriate GOPATH (environment variable) set.

Use `go get` to install the latest version of both the CLI and the library.
	$ go get -v github.com/venicegeo/bf-algo/...

To install
	$ go install github.com/venicegeo/bf-algo/...
Alternate install:
	navigate to GOPATH/src/github.com/venicegeo/bf-algo/bf-service
	$ go install .

To Run
	GOPATH/bin/bf-service

## Using

When run, it will produce no visible output.  This is normal.  At this point, it is providing a web service.  The following meaningful links are available:
	http://localhost:8080/dummyAlgo : "initiates the process".  If you don't do this first, the rest will produce error messages
	http://localhost:8080/checkStatus : tells you whether the process is complete.  Process will not actually complete unless checked a certain number of times
	http://localhost:8080/getResult : Once the process is complete, hands back a GeoJson of the results

This is currently quite limited in scope and also embarrassingly crude, but it's what's there right now.