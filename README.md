# URL Shortener

This is a very simple URL shortener using Go lang.

## Features
  - Make POST request with any URL to shorten it
  - Make a GET request with the shortened URL to retrieve the Long URL


## Installation:

This project uses base62 package for encoding the URL Id. Therefore import and install it using the command below:

```sh
$ go get github.com/pilu/go-base62
```

Then run the Url Shortener from its directory and it will start listening on port 8082

```sh
$ go run urlshortener.go
2016/10/31 17:23:18 Listening on port 8082 ...
```

## Requesting to shorten URL
Long URL can be posted as JSON formatted data to the URL shortener on path "/shorten"
. See example:
```sh
$ curl -sX POST -H 'Content-Type: application/json' 'localhost:8082/shorten' -d '{"url": "http://a.very.long.url"}'
HTTP 200
{"Short":"http://localhost/2Bi"}
```

## Retrieving the original URL
Original URL can be retrieved by submitting JSON formatted short URL to the URL shortener on path "/original" via a GET call
. See example:
```sh
$ curl -sX GET -H 'Content-Type: application/json' 'localhost:8082/original' -d '{"short": "http://localhost/2Bi"}'
HTTP 200
{"Original":"http://a.very.long.url"}
```

## Notes

This application is using non-persistent storage to store URL mappings. therefore once the application is stopped. you can no longer retrieve the URLs that were shortened earlier.