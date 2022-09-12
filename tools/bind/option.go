package bind

import "github.com/lingyao2333/go-basics-utils/options"

var defaultParse = []Binding{Query, FormPost, Header}

func DefaultParse(contentType string) []Binding {
	return append(defaultParse, contentTypeSelect(contentType))
}

func SetSelectorParse(bind []Binding) options.Option {
	return func(o interface{}) {
		o.(*all).selectorParse = bind
	}
}
