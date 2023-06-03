# Ready

Simple tool to loop through a set of checks to let you know if in a ready state.

Will also wait (default up to 30s) to see if all conditions are met.

## Basic Usage

Call it by binary and do something like below

```shell
# will wait up to 15 seconds for postgres and redis
ready --timeout 15 --host-ports localhost:5432,localhost:6379 run
```

You can also call it via docker:

```shell
docker run \
  --net=host \
  -it ghcr.io/madhuravius/ready:latest \
  --timeout 15 \
  --host-ports localhost:5432,localhost:6379 \
  run
```

Standard out text:

```sh
> ready
NAME:
   ready - wait for a group of hosts and ports to be ready

USAGE:
   ready [global options] command [command options] [arguments...]

COMMANDS:
   run,     start ready
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug                                    if enabled, will print out logs (default: true)
   --host-ports value [ --host-ports value ]  as a csv, specify a range of hosts and ports to check (ex: "localhost:3000,test:1234" )
   --timeout value                            as an integer, maximum number of seconds to wait and error if ready checks do not all pass by (default: 30)
   --help, -h                                 show help
```

## Why?

I always used scripts like this one:

```sh
# !/usr/bin/env bash
set -e 

if [ -z "$1" -o -z "$2" ]
then
    echo "Usage: ./service_started.sh HOST PORT"
    exit 1
fi
echo "Waiting for port $1:$2 to become available..."
while ! nc -z $1 $2 2>/dev/null
do
    let elapsed=elapsed+1
    if [ "$elapsed" -gt 30 ] 
    then
        echo "TIMED OUT !"
        exit 1
    fi  
    sleep 1;
done

echo "READY !"
```

Which I rehash repeatedly and often get elaborate when more than a few
processes are involved.