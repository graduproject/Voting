package main

import (
	"encoding/json"
)

// User is ...
type User struct {
	ID				string `json="id"`
	Pw  			string `json="pw"`
	PhoneNumber 	string `json="phonenumber"`
	Email			string `json="email"`
	IsAdmin     	bool   `json="isadmin"`
	IsLogIn			bool   `json="islogin"`
}

var UserSlice []User // 유저 목록

// CreateUser creates User structure
func CreateUser() { // 유저 구조체 생성(회원가입)

}

// RegisterUser registers User data
func (u *User) RegisterUser() { // 유저 데이터 등록

}

// ModifyUser modifies User data
func (u *User) ModifyUser() { // 등록된 유저의 정보 수정

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