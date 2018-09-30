function addRow(tableID, num){
    if(num == 1){
        console.log(num);
        var table=document.getElementById(tableID);
        var rowLen=table.rows.length;
        var row=table.insertRow();
        if(rowLen>6){
            alert("후보는 6개까지 추가가능합니다.")
            return 0;
    }
    row.insertCell(0).innerHTML=rowLen;
    row.insertCell(1).innerHTML="<input type=text name=candidate-name>";
    }
    else{
        console.log(num);
        var table=document.getElementById(tableID);
        var rowLen=table.rows.length;
        console.log(__dirname + rowLen);
        if(rowLen>=2){
            console.log(__dirname);
            location.href="/admin-main";
        }
        else{
            alert("후보자는 2명이상 입력되야합니다.");
        }
    }
}

function delRow(tableID){
    var table=document.getElementById(tableID);
    var rowLen=table.rows.length-1;
    if(rowLen>0){
        table.deleteRow(rowLen);
    }
}

function complete(tableID){
    
}