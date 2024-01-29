// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user_center_scene.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on Scene with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Scene) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Scene with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SceneMultiError, or nil if none found.
func (m *Scene) ValidateAll() error {
	return m.validate(true)
}

func (m *Scene) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Keyword

	// no validation rules for Name

	// no validation rules for Description

	for idx, item := range m.GetAgreements() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SceneValidationError{
						field:  fmt.Sprintf("Agreements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SceneValidationError{
						field:  fmt.Sprintf("Agreements[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SceneValidationError{
					field:  fmt.Sprintf("Agreements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return SceneMultiError(errors)
	}

	return nil
}

// SceneMultiError is an error wrapping multiple validation errors returned by
// Scene.ValidateAll() if the designated constraints aren't met.
type SceneMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SceneMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SceneMultiError) AllErrors() []error { return m }

// SceneValidationError is the validation error returned by Scene.Validate if
// the designated constraints aren't met.
type SceneValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SceneValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SceneValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SceneValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SceneValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SceneValidationError) ErrorName() string { return "SceneValidationError" }

// Error satisfies the builtin error interface
func (e SceneValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScene.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SceneValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SceneValidationError{}

// Validate checks the field values on GetSceneByKeywordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSceneByKeywordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSceneByKeywordRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSceneByKeywordRequestMultiError, or nil if none found.
func (m *GetSceneByKeywordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSceneByKeywordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKeyword()) < 1 {
		err := GetSceneByKeywordRequestValidationError{
			field:  "Keyword",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetSceneByKeywordRequestMultiError(errors)
	}

	return nil
}

// GetSceneByKeywordRequestMultiError is an error wrapping multiple validation
// errors returned by GetSceneByKeywordRequest.ValidateAll() if the designated
// constraints aren't met.
type GetSceneByKeywordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSceneByKeywordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSceneByKeywordRequestMultiError) AllErrors() []error { return m }

// GetSceneByKeywordRequestValidationError is the validation error returned by
// GetSceneByKeywordRequest.Validate if the designated constraints aren't met.
type GetSceneByKeywordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSceneByKeywordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSceneByKeywordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSceneByKeywordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSceneByKeywordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSceneByKeywordRequestValidationError) ErrorName() string {
	return "GetSceneByKeywordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSceneByKeywordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSceneByKeywordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSceneByKeywordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSceneByKeywordRequestValidationError{}

// Validate checks the field values on PageSceneRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PageSceneRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageSceneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageSceneRequestMultiError, or nil if none found.
func (m *PageSceneRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageSceneRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageSceneRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageSceneRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 50]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if len(errors) > 0 {
		return PageSceneRequestMultiError(errors)
	}

	return nil
}

// PageSceneRequestMultiError is an error wrapping multiple validation errors
// returned by PageSceneRequest.ValidateAll() if the designated constraints
// aren't met.
type PageSceneRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageSceneRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageSceneRequestMultiError) AllErrors() []error { return m }

// PageSceneRequestValidationError is the validation error returned by
// PageSceneRequest.Validate if the designated constraints aren't met.
type PageSceneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageSceneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageSceneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageSceneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageSceneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageSceneRequestValidationError) ErrorName() string { return "PageSceneRequestValidationError" }

// Error satisfies the builtin error interface
func (e PageSceneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageSceneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageSceneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageSceneRequestValidationError{}

// Validate checks the field values on PageSceneReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageSceneReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageSceneReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageSceneReplyMultiError,
// or nil if none found.
func (m *PageSceneReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageSceneReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PageSceneReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageSceneReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageSceneReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageSceneReplyMultiError(errors)
	}

	return nil
}

// PageSceneReplyMultiError is an error wrapping multiple validation errors
// returned by PageSceneReply.ValidateAll() if the designated constraints
// aren't met.
type PageSceneReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageSceneReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageSceneReplyMultiError) AllErrors() []error { return m }

// PageSceneReplyValidationError is the validation error returned by
// PageSceneReply.Validate if the designated constraints aren't met.
type PageSceneReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageSceneReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageSceneReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageSceneReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageSceneReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageSceneReplyValidationError) ErrorName() string { return "PageSceneReplyValidationError" }

