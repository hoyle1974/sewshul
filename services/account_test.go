package services

import (
	"net"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountAndLogin(t *testing.T) {
	appCtx := NewTestAppCtx()

	username := uuid.New().String()
	password := uuid.New().String()

	accountId1, err := CreateAccount(appCtx, username, password)
	if err != nil {
		t.Error(err)
	}

	loggedInAccountId1, err := Login(appCtx, username, password, net.ParseIP("192.168.1.1"), 1234)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, accountId1, loggedInAccountId1)
}
