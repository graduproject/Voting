package main

import (
	"encoding/json"
)

// Voting is ...
type Voting struct {
	VotingName		string		`json="votingName"`
	UserID			[]string	`json="userid"`
	Candidate		[]string	`json="candidate"`
	poll			[]int 		`json="poll"`
	VotingNumber    int         `json="VotingNumber"`	
	StartTime		int			`json="starttime"`
	EndTime			int 		`json="endtime"`
	CurrentState 	int			`json="currentstate"`  // state 0 : 투표 시작 전, 1 : 투표 가능, 2 : 투표 종료
}

// VotingSlice is ...
var VotingSlice []Voting // 투표 목록

// CreateVote creates Voting structure
func CreateVote() { // Voting 구조체 생성
	v := Voting{}
	VotingSlice = append(VotingSlice, v)
}

func VotingInit() {
	num := len(VotingSlice)
	VotingSlice[num].VotingName = ""
	VotingSlice[num].VotingNumber = num
	VotingSlice[num].StartTime =
	VotingSlice[num].EndTime =
	VotingSlice[num].CurrentState = 0

}

// RegisterCandidate register candidate in Voting structure
func (v *Voting) RegisterCandidate() { // 후보 등록

}

// GetCandidate gets candidate in Voting structure
func (v *Voting) GetCandidate() { // 후보자 확인

}

// ModifyCandidate modifies candidate in Voting structure
func (v *Voting) ModifyCandidate() { // 투표 시작 전에만 수정 가능

}

// ViewPoll views Poll in Voting structure
func (v *Voting) ViewPoll() { // 득표 확인

}

// Vote increases Poll belong to selected candidate
func (v *Voting) Vote() { // 투표

}

// ChangeState change Voting structure's CurrentState
func (v *Voting) ChangeState() { // Voting 상태 변화

}

// ViewCompleteVoting views completed Voting
func ViewCompleteVoting() { // 완료된 투표 목록 조회

}