package exception

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
)

// Exception the Exception type
type Exception struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Context interface{} `json:"context"`
}

// New Create a new exception instance
func New(message string, code int, args ...interface{}) *Exception {
	content := fmt.Sprintf(message, args...)
	return &Exception{
		Message: content,
		Code:    code,
	}
}

// Err Create an exception instance from the error
func Err(err error, code int) *Exception {
	return &Exception{
		Message: err.Error(),
		Code:    code,
	}
}

// Catch Exception catch and recovered
func Catch(recovered interface{}, err ...error) error {
	if recovered == nil {
		if len(err) > 0 {
			return err[0]
		}
		return nil
	} else if err, ok := recovered.(string); ok {
		return fmt.Errorf("%s", err)
	} else if err, ok := recovered.(Exception); ok {
		return fmt.Errorf("%s", err.Message)
	} else if err, ok := recovered.(*Exception); ok {
		return fmt.Errorf("%s", err.Message)
	}
	return fmt.Errorf("%s", recovered)
}

// CatchPrint Catch the exception and print it
func CatchPrint() {
	if r := recover(); r != nil {
		switch r.(type) {
		case *Exception:
			color.Red(r.(*Exception).Message)
			r.(*Exception).Print()
			break
		case string:
			color.Red(r.(string))
			break
		case error:
			color.Red(r.(error).Error())
			break
		default:
			color.Red("%#v\n", r)
		}
	}
}

// CatchDebug Catch the exception and print debug info
func CatchDebug() {
	if r := recover(); r != nil {
		switch r.(type) {
		case *Exception:
			color.Red(r.(*Exception).Message)
			r.(*Exception).Print()
			break
		case string:
			color.Red(r.(string))
			break
		case error:
			color.Red(r.(error).Error())
			break
		default:
			color.Red("%#v\n", r)
		}

		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
	}
}

// Ctx Add the context for the exception.
func (exception *Exception) Ctx(context interface{}) *Exception {
	exception.Context = context
	return exception
}

// Print print the exception
func (exception Exception) Print() {
	f := colorjson.NewFormatter()
	f.Indent = 2
	var res interface{}
	txt, _ := json.Marshal(exception)
	json.Unmarshal(txt, &res)
	s, _ := colorjson.Marshal(res)
	fmt.Println(string(s))
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
