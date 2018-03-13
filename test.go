// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-03-13 15:39

package main

import (
	"log"
	excel "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx := excel.NewFile()
	//index := xlsx.NewSheet("work1")
	//xlsx.SetActiveSheet(index)

	err := xlsx.SaveAs("/Users/pfwu/Documents/tmp.xlsx")
	if err != nil {
		log.Println(err)
	}
}
