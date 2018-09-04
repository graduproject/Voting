package main

import (
	"sort"
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

type pair struct {
	key   string
	value int
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
		fmt.Print(key, val, " ")
	}
	fmt.Println()
}

// deleteCandidate deletes candidate in Voting structure
func (v *Voting) deleteCandidate(cd string) { // cd는 후보
	/*if v.StartTime < time.Now().Unix() { // 시작전에만 삭제 가능
		fmt.Println("후보를 삭제할 수 없습니다")
		return
	}*/
	delete(v.Candidate, cd)
}

func (v *Voting) check(id string) bool {
	var b bool
	for _, i := range v.UserID {
		if i == id {
			b = false
			break
		}
	}
	return b
}

// vote increases Poll belong to selected candidate
func (v *Voting) vote(cd string) { // 투표, cd는 후보
	// TODO : 아이디 받아오기
	// id := ~~~~
	// TODO : 중복 체크를 어떻게 할지 생각해보기(매번 탐색?)
	/* if v.check(id) {
		v.Candidate[cd]++
		v.saveCompleteID(id)
	} else {

	} */
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


// TODO : value로 sorting이 되긴 하는데 sorting을 해서 정리할 필요가 있을지 생각해보기
//        투표 결과를 보여줄때 번호 순으로 보여 준다 생각하면 정리할 필요가 없을 것 같음
/* sort by value 긁어온 코드
// TODO : 코드 쓰게되면 정리하기
func rankByWordCount(wordFrequencies map[string]int) PairList{
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
	  pl[i] = Pair{k, v}
	  i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
  }
  
  type Pair struct {
	Key string
	Value int
  }
  
  type PairList []Pair
  
  func (p PairList) Len() int { return len(p) }
  func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
  func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
*/

// saveCompleteID saves ID
func (v *Voting) saveCompleteID(id string) {
	v.UserID = append(v.UserID, id)
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
		votingSlice[0].vote("김현우")
		votingSlice[0].vote("김현우")
		votingSlice[0].vote("김도정")
		votingSlice[0].getCandidate()
		// fmt.Println(rankByWordCount(votingSlice[0].Candidate))
		votingSlice[0].CurrentState = 2
}