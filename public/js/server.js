'use strict';

var express = require('express');
var http = require('http');
var url = require('url');
var app = express();
var path = require("path")
var fs = require('fs');

//SHOULD BE CHANGED//
// var mysql = require('mysql');
//SHOULD BE CHANGED//

// var client = mysql.createConnection({
//     user : 'root',
//     password : 'apmsetup',
//     database : 'Company'
// });

app.use(express.static(path.join(__dirname + '/../../public')));
app.use(express.json());
app.use(express.urlencoded());
// app.use(app.router);

app.get('/', function (req, res) {
    fs.readFile(path.join(__dirname + '/../index.html'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html' });
        res.end(data);
    });
});

app.get('/vote', function (req, res) {
    fs.readFile(path.join(__dirname + '/../app/voting_authen.html'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html' });
        res.end(data);
    });
});

app.get('/images', function (req, res) {
    fs.readFile(path.join(__dirname + '/images/main.png'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html' });
        res.end(data);
    });
});

app.get('/choose_menu', function (req, res) {
    fs.readFile(path.join(__dirname + '/../app/Administrator/choose_menu.html'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html' });
        res.end(data);
    });
});

app.get('/admin', function (req, res) {
    fs.readFile(path.join(__dirname + '/../login.html'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html' });
        res.end(data);
    });
});

app.get('/enroll_vote', function (req, res) {
    fs.readFile(path.join(__dirname + '/../app/Administrator/vote_enroll.html'), function (error, data) {
        res.writeHead(200, { 'Content-Type': 'text/html' });
        res.end(data);
    });
});


adminPartyApp.post('?????', function(req, res) {
    logger.info('JOIN CHANNEL');
    var chaincodeName = req.params.channelName;
    var peersId = req.body.peers || [];
    var peers   = peersId.map(getPeerHostByCompositeID);
    logger.debug('channelName : ' + chaincodeName);
    logger.debug('peers : ' + peers);

    if (!chaincodeName) {
        res.error(getErrorMessage('\'channelName\''));
        return;
    }
    if (!peers || peers.length === 0) {
        res.error(getErrorMessage('\'peers\''));
        return;
    }

    res.promise(
        joinChannel.joinChannel(peers, chaincodeName, USERNAME, ORG)
    );
});

app.listen(3000, function () {
    console.log(__dirname);
});
