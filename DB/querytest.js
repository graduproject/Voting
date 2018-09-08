var Client = require('mongodb').MongoClient;

Client.connect('mongodb://localhost:3000/member', function(error, db){
    if(error) {
        console.log(error);
    } else {
        var query = {ID:'One'};
        // 2. 읽어올 Field 선택
        var projection = {name:1, id:0};
        // 3. find() 함수에 작성된 query와 projection을 입력
        var cursor = db.collection('member').find(query,projection);
        cursor.each(function(err,doc){
            if(err){
                console.log(err);
            }else{
                if(doc != null){
                    console.log(doc);
                }
            }
        });
        db.close();
    }
});