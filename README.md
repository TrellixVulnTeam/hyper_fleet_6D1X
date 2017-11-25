# hyper_fleet

Before starting Blockchain:
* docker rm -f $(docker ps -aq)
* docker network prune
* docker rmi docker rmi dev-peer0.org1.example.com-cargo_condition-1.0-d1ad635b3892e37c210060b46f68a20ab2816d0ca2d1d9b7252091c2a3a349cf
* rm hfc-key-store/*


Run scripts to make environment:
* ./startFabric.sh
* docker logs -f ca.example.com
* node enrollAdmin.js
* node registerUser.js
