// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-03-12 21:14

package communityspider

import (
	"log"
	"strconv"
	"strings"

	excel "github.com/360EntSecGroup-Skylar/excelize"
	"os"
)

func ExcelFile() *excel.File {
	if err := fileCheck(); err != nil{
		log.Printf("can not create xlsx file: %s", err)
		os.Exit(1)
	}

	if xls, err := excel.OpenFile(xlsxFile); err == nil {
		return xls
	}
	return nil
}

func OutputFile(index int, data HomeDetail) {
	//file, _ = excel.OpenFile(ytInfo.file)
	file := ExcelFile()
	i := strconv.Itoa(index + 2)

	file.SetCellStr(ytInfo.fileSheet, "A"+i, data.Name)
	file.SetCellStr(ytInfo.fileSheet, "B"+i, data.UntilPrice)
	file.SetCellStr(ytInfo.fileSheet, "C"+i, data.TotalPrice)
	file.SetCellStr(ytInfo.fileSheet, "D"+i, data.Payment)
	file.SetCellStr(ytInfo.fileSheet, "E"+i, strings.Join(data.HouseStyle, ","))
	file.SetCellStr(ytInfo.fileSheet, "F"+i, data.Heating)
	file.SetCellStr(ytInfo.fileSheet, "G"+i, data.Water)
	file.SetCellStr(ytInfo.fileSheet, "H"+i, data.Electricity)
	file.SetCellStr(ytInfo.fileSheet, "I"+i, data.Property)
	file.SetCellStr(ytInfo.fileSheet, "J"+i, data.Traffic)
	file.SetCellStr(ytInfo.fileSheet, "K"+i, data.Education)
	file.SetCellStr(ytInfo.fileSheet, "L"+i, data.CarRadio)
	file.SetCellStr(ytInfo.fileSheet, "M"+i, data.Score)
	file.SetCellStr(ytInfo.fileSheet, "N"+i, data.Dev)
	file.SetCellStr(ytInfo.fileSheet, "O"+i, data.NewOpening)
	file.SetCellStr(ytInfo.fileSheet, "P"+i, data.SalesStatus)
	file.SetCellStr(ytInfo.fileSheet, "Q"+i, data.Bank)
	file.SetCellStr(ytInfo.fileSheet, "R"+i, data.Gift)
	file.SetCellStr(ytInfo.fileSheet, "S"+i, data.Address)
	file.SetCellStr(ytInfo.fileSheet, "T"+i, data.Tel)
	file.SetCellStr(ytInfo.fileSheet, "U"+i, data.Link)

	err := file.Save()
	if err != nil {
		log.Println(err)
	}
}

//func main() {
//    xlsx, err := excelize.OpenFile("/Users/pfwu/Documents/西安雁塔小区信息.xlsx")
//    if err != nil {
//        fmt.Println(err)
//    }
//    //xlsx.SetCellStr("work1", "A1", "Hello")
//    //err = xlsx.Save()
//    cell := xlsx.GetCellValue("work1", "A2")
//    fmt.Println(cell)
//}
