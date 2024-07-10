package factory

import (
	"reflect"
	"testing"
)

// 简单工厂
type IParser interface {
	Parse(data []byte)
}

type jsonE struct{}

func (j jsonE) Parse(data []byte) {
	panic("implement me")
}

type yamlE struct{}

func (y yamlE) Parse(data []byte) {
	panic("implement me")
}

func NewIParser(t string) IParser {
	switch t {
	case "json":
		return jsonE{}
	case "yaml":
		return yamlE{}
	}
	return nil
}

func TestEasy(t *testing.T) {
	tests := []struct {
		name string
		args string
		want IParser
	}{
		{
			name: "json",
			args: "json",
			want: jsonE{},
		},
		{
			name: "yaml",
			args: "yaml",
			want: yamlE{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIParser(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIParser: %v, want %v\n", got, tt.want)
			}
		})
	}
}
