package types

import (
	"fmt"
)

const (
	INITIAL_ID       = 0
	INITIAL_REVISION = 1
	DEFAULT_NAME     = "anonymous"
)

type IDENTIFICATION int

func NewIDENTIFICATION(id int) (IDENTIFICATION, error) {
	if id < 0 {
		return IDENTIFICATION(id), fmt.Errorf("ID is less than or equal to zero")
	}
	return IDENTIFICATION(id), nil
}

type Roll int

const (
	AdminRoll Roll = iota + 1
	MembersRoll
)

func NewRoll(i int) (Roll, error) {
	t := Roll(i)
	switch t {
	case AdminRoll:
		return AdminRoll, nil
	case MembersRoll:
		return MembersRoll, nil
	}

	return t, fmt.Errorf("invalid user roll")
}
