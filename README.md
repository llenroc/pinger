# Local Pinger [![Build Status](https://travis-ci.org/toggl/pinger.svg?branch=master)](https://travis-ci.org/toggl/pinger)

A small utility to periodically check different services and post the status JSON to your own analytics or a 3rd party
server. You can read about reasoning for creating this at [toggl blog](http://blog.toggl.com/2014/04/5-tips-servers-service-monitoring/)
## Usage

Local Pinger compiles into an executable binary. You'll need to set the ping period and an URL for jobs listing/results saving.<br/>
Example command line:
```shell
local_pinger -period=10 -jobs_url="http://localhost/http_checks"
```

### Jobs listing/results saving

You ought to provide a list of services from a service as HTTP GET, that needs 'pinging'.<br/>
Format is simple JSON:
```json
[
  "http://service1/status",
  "http://service2/status/db",
  "http://172.16.0.1/cache.manifest"
]
```

In current version the results saving is a POST to the same URL.

## Dummy data
There is a dummy jobs listing server under [src/dummy](src/dummy), runnable via ```make dummy```. It starts up a http server on :3000 that suplies an example JSON for the Local Pinger

## Contributing
Want to contribute? Great! Just fork the project, make your changes and open a Pull Request
