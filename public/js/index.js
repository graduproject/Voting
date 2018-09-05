'use strict';

var express = require('express');
var app = express();
var fs = require('fs');
var path = require("path")

app.get('/', function(req, res){
    res.sendFile(path.join(__dirname + '/../index.html'));
    
});
app.get('/images', function(req, res){
    res.readFile('main.png');
});


app.listen(3000, function() {
    console.log("Go!");
  });