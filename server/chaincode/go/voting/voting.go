package main

import (
	"fmt"
	//"encoding/json"
	"time"
)

// Voting is ...
type Voting struct {
	VotingName      string      `json="votingname"`
	UserID          []string    `json="userid"`
	Candidate       map[string]int
	VotingNumber    int         `json="VotingNumber"`	
	StartTime       int64       `json="starttime"`
	EndTime         int64       `json="endtime"`
	CurrentState    int         `json="currentstate"`  // state 0 : 투표 시작 전,   1 : 투표 가능,   2 : 투표 종료
}

// votingSlice is ...
var votingSlice []Voting // 투표 목록

// createVote creates Voting structure
func createVote(name string, startTime int64, endTime int64) { // Voting 구조체 생성
	// TODO : 시작 시간와 끝나는 시간 Unix시간으로 바꿔서 받아야함
	v := Voting{Candidate: make(map[string]int)}
	votingSlice = append(votingSlice, v)
	votingInit(name, startTime, endTime)
}

// votingInit is ...
func votingInit(name string, startTime int64, endTime int64) {
	// TODO : 값 받아오기
	// ======================
	num := len(votingSlice) - 1
	votingSlice[num].VotingName = name
	votingSlice[num].VotingNumber = num + 1
	votingSlice[num].StartTime = startTime
	votingSlice[num].EndTime = endTime
	votingSlice[num].CurrentState = 0
}

// registerCandidate register candidate in Voting structure
func (v *Voting) registerCandidate(cd string) { // 후보 등록, cd는 후보 이름
	v.Candidate[cd] = 0
}

// getCandidate gets candidate in Voting structure
func (v *Voting) getCandidate() { // 후보 및 표 확인
	for key, val := range v.Candidate {
		fmt.Println(key, val)
	}
}

// deleteCandidate deletes candidate in Voting structure
func (v *Voting) deleteCandidate(cd string) { // cd는 후보
	/*if v.StartTime < time.Now().Unix() { // 시작전에만 삭제 가능
		fmt.Println("후보를 삭제할 수 없습니다")
		return
	}*/
	delete(v.Candidate, cd)
}

// vote increases Poll belong to selected candidate
func (v *Voting) vote(cd string) { // 투표, cd는 후보
	v.Candidate[cd]++
}

// changeState change Voting structure's CurrentState
func (v *Voting) changeState() { // Voting 상태 변화
	// TODO : 어떻게 할지 생각해봐야함
	if v.StartTime < time.Now().Unix() && v.EndTime > time.Now().Unix() { // 투표 시작
		v.CurrentState = 1
	} else if v.EndTime < time.Now().Unix() { // 투표가 끝난 상태
		v.CurrentState = 2
		// TODO : 표가 가장 많은 후보를 Winner로
	}
}

// viewCompleteVoting views completed Voting
func viewCompleteVoting() { // 전체 투표 목록 중 완료된 투표 조회
	for i := 0; i < len(votingSlice); i++ {
		if votingSlice[i].CurrentState == 2 { // 상태가 2인 투표들은 투표가 완료된 것들
			fmt.Println(votingSlice[i])
		}
	}
}

// saveCompleteID saves ID
func (v *Voting) saveCompleteID() {
	// TODO : 투표 완료한 아이디 추가
}

func main() { // Test
		createVote("First", time.Now().Unix(), time.Now().Unix() + 86400)
		votingSlice[0].registerCandidate("이상현")
		votingSlice[0].registerCandidate("김도정")
		votingSlice[0].registerCandidate("김현우")
		votingSlice[0].registerCandidate("유상욱")
		votingSlice[0].registerCandidate("최현빈")
		votingSlice[0].vote("이상현")
		votingSlice[0].vote("이상현")
		votingSlice[0].vote("이상현")
		votingSlice[0].vote("이상현")
		votingSlice[0].deleteCandidate("유상욱")
		votingSlice[0].getCandidate()
		fmt.Println(votingSlice[0])
		votingSlice[0].CurrentState = 2
}