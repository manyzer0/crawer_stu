package parses

import (
	"crawlergo/engine"
	"regexp"
)

var NicknameReg = regexp.MustCompile(`<a href="http://album.zhenai.com/u/\d{5,}" target="_blank">([^<]+)</a>`)
var cityurl = regexp.MustCompile(`http://www.zhenai.com/zhenghun/shenzhen/\d{1,}`)

func PaserUser(body []byte) engine.ParseResult {

	names := NicknameReg.FindAllStringSubmatch(string(body), -1)
	// urls := cityurl.FindAllString(string(body), -1)
	result := engine.ParseResult{}
	for _, name := range names {
		result.Items = append(result.Items, name[1])
	}
	request := engine.Request{}
	citys := cityurl.FindAllString(string(body), -1)
	for _, i := range citys {
		if _, ok := engine.Cityurls[i]; ok {
			continue
		}
		request.Url = i
		request.ParseFunc = PaserUser
		result.Request = append(result.Request, request)
	}

	return result
}
