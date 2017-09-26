
#!/bin/bash

workingDir="/home/uname/sshtunnel"
managerURL="manager@100.100.100.100"
apiURL="http://100.100.100.100:9000/register"

freePort=`cat $workingDir/port`
currentIp=`cat $workingDir/ip`
iface='eth0'
4GDialScript=$workingDir/dial
mac=`cat /sys/class/net/$iface/address`
uname=`whoami`

check_ssh()
{	
	COMMAND="ssh -N -o ExitOnForwardFailure=yes -R $freePort:localhost:22 $managerURL"
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

	COMMAND="ssh -N -o ExitOnForwardFailure=yes -R $freePort:localhost:22 $managerURL"
	ssh -N -o ExitOnForwardFailure=yes -R $freePort:localhost:22 manager@139.59.88.117 &
	check_ssh
	
	if [ $? -eq 0 ]
	then
check if varshel 		echo 0
	else
		echo 1
	fi

}


#Check if interface is up else, turn it on
#Link to 4G Dial script or 
if [ -z $4GDialScript ]
then
$4GDialScript
while [ $? -ne 0 ]
do
sleep 100
sudo killall dial
$4GDialScript
done
fi

#Check if device is online
sudo wget  -q --spider http://google.com
count=1
while [ $? -ne 0 ] && [ $count -lt 5 ]
do
	((count++))
	sleep 5
	sudo wget --tries=10 -q --spider http://google.com
	
done


if [ $? -eq 0 ]
then


	oldIp=`cat $workingDir/ip`
	newIp=`wget ipinfo.io/ip -O - -q`
	oldIp=`cat $workingDir/ip`

	if [ "$newIp" != "$oldIp" ]
	then
		echo "IP Changed"
		echo "$newIp" > $workingDir/ip
		cmd='{"ip":"'$newIp'","mac":"'$mac'","username":"'$uname'"}'
		freePort=`curl -s --header 'content-type: application/json' -X  POST  -d $cmd $apiURL | jq -r '.port'  `
		echo "$freePort" > $workingDir/port
		pkill -f "ssh -N -o ExitOnForwardFailure=yes -R"
		establish_ssh
		if [ $? -eq 0 ]
	        then
	        	exit 0
	        else
			exit 1
		fi
	fi

	COMMAND="ssh -N -o ExitOnForwardFailure=yes -R $freePort:localhost:22 $managerURL"
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













		 
	


