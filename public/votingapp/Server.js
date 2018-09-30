var express = require('express');
var bodyParser = require('body-parser');
var path = require('path');
var app = express();
var parser = require('./utils/parser');
var cmd = require('./chaincode/voting_chaincode');

app.use(bodyParser.urlencoded({extended: true}));
app.set('view engine', 'ejs'); 
app.set('views','../view');

app.use(express.static(path.join(__dirname, '../../public')));

app.get('/', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/login.html'));
});

app.get('/home', function(req, res){
    var ans = cmd.queryAllVote("4").slice();
	for(var i = 0; i < ans.length; i = i + 3){
		ans[i + 1] = parser.POSIXtoDATE(ans[i + 1]);
		ans[i + 2] = parser.POSIXtoDATE(ans[i + 2]);
	}
	res.render('User/main',{vote : ans});
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
    res.sendFile(path.join(__dirname, '../view/Admin/admin-login.html'));
});

app.get('/admin-main', function(req, res){
    res.sendFile(path.join(__dirname, '../view/Admin/admin-main.html'));
});

app.get('/admin-mypage', function(req, res){
    res.sendFile(path.join(__dirname, '../view/Admin/admin-mypage.html'));
});

app.get('/admin-vote_result', function(req, res){
    res.sendFile(path.join(__dirname, '../view/Admin/admin-vote_result.html'));
});

app.get('/add-candidate', function(req, res){
    res.sendFile(path.join(__dirname, '../view/Admin/add-candidate.html'));
});

app.get('/vote-create', function(req, res){
    res.render('Admin/vote-create');
});

app.post('/vote-create', function(req, res){
    console.log(req.body);
});

app.get('/vote-manage', function(req, res){
    res.sendFile(path.join(__dirname, '../view/Admin/vote-manage.html'));
});

app.listen(3000, function() {
    console.log("Go!");
});
