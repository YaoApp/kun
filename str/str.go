package str

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"strings"

	"github.com/yaoapp/kun/any"
)

// String type of string
type String string

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// Of make a new string
func Of(value interface{}) String {
	return String(fmt.Sprintf("%v", value))
}

// Bind apply data to the given template
func Bind(input string, data interface{}) string {
	row := any.Of(data).Map().Dot()
	reg := regexp.MustCompile("{{[ ]*([^\\s]+[ ]*)}}")
	matchs := reg.FindAllStringSubmatch(input, -1)
	replaces := map[string]string{}
	for _, match := range matchs {
		name := match[0]
		if _, has := replaces[name]; has {
			continue
		}
		key := match[1]
		replaces[name] = ""
		if row.Has(key) && row.Get(key) != nil {
			replaces[name] = any.Of(row.Get(key)).CString()
		}
	}

	for name, value := range replaces {
		input = strings.ReplaceAll(input, name, value)
	}

	return input
}

// Bind apply data to the given template
func (s String) Bind(data interface{}) string {
	return Bind(string(s), data)
}

// After The After method returns everything after the given value in a string
func After() {}

// After The After method returns everything after the given value in a string
func (s String) After() {}

// AfterLast The AfterLast method returns everything after the last occurrence of the given value in a string
func AfterLast() {}

// AfterLast The AfterLast method returns everything after the last occurrence of the given value in a string
func (s String) AfterLast() {}

// Append The Append method appends the given values to the string
func (s String) Append() {}

// ASCII The ASCII method will attempt to transliterate the string into an ASCII value
func ASCII() {}

// ASCII The ASCII method will attempt to transliterate the string into an ASCII value
func (s String) ASCII() {}

// BaseName The BaseName method will return the trailing name component of the given string
func BaseName() {}

// BaseName The BaseName method will return the trailing name component of the given string
func (s String) BaseName() {}

// Before The Before method returns everything before the given value in a string
func Before() {}

// Before The Before method returns everything before the given value in a string
func (s String) Before() {}

// BeforeLast The BeforeLast method returns everything before the last occurrence of the given value in a string
func BeforeLast() {}

// BeforeLast The BeforeLast method returns everything before the last occurrence of the given value in a string
func (s String) BeforeLast() {}

// Between The Between method returns the portion of a string between two values
func Between() {}

// Between The Between method returns the portion of a string between two values
func (s String) Between() {}

// Camel The Camel method converts the given string to camelCase
func Camel() {}

// Camel The Camel method converts the given string to camelCase
func (s String) Camel() {}

// Contains The Contains method determines if the given string contains the given value. This method is case sensitive
func Contains() {}

// Contains The Contains method determines if the given string contains the given value. This method is case sensitive
func (s String) Contains() {}

// ContainsAll The ContainsAll method determines if the given string contains all of the values in a given array
func ContainsAll() {}

// ContainsAll The ContainsAll method determines if the given string contains all of the values in a given array
func (s String) ContainsAll() {}

// DirName The DirName method returns the parent directory portion of the given string
func DirName() {}

// DirName The DirName method returns the parent directory portion of the given string
func (s String) DirName() {}

// EndsWith The EndsWith method determines if the given string ends with the given value
func EndsWith() {}

// EndsWith The EndsWith method determines if the given string ends with the given value
func (s String) EndsWith() {}

// Exactly The Exactly method determines if the given string is an exact match with another string
func Exactly() {}

// Exactly The Exactly method determines if the given string is an exact match with another string
func (s String) Exactly() {}

// Explode The Explode method splits the string by the given delimiter and returns a collection containing each section of the split string:
func Explode() {}

// Explode The Explode method splits the string by the given delimiter and returns a collection containing each section of the split string:
func (s String) Explode() {}

// Finish The Finish method adds a single instance of the given value to a string if it does not already end with that value
func Finish() {}

// Finish The Finish method adds a single instance of the given value to a string if it does not already end with that value
func (s String) Finish() {}

// Is The Is method determines if a given string matches a given pattern. Asterisks may be used as wildcard values
func Is() {}

// Is The Is method determines if a given string matches a given pattern. Asterisks may be used as wildcard values
func (s String) Is() {}

// IsASCII The IsASCII method determines if a given string is 7 bit ASCII
func IsASCII() {}

// IsASCII The IsASCII method determines if a given string is 7 bit ASCII
func (s String) IsASCII() {}

