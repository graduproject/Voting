package main

import (
	"strings"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// User is ...
type User struct {
	PW  			string `json="pw"`
	IDNumber        string `json="idnumber"`
	PhoneNumber 	string `json="phonenumber"`
	Email			string `json="email"`
	IsAdmin     	bool   `json="isadmin"`
}

type UserChaincode struct {
	stub     shim.ChaincodeStubInterface
	function string
	args     []string
}

var UserSlice []User // 유저 목록
var withdrawalSlice []string // 탈퇴한 회원의 주민등록번호 모음

func (u *UserChaincode) call() pb.Response {
	function := u.function

	callMap := map[string]func() pb.Response {
		"signup":               u.signup,
		"signin":               u.signin,
		"modifyUser":           u.modifyUser,
		"getUserInfo":          u.getUserInfo,
		"deleteUser":           u.deleteUser,
	}

	h := callMap[function]
	if h != nil {
		return callMap[function]()
	}

	res := make([]string, 0)
	for k := range callMap {
		res = append(res, `"`+k+`"`)
	}

	return shim.Error("Invalid invoke function name. Expecting " + strings.Join(res, ", "))
}

// Init ...
func (u *UserChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("User Init")

	/*
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}*/

	return shim.Success(nil)
}

// Invoke ...
func (u *UserChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	u.function = function
	u.args = args
	u.stub = stub

	return u.call()
}

// signup creates User structure
func (u *UserChaincode) signup() pb.Response { // 유저 구조체 생성(회원가입)
	args := u.args // ID, PW, IDNumber(주민번호), PhoneNumber, Email
	
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	user := User{PW: args[1], IDNumber: args[2], PhoneNumber: args[3], Email: args[4], IsAdmin: false}
	userAsBytes, _ := json.Marshal(user)
	u.stub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

// signin is ...
func (u *UserChaincode) signin() pb.Response { // 로그인
	args := u.args // ID, PW
	
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	id := args[0]
	pw := args[1]
	
	user := User{}
	userAsBytes, err := u.stub.GetState(id)
	if err != nil {
		return shim.Error(id + " is not registered.")
	}
	json.Unmarshal(userAsBytes, &user)

	if user.PW != pw {
		return shim.Error("Incorrect password.")
	}

	return shim.Success(nil)
}

// ModifyUser modifies User data
func (u *UserChaincode) modifyUser() pb.Response { // 등록된 유저의 정보 수정
	args := u.args // ID, PW, PhoneNumber, Email
	
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	
	user := User{}
	userAsBytes, _ := u.stub.GetState(args[0])
	json.Unmarshal(userAsBytes, &user)

	user.PW = args[1]
	user.PhoneNumber = args[2]
	user.Email = args[3]
	
	userAsBytes, _ = json.Marshal(user)
	u.stub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

// getUserInfo gets a User data
func (u *UserChaincode) getUserInfo() pb.Response { // 유저 정보 조회 
	args := u.args // ID

	if len(args) != 1 { 
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[0]
	userAsBytes, _ := u.stub.GetState(id)

	return shim.Success(userAsBytes)
}

// DeleteUser deletes User data
func (u *UserChaincode) deleteUser() pb.Response { // 유저 데이터 삭제(회원탈퇴)
	args := u.args // ID

	if len(args) != 1 { 
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[0]
	err := u.stub.DelState(id)
	if err != nil {
		return shim.Error(id + " is not registered.")
	}

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(UserChaincode))
	if err != nil {
		fmt.Printf("Error starting User chaincode: %s", err)
	}
}