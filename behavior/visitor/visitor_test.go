package visitor

import (
	"fmt"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

// 访问者模式

// 访问者
type Visitor interface {
	Visit(IResourceFile) error
}

type IResourceFile interface {
	Accept(Visitor) error
}

func NewResourceFile(filepath string) (IResourceFile, error) {
	switch path.Ext(filepath) {
	case ".ppt":
		return &PPTFile{path: filepath}, nil
	case ".pdf":
		return &PDFFile{path: filepath}, nil
	default:
		return nil, fmt.Errorf("not found file type: %s", filepath)
	}
}

type PDFFile struct {
	path string
}

func (p *PDFFile) Accept(vis Visitor) error {
	return vis.Visit(p)
}

type PPTFile struct {
	path string
}

func (p *PPTFile) Accept(vis Visitor) error {
	return vis.Visit(p)
}

// 实现压缩功能
type Compressor struct{}

func (c *Compressor) VisitPPTFile(f *PPTFile) error {
	fmt.Println("this is ppt file")
	return nil
}

func (c *Compressor) VisitPDFFile(f *PDFFile) error {
	fmt.Println("this is pdf file")
	return nil
}

// Visit 实现访问者模式方法
// 我们可以发现由于没有函数重载，我们只能通过断言来根据不同的类型调用不同函数
// 但是我们即使不采用访问者模式，我们其实也是可以这么操作的
// 并且由于采用了类型断言，所以如果需要操作的对象比较多的话，这个函数其实也会膨胀的比较厉害
// 后续可以考虑按照命名约定使用 generate 自动生成代码
// 或者是使用反射简化
func (c *Compressor) Visit(r IResourceFile) error {
	switch f := r.(type) {
	case *PPTFile:
		return c.VisitPPTFile(f)
	case *PDFFile:
		return c.VisitPDFFile(f)
	default:
		return fmt.Errorf("not fount resource type: %#V", r)
	}
}

func TestVisit(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr string
	}{
		{
			name: "pdf",
			path: "./xx.pdf",
		},
		{
			name: "ppt",
			path: "./xx.ppt",
		},
		{
			name:    "404",
			path:    "./xx.xx",
			wantErr: "not found file type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := NewResourceFile(tt.path)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			compressor := &Compressor{}
			f.Accept(compressor)
		})
	}
}
