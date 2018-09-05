package main

import (
	//"sort"
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

func changeToUnixTime(str string) int64 { // string으로 받은 시간을 Unix 시간으로 바꿔준다
	layout := "01/02/2006 3:04:05 PM" 
	t, _ := time.Parse(layout, str)
	tUTC := t.Unix() - 32400  // 받은 시간은 KST, Unix() 시간은 UTC
	return tUTC
}

// createVote creates Voting structure
func createVote(name string, startTime string, endTime string) { // Voting 구조체 생성
	v := Voting{Candidate: make(map[string]int)}
	votingSlice = append(votingSlice, v)
	votingInit(name, changeToUnixTime(startTime), changeToUnixTime(endTime))
}

// votingInit is ...
func votingInit(name string, startTime int64, endTime int64) { // 투표 초기값 입력
	num := len(votingSlice) - 1
	votingSlice[num].VotingName = name
	votingSlice[num].VotingNumber = num + 1
	votingSlice[num].StartTime = startTime
	votingSlice[num].EndTime = endTime
	votingSlice[num].CurrentState = 0
}

// registerCandidate register candidate in Voting structure
func (v *Voting) registerCandidate(cd string) { // 후보 등록, cd는 후보 이름
	if v.CurrentState == 0 { // 투표 시작 전에만 후보 등록 가능
		v.Candidate[cd] = 0
	} else {
		fmt.Println("후보를 등록할 수 없습니다")
	}
}

// getCandidate gets candidate in Voting structure
func (v *Voting) getCandidate() { // 후보 및 표 확인
	for key, val := range v.Candidate {
		fmt.Print(key, " ", val, " ")
	}
	fmt.Println()
}

// deleteCandidate deletes candidate in Voting structure
func (v *Voting) deleteCandidate(cd string) { // cd는 후보
	if v.CurrentState == 0 { // 시작전에만 삭제 가능
		fmt.Println("후보를 삭제할 수 없습니다")
		return
	}
	delete(v.Candidate, cd)
}

func (v *Voting) checkID(id string) bool { // 투표를 이미 한 id인지 체크
	b := true
	for _, i := range v.UserID {
		if i == id {
			b = false
			break
		}
	}
	return b
}

func (v *Voting) checkCandidateExist(cd string) bool { // 후보가 존재하는지 확인 
	_, exist := v.Candidate[cd]
	return exist
}

// vote increases Poll belong to selected candidate
func (v *Voting) vote(cd string, userID string) { // 투표, cd는 후보
	id := userID
	if v.CurrentState == 0 { // 투표 시작 전
		fmt.Println("아직 투표할 수 없습니다")
	} else if v.CurrentState == 1 && v.checkCandidateExist(cd) { // 투표를 할수 있는 상태 && 후보가 존재하면 -> 투표
		if v.checkID(id) {
			fmt.Println("확인")
			v.Candidate[cd] = v.Candidate[cd] + 1
			v.saveCompleteID(id)
		} else { // 투표가 끝난 후
			fmt.Println("중복")
		}
	} else {
		fmt.Println("투표가 끝났습니다")
	}
}

// changeState change Voting structure's CurrentState
func changeState() { // Voting 상태 변화 실시간으로 체크해서 투표의 상태를 변경한다(모든 투표를 대상으로 확인)
	for i := range votingSlice {
		if votingSlice[i].StartTime < time.Now().Unix() && votingSlice[i].EndTime > time.Now().Unix() { // 투표 시작
			votingSlice[i].CurrentState = 1
		} else if votingSlice[i].EndTime < time.Now().Unix() { // 투표가 끝난 상태
			votingSlice[i].CurrentState = 2
		}
	}
}

func earlyComplete(num int) { // 투표 번호를 받아서 투표 조기 종료
	votingSlice[num - 1].CurrentState = 2
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
func (v *Voting) saveCompleteID(id string) { // 투표 완료한 아이디 저장
	v.UserID = append(v.UserID, id)
}

func main() { // Test
	createVote("First", "09/05/2018 6:40:00 PM", "09/05/2018 6:41:00 PM")
	for {
		changeState()
		votingSlice[0].registerCandidate("이상현")
		votingSlice[0].registerCandidate("김도정")
		votingSlice[0].registerCandidate("김현우")
		votingSlice[0].registerCandidate("유상욱")
		votingSlice[0].registerCandidate("최현빈")
		fmt.Println(votingSlice[0])
		votingSlice[0].vote("이상현", "a")
		votingSlice[0].vote("이상현", "b")
		changeState()
		votingSlice[0].vote("이상현", "c")
		votingSlice[0].vote("김현우", "d")
		votingSlice[0].vote("김현우", "a")
		votingSlice[0].vote("김도정", "b")
		votingSlice[0].getCandidate()
		fmt.Println(votingSlice[0])
		votingSlice[0].vote("이상현", "e")
		fmt.Println(votingSlice[0])
		fmt.Println("==========================================")
		fmt.Println()
		fmt.Println(votingSlice)
		time.Sleep(10 * time.Second)
	}
}