// Error satisfies the builtin error interface
func (e PageSceneReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageSceneReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageSceneReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageSceneReplyValidationError{}

// Validate checks the field values on AddSceneRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddSceneRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddSceneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddSceneRequestMultiError, or nil if none found.
func (m *AddSceneRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddSceneRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKeyword()) < 1 {
		err := AddSceneRequestValidationError{
			field:  "Keyword",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := AddSceneRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := AddSceneRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddSceneRequestMultiError(errors)
	}

	return nil
}

// AddSceneRequestMultiError is an error wrapping multiple validation errors
// returned by AddSceneRequest.ValidateAll() if the designated constraints
// aren't met.
type AddSceneRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddSceneRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddSceneRequestMultiError) AllErrors() []error { return m }

// AddSceneRequestValidationError is the validation error returned by
// AddSceneRequest.Validate if the designated constraints aren't met.
type AddSceneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddSceneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddSceneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddSceneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddSceneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddSceneRequestValidationError) ErrorName() string { return "AddSceneRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddSceneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddSceneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddSceneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddSceneRequestValidationError{}

// Validate checks the field values on AddSceneReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddSceneReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddSceneReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddSceneReplyMultiError, or
// nil if none found.
func (m *AddSceneReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AddSceneReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddSceneReplyMultiError(errors)
	}

	return nil
}

// AddSceneReplyMultiError is an error wrapping multiple validation errors
// returned by AddSceneReply.ValidateAll() if the designated constraints
// aren't met.
type AddSceneReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddSceneReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddSceneReplyMultiError) AllErrors() []error { return m }

// AddSceneReplyValidationError is the validation error returned by
// AddSceneReply.Validate if the designated constraints aren't met.
type AddSceneReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddSceneReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddSceneReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddSceneReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddSceneReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddSceneReplyValidationError) ErrorName() string { return "AddSceneReplyValidationError" }

// Error satisfies the builtin error interface
func (e AddSceneReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddSceneReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddSceneReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddSceneReplyValidationError{}

// Validate checks the field values on UpdateSceneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSceneRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSceneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSceneRequestMultiError, or nil if none found.
func (m *UpdateSceneRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSceneRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateSceneRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetKeyword()) < 1 {
		err := UpdateSceneRequestValidationError{
			field:  "Keyword",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UpdateSceneRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := UpdateSceneRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateSceneRequestMultiError(errors)
	}

	return nil
}

// UpdateSceneRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateSceneRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateSceneRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSceneRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSceneRequestMultiError) AllErrors() []error { return m }

// UpdateSceneRequestValidationError is the validation error returned by
// UpdateSceneRequest.Validate if the designated constraints aren't met.
type UpdateSceneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSceneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSceneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSceneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSceneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSceneRequestValidationError) ErrorName() string {
	return "UpdateSceneRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSceneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSceneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSceneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSceneRequestValidationError{}

// Validate checks the field values on DeleteSceneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSceneRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSceneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSceneRequestMultiError, or nil if none found.
func (m *DeleteSceneRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSceneRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteSceneRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteSceneRequestMultiError(errors)
	}

	return nil
}

// DeleteSceneRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteSceneRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteSceneRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSceneRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSceneRequestMultiError) AllErrors() []error { return m }

// DeleteSceneRequestValidationError is the validation error returned by
// DeleteSceneRequest.Validate if the designated constraints aren't met.
type DeleteSceneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSceneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSceneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSceneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSceneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSceneRequestValidationError) ErrorName() string {
	return "DeleteSceneRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSceneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSceneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSceneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSceneRequestValidationError{}

// Validate checks the field values on Scene_Agreement with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Scene_Agreement) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Scene_Agreement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Scene_AgreementMultiError, or nil if none found.
func (m *Scene_Agreement) ValidateAll() error {
	return m.validate(true)
}

func (m *Scene_Agreement) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return Scene_AgreementMultiError(errors)
	}

	return nil
}

// Scene_AgreementMultiError is an error wrapping multiple validation errors
// returned by Scene_Agreement.ValidateAll() if the designated constraints
// aren't met.
type Scene_AgreementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Scene_AgreementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Scene_AgreementMultiError) AllErrors() []error { return m }

// Scene_AgreementValidationError is the validation error returned by
// Scene_Agreement.Validate if the designated constraints aren't met.
type Scene_AgreementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Scene_AgreementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Scene_AgreementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Scene_AgreementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Scene_AgreementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Scene_AgreementValidationError) ErrorName() string { return "Scene_AgreementValidationError" }

// Error satisfies the builtin error interface
func (e Scene_AgreementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScene_Agreement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Scene_AgreementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Scene_AgreementValidationError{}