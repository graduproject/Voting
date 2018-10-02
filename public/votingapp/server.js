var express = require('express');
var bodyParser = require('body-parser');
var path = require('path');
var app = express();
var parser = require('./utils/parser');
var cmd = require('./chaincode/voting_chaincode');
var ucmd = require('./chaincode/user_chaincode');
var vc_ctrl = require('./controller/vote_create.js');
var fs = require('fs');
var v_idx = fs.readFileSync('./controller/index.inp');
var th = require('./utils/time_handler');

var endSet = new Set();

app.use(bodyParser.urlencoded({extended: true}));
app.set('view engine', 'ejs'); 
app.set('views','../view');

app.use(express.static(path.join(__dirname, '../../public')));

app.get('/', function(req, res){
    res.render('User/login');
});

app.post('/', function(req, res){
	console.log(req.body);
});

app.get('/home', function(req, res){
    var ans = cmd.queryAllVote(v_idx.toString()).slice();
	var st = [];
	for(var i = 0; i < ans.length; i = i + 3){
		ans[i + 1] = parser.POSIXtoDATE(ans[i + 1]);
		ans[i + 2] = parser.POSIXtoDATE(ans[i + 2]);
		st.push(th.puttimeState(ans[i+2]));
	}
	res.render('User/main',{vote : ans, st : st});
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
    res.render('Admin/admin-main')
});

app.get('/admin-mypage', function(req, res){
    res.sendFile(path.join(__dirname, '../view/Admin/admin-mypage.html'));
});

app.get('/admin-vote_result', function(req, res){
    res.sendFile(path.join(__dirname, '../view/Admin/admin-vote_result.html'));
});

app.get('/add-candidate', function(req, res){
    res.render('Admin/add-candidate')
});

app.post('/add-candidate', function(req,res){
	var cand = req.body['candidate-name'];	
	vc_ctrl.getCandid(cand);
    if(vc_ctrl.iscandidTaken() >= 2) {
		var cand_info = vc_ctrl.putCandid();
		var vote_date = vc_ctrl.putDate();
		var vote_title = vc_ctrl.putTitle();
		console.log(cand_info[1]);
		for(var i = 0; i < cand_info.length; i++)
			cmd.registerCandidate(v_idx, cand_info[i]);
		cmd.createVoting(v_idx, vote_title, vote_date[0], vote_date[1]);
		fs.writeFileSync('./controller/index.inp', ++v_idx, 'utf8');
		//res.redirect('/admin-main');
	}		
	
	res.redirect('/admin-main');
});

app.get('/vote-create', function(req, res){
    res.render('Admin/vote-create');
});

app.post('/vote-create', function(req, res){
    var title = req.body['title'].slice();
    var yy = req.body['year'].slice();
    var mm = req.body['month'].slice();
    var dd = req.body['day'].slice();
	var tt = req.body['time'].slice();
	vc_ctrl.getVoteinfo(title,yy,mm,dd,tt);
	res.redirect('/add-candidate');
});

app.get('/vote-manage', function(req, res){
    var ans = cmd.queryNotCompleteVote(v_idx.toString()).slice();
	var st = [];
	for(var i = 0; i < ans.length; i = i + 3){
		ans[i + 1] = parser.POSIXtoDATE(ans[i + 1]);
		ans[i + 2] = parser.POSIXtoDATE(ans[i + 2]);
		st.push(th.puttimeState(ans[i+2]));
	}
	for(var i = 0; i < st.length; i++)
		st[i] = "ON";
    
	var ans2 = cmd.queryNotCompleteVote(v_idx.toString()).slice();
	var st2 = [];
	for(var i = 0; i < ans.length; i = i + 3){
		ans2[i + 1] = parser.POSIXtoDATE(ans2[i + 1]);
		ans2[i + 2] = parser.POSIXtoDATE(ans2[i + 2]);
		st2.push(th.puttimeState(ans2[i+2]));
	}
	for(var i = 0; i < st2.length; i++)
		st2[i] = "DONE";
    	
	res.render('Admin/vote-manage',{ans : ans, st : st, ans2 : ans2, st2 : st2});
});

app.post('/vote-manage', function(req, res){
	var eid = req.body['idx'];
	cmd.earlyComplete(eid);
	res.redirect('/vote-manage');
});

app.listen(3000, function() {
    console.log("Go!");
});
