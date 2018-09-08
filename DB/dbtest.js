var Client = require('mongodb').MongoClient;

Client.connect('mongodb://localhost:3000/member', function(error, db){
    if(error) {
        console.log(error);
    } else {
        // 1. 입력할 documents 를 미리 생성
        var one = {ID:'One', password: '1234'};
        var two = {ID:'Two', password : '1234'};
        var three = {ID:'Three', password : '1234'};

        // 2. insertMany( ) 함수에 배열 형태로 입력
        db.collection('member').insertMany([one, two, three]);
        db.close();
    }
});