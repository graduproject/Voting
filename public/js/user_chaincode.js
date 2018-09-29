var parser = require('../utils/parser.js');

exports.signup= function(ID, PW, IDNumber, PhoneNumber, Email){
	var args = Array.from(arguments);
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.signin = function(ID, PW){
	var args = Array.from(arguments);
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.modifyUser = function(ID,PW,PhoneNumber,Email){
	var args = Array.from(arguments);
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.getUserInfo = function(ID){
	var args = Array.from(arguments);
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}

exports.deleteUser = function(ID){
	var args = Array.from(arguments);
	var syscmd = parser.cmd_parse(args);
	syscmd.toString();	
	
	return syscmd;
}
