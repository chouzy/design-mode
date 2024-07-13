package strategy

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 策略模式

// 存储策略
type StorageStrategy interface {
	Save(name string, data []byte) error
}

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}
	return s, nil
}

// 保存到文件
type fileStorage struct{}

func (f *fileStorage) Save(name string, data []byte) error {
	return os.WriteFile(name, data, os.ModeAppend)
}

// 加密保存到文件
type encryptFileStorage struct{}

func (e *encryptFileStorage) Save(name string, data []byte) error {
	data, err := encrypt(data)
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	// 实现加密算法
	return data, nil
}

func TestStrategy(t *testing.T) {
	data, sensitive := getData()
	strategyType := "file"
	if sensitive {
		strategyType = "encrypt_file"
	}

	storage, err := NewStorageStrategy(strategyType)
	assert.NoError(t, err)
	assert.NoError(t, storage.Save("./text.txt", data))
}

func getData() ([]byte, bool) {
	return []byte("test, data"), false
}
