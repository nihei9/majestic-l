package majestic

import (
	"fmt"
)

type MatchableValue interface {
	Match(target interface{}) bool
}

type IntValue struct {
	value int
}

func NewIntValue(value int) MatchableValue {
	return &IntValue{
		value: value,
	}
}

func (v *IntValue) Match(target interface{}) bool {
	t, ok := target.(int)
	if !ok {
		return false
	}
	return v.value == t
}

type Int8Value struct {
	value int8
}

func NewInt8Value(value int8) MatchableValue {
	return &Int8Value{
		value: value,
	}
}

func (v *Int8Value) Match(target interface{}) bool {
	t, ok := target.(int8)
	if !ok {
		return false
	}
	return v.value == t
}

type Int16Value struct {
	value int16
}

func NewInt16Value(value int16) MatchableValue {
	return &Int16Value{
		value: value,
	}
}

func (v *Int16Value) Match(target interface{}) bool {
	t, ok := target.(int16)
	if !ok {
		return false
	}
	return v.value == t
}

type Int32Value struct {
	value int32
}

func NewInt32Value(value int32) MatchableValue {
	return &Int32Value{
		value: value,
	}
}

func (v *Int32Value) Match(target interface{}) bool {
	t, ok := target.(int32)
	if !ok {
		return false
	}
	return v.value == t
}

type Int64Value struct {
	value int64
}

func NewInt64Value(value int64) MatchableValue {
	return &Int64Value{
		value: value,
	}
}

func (v *Int64Value) Match(target interface{}) bool {
	t, ok := target.(int64)
	if !ok {
		return false
	}
	return v.value == t
}

type UintValue struct {
	value uint
}

func NewUintValue(value uint) MatchableValue {
	return &UintValue{
		value: value,
	}
}

func (v *UintValue) Match(target interface{}) bool {
	t, ok := target.(uint)
	if !ok {
		return false
	}
	return v.value == t
}

type Uint8Value struct {
	value uint8
}

func NewUint8Value(value uint8) MatchableValue {
	return &Uint8Value{
		value: value,
	}
}

func (v *Uint8Value) Match(target interface{}) bool {
	t, ok := target.(uint8)
	if !ok {
		return false
	}
	return v.value == t
}

type Uint16Value struct {
	value uint16
}

func NewUint16Value(value uint16) MatchableValue {
	return &Uint16Value{
		value: value,
	}
}

func (v *Uint16Value) Match(target interface{}) bool {
	t, ok := target.(uint16)
	if !ok {
		return false
	}
	return v.value == t
}

type Uint32Value struct {
	value uint32
}

func NewUint32Value(value uint32) MatchableValue {
	return &Uint32Value{
		value: value,
	}
}

func (v *Uint32Value) Match(target interface{}) bool {
	t, ok := target.(uint32)
	if !ok {
		return false
	}
	return v.value == t
}

type Uint64Value struct {
	value uint64
}

func NewUint64Value(value uint64) MatchableValue {
	return &Uint64Value{
		value: value,
	}
}

func (v *Uint64Value) Match(target interface{}) bool {
	t, ok := target.(uint64)
	if !ok {
		return false
	}
	return v.value == t
}

type StringValue struct {
	value string
}

func NewStringValue(value string) MatchableValue {
	return &StringValue{
		value: value,
	}
}

func (v *StringValue) Match(target interface{}) bool {
	t, ok := target.(string)
	if !ok {
		return false
	}
	return v.value == t
}

func ConvertToMatchableValue(src interface{}) (MatchableValue, error) {
	var mv MatchableValue
	switch v := src.(type) {
	case int:
		mv = NewIntValue(v)
	case int8:
		mv = NewInt8Value(v)
	case int16:
		mv = NewInt16Value(v)
	case int32:
		mv = NewInt32Value(v)
	case int64:
		mv = NewInt64Value(v)
	case uint:
		mv = NewUintValue(v)
	case uint8:
		mv = NewUint8Value(v)
	case uint16:
		mv = NewUint16Value(v)
	case uint32:
		mv = NewUint32Value(v)
	case uint64:
		mv = NewUint64Value(v)
	case string:
		mv = NewStringValue(v)
	default:
		return nil, fmt.Errorf("Failed to convert to MatchableValue")
	}

	return mv, nil
}
