package services

type AccountId string

func NilAccountId() AccountId {
	return NewAccountId("")
}

func (a AccountId) String() string {
	return string(a)
}

func NewAccountId(id string) AccountId {
	return AccountId(id)
}
