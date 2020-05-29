# GoZilla
An opinionated command line for Bugzilla written in Go.

## Build it

```
go build -o gozilla
```

## Use it

```shell
./gozilla
NAME:
   GoZilla - An opinionated command line for Bugzilla written in Go.

USAGE:
   gozilla [global options] command [command options] [arguments...]

COMMANDS:
   bug      Fetch a Bugzilla issue by its ID.
   version  Shows the version for the Bugzilla server.
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

### Bugzilla Server Version

```shell
./gozilla version
2020/05/29 15:44:30 https://bugzilla.redhat.com/rest/version
{"version":"5.0.4.rh44"}
```

### GoZilla Bug Command Help

```shell
./gozilla bug --help
NAME:
   gozilla bug - Fetch a Bugzilla issue by its ID.

USAGE:
   gozilla bug [command options] [arguments...]

OPTIONS:
   --id value  
   --help, -h  show help (default: false)
```

### GoZilla Bug Command

```shell
./gozilla bug --id 1839234
2020/05/29 15:46:45 https://bugzilla.redhat.com/rest/bug/1839234
{"id":1839234,"product":"Fedora","summary":"sqlite-3.32.0 is available","severity":"unspecified"}
```