// IsEmpty The isEmpty method determines if the given string is empty
func (s String) IsEmpty() {}

// IsNotEmpty The IsNotEmpty method determines if the given string is not empty
func (s String) IsNotEmpty() {}

// IsUUID The IsUUID method determines if the given string is a valid UUID
func IsUUID() {}

// IsUUID The IsUUID method determines if the given string is a valid UUID
func (s String) IsUUID() {}

// Kebab The Kebab method converts the given string to kebab-case
func Kebab() {}

// Kebab The Kebab method converts the given string to kebab-case
func (s String) Kebab() {}

// Length The Length method returns the length of the given string
func Length(s string) int {
	return len(s)
}

// Length The Length method returns the length of the given string
func (s String) Length() int {
	return Length(string(s))
}

// Limit The limit method truncates the given string to the specified length
func Limit() {}

// Limit The limit method truncates the given string to the specified length
func (s String) Limit() {}

// Lower The Lower method converts the given string to lowercase
func Lower() {}

// Lower The Lower method converts the given string to lowercase
func (s String) Lower() {}

// Ltrim The Ltrim method trims the left side of the string
func Ltrim() {}

// Ltrim The Ltrim method trims the left side of the string
func (s String) Ltrim() {}

// Markdown The Markdown method converts GitHub flavored Markdown into HTML
func Markdown() {}

// Markdown The Markdown method converts GitHub flavored Markdown into HTML
func (s String) Markdown() {}

// Match The Match method will return the portion of a string that matches a given regular expression pattern
func Match() {}

// Match The Match method will return the portion of a string that matches a given regular expression pattern
func (s String) Match() {}

// MatchAll The MatchAll method will return a collection containing the portions of a string that match a given regular expression pattern
func MatchAll() {}

// MatchAll The MatchAll method will return a collection containing the portions of a string that match a given regular expression pattern
func (s String) MatchAll() {}

// OrderedUUID The OrderedUUID method generates a "timestamp first" UUID that may be efficiently stored in an indexed database column.
// Each UUID that is generated using this method will be sorted after UUIDs previously generated using the method
func OrderedUUID() {}

// PadBoth The PadBoth method padding both sides of a string with another string until the final string reaches a desired length
func PadBoth() {}

// PadBoth The PadBoth method padding both sides of a string with another string until the final string reaches a desired length
func (s String) PadBoth() {}

// PadLeft The PadLeft method padding the left side of a string with another string until the final string reaches a desired length
func PadLeft() {}

// PadLeft The PadLeft method padding the left side of a string with another string until the final string reaches a desired length
func (s String) PadLeft() {}

// PadRight The PadRight method padding the right side of a string with another string until the final string reaches a desired length
func PadRight() {}

// PadRight The PadRight method padding the right side of a string with another string until the final string reaches a desired length
func (s String) PadRight() {}

// Pipe The pipe method allows you to transform the string by passing its current value to the given callable
func Pipe() {}

// Pipe The pipe method allows you to transform the string by passing its current value to the given callable
func (s String) Pipe() {}

// Random The Random method generates a random string of the specified length.
func Random() {}

// Prepend The Prepend method prepends the given values onto the string
func Prepend() {}

// Prepend The Prepend method prepends the given values onto the string
func (s String) Prepend() {}

// Remove The Remove method removes the given value or array of values from the string:
func Remove() {}

// Remove The Remove method removes the given value or array of values from the string:
func (s String) Remove() {}

