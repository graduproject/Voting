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
	//Parser the Output of Ledger
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

exports.POSIXtoDATE = function(timestamp){
	var d = new Date(timestamp * 1000), // Convert the passed timestamp to millisecond
        yyyy = d.getFullYear(),
        mm = ('0' + (d.getMonth() + 1)).slice(-2),  // Months are zero based. Add leading 0.
        dd = ('0' + d.getDate()).slice(-2),         // Add leading 0.
        hh = d.getHours(),
        h = hh,
        min = ('0' + d.getMinutes()).slice(-2),     // Add leading 0.
        ampm = 'AM',
        time;

    if (hh > 12) {
        h = hh - 12;
        ampm = 'PM';
    } else if (hh === 12) {
        h = 12;
        ampm = 'PM';
    } else if (hh == 0) {
        h = 12;
    }

    // ie: 2014-03-24, 3:00 PM
    time = yyyy + '-' + mm + '-' + dd + ', ' + h + ':' + min + ' ' + ampm;
    return time
}
