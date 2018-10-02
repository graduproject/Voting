var sys = require('sys');
var exec = require('child_process').exec;
var parser = require('../utils/parser.js');
var fs = require('fs');
var path = require('path');
var child;

exports.signup= function(ID, PW, IDNumber, PhoneNumber, Email){
	var args = ['signup'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.signin = function(ID, PW){
	var args = ['signin'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_query(args);
	syscmd.toString();	
	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.modifyUser = function(ID,PW,PhoneNumber,Email){
	var args = ['modifyUser'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.getUserInfo = function(ID){
	var args = ['getUserInfo'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_query(args);
	syscmd.toString();	
	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.deleteUser = function(ID){
	var args = ['deleteUser'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	child = exec(syscmd, function(error, stdout, stderr){});
}

exports.isAdmin = function(ID){
	var args = ['isAdmin'];
	args = args.concat(Array.from(arguments));
	var syscmd = parser.cmd_query(args);
	syscmd.toString();	

	var ststr;

	child = exec(syscmd, function(error, stdout, stderr){
		fs.writeFile('./chaincode/Argumentation_B.inp', stderr,'utf8');
	});
	ststr = fs.readFileSync(__dirname + '/Argumentation_B.inp', 'utf8');
	return parser.read_parse(ststr);
}

