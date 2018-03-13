// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-03-12 21:14

package communityspider

import (
	"os"
	"log"
	"strconv"
	"strings"

	excel "github.com/360EntSecGroup-Skylar/excelize"
)

func fileCheck() error {
	f, err := os.Open(xlsxFile)
	defer f.Close()

	if err != nil {
		xlsx := excel.NewFile()
		err := xlsx.SaveAs(xlsxFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func OutputFile(index int, data HomeDetail) {
	if err := fileCheck(); err != nil {
		log.Println(err)
		os.Exit(4)
	}

	if file, err := excel.OpenFile(xlsxFile); err == nil {
		i := strconv.Itoa(index + 2)

		file.SetCellStr(fileSheet, "A"+i, data.Name)
		file.SetCellStr(fileSheet, "B"+i, data.UntilPrice)
		file.SetCellStr(fileSheet, "C"+i, data.TotalPrice)
		file.SetCellStr(fileSheet, "D"+i, data.Payment)
		file.SetCellStr(fileSheet, "E"+i, strings.Join(data.HouseStyle, ","))
		file.SetCellStr(fileSheet, "F"+i, data.Heating)
		file.SetCellStr(fileSheet, "G"+i, data.Water)
		file.SetCellStr(fileSheet, "H"+i, data.Electricity)
		file.SetCellStr(fileSheet, "I"+i, data.Property)
		file.SetCellStr(fileSheet, "J"+i, data.Traffic)
		file.SetCellStr(fileSheet, "K"+i, data.Education)
		file.SetCellStr(fileSheet, "L"+i, data.CarRadio)
		file.SetCellStr(fileSheet, "M"+i, data.Score)
		file.SetCellStr(fileSheet, "N"+i, data.Dev)
		file.SetCellStr(fileSheet, "O"+i, data.NewOpening)
		file.SetCellStr(fileSheet, "P"+i, data.SalesStatus)
		file.SetCellStr(fileSheet, "Q"+i, data.Bank)
		file.SetCellStr(fileSheet, "R"+i, data.Gift)
		file.SetCellStr(fileSheet, "S"+i, data.Address)
		file.SetCellStr(fileSheet, "T"+i, data.Tel)
		file.SetCellStr(fileSheet, "U"+i, data.Link)

		err := file.Save()
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Printf("can not output to xlsx file: %s", err)
		os.Exit(4)
	}
}