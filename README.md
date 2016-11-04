# go_cgi_demo
Quick Demo Using Golang Binaries On Shared Host Using CGI-BIN

NOTE before compiling any db example you need to replace USERNAME, PASSWORD, DATABASE, and table with valid strings for your server configuration to allow the script to connect and read/write from/to a valid table.

Either execute the file ./build or run each command in the file individually to compile the examples:
```bash
env GOOS=linux go build -ldflags "-s -w" -o index.cgi simple/index.go
env GOOS=linux go build -ldflags "-s -w" -o db_read.cgi db_read/db_read.go
```

This will generate a cgi for each example... You should move them to your shared host under the cgi-bin folder. If this folder does not already exist, create it. You may need to alter your host file to point to this directory, though it is a standard location on most hosts.
```
www/cgi-bin/index.cgi
```

Change the permissions to 755 for the file once its on the host.

This is confirmed to work on a Linux Apache Shared Hosting setup.
