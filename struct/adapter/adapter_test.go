package adapter

import (
	"fmt"
	"testing"
)

// 适配器模式

// 创建云主机
type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

// aws sdk
type AwsClient struct{}

// 启动实例
func (c *AwsClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu: %f, mem: %f", cpu, mem)
	return nil
}

// 适配器
type AwsClientAdapter struct {
	Client AwsClient
}

// 启动实例
func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.RunInstance(cpu, mem)
	return nil
}

// aliyun sdk
type AliyunClient struct{}

// 启动实例
func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aws client run success cpu: %d, mem: %d", cpu, mem)
	return nil
}

// 适配器
type AliyunClientAdapter struct {
	Client AliyunClient
}

// 启动实例
func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}

func TestAwsClient(t *testing.T) {
	// 确保adapter实现了接口
	var a ICreateServer = &AwsClientAdapter{
		Client: AwsClient{},
	}
	a.CreateServer(1.0, 2.0)
}

func TestAliyun(t *testing.T) {
	// 确保adapter实现了接口
	var a ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}
	a.CreateServer(1.0, 2.0)
}
