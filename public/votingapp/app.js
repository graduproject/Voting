'use strict';

var express = require('express'),
    bodyParser = require('body-parser'),
    path = require('path');

var app = express();

app.use(bodyParser.urlencoded({extended: true}));

console.log(__dirname + '/../');

app.get('/', function(req, res){
    res.sendFile(path.join(__dirname, '../view/user/main.html'));
});

app.listen(3000, function() {
    console.log("Go!");
  });