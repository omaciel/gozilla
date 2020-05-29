# gozilla
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
{"bugs":[{"priority":"unspecified","cf_last_closed":null,"assigned_to_detail":{"email":"odubaj","real_name":"Ondrej Dubaj","name":"odubaj","id":428719},"blocks":[],"creator":"Upstream Release Monitoring","last_change_time":"2020-05-29T13:54:01Z","is_cc_accessible":true,"keywords":["FutureFeature","Triaged"],"creator_detail":{"email":"upstream-release-monitoring","real_name":"Upstream Release Monitoring","name":"upstream-release-monitoring","id":282165},"cc":["Itamar Reis Peixoto","Jan Staněk","Ondrej Dubaj","Petr Kubat","Wilmer Jaramillo"],"url":"","assigned_to":"Ondrej Dubaj","groups":[],"see_also":[],"id":1839234,"whiteboard":"","creation_time":"2020-05-22T19:47:16Z","qa_contact":"Fedora Extras Quality Assurance","depends_on":[],"dupe_of":null,"docs_contact":"","qa_contact_detail":{"email":"extras-qa","real_name":"Fedora Extras Quality Assurance","name":"extras-qa","id":171387},"resolution":"","classification":"Fedora","alias":[],"cf_doc_type":"If docs needed, set a value","op_sys":"Unspecified","target_release":["---"],"status":"MODIFIED","cf_type":"---","cc_detail":[{"email":"itamar","real_name":"Itamar Reis Peixoto","name":"itamar","id":182415},{"email":"jstanek","real_name":"Jan Staněk","name":"jstanek","id":348991},{"email":"odubaj","real_name":"Ondrej Dubaj","name":"odubaj","id":428719},{"email":"pkubat","real_name":"Petr Kubat","name":"pkubat","id":355797},{"email":"wilmer5","real_name":"Wilmer Jaramillo","name":"wilmer5","id":262664}],"cf_clone_of":null,"summary":"sqlite-3.32.0 is available","is_open":true,"platform":"Unspecified","severity":"unspecified","cf_environment":"","version":["rawhide"],"deadline":null,"component":["sqlite"],"cf_fixed_in":"","is_creator_accessible":true,"is_confirmed":true,"target_milestone":"---","product":"Fedora","cf_release_notes":""}],"faults":[]}
```