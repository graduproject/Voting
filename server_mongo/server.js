var mongoose = require('mongoose');
mongoose.connect('mongodb://localhose:27017/testDB');
var db = mongoose.connection;

db.on('error', function(){
    console.log('Connection Failed');
});

db.once('open', function(){
    console.log('Connected');
});

var member = mongoose.Schema({
    ident : 'string',
    password : 'string'
})

var Member = mongoose.model('Schema', member);
var newMember = new Member({ident: 'helloworld', password : '1q2w3e'});
newMember.save(function(error, data){
    if(error){
        console.log(error);
    }else{
        console.log('Saved');
    }
});

Member.find(function(error, members){
    console.log('---Read all---');
    if(error){
        console.log(error);
    }else{
        console.log(members);
    }
})

Member.findOne({_id:'123456789'}, function(error,member){
    console.log('--- Read one ---');
    if(error){
        console.log(error);
    }else{
        console.log(member);
    }
});

Member.findById({_id:'123456789'}, function(error,member){
    console.log('--- Update(PUT) ---');
    if(error){
        console.log(error);
    }else{
        member.name = '--modified--';
        member.save(function(error,modified_member){
            if(error){
                console.log(error);
            }else{
                console.log(modified_member);
            }
        });
    }
});

Member.remove({_id:'123456788'}, function(error,output){
    console.log('--- Delete ---');
    if(error){
        console.log(error);
    }
    console.log('--- deleted ---');
});
