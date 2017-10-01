#!/bin/bash



usage() { echo "Usage: $0 [-u <string>] [-p <string>]" 1>&2; exit 1; }

while getopts ":u:p:" o; do
    case "${o}" in
        u)
            u=${OPTARG}
            
            ;;
        p)
            p=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done
shift $((OPTIND-1))


echo "u = ${u}"
echo "p = ${p}"


nohup gotty -w tmux new-session -s ${p} ssh manager@139.59.88.117 ssh ${u}@localhost -p ${p} &

