var sys = require('sys');
var exec = require('child_process').exec;
var temp = require('./js/command_config/chaincode_service.js');
var child;

var str = temp.createVoting("createVoting","1","first-vote","09/19/2018 3:20:00 PM","09/19/2018 8:20:00 PM");
sys.print(str.toString() + "\n");

child = exec(str.toString(), function(error, stdout, stderr){
	sys.print('stdout: ' + stdout);
	sys.print('stderr: ' + stderr);
	if(error != null){
		console.log('exec error: ' + error);
	}
});



