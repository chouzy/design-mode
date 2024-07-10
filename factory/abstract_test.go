package factory

import (
	"reflect"
	"testing"
)

// 抽象工厂
type IJson interface {
	Parse(data []byte)
}

type json struct{}

func (j json) Parse(data []byte) {
	panic("parse json")
}

type IYaml interface {
	Parse(data []byte)
}

type yaml struct{}

func (y yaml) Parse(data []byte) {
	panic("parse yaml")
}

// 工厂方法接口
type IParseFactory interface {
	CreateJson() IJson
	CreateYaml() IYaml
}

type ParseFactory struct{}

func (p ParseFactory) CreateJson() IJson {
	return json{}
}

func (p ParseFactory) CreateYaml() IYaml {
	return yaml{}
}

func TestJson(t *testing.T) {
	tests := []struct {
		name string
		want IJson
	}{
		{
			name: "json",
			want: json{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := ParseFactory{}
			if got := j.CreateJson(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json: %v, want: %v\n", got, tt.want)
			}
		})
	}
}

func TestYaml(t *testing.T) {
	tests := []struct {
		name string
		want IYaml
	}{
		{
			name: "yaml",
			want: yaml{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := ParseFactory{}
			if got := j.CreateYaml(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yaml: %v, want: %v\n", got, tt.want)
			}
		})
	}
}
