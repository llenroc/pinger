# Local Pinger

A small utility to periodically check different services and post the status JSON to a 3rd party.

## Usage

Local Pinger compiles into an executable binary. You'll need to set the ping period and an URL for jobs listing/results saving.
example command line:
```local_pinger -period=10 -jobs_url="http://localhost/http_checks"```

### Jobs listing/results saving

You ought to provide a list of services from a, well, service as HTTP GET, that needs 'pinging', format is simple JSON:
["http://service1/status","http://service2/status/db","http://172.16.0.1/cache.manifest"]

In current version the results saving is a POST to the same URL.

## Dummy data
There is a dummy jobs listing server under src/dummy, runnable via make dummy. Starts up a http server on :3000 that suplies an example JSON for the Local Pinger

## Contributing
Want to contribute? Great! Just fork the project, make your changes and open a Pull Request