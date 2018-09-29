var sys = require('sys');
var exec = require('child_process').exec;
var parser = require('../utils/parser.js');

var child;

exports.createVoting = function(vnum, vname, vst_time, ved_time){
	var args = ['createVoting']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.changeState = function(last_vnum){
	var args = ['changeState']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.registerCandidate = function(vnum,candidNum){
	var args = ['registerCandidate']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.vote = function(vnum, candidName, userID){
	var args = ['vote']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.queryAllVote = function(last_vnum){
	var result = [];
	var args = ['queryAllVote']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	child = exec(syscmd, function(error, stdout, stderr){
		var temp = stderr;
		result = parser.read_parse(temp);
	});

	return result;
}

exports.queryCompleteVote = function(last_vnum){
	var args = ['queryCompleteVote']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.earlyComplete = function(vnum){
	var args = ['earlyComplete']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.deleteCandidate = function(vnum, candidName){
	var args = ['deleteCandidate']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.queryNotCompleteVote = function(last_vnum){
	var args = ['queryNotCompleteVote']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.queryCandidateWithPoll = function(last_vnum){
	var args = ['queryCandidateWithPoll']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.queryCandidate = function(vnum){
	var args = ['queryCandidate']
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}


