var docking = "docker exec cli ";
var invoke = "peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile ";
var pemPath = "/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem ";
var etc = "-C mychannel -n mycc -c ";

function parse(parameter){
	var args = Array.from(arguments);
	var temp = [];
		
	for(var i = 0; i < arguments.length; i++){
		temp += args[i];
	}	

	temp = "\'{\"Args\":" + temp + "}\' ";
	temp = docking + invoke + pemPath + temp;
	console.log(temp);	
	return temp;
}

exports.createVoting = function(vnum, vname, vst_time, ved_time){
	var args = Array.from(arguments);
	var syscmd = parse(args);

	return args;
}

var temp = ["1", "vnum", "201324439", "201324438"];
parse(temp);

