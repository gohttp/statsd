
# statsd

 Statsd metrics for your HTTP server.

 View the [docs](http://godoc.org/github.com/gohttp/statsd).

## Metrics

 - `requests` (counter) request count
 - `response.ok` (counter) successful requests
 - `response.errors.client` (counter) client errors
 - `response.errors.server` (counter) server errors
 - `response.duration` (timer) request duration
 - `response.size` (timer) response size in bytes