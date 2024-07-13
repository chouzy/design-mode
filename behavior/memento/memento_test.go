package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 备忘录模式

// 如果输入 :list 则显示当前保存的内容
// 如果输入 :undo 则删除上一次的输入
// 如果输入其他的内容则追加保存

// 快照，用于存储数据快照
// 对于快照，只能获取，不能修改
type Snapshot struct {
	content string
}

func (s *Snapshot) GetText() string {
	return s.content
}

// 用于保存数据
type InputText struct {
	content string
}

func (i *InputText) Append(context string) {
	i.content += context
}

func (i *InputText) GetText() string {
	return i.content
}

func (i *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: i.content}
}

func (i *InputText) Restore(s *Snapshot) {
	i.content = s.GetText()
}

func TestDemo(t *testing.T) {
	in := &InputText{}
	snapshots := []*Snapshot{}

	tests := []struct {
		input string
		want  string
	}{
		{
			input: ":list",
			want:  "",
		},
		{
			input: "hello",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
		{
			input: "world",
			want:  "",
		},
		{
			input: ":list",
			want:  "helloworld",
		},
		{
			input: ":undo",
			want:  "",
		},
		{
			input: ":list",
			want:  "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			switch tt.input {
			case ":list":
				assert.Equal(t, tt.want, in.GetText())
			case ":undo":
				in.Restore(snapshots[len(snapshots)-1])
				snapshots = snapshots[:len(snapshots)-1]
			default:
				snapshots = append(snapshots, in.Snapshot())
				in.Append(tt.input)
			}
		})
	}
}
