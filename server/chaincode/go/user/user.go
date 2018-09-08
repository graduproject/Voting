package main

import (
	// "encoding/json"
)

// User is ...
type User struct {
	ID				string `json="id"`
	PW  			string `json="pw"`
	IDNumber        string `json="idnumber"`
	PhoneNumber 	string `json="phonenumber"`
	Email			string `json="email"`
	IsAdmin     	bool   `json="isadmin"`
	IsLogIn			bool   `json="islogin"`
}

var UserSlice []User // 유저 목록
var withdrawalSlice []string // 탈퇴한 회원의 주민등록번호 모음

// CreateUser creates User structure
func CreateUser(id string, pw string, idNumber string, phone string, mail string) { // 유저 구조체 생성(회원가입)
	u := User{ID: id, PW: pw, PhoneNumber: phone, IDNumber: idNumber, Email: mail, IsAdmin: false, IsLogIn: false}
	UserSlice = append(UserSlice, u)
}

// ModifyUser modifies User data
func (u *User) ModifyUser(pw string, phone string, mail string) { // 등록된 유저의 정보 수정
	u.PW = pw
	u.PhoneNumber = phone
	u.Email = mail
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