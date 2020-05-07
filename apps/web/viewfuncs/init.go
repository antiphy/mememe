package viewfuncs

import (
	"errors"
	"math/rand"
	"time"

	"github.com/flosch/pongo2"
)

func init() {
	// Account
	pongo2.RegisterFilter("AccountStatusDisplay", accountStatusDisplay)
	pongo2.RegisterFilter("AccountGroupDisplay", accountGroupDisplay)
	// Datetime
	pongo2.RegisterFilter("DateTimeDisplay", dateTimeDisplay)
	pongo2.RegisterFilter("TimestampDisplay", timestampDisplay)
	// Math
	pongo2.RegisterFilter("RandInt", randInt)
}

func accountStatusDisplay(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	i, ok := in.Interface().(int8)
	if !ok {
		return nil, &pongo2.Error{
			Sender:    "filter.accountStatusDisplay",
			OrigError: errors.New("in not int8"),
		}
	}
	switch i {
	case 10:
		return pongo2.AsValue("活跃"), nil
	case 1:
		return pongo2.AsValue("正常"), nil
	case 0:
		return pongo2.AsValue("未激活"), nil
	case -1:
		return pongo2.AsValue("非活跃"), nil
	case -10:
		return pongo2.AsValue("禁言"), nil
	default:
		return pongo2.AsValue("未知"), nil
	}
}

func accountGroupDisplay(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	i, ok := in.Interface().(int8)
	if !ok {
		return nil, &pongo2.Error{
			Sender:    "filter.accountGroupDisplay",
			OrigError: errors.New("in not int8"),
		}
	}
	switch i {
	case 127:
		return pongo2.AsValue("管理员"), nil
	case 1:
		return pongo2.AsValue("普通"), nil
	case -1:
		return pongo2.AsValue("消极"), nil
	case -126:
		return pongo2.AsValue("禁言"), nil
	default:
		return pongo2.AsValue("未知"), nil
	}
}

func dateTimeDisplay(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	t, ok := in.Interface().(time.Time)
	if !ok {
		return nil, &pongo2.Error{
			Sender:    "filter.dateTimeDisplay",
			OrigError: errors.New("in not time"),
		}
	}
	now := time.Now()
	if t.Year() == now.Year() {
		if t.Month() == now.Month() && t.Day() == now.Day() {
			return pongo2.AsValue(t.Format("今天 15:04:05")), nil
		}
		return pongo2.AsValue(t.Format("01-02 15:04:05")), nil
	}
	return pongo2.AsValue(t.Format("06-01-02 15:04:05")), nil
}

func timestampDisplay(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	ts, ok := in.Interface().(int64)
	if !ok {
		return nil, &pongo2.Error{
			Sender:    "filter.timestampDisplay",
			OrigError: errors.New("in not int64"),
		}
	}
	now := time.Now()
	t := time.Unix(ts, 0)
	if t.Year() == now.Year() {
		if t.Month() == now.Month() && t.Day() == now.Day() {
			return pongo2.AsValue(t.Format("今天 15:04:05")), nil
		}
		return pongo2.AsValue(t.Format("01-02 15:04:05")), nil
	}
	return pongo2.AsValue(t.Format("06-01-02 15:04:05")), nil
}

func randInt(in, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	max, ok := in.Interface().(int)
	if !ok {
		return nil, &pongo2.Error{
			Sender:    "filter.randInt",
			OrigError: errors.New("in not int"),
		}
	}
	return pongo2.AsValue(rand.Intn(max) + 1), nil
}
