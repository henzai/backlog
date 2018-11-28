package backlog

import (
	"fmt"
	"strconv"
)

type Users []User

type User struct {
	ID          int64  `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	RoleType    int    `json:"roleType"`
	Lang        string `json:"lang"`
	MailAddress string `json:"mailAddress"`
}

func (b *BackLog) GetUsers() (*Users, error) {
	backlogURL := fmt.Sprintf(b.baseURL, "users")
	var m Users
	err := Get(backlogURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (b *BackLog) GetUserID(id int) (*User, error) {
	endpoint := "users/" + strconv.Itoa(id)
	backlogURL := fmt.Sprintf(b.baseURL, endpoint)
	var m User
	err := Get(backlogURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (b *BackLog) GetMyUser() (*User, error) {
	backlogURL := fmt.Sprintf(b.baseURL, "users/myself")
	var m User
	err := Get(backlogURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
