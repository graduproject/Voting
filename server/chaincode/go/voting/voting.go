package main

import (
	"fmt"
	//"encoding/json"
	"time"
)

// Voting is ...
type Voting struct {
	VotingName      string      `json="votingName"`
	UserID          []string    `json="userid"`
	Candidate       []string    `json="candidate"`
	Poll            []int       `json="poll"`
	VotingNumber    int         `json="VotingNumber"`	
	StartTime       int64       `json="starttime"`
	EndTime         int64       `json="endtime"`
	CurrentState    int         `json="currentstate"`  // state 0 : 투표 시작 전, 1 : 투표 가능, 2 : 투표 종료
}

// VotingSlice is ...
var VotingSlice []Voting // 투표 목록

// CreateVote creates Voting structure
func CreateVote(name string, startTime int64, endTime int64) { // Voting 구조체 생성
	v := Voting{}
	VotingSlice = append(VotingSlice, v)
	VotingInit(name, startTime, endTime)
}

// VotingInit is ...
func VotingInit(name string, startTime int64, endTime int64) {
	// TODO : 값 받아오기
	// ======================
	num := len(VotingSlice) - 1
	VotingSlice[num].VotingName = name
	VotingSlice[num].VotingNumber = num + 1
	VotingSlice[num].StartTime = startTime
	VotingSlice[num].EndTime = endTime
	VotingSlice[num].CurrentState = 0
}

// RegisterCandidate register candidate in Voting structure
func (v *Voting) RegisterCandidate(cd string) { // 후보 등록
	VotingSlice[v.VotingNumber-1].Candidate = append(VotingSlice[v.VotingNumber-1].Candidate, cd)
	VotingSlice[v.VotingNumber-1].Poll = append(VotingSlice[v.VotingNumber-1].Poll, 0)
}

// GetCandidate gets candidate in Voting structure
func (v *Voting) GetCandidate() { // 후보자 확인
	for i := 0; i < len(VotingSlice[v.VotingNumber-1].Candidate); i++{
		fmt.Println(VotingSlice[v.VotingNumber-1].Candidate[i])
	}
}

// DeleteCandidate deletes candidate in Voting structure
func (v *Voting) DeleteCandidate(num int) {
	/*if v.StartTime < time.Now().Unix() {
		return
	}*/
	num--
	copy(VotingSlice[v.VotingNumber-1].Candidate[num:], VotingSlice[v.VotingNumber-1].Candidate[num+1:]) //
	VotingSlice[v.VotingNumber-1].Candidate[len(VotingSlice[v.VotingNumber-1].Candidate)-1] = ""
	VotingSlice[v.VotingNumber-1].Candidate = VotingSlice[v.VotingNumber-1].Candidate[:len(VotingSlice[v.VotingNumber-1].Candidate)-1]
}

// ViewPoll views Poll in Voting structure
func (v *Voting) ViewPoll() { // 득표 확인
	for i := 0; i < len(VotingSlice[v.VotingNumber-1].Poll); i++{
		fmt.Print(VotingSlice[v.VotingNumber-1].Poll[i], " ")
	}
	fmt.Println()
}

// Vote increases Poll belong to selected candidate
func (v *Voting) Vote(num int) { // 투표
	num--
	VotingSlice[v.VotingNumber-1].Poll[num]++
}

// ChangeState change Voting structure's CurrentState
func (v *Voting) ChangeState() { // Voting 상태 변화
	// TODO : 어떻게 할지 생각해봐야함	
}

// ViewCompleteVoting views completed Voting
func ViewCompleteVoting() { // 완료된 투표 목록 조회
	for i := 0; i < len(VotingSlice); i++ {
		if VotingSlice[i].CurrentState == 2 {
			fmt.Println(VotingSlice[i])
		}
	}
}

func main() {
		CreateVote("First", time.Now().Unix(), time.Now().Unix() + 86400)
		VotingSlice[0].RegisterCandidate("이상현")
		VotingSlice[0].RegisterCandidate("김도정")
		VotingSlice[0].RegisterCandidate("김현우")
		fmt.Println(len(VotingSlice[0].Candidate))
		VotingSlice[0].GetCandidate()
		VotingSlice[0].DeleteCandidate(2)
		fmt.Println(len(VotingSlice[0].Candidate))
		VotingSlice[0].GetCandidate()
		VotingSlice[0].Vote(1)
		VotingSlice[0].Vote(1)
		VotingSlice[0].Vote(1)
		VotingSlice[0].ViewPoll()
		fmt.Println(time.Now().Unix())
		fmt.Println(time.Now())
		fmt.Println(VotingSlice[0])
}