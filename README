
# Gateway Manager & Client



#What it does 
1.Remote device  requests a gateway manager for an open port and registers it's pub ip and mac address. 
2. Upon receipt of a free port on the manager server, remote device creates an autossh pipe which the manager can use to reverse ssh to the remote.

#Dependedncies :-
jq
curl
autossh

#Steps:-
1. Rename primary interface name
2. "touch ip" in the working directory. This is where public ip addres is cached
3. Use crontab -e to point to this script. Make it execute every hour or so. If ip remains same then a request won't be made


TODO :

1.Daemon to monitor change in pub ip and 

