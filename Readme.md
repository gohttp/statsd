
# statsd

 Statsd metrics for your HTTP server.

 View the [docs](http://godoc.org/github.com/gohttp/statsd).

## Metrics

 - `requests` (counter) request count
 - `request.size` (timer) request content-length
 - `response.ok` (counter) successful requests
 - `response.errors.client` (counter) client errors
 - `response.errors.server` (counter) server errors
 - `response.duration` (timer) request duration
 - `response.size` (timer) response size in bytes

## Links

 - https://github.com/statsd/client-interface - interface
 - https://github.com/statsd/client - client

# License

 MIT