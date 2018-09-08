package main

import (
	//"sort"
	"fmt"
	"encoding/json"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Voting is ...
type Voting struct {
	VotingName      string         `json="votingname"`
	UserID          []string       `json="userid"`
	Candidate       map[string]int `json-"candidate"`
	StartTime       int64          `json="starttime"`
	EndTime         int64          `json="endtime"`
	CurrentState    int            `json="currentstate"`  // state 0 : 투표 시작 전,   1 : 투표 가능,   2 : 투표 종료
}

type VotingChaincode struct {
	stub     shim.ChaincodeStubInterface
	function string
	args     []string
}

/*
func (t *UserChaincode) call() pb.Response {
	function := t.function
	
	callMap := map[string]func() pb.Response {
		"createVoting":             t.createVoting,
		""
	}
}
*/

// Init ...
func (v *VotingChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Voting Init")

	/*
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}*/

	return shim.Success(nil)
}

func (v *VotingChaincode) createVote() pb.Response { 
	args := v.args // 투표 번호, 이름, 시작 시간, 끝 시간
	
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	startTime := changeToUnixTime(args[2])
	endTime := changeToUnixTime(args[3])

	voting := Voting{VotingName: args[1], StartTime: startTime, EndTime: endTime, CurrentState: 0}
	votingAsBytes, _ := json.Marshal(voting)
	v.stub.PutState(args[0], votingAsBytes)

	return shim.Success(nil)
}

// TODO: 어떻게 처리 할 것인지 / 처음부터 끝까지 모든 투표를 다 받아와서 조건 확인 후 상태 변화
func (v *VotingChaincode) changeState() pb.Response {
	args := v.args // 마지막 투표 번호

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	
	// voteAsBytes, _ := v.stub.GetState(args)

	return shim.Success(nil)
}

func (v *VotingChaincode) registerCandidate() pb.Response {
	args := v.args // 투표 번호, 후보 이름

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	votingAsBytes, _ := v.stub.GetState(args[0])
	voting := Voting{}

	json.Unmarshal(votingAsBytes, &voting)
	voting.Candidate[args[1]] = 0

	votingAsBytes, _ = json.Marshal(voting)
	v.stub.PutState(args[0], votingAsBytes)

	return shim.Success(nil)
}

func (v *VotingChaincode) vote() pb.Response {
	args := v.args // 투표 번호, 후보 이름, 유저 아이디

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	id := args[2]

	votingAsBytes, _ := v.stub.GetState(args[0])
	voting := Voting{}

	json.Unmarshal(votingAsBytes, &voting)

	if voting.CurrentState == 0 { // 투표 시작 전
		fmt.Println("아직 투표할 수 없습니다")
	} else if voting.CurrentState == 1 && voting.checkCandidateExist(args[1]) { // 투표를 할수 있는 상태 && 후보가 존재하면 -> 투표
		if voting.checkID(id) { // 이미 투표 했을때
			fmt.Println("확인")
			voting.Candidate[args[1]] = voting.Candidate[args[1]] + 1
			voting.UserID = append(voting.UserID, id)
		} else { 
			fmt.Println("중복")
		}
	} else if voting.CurrentState == 2 { // 투표가 끝난 후
		fmt.Println("투표가 끝났습니다")
	}

	votingAsBytes, _ = json.Marshal(voting)
	v.stub.PutState(args[0], votingAsBytes)

	return shim.Success(nil)
}


// TODO: 데이터 받아와서 처리하는 부분 구현
func (v *VotingChaincode) queryAllVote() pb.Response { 
	args := v.args // 마지막 투표 번호

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	startKey := "1"
	endKey := args[0]

	results, err := v.stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer results.Close()

	for results.HasNext() {
		
	}


	return shim.Success()
}

func (v *VotingChaincode) earlyComplete() pb.Response {
	args := v.args

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	votingAsBytes, _ := v.stub.GetState(args[0])
	voting := Voting{}

	json.Unmarshal(votingAsBytes, &voting)
	voting.CurrentState = 2

	votingAsBytes, _ = json.Marshal(voting)
	v.stub.PutState(args[0], votingAsBytes)

	return shim.Success(nil)
}

// =========================================== 밑으로는 그냥 함수

func (v *Voting) checkCandidateExist(cd string) bool { // 후보가 존재하는지 확인 
	_, exist := v.Candidate[cd]
	return exist
} // vote()에서 후보가 존재하는지 확인하기 위해 사용

func (v *Voting) checkID(id string) bool { // 투표를 이미 한 ID인지 체크
	b := true
	for _, i := range v.UserID {
		if i == id {
			b = false
			break
		}
	}
	return b
} // vote()에서 이미 투표한 아이디인지 확인하기 위해 사용

func changeToUnixTime(str string) int64 { // string으로 받은 시간을 Unix 시간으로 바꿔준다
	layout := "01/02/2006 3:04:05 PM"
	t, _ := time.Parse(layout, str)
	tUTC := t.Unix() - 32400  // 받은 시간은 KST, Unix() 시간은 UTC기준이므로 비교를 위해 UTC시간으로 변경
	return tUTC
} // createVote에서 startTime과 endTime을 유닉스 시간으로 바꾸어 줄 때 사용

// getCandidate gets candidate in Voting structure
func (v *Voting) getCandidateWithPoll() { // 후보 및 표 확인 post
	for key, val := range v.Candidate {
		fmt.Print(key, " ", val, " ")
	}
	fmt.Println()
} // 유저 투표 결과 확인 .html, 투표 현황 조회 .html에서 사용

func (v *Voting) getCandidate() { // post
	for key := range v.Candidate {
		fmt.Print(key, " ")
	}
	fmt.Println()
} // 투표 하는 페이지 .html에서 후보 이름을 확인하기 위해 사용

// deleteCandidate deletes candidate in Voting structure
func (v *Voting) deleteCandidate(cd string) { // cd는 후보
	if v.CurrentState == 0 { // 시작전에만 삭제 가능
		fmt.Println("후보를 삭제할 수 없습니다")
		return
	}
	delete(v.Candidate, cd)
} // 관리자 후보입력 .html에서 사용 





func getAllVoting() { // 모든 투표 목록(관리자) post
	for i := range votingSlice {
		// post 투표이름
		fmt.Println(votingSlice[i].VotingName)
	}
} // 관리자 투표 관리 .html에서 투표 목록을 불러올 때 사용

func notCompleteVote() { // 끝나지 않은 투표들 보내주기 post
	for i := range votingSlice {
		if votingSlice[i].CurrentState == 1 {
			// Post
			fmt.Println(votingSlice[i].VotingName)
		}
	}
} // 유저 투표 목록 .html에서 끝나지 않은 투표들을 불러올 때 사용

// changeState change Voting structure's CurrentState
func changeState() { // Voting 상태 변화 실시간으로 체크해서 투표의 상태를 변경한다(모든 투표를 대상으로 확인)
	for i := range votingSlice {
		if votingSlice[i].CurrentState == 2 { // 투표가 끝난 상태면 더 이상 상태를 바꾸지 않음
			continue
		}
		if votingSlice[i].StartTime < time.Now().Unix() && votingSlice[i].EndTime > time.Now().Unix() { // 투표 시작
			votingSlice[i].CurrentState = 1
		} else if votingSlice[i].EndTime < time.Now().Unix() { // 투표가 끝난 상태
			votingSlice[i].CurrentState = 2
		}
	}
} // 일정시간마다 동작해 시작 시간과 끝 시간에 따라 투표들의 상태를 변경

// viewCompleteVoting views completed Voting
func viewCompleteVoting() { // 전체 투표 목록 중 완료된 투표 조회 post
	for i := 0; i < len(votingSlice); i++ {
		if votingSlice[i].CurrentState == 2 { // 상태가 2인 투표들은 투표가 완료된 것들
			fmt.Println(votingSlice[i].VotingName)
		}
	}
} // 사용자 완료된 투표 목록 .html에서 완료된 투표 목록을 불러오기 위해 사용