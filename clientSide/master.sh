
#!/bin/bash

workingDir="/home/username/clientSide"
managerURL="manager@123.123.123.123"
apiURL="http://123.123.123.123:3000/api/register"

freePort=`cat port`
COMMAND="ssh -N -o ExitOnForwardFailure=yes -R $freePort:localhost:22 $managerURL"
currentIp=`cat ip`
iface='enp1s0'
mac=`cat /sys/class/net/$iface/address`
uname=`whoami`

check_ssh()
{
	if ssh -t $managerURL netcat -z -v -w 10  localhost $freePort | grep -q "succeeded"
	then
        	echo 0
	else
		echo "Tunnel not established"
		echo 1
	fi
}



establish_ssh()
{

	ssh -N -o ExitOnForwardFailure=yes -R $freePort:localhost:22 manager@139.59.88.117 &
	check_ssh
	
	if [ $? -eq 0 ]
	then
		echo 0
	else
		echo 1
	fi

}


#Check if device is online
wget --tries=10 -q --spider http://google.com
if [ $? -eq 0 ]
then


	oldIp=`cat $workingDir/ip`
	newIp=`wget ipinfo.io/ip -O - -q`
	oldIp=`cat $workingDir/ip`

	if [ "$newIp" != "$oldIp" ]
	then
		echo "IP Changed"
		`echo "$newIp" > ip`
		cmd='{"ip":"'$newIp'","mac":"'$mac'","username":"'$uname'"}'
		freePort=`curl -s --request POST --url $apiURL --header 'content-type: application/json' --data $cmd | jq -r '.port'  `
		`echo "$freePort" > port`
		pkill -f "ssh -N -o ExitOnForwardFailure=yes -R"
		establish_ssh
		if [ $? -eq 0 ]
	        then
	        	exit 0
	        else
			exit 1
		fi
	fi


	if pgrep -f -x "$COMMAND" > /dev/null 
	then
		check_ssh
		if [ $? -eq 0 ] 
		then	
			echo "SSH ESTABLISHED"
			exit 0
		else
			echo "PROCESS EXISTS BUT SSH NOT ESTABLISHED"
			pkill -f "ssh -N -o ExitOnForwardFailure=yes -R"
			establish_ssh
			if [ $? -eq 0 ]
			then 
				exit 0
			else
				exit 1
			fi
		fi

	else 
		echo "SSH NOT ESTABLISHED"
		establish_ssh
		if [ $? -eq 0 ]
	        then
	            exit 0
	        else
				exit 1
		fi
	fi


else
	exit 1

fi













		 
	


