var sys = require('sys');
var exec = require('child_process').exec;
var parser = require('../utils/parser.js');
var fs = require('fs');
var path = require('path');
var child;

exports.createVoting = function(vnum, vname, vst_time, ved_time){
	var args = ['createVoting'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.changeState = function(last_vnum){
	var args = ['changeState'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.registerCandidate = function(vnum,candidNum){
	var args = ['registerCandidate'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.vote = function(vnum, candidName, userID){
	var args = ['vote'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.queryAllVote = function(last_vnum){
	var args = ['queryAllVote'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	var ststr;

	child = exec(syscmd, function(error, stdout, stderr){
		console.log(syscmd);
		console.log(stderr);
		fs.writeFile('./chaincode/Argumentation.inp', stderr,'utf8');
	});
	ststr = fs.readFileSync(__dirname + '/Argumentation.inp', 'utf8');
	return parser.read_parse(ststr);
}

exports.queryCompleteVote = function(last_vnum){
	var args = ['queryCompleteVote'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.earlyComplete = function(vnum){
	var args = ['earlyComplete'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.deleteCandidate = function(vnum, candidName){
	var args = ['deleteCandidate'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.queryNotCompleteVote = function(last_vnum){
	var args = ['queryNotCompleteVote'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.queryCandidateWithPoll = function(last_vnum){
	var args = ['queryCandidateWithPoll'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.queryCandidate = function(vnum){
	var args = ['queryCandidate'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	child = exec(syscmd, function(error, stdout, stderr){});
}


