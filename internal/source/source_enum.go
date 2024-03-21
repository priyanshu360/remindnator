// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package source

import (
	"errors"
	"fmt"
)

const (
	// SourceEnumGoogleTask is a SourceEnum of type GoogleTask.
	SourceEnumGoogleTask SourceEnum = iota
	// SourceEnumGoogleCalendar is a SourceEnum of type GoogleCalendar.
	SourceEnumGoogleCalendar
	// SourceEnumSlackChannel is a SourceEnum of type SlackChannel.
	SourceEnumSlackChannel
)

var ErrInvalidSourceEnum = errors.New("not a valid SourceEnum")

const _SourceEnumName = "googleTaskgoogleCalendarSlackChannel"

var _SourceEnumMap = map[SourceEnum]string{
	SourceEnumGoogleTask:     _SourceEnumName[0:10],
	SourceEnumGoogleCalendar: _SourceEnumName[10:24],
	SourceEnumSlackChannel:   _SourceEnumName[24:36],
}

// String implements the Stringer interface.
func (x SourceEnum) String() string {
	if str, ok := _SourceEnumMap[x]; ok {
		return str
	}
	return fmt.Sprintf("SourceEnum(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x SourceEnum) IsValid() bool {
	_, ok := _SourceEnumMap[x]
	return ok
}

var _SourceEnumValue = map[string]SourceEnum{
	_SourceEnumName[0:10]:  SourceEnumGoogleTask,
	_SourceEnumName[10:24]: SourceEnumGoogleCalendar,
	_SourceEnumName[24:36]: SourceEnumSlackChannel,
}

// ParseSourceEnum attempts to convert a string to a SourceEnum.
func ParseSourceEnum(name string) (SourceEnum, error) {
	if x, ok := _SourceEnumValue[name]; ok {
		return x, nil
	}
	return SourceEnum(0), fmt.Errorf("%s is %w", name, ErrInvalidSourceEnum)
}

// MarshalText implements the text marshaller method.
func (x SourceEnum) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *SourceEnum) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseSourceEnum(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}