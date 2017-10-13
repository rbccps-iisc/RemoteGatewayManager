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



killall gotty
fuser -k 8080/tcp

nohup gotty -t -c "admin:1!CrazyPassword@123" -w tmux new-session -A -s ${p}  ssh ${u}@localhost -p ${p}  > nohup.out 2>&1 &

