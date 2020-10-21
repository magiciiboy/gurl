# gURL - A simplest version of cURL written in Go
A command-line tool that handles the HTTP requests  
This package uses HTTP/1.1 accordingly to RFC 7230.

## Usage
```bash
usage: gurl [-h|--help] -u|--url "<value>" [-p|--profile <integer>]
            [-v|--verbose]

            A simplest version of cURL written in Go

Arguments:

  -h  --help     Print help information
  -u  --url      URL to request
  -p  --profile  Profile n requests
  -v  --verbose  Print details
```

## How to run
### Quick run
```bash
go run cmd/gurl/gurl.go
```

## Resources
- Project Layout: https://github.com/golang-standards/project-layout