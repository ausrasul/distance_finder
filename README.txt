# Shortest distance finder

## Description

A web service that takes a source coordinates and a list of destinations and returns a list of
routes between source and each destination.
The list is sorted by time and distance (if time is equal).

Thus, it answers the question: which destination is closest to the source and
how fast one can get there by car.

# All the following steps runs in a docker container

## Build

```bash
$ make
```

## Testing

```bash
$ make test
```

## Usage

Run the service

```bash
$ make run
```

It runs in a docker container and listens on http port 8080

API requests are like this:
http://provided-url:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219

The response will look like this:
    HTTP/1.1 200 OK
    Content-Type: application/json
    {
      "source": "13.388860,52.517037",
      "routes": [
        {
          "destination": "13.397634,52.529407",
          "duration": 465.2,
          "distance": 1879.4
        }, {
          "destination": "13.428555,52.523219",
          "duration": 712.6,
          "distance": 4123
        }
      ]
    }

Where input parameters are:
- src - source location (customer's home), only one can be provided
- dst - destination location (pickup point), multiple can be provided

