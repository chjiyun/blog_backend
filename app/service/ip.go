package service

import (
	"blog_backend/app/util"
	"blog_backend/config"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

type IpInfo struct {
	Ip       string `json:"ip"`
	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省
	City     string `json:"city"`     // 市
	County   string `json:"county"`   // 县、区
	Region   string `json:"region"`   // 区域位置
	ISP      string `json:"isp"`      // 互联网服务提供商
}

func Ip2Region(ip string) (*IpInfo, error) {
	if ip == "" {
		return nil, errors.New("ip is required")
	}
	url := util.WriteString(config.Cfg.Ip2RegionUrl, "?ip=", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ipInfo IpInfo
	_ = jsoniter.Unmarshal(body, &ipInfo)
	return &ipInfo, nil
}
