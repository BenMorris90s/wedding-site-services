## Lambda API

API for connecting to the Guest Info table.

## Running Locally

```bash
  ./run_local.sh
```

Trigger lambda once running with:

```bash
  curl http://127.0.0.1:3000/rsvp
```

## Deployment

The settings in samconfig.toml are used by the deploy script. Change these if needed.

```bash
    ./deploy.sh
```

GET endpoint can be hit with

```bash
    curl https://23z5anklf9.execute-api.eu-west-2.amazonaws.com/Prod/rsvp
```

## Unit Tests

```bash
  cd ./upsert-rsvp-status 
  go test
```
Yet to be added to a CI step in Github, so can be run on a pre push git hook locally.
