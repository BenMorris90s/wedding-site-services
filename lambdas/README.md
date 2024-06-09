## Lambda API

API for connecting to the Guest Info table.

## Request 

# Running Locally

```bash
  ./run_local.sh
```

Trigger lambda once running with:

```bash
  curl http://127.0.0.1:3000/rsvp
```

## Deployment

Currently a work in progress and not implemented.

## Testing

```bash
  cd ./upsert-rsvp-status 
  go test
```
Yet to be added to a CI step in Github, so can be run on a pre push git hook locally.
