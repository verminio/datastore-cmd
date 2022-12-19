# Datastore Command Line

A command line tool for querying GCP Datastore, written in [Go](https://go.dev).

## Testing

`go test ./...`

## How to use

This command has a required flag `--project-id` to specify the GCP project to be queried.
`datastore-cmd [command] [operation] --project-id=gcp-project`

The output can be rendered as json by passing the flag `--output=json`.

### Supported Commands

#### namespace

The following operations are supported for the namespace command:
* `list` can be used in order to retrieve a list of namespaces