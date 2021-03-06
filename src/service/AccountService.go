package service

import (
	"errors"
	"strconv"
	"strings"

	"../config"
	"../domain"
)

// FindAccount : find accounts
func FindAccount(uid int32) (*domain.Account, error) {

	if uid <= 0 {
		return nil, errors.New("user not found")
	}

	db := config.GetDBConn()

	rows, err := db.Query(
		strings.Join([]string{"select ac.agent_id, ac.username, ac.nickname, ac.password from agent ag join account ac on ag.id = ac.agent_id",
			"where ag.id = " + strconv.Itoa(int(uid)),
			"and ac.is_del = 0 and ag.is_del = 0 and ac.ban = 0 and ag.ban = 0"}, " "))
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	// var account Account
	var (
		agentID  int32
		username string
		password string
		nickname string
		account  *domain.Account
	)

	rows.Next()
	if err = rows.Scan(&agentID, &username, &nickname, &password); err != nil {
		return nil, err
	}

	account = &domain.Account{
		AgentID:  agentID,
		Username: username,
		Nickname: nickname,
		Password: password,
	}
	return account, nil
}
