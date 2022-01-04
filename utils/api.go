package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lixiang4u/aliyundrive-sdk/models"
	_ "github.com/lixiang4u/aliyundrive-sdk/models"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	ApiLogin        = "https://passport.aliyundrive.com/newlogin/sms/login.do?appName=aliyun_drive&fromSite=52&_bx-v=2.0.31"
	LoginPage       = "https://passport.aliyundrive.com/mini_login.htm?lang=zh_cn&appName=aliyun_drive&appEntrance=web&styleType=auto&bizParams=&notLoadSsoView=false&notKeepLogin=false&isMobile=false&ad__pass__q__rememberLogin=false&ad__pass__q__forgotPassword=false&ad__pass__q__licenseMargin=false&ad__pass__q__loginType=normal&hidePhoneCode=true&rnd=0.32309972272841225"
	ApiFileList     = "https://api.aliyundrive.com/adrive/v3/file/list"
	ApiFileSearch   = "https://api.aliyundrive.com/adrive/v3/file/search"
	ApiFileDownload = "https://api.aliyundrive.com/v2/file/get_download_url"
	ApiCreateFolder = "https://api.aliyundrive.com/adrive/v2/file/createWithFolders"

	ContentTypeJSON = "application/json;charset=UTF-8"
	ContentTypeForm = "application/x-www-form-urlencoded"
	AuthToken       = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxZjY1MGE4YWM4ZTU0MjNhOGZkMDFlYzUzOTY1M2JjMiIsImN1c3RvbUpzb24iOiJ7XCJjbGllbnRJZFwiOlwiMjVkelgzdmJZcWt0Vnh5WFwiLFwiZG9tYWluSWRcIjpcImJqMjlcIixcInNjb3BlXCI6W1wiRFJJVkUuQUxMXCIsXCJTSEFSRS5BTExcIixcIkZJTEUuQUxMXCIsXCJVU0VSLkFMTFwiLFwiU1RPUkFHRS5BTExcIixcIlNUT1JBR0VGSUxFLkxJU1RcIixcIkJBVENIXCIsXCJPQVVUSC5BTExcIixcIklNQUdFLkFMTFwiLFwiSU5WSVRFLkFMTFwiLFwiQUNDT1VOVC5BTExcIl0sXCJyb2xlXCI6XCJ1c2VyXCIsXCJyZWZcIjpcImh0dHBzOi8vd3d3LmFsaXl1bmRyaXZlLmNvbS9cIixcImRldmljZV9pZFwiOlwiNzcyYzA3NGQ4M2NiNGJlODkxN2U4ZWVhNjBmODRhMDZcIn0iLCJleHAiOjE2NDEwNDk4MjUsImlhdCI6MTY0MTA0MjU2NX0.Er809bQ2FdCqm-9xQNoDFSi7ZVQrSG557HzZs4wlFew1uTztXv6cyCEQBerl0t-AugJGVlU-eUEzkSRUvl5MKDOYjn06A5Jg6eNBRdXhGAF2WkCJWozyF10ckKheuUaVsQyO2-dHMoSy_TIQHMN-ozNDz3MFrDVsQvvUbA8zIxc"
)

