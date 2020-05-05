package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// 正则处理提取
func infoFormatter(rawData string) ([]string, []string) {
	if rawData == "" {
		return nil, nil
	}
	var releasesNumber []string
	var releasesDate []string
	r, _ := regexp.Compile(`<ol class="list-row-container menu">[\s\S]+</ol>`)
	olData := r.FindAllStringSubmatch(rawData, -1)
	r, _ = regexp.Compile(`<span class="release-number"><a[^>]*>([^<]*)</a></span>`)
	releaseNum := r.FindAllStringSubmatch(olData[0][0], -1)
	for _, value := range releaseNum {
		releasesNumber = append(releasesNumber, strings.Replace(value[1], " ", "_", -1))
	}
	r, _ = regexp.Compile(`<span class="release-date">([^<]*)</span>`)
	releaseDate := r.FindAllStringSubmatch(olData[0][0], -1)
	for _, value := range releaseDate {
		releasesDate = append(releasesDate, value[1])
	}
	return releasesNumber, releasesDate
}

// 获取python版本
func getAllPythonVersions() string {
	res, err := http.Get("https://www.python.org/downloads/")
	if err != nil {
		fmt.Printf("请求版本错误: %v\n", err)
		return ""
	}
	result, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		fmt.Printf("请求版本错误: %v\n", err)
		return ""
	}
	return string(result)
}

// 返回本地版本
func GetLocalPythonVersions() ([]string, error) {
	pythonDir := viper.Get("PYVM_HOME").(string) + "/pythons"
	dirs, err := GetLocalDirs(pythonDir)
	if err != nil {
		return nil, err
	}
	return dirs, nil
}

// 返回版本信息
func BackPythonVersions() ([]string, []string) {
	return infoFormatter(getAllPythonVersions())
}


// 版本号提取
func BackVersionNum(version string) string {
	r, _ := regexp.Compile(`((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){2}`)
	return r.FindAllStringSubmatch(version, -1)[0][0]
}