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

func createAccount(t *testing.T, appCtx AppCtx) (AccountId, string, string) {
	username := "USER:" + uuid.New().String()
	password := "PASS:" + uuid.New().String()
	accountId, err := CreateAccount(appCtx, username, password)
	if err != nil {
		t.Error(err)
	}
	return accountId, username, password
}

func TestSocialList(t *testing.T) {
	appCtx := NewTestAppCtx()

	// Create some accounts
	accountIds := make([]AccountId, 10)
	usernames := make([]string, 10)
	passwords := make([]string, 10)

	for i := 0; i < len(accountIds); i++ {
		accountIds[i], usernames[i], passwords[i] = createAccount(t, appCtx)
		if i != 0 {
			slt := SocialListType_BLOCKED
			if i%2 == 0 {
				slt = SocialListType_FOLLOWS
			}
			err := AddToSocialList(appCtx, accountIds[0], slt, accountIds[i])
			if err != nil {
				t.Error(err)
			}
		}
	}

	blocked, err := GetSocialList(appCtx, accountIds[0], SocialListType_BLOCKED)
	if err != nil {
		t.Error(err)
	}
	assert.Contains(t, blocked, accountIds[1])
	assert.Contains(t, blocked, accountIds[3])
	assert.Contains(t, blocked, accountIds[5])
	assert.Contains(t, blocked, accountIds[7])
	assert.Contains(t, blocked, accountIds[9])

	following, err := GetSocialList(appCtx, accountIds[0], SocialListType_FOLLOWING)
	if err != nil {
		t.Error(err)
	}
	assert.Empty(t, following)

	follows, err := GetSocialList(appCtx, accountIds[0], SocialListType_FOLLOWS)
	if err != nil {
		t.Error(err)
	}
	assert.Contains(t, follows, accountIds[2])
	assert.Contains(t, follows, accountIds[4])
	assert.Contains(t, follows, accountIds[6])
	assert.Contains(t, follows, accountIds[8])

	err = AddToSocialList(appCtx, accountIds[2], SocialListType_FOLLOWS, accountIds[0])
	if err != nil {
		t.Error(err)
	}

	following, err = GetSocialList(appCtx, accountIds[0], SocialListType_FOLLOWING)
	if err != nil {
		t.Error(err)
	}
	assert.Contains(t, following, accountIds[2])

}
