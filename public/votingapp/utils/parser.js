var docking = "docker exec cli ";
var invoke = "peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile ";
var pemPath = "/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem ";
var etc = "-C mychannel -n mycc -c ";

function isBlank(value){
	return value != '';
}

function isRest(value){
	return value != ',';
}

exports.cmd_parse = function(parameter){
	var args = Array.from(arguments[0]);
	var temp = "[";

	for(var i = 0; i < args.length; i++){
		temp += '\"' + args[i] + '\"';
		if(i != args.length - 1){
			temp += ',';
		}
	}
	temp += "]";
	temp = "\'{\"Args\":" + temp + "}\' ";
	temp = docking + invoke + pemPath + etc + temp;
	return temp;
}

exports.read_parse = function(parameter){
	var str = parameter.toString();
	var token = "\\";
	var token_index = str.indexOf(token);
	var str_temp = str.slice(token_index, str.length - 1);
	str_temp = str_temp.split('\\');
	str_temp.pop();
	str_temp = str_temp.filter(isBlank);
	for(var i = 0; i < str_temp.length; i++){
		var regex = '\"';
		str_temp[i] = str_temp[i].replace(regex, "");
	}
	str_temp = str_temp.filter(isRest);
	str_temp = str_temp.filter(isBlank);
	return str_temp;
}