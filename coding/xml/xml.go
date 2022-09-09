package xml

import (
	"encoding/xml"
	"reflect"

	"github.com/lingyao2333/go-basics-utils/coding"
)

const Name = "xml"

func init() {
	_ = coding.RegisterCode(code{})
}

type code struct{}

func (c code) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (c code) Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}
	return xml.Unmarshal(data, v)
}

func (c code) Name() string {
	return Name
}
