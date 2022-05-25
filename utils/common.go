package utils

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

//MaxInt64 returns the max of two int64s
// Can't believe I have to define this.
func MaxInt64(a int64, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

//CheckNullAndSetString b if b is not null, else it returns a for string
func CheckNullAndSetString(a *string, b *string) *string {
	if b != nil {
		return b
	}
	return a
}

//CheckNullAndSetInt64 b if b is not null, else it returns a for int64
func CheckNullAndSetInt64(a *int64, b *int64) *int64 {
	if b != nil {
		return b
	}
	return a
}

//CheckNullAndSetFloat64 b if b is not null, else it returns a for Float64
func CheckNullAndSetFloat64(a *float64, b *float64) *float64 {
	if b != nil {
		return b
	}
	return a
}

//CheckNullAndSetBool b if b is not null, else it returns a for bool
func CheckNullAndSetBool(a *bool, b *bool) *bool {
	if b != nil {
		return b
	}
	return a
}

//GetRandomUUID returns a random uuid
func GetRandomUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// Remove removes the string at index i in a slice.
// Note that it doesn't preserve the order. So don't use when order is important
func Remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

//SerializeStringSliceToString serializes a slice of string to comma separate string
func SerializeStringSliceToString(list []string) string {
	res := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(list)), ","), "[]")
	return res
}

//StringToStringSlice returns a slice of string from a single string of comma separated strings
func StringToStringSlice(s *string) []string {
	var res []string
	if s != nil {
		res = strings.Split(*s, ",")
	}
	return res
}

// EncodeURL will encode string to url
func EncodeURL(u string) string {
	t := &url.URL{Path: u}
	return t.String()
}
