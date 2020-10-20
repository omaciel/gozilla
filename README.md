![GoZilla CI](https://github.com/omaciel/gozilla/workflows/GoZilla%20CI/badge.svg)
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
Search for multiple bug ids
```shell
./gozilla bug --id 1839234,1839235
2020/10/20 08:26:43 https://bugzilla.redhat.com/rest/bug/1839234
{"id":1839234,"product":"Fedora","summary":"sqlite-3.32.0 is available","severity":"unspecified"}
2020/10/20 08:26:49 https://bugzilla.redhat.com/rest/bug/1839235
{"id":1839235,"product":"Red Hat OpenStack","Manage/Unmanage support for CephFS drivers","severity":"medium"}
```
Search with keywords
```shell
./gozilla bug --keywords=FutureFeature,Triaged --limit=3
2020/10/20 08:30:05 https://bugzilla.redhat.com/rest/bug?keywords=FutureFeature,Triaged&limit=3
{"id": 1120141, "product": "Fedora","summary": "perl-Inline-0.62 is available","severity": "unspecified"}
{"id": 1785814,"product": "Fedora","summary": "svt-av1-0.8.0 is available","severity": "unspecified"}
{"id": 1512147,"product": "Fedora","summary": "source-to-image-1.3.1 is available","severity": "unspecified",
```
