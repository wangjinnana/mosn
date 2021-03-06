// Code generated by protoc-gen-validate
// source: envoy/type/matcher/number.proto
// DO NOT EDIT!!!

package matcher

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on DoubleMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DoubleMatcher) Validate() error {
	if m == nil {
		return nil
	}

	switch m.MatchPattern.(type) {

	case *DoubleMatcher_Range:

		if v, ok := interface{}(m.GetRange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DoubleMatcherValidationError{
					Field:  "Range",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *DoubleMatcher_Exact:
		// no validation rules for Exact

	default:
		return DoubleMatcherValidationError{
			Field:  "MatchPattern",
			Reason: "value is required",
		}

	}

	return nil
}

// DoubleMatcherValidationError is the validation error returned by
// DoubleMatcher.Validate if the designated constraints aren't met.
type DoubleMatcherValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e DoubleMatcherValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDoubleMatcher.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = DoubleMatcherValidationError{}
