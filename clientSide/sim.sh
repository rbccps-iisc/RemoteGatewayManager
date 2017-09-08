#!/bin/bash


oldIp=`cat ip`

newIp=`curl -s ipinfo.io/ip`

if [ "$newIp" != "$oldIp" ]; then 
	echo "IP Changed"
	`echo "$newIp" > ip`
	mac=`cat /sys/class/net/enp1s0/address`
	cmd='{"ip":"'$newIp'","mac":"'$mac'"}'
	echo $cmd
	curl --request POST --url http://139.59.88.117:3000/api/register --header 'content-type: application/json' --data $cmd
fi;

