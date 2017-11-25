# hyper_fleet

Before starting Blockchain:
* docker rm -f $(docker ps -aq)
* docker network prune
* find old docker images ("docker images") and remove them ("docker rmi <image-name>")
* rm hfc-key-store/\*


Run scripts to start app:
* cd ./setup ; ./startFabric.sh
* cd .. ; node enrollAdmin.js
* node registerUser.js
* node server.js
