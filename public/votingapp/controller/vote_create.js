var info = new Object();

exports.getCandid = function(parameter){
	info.candid = parameter;
	console.log("candid " + candid);
	candidTaken = 1;
}

exports.getVoteinfo = function(title,yy,mm,dd){
	info.title = title;
	info.yy = yy;
	info.mm = mm;
	info.dd = dd;
}

exports.iscandidTaken = function(){
	return candidTaken;
}

exports.putCandid = function(){
	return candid;
}


