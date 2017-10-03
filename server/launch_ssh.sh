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



tmux kill-session -t ${p}
killall gotty
fuser -k 8080/tcp

nohup gotty  -w tmux new-session -s ${p}  ssh ${u}@localhost -p ${p} &