func getHeader() map[string]string {
	var m = make(map[string]string)
	m["authority"] = "passport.aliyundrive.com"
	m["user-agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"
	m["content-type"] = "application/json;charset=UTF-8"
	m["origin"] = "https://www.aliyundrive.com"
	m["referer"] = "https://www.aliyundrive.com/"
	m["accept"] = "application/json, text/plain, */*"
	m["accept-encoding"] = "gzip, deflate, br"
	m["accept-language"] = "zh-CN,zh;q=0.9"
	m["authorization"] = fmt.Sprintf("Bearer %s", AuthToken)
	m["x-device-id"] = "uEfvGXwdOX8CASuB8RRHGAlt"
	return m
}

func getLoginForm() map[string]string {
	var m = make(map[string]string)
	m["loginId"] = "18552072610"
	m["phoneCode"] = "86"
	m["countryCode"] = "CN"
	m["smsCode"] = "111111"
	m["smsToken"] = "idc_1A9928C8A1BFC45FF99FFCD7C49436643"
	m["keepLogin"] = "false"
	m["ua"] = "140#l6urBM9KzzZXizo2LikQA3SogRPDLUnPrjob41UJwT9DvRvmM0oY13KIs9IK5UP0TvPVzeHjlpVUJ1JihnOi2JUH2Cxh+Dgqlbzxk8k8gGwfzznSO2aIltQzzPzbVXlqlbr+S+FKwI7kz/rdkZmPpzUBeIrYdsoqtuhu2PcLVthqabzi228nHknwzPwbbOrdxz4F2P35u3QpzFKw21GXl3gtPIKuV26AAzKx2P+mDp5+zQKfzWMF1YRLMZ2aMXLi5ye7BccNmVFtCXzbg10ntWLV+H8Sjgo4w00hcHrbttQltyf7pY/68NhzgkzHAJsFTO9ArHp0IiAxct6wxZwiNNqMZKDkWIlSyoE5y6JxXDaBbFouzVYYOf0Vk01Ott3sUpP/wP9nyRTyfKVHJybHc2J1hznna0l5dRwzdVzKraZHt0fqFj9skoanHxjcVzKRliswKqZEnKq3WaakHw6/wJi4iBOWzuYNEelphtgB0wH1Xr8iaTVc+5b18aesbCsIjyd4OyLxGMHY7bwLh2MlQravawfl+6UFlPJ+aST36BeJpCfdtyr5JTaKapNF2HUEIEgT65Fa7Kb+O65YMDXTb3qW0nObec0Wa3er/JvIcafavoX+olHBbceUIwvM9DdxJWxXrgr6gwijFOcPVa54a1TlmoAlxOsw8p+j2CANJOXRfsa3XW/Q9oILSL8onRrG365HC0eRHq120WHSd3bju1yEWLvhWguQB0vXbMlZyIg29Rdhsxg4bDblNef7OMlh94AEzbBJiPqbAu5tY4TJbzQCoMkOil9erDTYYHqAhZjr4kx609Ako24i+Tnua80Tk96YK+6H2C2EJ4A/0wSd/v5CRHFYdM7wU334cxH2HxVFHpBstV/oHS6hMXvbt3krivFghe0SPQIJVCEtyWVJ7qOWBblauUFbfBxNVq/NQQjvJr7mOZX+5hXT5db72bMBL3bvfIS9DQBNXOkMTV/dMs9yehipEXMTNIcsJnHotHjCjHE7ux8BC1hLwN4HSgWxWustdHyWbLDjnxvDEOWNkPkY7M92H/sq4Y8bGalIRSXdAfjlluwxH9N+NehDmfI2AC812LuRVba4O82tSdOu+ZCN2VqVJGNBmR8TJSelgim4PUsz0ygIcf48EVAcUfChkEBTZVkm2mzJGBoZHsFBpK9qBViHCjNF2C9lX59wkpNKE5saREa2+qTAGV76i+mAZcyNfeXXxHU7YXMf086L80NDmnQmqabXCLc425VZeH6YrktV//ut7AFIeKP5gKj1jfZ49iUMo9SkgA+AInDuOT4ES0+ma3cwi219cZJci3xlpdOCNUHnyuRFPxIuhCc3WPoT6MQn+gp3QoYF9sYWwE6CGI8s1RqGFpGGIBPefSh+JG1sZfHDu42/f+aXRfFhiXza58Pvh5e1ZT+llWSVNbVMOpkPNl3w0pxv51L7VGfT2otQCkL+zxwFHE5gLgDbLKJv3GKh0Uueckh6jRxb8Wpdx4BPlBRx8/imMSSy+NVPmGPG5ApKg1FAOySltvtagO2nCVlx"
	m["appName"] = "aliyun_drive"
	m["appEntrance"] = "web"
	m["appName"] = "appEntrance"

	return m
}

// gzip处理
func ResponseUnzip(resp *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// gzip处理
	if resp.Header.Get("Content-Encoding") == "gzip" {
		buffer := bytes.NewBuffer(b)
		r, err := gzip.NewReader(buffer)
		if err != nil {
			return nil, err
		}
		b, err = ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

// 解析登陆页面配置数据
func GetLoginFormConfig() (gjson.Result, error) {
	resp, err := http.Get(LoginPage)
	if err != nil {
		return gjson.Result{}, err
	}
	b, err := ResponseUnzip(resp)
	if err != nil {
		return gjson.Result{}, err
	}

	reg, err := regexp.Compile(`window.viewData = ([\s\S]+);`)
	if err != nil {
		return gjson.Result{}, err
	}

	for _, str := range strings.Split(string(b), "\n") {
		findList := reg.FindStringSubmatch(str)
		if len(findList) <= 1 {
			continue
		}

		return gjson.Parse(findList[1]), nil
	}

	return gjson.Result{}, errors.New("not Found")
}

func Login(username, password string, loginConf models.LoginConfig) {
	var p = url.Values{}

	p.Add("loginId", username)
	p.Add("phoneCode", "86")
	p.Add("countryCode", "CN")
	p.Add("smsCode", password)
	p.Add("smsToken", "idc_1A9928C8A1BFC45FF99FFCD7C49436643")
	p.Add("ua", "140#l6urBM9KzzZXizo2LikQA3SogRPDLUnPrjob41UJwT9DvRvmM0oY13KIs9IK5UP0TvPVzeHjlpVUJ1JihnOi2JUH2Cxh+Dgqlbzxk8k8gGwfzznSO2aIltQzzPzbVXlqlbr+S+FKwI7kz/rdkZmPpzUBeIrYdsoqtuhu2PcLVthqabzi228nHknwzPwbbOrdxz4F2P35u3QpzFKw21GXl3gtPIKuV26AAzKx2P+mDp5+zQKfzWMF1YRLMZ2aMXLi5ye7BccNmVFtCXzbg10ntWLV+H8Sjgo4w00hcHrbttQltyf7pY/68NhzgkzHAJsFTO9ArHp0IiAxct6wxZwiNNqMZKDkWIlSyoE5y6JxXDaBbFouzVYYOf0Vk01Ott3sUpP/wP9nyRTyfKVHJybHc2J1hznna0l5dRwzdVzKraZHt0fqFj9skoanHxjcVzKRliswKqZEnKq3WaakHw6/wJi4iBOWzuYNEelphtgB0wH1Xr8iaTVc+5b18aesbCsIjyd4OyLxGMHY7bwLh2MlQravawfl+6UFlPJ+aST36BeJpCfdtyr5JTaKapNF2HUEIEgT65Fa7Kb+O65YMDXTb3qW0nObec0Wa3er/JvIcafavoX+olHBbceUIwvM9DdxJWxXrgr6gwijFOcPVa54a1TlmoAlxOsw8p+j2CANJOXRfsa3XW/Q9oILSL8onRrG365HC0eRHq120WHSd3bju1yEWLvhWguQB0vXbMlZyIg29Rdhsxg4bDblNef7OMlh94AEzbBJiPqbAu5tY4TJbzQCoMkOil9erDTYYHqAhZjr4kx609Ako24i+Tnua80Tk96YK+6H2C2EJ4A/0wSd/v5CRHFYdM7wU334cxH2HxVFHpBstV/oHS6hMXvbt3krivFghe0SPQIJVCEtyWVJ7qOWBblauUFbfBxNVq/NQQjvJr7mOZX+5hXT5db72bMBL3bvfIS9DQBNXOkMTV/dMs9yehipEXMTNIcsJnHotHjCjHE7ux8BC1hLwN4HSgWxWustdHyWbLDjnxvDEOWNkPkY7M92H/sq4Y8bGalIRSXdAfjlluwxH9N+NehDmfI2AC812LuRVba4O82tSdOu+ZCN2VqVJGNBmR8TJSelgim4PUsz0ygIcf48EVAcUfChkEBTZVkm2mzJGBoZHsFBpK9qBViHCjNF2C9lX59wkpNKE5saREa2+qTAGV76i+mAZcyNfeXXxHU7YXMf086L80NDmnQmqabXCLc425VZeH6YrktV//ut7AFIeKP5gKj1jfZ49iUMo9SkgA+AInDuOT4ES0+ma3cwi219cZJci3xlpdOCNUHnyuRFPxIuhCc3WPoT6MQn+gp3QoYF9sYWwE6CGI8s1RqGFpGGIBPefSh+JG1sZfHDu42/f+aXRfFhiXza58Pvh5e1ZT+llWSVNbVMOpkPNl3w0pxv51L7VGfT2otQCkL+zxwFHE5gLgDbLKJv3GKh0Uueckh6jRxb8Wpdx4BPlBRx8/imMSSy+NVPmGPG5ApKg1FAOySltvtagO2nCVlx")
	p.Add("navUserAgent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	p.Add("appName", "aliyun_drive")
	p.Add("appEntrance", "web")
	p.Add("_csrf_token", loginConf.LoginForm.CsrfToken)
	p.Add("umidToken", loginConf.LoginForm.UmIdToken)
	p.Add("isMobile", strconv.FormatBool(loginConf.LoginForm.IsMobile))
	p.Add("lang", loginConf.LoginForm.Lang)
	p.Add("returnUrl", loginConf.LoginForm.ReturnUrl)
	p.Add("hsiz", loginConf.LoginForm.Hsiz)
	p.Add("fromSite", strconv.Itoa(loginConf.LoginForm.FromSite))
	p.Add("bizParams", loginConf.LoginForm.BizParams)

	resp, err := http.PostForm(ApiLogin, p)
	b, err := ResponseUnzip(resp)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(b))

}

func ApiPost(url string, body io.Reader, header map[string]string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := ResponseUnzip(resp)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func FileList(driveId, parentFileId string) ([]byte, error) {
	var p = make(map[string]interface{})
	p["all"] = false
	p["drive_id"] = driveId
	p["fields"] = "*"
	p["image_thumbnail_process"] = "image/resize,w_400/format,jpeg"
	p["image_url_process"] = "image/resize,w_1920/format,jpeg"
	p["limit"] = 100
	p["order_by"] = "updated_at"
	p["order_direction"] = "DESC"
	p["parent_file_id"] = parentFileId
	p["url_expire_sec"] = 1600
	p["video_thumbnail_process"] = "video/snapshot,t_1000,f_jpg,ar_auto,w_300"

	paramJson, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	b, err := ApiPost(ApiFileList, strings.NewReader(string(paramJson)), getHeader())
	if err != nil {
		return nil, err
	}

	return b, nil
}

func FileSearch(driveId string) ([]byte, error) {
	var p = make(map[string]interface{})
	p["drive_id"] = driveId
	p["image_thumbnail_process"] = "image/resize,w_400/format,jpeg"
	p["image_url_process"] = "image/resize,w_1920/format,jpeg"
	p["limit"] = 100
	p["order_by"] = "created_at DESC"
	p["query"] = "type = \"file\""
	p["video_thumbnail_process"] = "video/snapshot,t_1000,f_jpg,ar_auto,w_300"

	paramJson, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	b, err := ApiPost(ApiFileSearch, strings.NewReader(string(paramJson)), getHeader())
	if err != nil {
		return nil, err
	}
	return b, nil
}

func FileDownloadUrl(driveId, fileId string) ([]byte, error) {
	var p = make(map[string]interface{})
	p["drive_id"] = driveId
	p["file_id"] = fileId

	paramJson, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	b, err := ApiPost(ApiFileDownload, strings.NewReader(string(paramJson)), getHeader())
	if err != nil {
		return nil, err
	}
	return b, nil
}

func FileDownload(url string) {
	//
}

func CreateFolder(driveId, parentFileId, name string) ([]byte, error) {
	var p = make(map[string]interface{})
	p["check_name_mode"] = "refuse"
	p["drive_id"] = driveId
	p["name"] = name
	p["parent_file_id"] = parentFileId
	p["type"] = "folder"

	paramJson, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	b, err := ApiPost(ApiFileDownload, strings.NewReader(string(paramJson)), getHeader())
	if err != nil {
		return nil, err
	}
	return b, nil
}
