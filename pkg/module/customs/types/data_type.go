package types

import (
	"fmt"

	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
)

const (
	INITIAL_ID       = 0
	INITIAL_REVISION = 1
	DEFAULT_NAME     = "anonymous"
)

type IDENTIFICATION int

func NewIDENTIFICATION(id int) (IDENTIFICATION, error) {
	if id < 0 {
		return IDENTIFICATION(0), fmt.Errorf(errors.ERR0000)
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
