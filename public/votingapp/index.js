var sys = require('sys');
var exec = require('child_process').exec;
var temp = require('./chaincode/voting_chaincode.js');
var parser = require('./utils/parser.js');
var fs = require('fs');
var child;
var tem;
//var str = temp.createVoting("1","first-vote","09/19/2018 3:20:00 PM","09/19/2018 8:20:00 PM");
// var str = temp.queryAllVote("queryAllVote", "4");

 var result = '';
// child = exec(str, function(error, stdout, stderr){
// 	tem = stderr;
// 	var s = parser.read_parse(tem);
	//process.stdout.write(s.toString());
	//if(error != null){
 	//	console.log('exec error: ' + error);
 	//}
// });

var data = "FUCKFUCK";
fs.writeFile('WriteASync.txt', data ,'utf8', function(error, data){
	if (error) {throw error};
	console.log("ASync Write Complete");
  });
 data = "TOCKTOCK"; 
  fs.writeFile('WriteASync.txt', data ,'utf8', function(error, data){
	if (error) {throw error};
	console.log("ASync Write Complete");
  });




