package reflection

import (
	"reflect"
	"strings"
)

func Reflection(x any) string {
	var buffer strings.Builder
	buffer.WriteString("type: ")
	buffer.WriteString(reflect.TypeOf(x).String())
	return buffer.String()
}
