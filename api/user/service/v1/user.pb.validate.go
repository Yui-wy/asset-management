// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/user/service/v1/user.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on GetUserReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetUserReqValidationError is the validation error returned by
// GetUserReq.Validate if the designated constraints aren't met.
type GetUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReqValidationError) ErrorName() string { return "GetUserReqValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReqValidationError{}

// Validate checks the field values on GetUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUserReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Username

	return nil
}

// GetUserReplyValidationError is the validation error returned by
// GetUserReply.Validate if the designated constraints aren't met.
type GetUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReplyValidationError) ErrorName() string { return "GetUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReplyValidationError{}

// Validate checks the field values on CreateUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Username

	// no validation rules for Password

	return nil
}

// CreateUserReqValidationError is the validation error returned by
// CreateUserReq.Validate if the designated constraints aren't met.
type CreateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReqValidationError) ErrorName() string { return "CreateUserReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReqValidationError{}

// Validate checks the field values on CreateUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateUserReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Username

	return nil
}

// CreateUserReplyValidationError is the validation error returned by
// CreateUserReply.Validate if the designated constraints aren't met.
type CreateUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReplyValidationError) ErrorName() string { return "CreateUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReplyValidationError{}

// Validate checks the field values on ListUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PageNum

	// no validation rules for PageSize

	return nil
}

// ListUserReqValidationError is the validation error returned by
// ListUserReq.Validate if the designated constraints aren't met.
type ListUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReqValidationError) ErrorName() string { return "ListUserReqValidationError" }

// Error satisfies the builtin error interface
func (e ListUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReqValidationError{}

// Validate checks the field values on ListUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListUserReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListUserReplyValidationError is the validation error returned by
// ListUserReply.Validate if the designated constraints aren't met.
type ListUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReplyValidationError) ErrorName() string { return "ListUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReplyValidationError{}

// Validate checks the field values on DeleteUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeleteUserReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteUserReqValidationError is the validation error returned by
// DeleteUserReq.Validate if the designated constraints aren't met.
type DeleteUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserReqValidationError) ErrorName() string { return "DeleteUserReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserReqValidationError{}

// Validate checks the field values on DeleteUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteUserReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Ok

	return nil
}

// DeleteUserReplyValidationError is the validation error returned by
// DeleteUserReply.Validate if the designated constraints aren't met.
type DeleteUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserReplyValidationError) ErrorName() string { return "DeleteUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserReplyValidationError{}

// Validate checks the field values on VerifyPasswordReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *VerifyPasswordReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Username

	// no validation rules for Password

	return nil
}

// VerifyPasswordReqValidationError is the validation error returned by
// VerifyPasswordReq.Validate if the designated constraints aren't met.
type VerifyPasswordReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyPasswordReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyPasswordReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyPasswordReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyPasswordReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyPasswordReqValidationError) ErrorName() string {
	return "VerifyPasswordReqValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyPasswordReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyPasswordReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyPasswordReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyPasswordReqValidationError{}

// Validate checks the field values on VerifyPasswordReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *VerifyPasswordReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Ok

	return nil
}

// VerifyPasswordReplyValidationError is the validation error returned by
// VerifyPasswordReply.Validate if the designated constraints aren't met.
type VerifyPasswordReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyPasswordReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyPasswordReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyPasswordReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyPasswordReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyPasswordReplyValidationError) ErrorName() string {
	return "VerifyPasswordReplyValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyPasswordReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyPasswordReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyPasswordReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyPasswordReplyValidationError{}

// Validate checks the field values on UpdatePasswordReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdatePasswordReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Password

	return nil
}

// UpdatePasswordReqValidationError is the validation error returned by
// UpdatePasswordReq.Validate if the designated constraints aren't met.
type UpdatePasswordReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePasswordReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePasswordReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePasswordReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePasswordReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePasswordReqValidationError) ErrorName() string {
	return "UpdatePasswordReqValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePasswordReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePasswordReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePasswordReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePasswordReqValidationError{}

// Validate checks the field values on UpdatePasswordReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdatePasswordReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Username

	return nil
}

// UpdatePasswordReplyValidationError is the validation error returned by
// UpdatePasswordReply.Validate if the designated constraints aren't met.
type UpdatePasswordReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePasswordReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePasswordReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePasswordReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePasswordReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePasswordReplyValidationError) ErrorName() string {
	return "UpdatePasswordReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePasswordReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePasswordReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePasswordReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePasswordReplyValidationError{}

// Validate checks the field values on ListUserReply_User with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListUserReply_User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Username

	return nil
}

// ListUserReply_UserValidationError is the validation error returned by
// ListUserReply_User.Validate if the designated constraints aren't met.
type ListUserReply_UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReply_UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReply_UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReply_UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReply_UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReply_UserValidationError) ErrorName() string {
	return "ListUserReply_UserValidationError"
}

// Error satisfies the builtin error interface
func (e ListUserReply_UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReply_User.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReply_UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReply_UserValidationError{}
