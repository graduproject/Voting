package main

import (
	// "encoding/json"
	"fmt"
	"strconv"
	"encoding/json"
	"time"
	"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// User is ...
type User struct {
	ID				string `json="id"`
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

func (u *UserChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("User Init")

	/*
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}*/

	return shim.Success(nil)
}

func (u *UserChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	u.function = function
	u.args = args
	u.stub = stub

	return u.call()
}

func (u *UserChaincode) call() pb.Response {

	return shim.Error("")
}

// CreateUser creates User structure
func (u *UserChaincode) createUser() pb.Response { // 유저 구조체 생성(회원가입)
	args := u.args // ID, PW, IDNumber(주민번호), PhoneNumber, Email
	
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	user := User{PW: args[1], IDNumber: args[2], PhoneNumber: args[3], Email: args[4], IsAdmin: false}
	userAsBytes, _ := json.Marshal(user)
	u.stub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}


// ModifyUser modifies User data
func (u *UserChaincode) ModifyUser() pb.Response { // 등록된 유저의 정보 수정
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

func signUp() {
	
}

// getUserInfo gets a User data
func (u *User) getUserInfo() { // 유저 정보 조회
	
}

// LogIn is ...
func (u *User) LogIn() { // 로그인

}

// LogOut is ...
func (u *User) LogOut() { // 로그아웃

}

// DeleteUser deletes User data
func DeleteUser() { // 유저 데이터 삭제(회원탈퇴)
	
}

func main() {

}