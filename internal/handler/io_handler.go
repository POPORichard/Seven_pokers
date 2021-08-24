package handler

import (
	"Seven_pokers/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)
// ReadDataToModel 读取外部文件
func ReadDataToModel(path string)[]model.Data {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return nil
	}

	var inputData model.InputData
	err = json.Unmarshal(bytes, &inputData)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return nil
	}


	data := inputData.Matches
	return data
}