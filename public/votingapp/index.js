var sys = require('sys');
var exec = require('child_process').exec;
var temp = require('./js/command_config/voting_chaincode.js');
var parser = require('./js/utils/parser.js');
var child;
var tem;
//var str = temp.createVoting("createVoting","1","first-vote","09/19/2018 3:20:00 PM","09/19/2018 8:20:00 PM");
var str = temp.queryAllVote("queryAllVote", "4");
sys.print(str.toString() + "\n");

child = exec(str, function(error, stdout, stderr){
	tem = stderr;
	parser.read_parse(tem);
	if(error != null){
	//	console.log('exec error: ' + error);
	}
});


