// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-03-12 21:14

package main

import (
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
    xlsx, err := excelize.OpenFile("/Users/pfwu/Documents/西安雁塔小区信息.xlsx")
    if err != nil {
        fmt.Println(err)
    }
    //xlsx.SetCellStr("work1", "A1", "Hello")
    //err = xlsx.Save()
    cell := xlsx.GetCellValue("work1", "A2")
    fmt.Println(cell)
}
