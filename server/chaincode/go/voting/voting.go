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

// 투표 생성 및 초기화(관리자 페이지)
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
// 투표 startTime, endTime을 체크해 투표 가능 상태를 변화
func (v *VotingChaincode) changeState() pb.Response {
	args := v.args // 마지막 투표 번호

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	
	// voteAsBytes, _ := v.stub.GetState(args)

	return shim.Success(nil)
}

// 투표 후보자 등록 (관리자 페이지)
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

// 투표 (사용자 페이지)
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
// 존재하는 모든 투표 불러오기 (관리자페이지)
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


	return shim.Success(nil)
}

// TODO: 구현
// 완료된 투표 불러오기 (유저페이지)
func (v *VotingChaincode) queryCompleteVote() pb.Response {
	return shim.Success(nil)
}

// endTime전에 투표 종료 (관리자페이지)
func (v *VotingChaincode) earlyComplete() pb.Response {
	args := v.args // 투표 번호

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

// 등록된 후보 삭제 (관리자 페이지)
func (v *VotingChaincode) deleteCandidate() pb.Response {
	args := v.args // 투표 번호, 후보 이름
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	votingAsBytes, _ := v.stub.GetState(args[0])
	voting := Voting{}

	json.Unmarshal(votingAsBytes, &voting)

	delete(voting.Candidate, args[1])

	votingAsBytes, _ = json.Marshal(voting)
	v.stub.PutState(args[0], votingAsBytes)

	return shim.Success(nil)
}

// TODO: 구현
// 완료되지 않은 투표 목록 불러오기 (사용자 페이지)
func (v *VotingChaincode) queryNotCompleteVote() pb.Response {
	return shim.Success(nil)
}

// TODO: 구현
// 후보와 표 불러오기 (사용자 페이지, 관리자 페이지)
func (v *VotingChaincode) queryCandidateWithPoll() pb.Response {
	return shim.Success(nil)
}

// TODO: 구현
// 후보 불러오기 (사용자 페이지)
func (v *VotingChaincode) queryCandidate() pb.Response {
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