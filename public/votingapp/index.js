var sys = require('sys');
var exec = require('child_process').exec;
var temp = require('./chaincode/voting_chaincode.js');
var parser = require('./utils/parser.js');
var fs = require('fs');
var child;
var tem;
//var str = temp.createVoting("1","first-vote","09/19/2018 3:20:00 PM","09/19/2018 8:20:00 PM");
var str = temp.queryAllVote("4");

 var result = '';
// child = exec(str, function(error, stdout, stderr){
// 	tem = stderr;
// 	var s = parser.read_parse(tem);
	//process.stdout.write(s.toString());
	//if(error != null){
 	//	console.log('exec error: ' + error);
 	//}
// });
