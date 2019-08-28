## To Build

```
make
```

## To test
```
make test
```

## To build multi platform release binaries
```
make release
```

## To run

```
$ ./lift

lift is a tool for enriching your application so it can be deployed to multiple cloud platforms with minimal effort.

Usage:
  lift [command]

Available Commands:
  help        Help about any command
  platform    Platform commands

Flags:
  -h, --help   help for lift

Use "lift [command] --help" for more information about a command.
```

```
$ ./lift platform
Commands related to platform operations

Usage:
  lift platform [flags]
  lift platform [command]

Available Commands:
  list        Platform list

Flags:
  -h, --help   help for platform

Use "lift platform [command] --help" for more information about a command.
```

```
$ ./lift platform list
+-----------------------+--------+--------------+---------+-------------+
|         NAME          | ALIAS  |     TYPE     | PROFILE | DESCRIPTION |
+-----------------------+--------+--------------+---------+-------------+
| gke-sandbox-cschaefer | ci-k8s | kubernetes   | qa      | GKE         |
| cf-sandbox-cschaefer  | ci-cf  | cloudfoundry | qa      | CF          |
+-----------------------+--------+--------------+---------+-------------+
```

Note: Output is hardcoded