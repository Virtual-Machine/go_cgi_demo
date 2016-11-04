# go_cgi_demo
Quick Demo Using Golang Binaries On Shared Host Using CGI-BIN

Either execute the file ./build or run the command:
env GOOS=linux go build -ldflags "-s -w" -o index.cgi index.go

The output file can then be put on your shared host:
www/cgi-bin/index.cgi

Change the permissions to 755 for the file once its on the host.

This is confirmed to work on a Linux Apache Shared Hosting setup.
