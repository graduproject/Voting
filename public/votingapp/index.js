var sys = require('sys');
var exec = require('child_process').exec;
var temp = require('./chaincode/voting_chaincode.js');
var parser = require('./utils/parser.js');
var child;
var tem;
//var str = temp.createVoting("createVoting","1","first-vote","09/19/2018 3:20:00 PM","09/19/2018 8:20:00 PM");
var str = temp.queryAllVote("queryAllVote", "4");
sys.print(str.toString() + "\n");

child = exec(str, function(error, stdout, stderr){
	tem = stderr;
	var s = parser.read_parse(tem);
	console.log(s);
	if(error != null){
	//	console.log('exec error: ' + error);
	}
});


