package exception

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"github.com/yaoapp/kun/any"
)

// Mode the mode of the application
var Mode = "production"

var (
	devWriter   io.Writer = os.Stdout
	devWriterMu sync.RWMutex
)

// SetWriter replaces the output target for debug/exception messages.
// Pass nil to reset to os.Stdout.
func SetWriter(w io.Writer) {
	devWriterMu.Lock()
	if w == nil {
		devWriter = os.Stdout
	} else {
		devWriter = w
	}
	devWriterMu.Unlock()
}

// GetWriter returns the current output target.
func GetWriter() io.Writer {
	devWriterMu.RLock()
	defer devWriterMu.RUnlock()
	return devWriter
}

var reEx = regexp.MustCompile(`Exception\|(\d+):(.*)`)
var reErr = regexp.MustCompile(`Error: (.*)`)

// Exception the Exception type
type Exception struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Context interface{} `json:"context"`
}

// New Create a new exception instance
func New(message string, code int, args ...interface{}) *Exception {
	content := fmt.Sprintf(message, args...)
	match := reEx.FindStringSubmatch(content)
	if len(match) > 0 {
		code = any.Of(match[1]).CInt()
		content = strings.TrimSpace(match[2])
	}
	return &Exception{Message: content, Code: code}
}

// Trim the exception message
func Trim(err error) string {
	message := err.Error()
	match := reEx.FindStringSubmatch(message)
	if len(match) > 0 {
		return strings.TrimSpace(match[2])
	}

	// Trim the Error:
	match = reErr.FindStringSubmatch(message)
	if len(match) > 0 {
		return strings.TrimSpace(match[1])
	}
	return message
}

// Err Create an exception instance from the error
func Err(err error, code int) *Exception {
	return New(err.Error(), code)
}

// Catch Exception catch and recovered
func Catch(recovered interface{}, err ...error) error {

	if recovered == nil {
		if len(err) > 0 {
			messages := []string{}
			for _, e := range err {
				if e != nil {
					messages = append(messages, e.Error())
				}
			}

			if len(messages) == 0 {
				return nil
			}

			return fmt.Errorf("%s", strings.Join(messages, ", "))
		}
		return nil
	} else if v, ok := recovered.(string); ok {
		return fmt.Errorf("%s", v)

	} else if v, ok := recovered.(Exception); ok {
		return fmt.Errorf("Exception|%d: %s", v.Code, v.Message)

	} else if v, ok := recovered.(*Exception); ok {
		return fmt.Errorf("Exception|%d: %s", v.Code, v.Message)
	}

	return fmt.Errorf("%s", recovered)
}

// DebugPrint print the message only in development mode
func DebugPrint(err error, message string, args ...interface{}) {
	if Mode == "development" {
		var buf strings.Builder
		red := color.New(color.FgRed)
		red.Fprintln(&buf, "\n----------------------------------")
		ex := Err(err, 500)
		red.Fprintf(&buf, "Exception: %d %s\n", ex.Code, ex.Message)
		red.Fprintln(&buf, "----------------------------------")
		fmt.Fprintf(&buf, message, args...)
		fmt.Fprintln(&buf)
		if Mode == "development" {
			color.New(color.FgYellow).Fprintln(&buf, "Trace Recovered:")
			fmt.Fprintf(&buf, "%s\n", debug.Stack())
		}
		GetWriter().Write([]byte(buf.String()))
	}
}

// CatchPrint Catch the exception and print it
func CatchPrint() {
	if r := recover(); r != nil {
		var buf strings.Builder
		red := color.New(color.FgRed)
		red.Fprintln(&buf, "Exception:")
		switch r.(type) {
		case *Exception:
			red.Fprintln(&buf, r.(*Exception).Message)
			printExceptionJSON(&buf, *r.(*Exception))
		case string:
			red.Fprintln(&buf, r.(string))
		case error:
			red.Fprintln(&buf, r.(error).Error())
		default:
			red.Fprintf(&buf, "%#v\n", r)
		}
		GetWriter().Write([]byte(buf.String()))
	}
}

// CatchDebug Catch the exception and print debug info
func CatchDebug() {
	if r := recover(); r != nil {
		var buf strings.Builder
		red := color.New(color.FgRed)
		red.Fprintln(&buf, "Exception:")
		switch r.(type) {
		case *Exception:
			red.Fprintln(&buf, r.(*Exception).Message)
			printExceptionJSON(&buf, *r.(*Exception))
		case string:
			red.Fprintln(&buf, r.(string))
		case error:
			red.Fprintln(&buf, r.(error).Error())
		default:
			red.Fprintf(&buf, "%#v\n", r)
		}
		fmt.Fprintln(&buf, "stacktrace from panic: \n"+string(debug.Stack()))
		GetWriter().Write([]byte(buf.String()))
	}
}

// Ctx Add the context for the exception.
func (exception *Exception) Ctx(context interface{}) *Exception {
	exception.Context = context
	return exception
}

// Print print the exception
func (exception Exception) Print() {
	var buf strings.Builder
	printExceptionJSON(&buf, exception)
	GetWriter().Write([]byte(buf.String()))
}

func printExceptionJSON(buf *strings.Builder, exception Exception) {
	f := colorjson.NewFormatter()
	f.Indent = 2
	var res interface{}
	txt, _ := json.Marshal(exception)
	json.Unmarshal(txt, &res)
	s, _ := colorjson.Marshal(res)
	fmt.Fprintln(buf, string(s))
}

// Throw Throw the exception and terminal progress.
func (exception Exception) Throw() {
	panic(exception)
}

// String interface
func (exception Exception) String() string {
	txt, _ := json.Marshal(exception)
	return string(txt)
}
