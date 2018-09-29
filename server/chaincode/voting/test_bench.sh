#!/bin/bash

MODE=$1



function createVoting() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["createVoting","1","first-vote","09/19/2018 3:20:00 PM","09/19/2018 8:20:00 PM"]}'
	
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["createVoting","2","second-vote","09/13/2018 6:20:00 PM","09/13/2018 8:20:00 PM"]}'

	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["createVoting","3","third-vote","09/13/2018 6:20:00 PM","09/13/2018 8:20:00 PM"]}'
}

function registerCandidate() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["registerCandidate","1","Lee"]}'

	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["registerCandidate","1","Kim"]}'

	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["registerCandidate","1","Park"]}'
}

function vote() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["vote","1","Lee","jdsd2233"]}'

	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["vote","1","Kim","jdsd2587"]}'

	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["vote","1","Park","jdsd1234"]}'
}

function changeState() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["changeState","3"]}'
}

function deleteCandidate() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["deleteCandidate","1","Kim"]}'
}

function earlyComplete() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["earlyComplete","1"]}'
}

function queryAllVote() {
	peer chaincode query -C mychannel -n mycc -c '{"Args":["queryAllVote","3"]}'
}

function queryCandidateWithPoll() {
	peer chaincode query -C mychannel -n mycc -c '{"Args":["queryCandidateWithPoll","1"]}'
}

function queryNotCompleteVote() {
	peer chaincode query -C mychannel -n mycc -c '{"Args":["queryNotCompleteVote","3"]}'
}

function queryCandidate() {
	peer chaincode query -C mychannel -n mycc -c '{"Args":["queryCandidate","1"]}'
}

function queryCompleteVote() {
	peer chaincode query -C mychannel -n mycc -c '{"Args":["queryCompleteVote","3"]}'
}

function signup() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc1 -c '{"Args":["signup","jdsd2587","1234","9309131840518","01063717762","jdsd2233@gmail.com"]}'
}

function getUserInfo() {
	peer chaincode query -C mychannel -n mycc1 -c '{"Args":["getUserInfo","jdsd2587"]}'
}

function deleteUser() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc1 -c '{"Args":["deleteUser","jdsd2587"]}'
}

function signin() {
	peer chaincode query -C mychannel -n mycc1 -c '{"Args":["signin","jdsd2587","1234"]}'
}

function modifyUser() {
	peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc1 -c '{"Args":["modifyUser","jdsd2587","1111","01000000000","aa@aa.aa"]}'
}


if [ "${MODE}" == "createVoting" ]; then
  createVoting
elif [ "${MODE}" == "changeState" ]; then
  changeState
elif [ "${MODE}" == "registerCandidate" ]; then
  registerCandidate
elif [ "${MODE}" == "vote" ]; then
  vote
elif [ "${MODE}" == "queryAllVote" ]; then
  queryAllVote
elif [ "${MODE}" == "queryCompleteVote" ]; then
  queryCompleteVote
elif [ "${MODE}" == "earlyComplete" ]; then
  earlyComplete
elif [ "${MODE}" == "deleteCandidate" ]; then
  deleteCandidate
elif [ "${MODE}" == "queryNotCompleteVote" ]; then
  queryNotCompleteVote
elif [ "${MODE}" == "queryCandidateWithPoll" ]; then
  queryCandidateWithPoll
elif [ "${MODE}" == "queryCandidate" ]; then
  queryCandidate
elif [ "${MODE}" == "signup" ]; then
  signup
elif [ "${MODE}" == "signin" ]; then
  signin
elif [ "${MODE}" == "modifyUser" ]; then
  modifyUser
elif [ "${MODE}" == "getUserInfo" ]; then
  getUserInfo
elif [ "${MODE}" == "deleteUser" ]; then
  deleteUser
else
  echo "no"
  exit 1
fi
