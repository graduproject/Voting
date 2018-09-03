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
	Candidate       []string    `json="candidate"`
	Poll            []int       `json="poll"`
	VotingNumber    int         `json="VotingNumber"`	
	StartTime       int64       `json="starttime"`
	EndTime         int64       `json="endtime"`
	CurrentState    int         `json="currentstate"`  // state 0 : 투표 시작 전,   1 : 투표 가능,   2 : 투표 종료
	Winner          string      `json="winner"`        // 표를 가장 많이 받은 후보
}

// votingSlice is ...
var votingSlice []Voting // 투표 목록

// createVote creates Voting structure
func createVote(name string, startTime int64, endTime int64) { // Voting 구조체 생성
	// TODO : 시작 시간와 끝나는 시간 Unix시간으로 바꿔서 받아야함
	v := Voting{}
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
	votingSlice[num].Winner = "null"
}

// registerCandidate register candidate in Voting structure
func (v *Voting) registerCandidate(cd string) { // 후보 등록, cd는 후보 이름
	v.Candidate = append(votingSlice[v.VotingNumber-1].Candidate, cd)
	v.Poll = append(votingSlice[v.VotingNumber-1].Poll, 0)
}

// getCandidate gets candidate in Voting structure
func (v *Voting) getCandidate() { // 후보자 확인
	for i := 0; i < len(v.Candidate); i++{
		fmt.Println(v.Candidate[i])
	}
}

// deleteCandidate deletes candidate in Voting structure
func (v *Voting) deleteCandidate(num int) { // num은 후보 번호
	/*if v.StartTime < time.Now().Unix() { // 시작전에만 삭제 가능
		fmt.Println("후보를 삭제할 수 없습니다")
		return
	}*/
	num--
	copy(v.Candidate[num:], v.Candidate[num+1:]) // 후보 이름 삭제
	v.Candidate[len(v.Candidate)-1] = ""
	v.Candidate = v.Candidate[:len(v.Candidate)-1]

	copy(v.Poll[num:], v.Poll[num+1:]) // 후보 표 칸 삭제
	v.Poll[len(v.Poll)-1] = 0
	v.Poll = v.Poll[:len(v.Poll)-1]
}

// viewPoll views Poll in Voting structure
func (v *Voting) viewPoll() { // 득표 확인
	for i := 0; i < len(v.Poll); i++{
		fmt.Print(v.Poll[i], " ")
	}
	fmt.Println()
}

// vote increases Poll belong to selected candidate
func (v *Voting) vote(num int) { // 투표, num은 후보 번호
	num--
	v.Poll[num]++
}

// changeState change Voting structure's CurrentState
func (v *Voting) changeState() { // Voting 상태 변화
	// TODO : 어떻게 할지 생각해봐야함
	if v.StartTime < time.Now().Unix() && v.EndTime > time.Now().Unix() { // 투표 시작
		v.CurrentState = 1
	}

	if v.EndTime < time.Now().Unix() { // 투표가 끝난 상태
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

// selectWinner selects winner
func (v *Voting) selectWinner() {
	var draw []string
	var temp int
	var tempIndex int
	for i := 0; i < len(v.Poll); i++ { // 최대 표와 최대 표를 받은 후보를 임시로 저장
		if temp < v.Poll[i] {
			temp = v.Poll[i]
			tempIndex = i
		}
	}
	for i := 0; i < len(v.Poll); i++ { // 최대 표와 동률인 표 검색
		if temp == v.Poll[i] {
			draw = append(draw, v.Candidate[i])
		}
	}

	if len(draw) == 1 { // 동률이 없을 시 해당 후보가 당선
		v.Winner = v.Candidate[tempIndex]
	} else { // 동률이 있을 시 동률인 후보들 출력
		for i := 0; i < len(draw); i++ { 
			fmt.Print(draw[i], " ")
		}
		fmt.Println("draw")
	}
}

func main() { // Test
		createVote("First", time.Now().Unix(), time.Now().Unix() + 86400)
		votingSlice[0].registerCandidate("이상현")
		votingSlice[0].registerCandidate("김도정")
		votingSlice[0].registerCandidate("김현우")
		votingSlice[0].registerCandidate("유상욱")
		votingSlice[0].registerCandidate("최현빈")
		fmt.Println(len(votingSlice[0].Candidate))
		votingSlice[0].getCandidate()
		votingSlice[0].deleteCandidate(2)
		fmt.Println(len(votingSlice[0].Candidate))
		votingSlice[0].getCandidate()
		votingSlice[0].vote(1)
		votingSlice[0].vote(1)
		votingSlice[0].vote(1)
		votingSlice[0].vote(1)
		votingSlice[0].vote(4)
		votingSlice[0].vote(4)
		votingSlice[0].vote(4)
		votingSlice[0].viewPoll()
		votingSlice[0].selectWinner()
		fmt.Println("Winner : ", votingSlice[0].Winner)
		fmt.Println(votingSlice[0])
		votingSlice[0].CurrentState = 2

		fmt.Println("=====================================================")
		createVote("Second", time.Now().Unix(), time.Now().Unix() + 86400)
		votingSlice[1].registerCandidate("이상현")
		votingSlice[1].registerCandidate("김도정")
		votingSlice[1].registerCandidate("김현우")
		fmt.Println(len(votingSlice[1].Candidate))
		votingSlice[1].getCandidate()
		votingSlice[1].deleteCandidate(2)
		fmt.Println(len(votingSlice[1].Candidate))
		votingSlice[1].getCandidate()
		votingSlice[1].vote(2)
		votingSlice[1].vote(1)
		votingSlice[1].viewPoll()
		fmt.Println(time.Now().Unix())
		fmt.Println(time.Now())
		fmt.Println(votingSlice[1])
		votingSlice[1].CurrentState = 2
		viewCompleteVoting()
}