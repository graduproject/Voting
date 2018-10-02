exports.puttimeState = function(parameter){
	var old = new Date(parameter);
	var time = new Date();
	if(time < old) return "ON";
	else return "DONE";
}
