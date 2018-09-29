console.log('Client-side code running');

var chaincaller = require('../chaincode/voting_chaincode')

const button = document.getElementById('vote');
console.log("HELLOWORLD");
button.addEventListener('click', function(e) {
    var obj = document.getElementsByName('candidate');
    var checked_index = -1;
	var checked_value = '';
	for( i=0; i<obj.length; i++ ) {
		if(obj[i].checked) {
			checked_index = i;
			checked_value = obj[i].value;
		}
    }
    
	alert( '선택된 항목 인덱스: '+checked_index+'\n선택된 항목 값: '+checked_value );
});