var express = require('express');
var bodyParser = require('body-parser');
var path = require('path');
var app = express();
var cmd = require('./chaincode/voting_chaincode');

app.use(bodyParser.urlencoded({extended: true}));
app.set('view engine', 'ejs'); 
app.set('views','../view');

app.use(express.static(path.join(__dirname, '../../public')));

app.get('/', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/login.html'));
});

app.get('/candidate', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/candidate.html'));
});

app.get('/ended_vote', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/ended_vote.html'));
});

app.get('/findID', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/findID.html'));
});

app.get('/findpage', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/findpage.html'));
});

app.get('/findPW', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/findPW.html'));
});

app.get('/home', function(req, res){
    var ans = cmd.queryAllVote("queryAllVote","4");
    console.log(ans);
	res.render('User/main',{vote: ans});
});

app.get('/mypage', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/mypage.html'));
});

app.get('/register', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/register.html'));
});

app.get('/vote_result', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/vote_result.html'));
});

app.get('/admin-login', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/admin-login.html'));
});

app.get('/admin-main', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/admin-main.html'));
});

app.get('/admin-mypage', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/admin-mypage.html'));
});

app.get('/admin-vote_result', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/admin-vote_result.html'));
});

app.get('/add-candidate', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/add-candidate.html'));
});

app.get('/vote-create', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/vote-create.html'));
});

app.get('/vote-manage', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/vote-manage.html'));
});

app.listen(3000, function() {
    console.log("Go!");
});
