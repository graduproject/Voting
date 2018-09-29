var express = require('express'),
    bodyParser = require('body-parser'),
    path = require('path');

var app = express();

app.use(bodyParser.urlencoded({extended: true}));
app.set('view engine', 'ejs'); 
app.engine('html', require('ejs').renderFile);
app.use(express.static(path.join(__dirname, '../../public')));

app.get('/', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/main.html'));
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

app.listen(3000, function() {
    console.log("Go!");
});