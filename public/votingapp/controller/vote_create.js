var candid;
var candidTaken = 0;

exports.getCandid = function(parameter){
	candid = parameter;
	console.log("candid " + candid);
	candidTaken = 1;
}

exports.getVoteinfo = function(title,yy,mm,dd){
	
}

exports.iscandidTaken = function(){
	return candidTaken;
}

exports.putCandid = function(){
	return candid;
}


