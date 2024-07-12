package factory

import (
	"reflect"
	"testing"
)

// 工厂方法

type IRuleParserFactory interface {
	CreateParser() IParser
}

type jsonFactory struct{}

func (j jsonFactory) CreateParser() IParser {
	return jsonE{}
}

type yamlFactory struct{}

func (y yamlFactory) CreateParser() IParser {
	return yamlE{}
}

func NewParserFactory(t string) IRuleParserFactory {
	switch t {
	case "json":
		return jsonFactory{}
	case "yaml":
		return yamlFactory{}
	}
	return nil
}

func TestFactory(t *testing.T) {
	tests := []struct {
		name string
		args string
		want IRuleParserFactory
	}{
		{
			name: "json",
			args: "json",
			want: jsonFactory{},
		},
		{
			name: "yaml",
			args: "yaml",
			want: yamlFactory{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParserFactory(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParserFactory: %v, want %v\n", got, tt.want)
			}
		})
	}
}
