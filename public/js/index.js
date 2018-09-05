'use strict';

var express = require('express');
var app = express();
var path = require("path")
var fs = require('fs');

app.use(express.static(path.join(__dirname + '/../../public')));

app.get('/', function (req, res) {
    fs.readFile(path.join(__dirname + '/../index.html'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html; charset = utf-8' });
        res.end(data);
    });
});

app.get('/vote', function (req, res) {
    fs.readFile(path.join(__dirname + '/../index.html'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html; charset = utf-8' });
        res.end(data);
    });
});

app.get('/images', function (req, res) {
    fs.readFile(path.join(__dirname + '/../images/main.png'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html; charset = utf-8' });
        res.end(data);
    });
});

app.listen(3000, function () {
    console.log("Go!");
});
