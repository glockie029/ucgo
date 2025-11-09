package MethodPractice

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type ScanConfig struct {
	Targeturl string
	Threads   int
	Header    map[string]string
	Silent    bool
}

type ScanResult struct {
	IPAddress     string `json: "IP"`
	Vulnerability string `json: "Vulnerability"`
	Timestamp     time.Time
}

func (s ScanConfig) Summary() string {
	return fmt.Sprintf("Scanning %s with %d threads.", s.Targeturl, s.Threads)
}

func NewScanConfig(url string, threads int) (*ScanConfig, error) {
	if url == "" {
		return nil, errors.New("url cannot be blank")
	}
	if threads <= 0 {
		threads = 5
	}
	return &ScanConfig{
		Targeturl: url,
		Threads:   threads,
		Header:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func ScanMain() {

	cfg := ScanConfig{
		Targeturl: "http://www.baidu.com",
		Threads:   1,
		Header: map[string]string{
			"User-Agent": "test",
		},
		Silent: false,
	}
	result := ScanResult{
		IPAddress:     "127.0.0.1",
		Vulnerability: "SQL Injection",
		Timestamp:     time.Time{},
	}
	fmt.Println("=====默认检查=====")
	fmt.Printf("1.目标:%s\n2.并发数:%d\n3.UA:%s\n3.静默模式:%t\n", cfg.Targeturl, cfg.Threads, cfg.Header["User-Agent"], cfg.Silent)

	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Json Error", err)
	}
	fmt.Println("=====输出格式化的json日志文件=====")
	fmt.Printf("1.输出日志文件:%s\n", string(jsonData))

	cfgf, _ := NewScanConfig("http://www.baidu.com", 10)
	fmt.Println("=====调用工厂方法的结果=====")
	fmt.Printf("2.使用工厂方法并调用Smmary方法:%s\n", cfgf.Summary())

	var results []ScanResult
	results = append(results, ScanResult{
		IPAddress:     "192.168.3.1",
		Vulnerability: "RCE",
		Timestamp:     time.Time{},
	})
	results = append(results, ScanResult{
		IPAddress:     "192.168.3.2",
		Vulnerability: "Xss",
		Timestamp:     time.Time{},
	})
	fmt.Println("=====输出所有扫描的IP地址=====")
	for _, i := range results {
		fmt.Printf("IP地址:%s\n", i.IPAddress)
	}
}
