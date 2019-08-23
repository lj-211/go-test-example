package nettest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func _httpCodeError(code int) error {
	return errors.Errorf("http状态码%d", code)
}

func GetUrl(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	rs, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()

	if rs.StatusCode != http.StatusOK {
		return nil, _httpCodeError(rs.StatusCode)
	}

	data, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func JsonGetUrl(ctx context.Context, url string, reply interface{}) error {
	rdata, err := GetUrl(ctx, url)
	if err != nil {
		return errors.Wrap(err, "GET请求失败")
	}

	if reply != nil {
		jerr := json.Unmarshal(rdata, reply)
		if jerr != nil {
			return errors.Wrap(jerr, "解码json数据失败")
		}
	}

	return nil
}

func VerifyData(url string) error {
	data := &struct {
		Result string
	}{}
	err := JsonGetUrl(context.Background(), url, data)
	if err != nil {
		return errors.Wrap(err, "请求错误")
	}

	if data.Result != "OK" {
		return errors.Errorf("返回错误")
	}
	return nil
}