// Replace The Replace method replaces a given string within the string:
func Replace(s string, old string, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// Replace The Replace method replaces a given string within the string:
func (s String) Replace(old string, new string, n int) string {
	return strings.Replace(string(s), old, new, n)
}

// ReplaceArray The ReplaceArray method replaces a given value in the string sequentially using an array
func ReplaceArray() {}

// ReplaceArray The ReplaceArray method replaces a given value in the string sequentially using an array
func (s String) ReplaceArray() {}

// ReplaceFirst The ReplaceFirst method replaces the first occurrence of a given value in a string
func ReplaceFirst() {}

// ReplaceFirst The ReplaceFirst method replaces the first occurrence of a given value in a string
func (s String) ReplaceFirst() {}

// ReplaceLast The ReplaceLast method replaces the last occurrence of a given value in a string
func ReplaceLast() {}

// ReplaceLast The ReplaceLast method replaces the last occurrence of a given value in a string
func (s String) ReplaceLast() {}

// ReplaceMatches The replaceMatches method replaces all portions of a string matching a pattern with the given replacement string:
func ReplaceMatches() {}

// ReplaceMatches The replaceMatches method replaces all portions of a string matching a pattern with the given replacement string:
func (s String) ReplaceMatches() {}

// Rtrim The Rtrim method trims the right side of the given string
func Rtrim() {}

// Rtrim The Rtrim method trims the right side of the given string
func (s String) Rtrim() {}

// Slug The Slug method generates a URL friendly "slug" from the given string
func Slug() {}

// Slug The Slug method generates a URL friendly "slug" from the given string
func (s String) Slug() {}

// Snake The Snake method converts the given string to snake_case
func Snake(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Snake The Snake method converts the given string to snake_case
func (s String) Snake() string {
	return Snake(string(s))
}

// Split The Split method splits a string into a collection using a regular expression
func Split() {}

// Split The Split method splits a string into a collection using a regular expression
func (s String) Split() {}

// Start The Start method adds a single instance of the given value to a string if it does not already start with that value
func Start() {}

// Start The Start method adds a single instance of the given value to a string if it does not already start with that value
func (s String) Start() {}

// StartsWith The StartsWith method determines if the given string begins with the given value
func StartsWith() {}

// StartsWith The StartsWith method determines if the given string begins with the given value
func (s String) StartsWith() {}

// Studly The Studly method converts the given string to StudlyCase
func Studly() {}

// Studly The Studly method converts the given string to StudlyCase
func (s String) Studly() {}

// Substr The Substr method returns the portion of string specified by the start and length parameters
func Substr() {}

// Substr The Substr method returns the portion of string specified by the start and length parameters
func (s String) Substr() {}

// SubstrCount The Str::substrCount method returns the number of occurrences of a given value in the given string
func SubstrCount() {}

// Tap The Tap method passes the string to the given closure,
// allowing you to examine and interact with the string while not affecting the string itself.
// The original string is returned by the tap method regardless of what is returned by the closure
func Tap() {}

// Tap The Tap method passes the string to the given closure,
// allowing you to examine and interact with the string while not affecting the string itself.
// The original string is returned by the tap method regardless of what is returned by the closure
func (s String) Tap() {}

// Test  The Test method determines if a string matches the given regular expression pattern
func Test() {}

// Test The Test method determines if a string matches the given regular expression pattern
func (s String) Test() {}

// Title The Title method converts the given string to Title Case
func Title() {}

// Title The Title method converts the given string to Title Case
func (s String) Title() {}

// Trim The Trim method trims the given string
func Trim() {}

// Trim The Trim method trims the given string
func (s String) Trim() {}

// Ucfirst The Str::Ucfirst method returns the given string with the first character capitalized
func Ucfirst() {}

// Ucfirst The Str::Ucfirst method returns the given string with the first character capitalized
func (s String) Ucfirst() {}

// Upper The Upper method converts the given string to uppercase
func Upper() {}

// Upper The Upper method converts the given string to uppercase
func (s String) Upper() {}

// When The When method invokes the given closure if a given condition is true. The closure will receive the fluent string instance
func (s String) When() {}

// WhenEmpty The WhenEmpty method invokes the given closure if the string is empty. If the closure returns a value, that value will also be returned by the whenEmpty method. If the closure does not return a value, the fluent string instance will be returned
func (s String) WhenEmpty() {}

// UUID The UUID method generates a UUID (version 4)
func UUID() {}

// WordCount The WordCount function returns the number of words that a string contains
func WordCount() {}

// WordCount The WordCount function returns the number of words that a string contains
func (s String) WordCount() {}

// Words The Words method limits the number of words in a string.
// An additional string may be passed to this method via its third argument to specify which string should be appended to the end of the truncated string
func Words() {}

// Words The Words method limits the number of words in a string.
// An additional string may be passed to this method via its third argument to specify which string should be appended to the end of the truncated string
func (s String) Words() {}

// MarshalJSON for json MarshalJSON
func (s *String) MarshalJSON() ([]byte, error) {
	return []byte(*s), nil
}

// UnmarshalJSON for json UnmarshalJSON
func (s *String) UnmarshalJSON(data []byte) error {
	*s = String(strings.Trim(string(data), `"`))
	return nil
}

// Scan for db scan
func (s *String) Scan(src interface{}) error {
	*s = Of(src)
	return nil
}

// Value for db driver value
func (s *String) Value() (driver.Value, error) {
	return string(*s), nil
}
