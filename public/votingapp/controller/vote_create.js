var info = new Object();
var candidTaken = 1;

exports.getCandid = function(parameter){
	info.candid = parameter;
	candidTaken++;
}

exports.getVoteinfo = function(title,yy,mm,dd,tt){
	info.title = title;
	var dt = [];
	for(var i = 0; i < 2; i++){
		if(mm[i] < 10) mm[i] = "0" + mm[i].toString();
		else mm[i] = mm[i].toString();
		if(dd[i] < 10) dd[i] = "0" + dd[i].toString();
		else dd[i] = dd[i].toString();
		if(tt[i] > 12) {
			tt[i] = tt[i] - 12;
			tt[i] = tt[i].toString() + ":00:00 PM";
		}
		else tt[i] = tt[i].toString() + ":00:00 AM";
		dt.push(mm[i] + '/' + dd[i] + '/' + yy[i] + ' ' + tt[i]);
	}
	info.date = dt;
}

exports.putVoteinfo = function(){
	return info;
}

exports.iscandidTaken = function(){
	return candidTaken;
}

exports.putCandid = function(){
	return info.candid;
}

exports.putTitle = function(){
	return info.title;
}

exports.putDate = function(){
	return info.date;
}

//"createVoting","2","second-vote","09/13/2018 6:20:00 PM","09/13/2018 8:20:00 PM"]
