#!/bin/bash


oldIp=`cat ip`
iface='eth0'
newIp=`curl -s ipinfo.io/ip`

if [ "$newIp" != "$oldIp" ]; then 
	echo "IP Changed"
	`echo "$newIp" > ip`
	mac=`cat /sys/class/net/$iface/address`
	uname=`whoami`
	cmd='{"ip":"'$newIp'","mac":"'$mac'","username":"'$uname'"}'
	freePort=`curl -s --request POST --url http://139.59.88.117:3000/api/register --header 'content-type: application/json' --data $cmd | jq -r '.port'  `
	echo "Connection to Manage at port $freePort established"
	autossh -M 20000 -N -f -R $freePort:localhost:22 manager@139.59.88.117 &

fi;

