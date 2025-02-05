package builder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 建造者模式

type configOpt struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type config struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type OptFunc func(opt *configOpt)

func NewConfig(name string, opts ...OptFunc) (*config, error) {
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}

	option := &configOpt{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}

	for _, opt := range opts {
		opt(option)
	}

	if option.maxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	if option.maxTotal < option.maxIdle || option.minIdle > option.maxIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	return &config{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}

func TestBuilder(t *testing.T) {
	type args struct {
		name string
		opts []OptFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *config
		wantErr bool
	}{
		{
			name: "name empty",
			args: args{
				name: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				name: "test",
				opts: []OptFunc{
					func(option *configOpt) {
						option.minIdle = 2
					},
				},
			},
			want: &config{
				name:     "test",
				maxTotal: 10,
				maxIdle:  9,
				minIdle:  2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.name, tt.args.opts...)
			require.Equalf(t, tt.wantErr, err != nil, "error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
