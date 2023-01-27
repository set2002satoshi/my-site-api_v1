package errors

import "fmt"

const (
	ERR0000 = "ERROR0000"

	TYPE0001 = "TYPE0001"

	DB0001 = "DB0001"
	DB0002 = "DB0002"

	UNDEFINED = "UNDEFINED"
)

var ErrMap = map[string]string{
	ERR0000: "error code undetermined",

	TYPE0001: "ID is less than or equal to zero",

	DB0001: "failure db connect",
	DB0002: "failure migration",

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
		getMap() map[string]string
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
	c.addSet(s, s)
	return c
}

// 第一引数には現状のエラーを渡す。第二引数には追加するエラーを渡す
// そうすると命名通りエラーを組み合わせる事ができる。
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

// errorのキーをもらってエラー内容を埋め込んでる。
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
		for k := range cErr.getMap() {
			c.add(k)
		}
		return
	}
	if err != nil {
		c.addSet(UNDEFINED, err.Error())
	}
}

func (c *customError) getMap() map[string]string {
	return c.ErrorMap
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
