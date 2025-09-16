package phone

import (
	"fmt"
)

type PhoneBook struct {
	People map[string]Contact
}

func NewPhoneBook() *PhoneBook {
	pb := &PhoneBook{}

	pb.People = make(map[string]Contact)

	return pb
}

func (pb *PhoneBook) Add(name, phone string) {
	contact := Contact{
		Name:  name,
		Phone: phone,
	}

	pb.People[name] = contact
	fmt.Printf(">> 추가 완료: %s (%s)\n", name, phone)
}

func (pb *PhoneBook) Search(name string) (string, error) {
	contact, okphone := pb.People[name]
	if okphone {
		return contact.Phone, nil
	} else {
		return "", fmt.Errorf("%s의 번호를 찾을 수 없습니다\n", name)
	}
}

func (pb *PhoneBook) Delete(name string) {
	delete(pb.People, name)
}
