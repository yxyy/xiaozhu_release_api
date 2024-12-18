package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

// 校验 IP 地址是否有效
func isValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

// Ip ip-api.com
type Ip struct {
	Ip          string `json:"ip"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	City        string `json:"city"`
}

type Aip struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func IpInfo(ip string) (*Ip, error) {
	if !isValidIP(ip) {
		return nil, fmt.Errorf("无效的ip地址")
	}
	urls := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)
	body, err := httpGet(urls)
	if err != nil {
		return nil, err
	}

	ipInfo := new(Ip)
	if err = json.Unmarshal(body, &ipInfo); err != nil {
		return nil, err
	}

	return ipInfo, nil
}

// Tip 太平洋IP地址归属地查询接口
type Tip struct {
	Ip          string `json:"ip"`
	Country     string `json:"pro"`
	CountryCode string `json:"proCode"`
	City        string `json:"city"`
	CityCode    string `json:"cityCode"`
	Region      string `json:"region"`
	RegionCode  string `json:"regionCode"`
	Addr        string `json:"addr"`
	RegionNames string `json:"regionNames"`
	Err         string `json:"err"`
}

func (t *Tip) IpInfo(ip string) (*Ip, error) {
	if !isValidIP(ip) {
		return nil, fmt.Errorf("无效的ip地址")
	}
	urls := fmt.Sprintf("https://whois.pconline.com.cn/ipJson.jsp?ip=%s8&json=true", ip)
	body, err := httpGet(urls)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	ipInfo := &Ip{
		Country:     t.Country,
		CountryCode: t.CountryCode,
		Region:      t.Region,
		RegionName:  t.RegionNames,
		City:        t.City,
	}

	return ipInfo, nil
}

func httpGet(urls string) ([]byte, error) {
	resp, err := http.Get(urls)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求返回非 200 状态码: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)

}
