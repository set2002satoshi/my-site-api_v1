package errors

import "fmt"

const (
	ERR0000 = "ERROR0000"
	ERR0001 = "ERROR0001"
	ERR0002 = "ERROR0002"
	ERR0003 = "ERROR0003"

	REPO0001 = "REPO0001"
	REPO0002 = "REPO0002"
	REPO0003 = "REPO0003"
	REPO0004 = "REPO0004"
	REPO0005 = "REPO0005"
	REPO0006 = "REPO0006"
	REPO0007 = "REPO0007"
	REPO0008 = "REPO0008"
	REPO0009 = "REPO0009"
	REPO0010 = "REPO0010"
	REPO0011 = "REPO0011"
	REPO0012 = "REPO0012"
	REPO0013 = "REPO0013"
	REPO0014 = "REPO0014"
	REPO0015 = "REPO0015"

	TYPE0001 = "TYPE0001"

	DB0001 = "DB0001"
	DB0002 = "DB0002"

	EN0001 = "EN0001"
	EN0002 = "EN0002"
	EN0003 = "EN0003"
	EN0004 = "EN0004"
	EN0005 = "EN0005"

	SE0000 = "SE0000"
	SE0001 = "SE0001"
	SE0002 = "SE0002"
	SE0003 = "SE0003"

	UNDEFINED = "UNDEFINED"
)

var ErrMap = map[string]string{
	ERR0000: "error code undetermined",
	ERR0001: "failure BindJSON",
	ERR0002: "failure to Model",
	ERR0003: "couldn't get userID",

	REPO0001: "create user failed",
	REPO0002: "created user but user found",
	REPO0003: "that email address not unique",
	REPO0004: "failed get user",
	REPO0005: "record could not be deleted",
	REPO0006: "create blog failed",
	REPO0007: "created blog but blog found",
	REPO0008: "couldn't update user",
	REPO0009: "couldn't update user on blog",
	REPO0010: "failed get blogs",
	REPO0011: "failed get blog",
	REPO0012: "create category failed",
	REPO0013: "failed get category",
	REPO0014: "failed get categories",
	REPO0015: "create blog with category failed",

	TYPE0001: "id is less than or equal to zero",

	DB0001: "failure db connect",
	DB0002: "failure migration",

	EN0001: "this id is invalid id",
	EN0002: "password couldn't not be hashed",
	EN0003: "this roll is invalid roll",
	EN0004: "did not match revision number",
	EN0005: "couldn't up revision number",

	SE0001: "couldn't load secret",
	SE0002: "invalid password",
	SE0003: "couldn't issue token",

	UNDEFINED: "undefined",
}

type (
	customError struct {
		ErrorMap map[string]string
	}
	iCustomError interface {
		error
		add(code string)
		addSet(code, message string)
		combine(err error)
		isContain(code string) bool
		isEmpty() bool
		getMap() map[string]string
		wrap(code, message string)
	}
)

var _ iCustomError = &customError{}

func NewCustomError(codes ...string) error {
	e := &customError{ErrorMap: map[string]string{}}
	for _, s := range codes {
		e.add(s)
	}
	return e
}

func New(s string) error {
	c := &customError{ErrorMap: map[string]string{}}
	c.addSet(UNDEFINED, s)
	return c
}

func Combine(origin, new error) error {
	if origin == nil && new == nil {
		return nil
	}
	if oErr, ok := origin.(iCustomError); ok {
		if new != nil {
			oErr.combine(new)
		}
		return oErr
	}
	if nErr, ok := new.(iCustomError); ok {
		if origin != nil {
			nErr.addSet(UNDEFINED, origin.Error())
		}
		return nErr
	}
	cErr := &customError{ErrorMap: map[string]string{}}
	cErr.combine(origin)
	cErr.combine(new)
	return cErr
}

func Wrap(err error, code, message string) error {
	if cErr, ok := err.(iCustomError); ok {
		cErr.wrap(code, message)
		return cErr
	}
	cErr := &customError{ErrorMap: map[string]string{}}
	cErr.wrap(code, message)
	return cErr
}

func Add(err error, code string) error {
	if cErr, ok := err.(iCustomError); ok {
		cErr.add(code)
		return cErr
	}
	cErr := &customError{ErrorMap: map[string]string{}}
	if err != nil {
		cErr.combine(err)
	}
	cErr.add(code)
	return cErr
}

func IsContain(err error, code string) bool {
	if cErr, ok := err.(iCustomError); ok {
		return cErr.isContain(code)
	}
	return false
}

func IsEmpty(err error) bool {
	if cErr, ok := err.(iCustomError); ok {
		return cErr.isEmpty()
	}
	return err == nil
}

func ToMap(err error) map[string]string {
	if cErr, ok := err.(iCustomError); ok {
		return cErr.getMap()
	}
	c := &customError{ErrorMap: map[string]string{}}
	if err != nil {
		c.addSet(UNDEFINED, err.Error())
	}
	return c.getMap()
}

func (c *customError) add(code string) {
	if _, ok := ErrMap[code]; ok {
		c.ErrorMap[code] = getMessage(code)
	}
}

func (c *customError) addSet(code, message string) {
	if _, ok := c.ErrorMap[code]; ok {
		c.ErrorMap[code] = fmt.Sprintf("%s, %s", c.ErrorMap[code], message)
		return
	}
	c.ErrorMap[code] = message
}
func (c *customError) combine(err error) {
	if cErr, ok := err.(iCustomError); ok {
		for k, v := range cErr.getMap() {
			c.addSet(k, v)
		}
		return
	}
	if err != nil {
		c.addSet(UNDEFINED, err.Error())
	}
}
func (c *customError) isContain(code string) bool {
	_, ok := c.ErrorMap[code]
	return ok
}

func (c *customError) isEmpty() bool {
	fmt.Println(len(c.ErrorMap))
	return len(c.ErrorMap) == 0
}

func (c *customError) getMap() map[string]string {
	return c.ErrorMap
}

func (c *customError) wrap(code, message string) {
	if _, ok := c.ErrorMap[code]; ok {
		c.ErrorMap[code] = fmt.Sprintf("%s=>%s", c.ErrorMap[code], message)
		return
	}
	if _, ok := ErrMap[code]; ok {
		c.ErrorMap[code] = fmt.Sprintf("%s=>%s", ErrMap[code], message)
		return
	}
	c.addSet(UNDEFINED, message)
}

func getMessage(code string) string {
	if s, ok := ErrMap[code]; ok {
		return s
	}
	return ""
}

func (c customError) Error() string {
	var msg string
	initFlag := true
	for _, v := range c.ErrorMap {
		if initFlag {
			msg = v
			initFlag = false
			continue
		}
		msg = fmt.Sprintf("%s, %s", msg, v)
	}

	return msg
}
