package main

import (
	"fmt"

	"github.com/zoklk/Go-prac/phonebook/phone"
)

func main() {
	pb := phone.NewPhoneBook()

	// 1. 추가
	pb.Add("zoklk1", "010-1111-1234")
	pb.Add("zoklk2", "010-2222-1234")
	pb.Add("zoklk3", "010-3333-1234")
	fmt.Println(pb)

	// 2. 검색 (성공)
	phoneNum, err := pb.Search("zoklk1")
	if err != nil {
		fmt.Printf("에러: %s\n", err)
	} else {
		fmt.Printf("zoklk1의 번호: %s\n", phoneNum)
	}

	// 2. 검색 (실패)
	phoneNum, err = pb.Search("zoklk4")
	if err != nil {
		fmt.Printf("에러: %s\n", err)
	} else {
		fmt.Printf("zoklk4의 번호: %s\n", phoneNum)
	}

	// 4. 삭제
	fmt.Println("\nzoklk2의 연락처를 삭제합니다...")
	pb.Delete("zoklk2")
	fmt.Println(pb)
}
