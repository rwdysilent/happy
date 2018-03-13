// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-03-11 15:52

package main

import (
	"log"
	"fmt"

	//"github.com/rwdysilent/gotools/whttp"
	"github.com/PuerkitoBio/goquery"
	excel "github.com/360EntSecGroup-Skylar/excelize"
	//"strconv"
	"strings"
	"strconv"
)

type HomeList struct {
	HomeLink []HomeLink
}

type HomeLink struct {
	Name string
	Link string
}

type HomeDetail struct {
	Name        string
	UntilPrice  string
	TotalPrice  string
	Payment     string
	HouseStyle  []string
	Heating     string
	Water       string
	Electricity string
	Property    string
	Traffic     string
	Education   string
	CarRadio    string
	Score       string
	Dev         string
	NewOpening  string
	SalesStatus string
	Bank        string
	Gift        string
	Address     string
	Tel         string
	Link        string
}

var (
	//筛选条件：住宅，在售/待售，均价15000以下
	// 雁塔
	baseUrl = "https://xa.fang.lianjia.com/"
	homeLink  = new(HomeLink)
	homeList  = new(HomeList)
	//ytUrl   = "https://xa.fang.lianjia.com/loupan/yanta/bap0eap15000nht1nhs1nhs3/#yanta"
	//ytYFUrl = "https://xa.fang.lianjia.com/loupan/yanta/nho0nho1nht1nhs1nhs3nhtt17/"

	ytInfo = struct {
		url string
		file string
		fileSheet string
	}{
		//15000
		//"https://xa.fang.lianjia.com/loupan/yanta/bap0eap15000nht1nhs1nhs3/#yanta",
		//18000
		"https://xa.fang.lianjia.com/loupan/yanta/bap0eap18000nht1nhs1nhs3pg3/",
		"/Users/pfwu/Documents/西安雁塔小区信息.xlsx",
		"18000-3",
	}
)

func LoadUrl(url string) (*goquery.Document) {
	if doc, err := goquery.NewDocument(url); err != nil {
		log.Printf("can not open url: %s", err)
	} else {
		return doc
	}
	return nil
}

func GetHomeList(url string) HomeList {
	doc := LoadUrl(url)

	doc.Find(".resblock-list-wrapper .resblock-list").Each(func(i int, s *goquery.Selection) {
		tag := s.Find("a")
		if name, ok := tag.Attr("title"); ok {
			homeLink.Name = name
			//fmt.Println(name)
		}
		if link, ok := tag.Attr("href"); ok {
			homeLink.Link = baseUrl + link
		}
		homeList.HomeLink = append(homeList.HomeLink, *homeLink)
	})

	//fmt.Printf("%s", homeList)
	return *homeList
}

func GetHomeDetail(doc *goquery.Document) HomeDetail {
	homeDetail := new(HomeDetail)
	//单价
	if price, err := doc.Find("div.box-left-top").Find("p.jiage").
		Find("span.junjia").Html(); err == nil {
		homeDetail.UntilPrice = price
	}

	//最新开盘时间
	homeDetail.NewOpening = doc.Find("div.bottom-info").Find("p.when").
		Find("span").Last().Text()

	//销售状态
	homeDetail.SalesStatus = doc.Find("div.dynamic-wrap-block").First().Find("a").Text()

	tag := doc.Find("div.views").Find("span.btn-s")
	//楼盘名
	if name, ok := tag.Attr("data-name"); ok {
		homeDetail.Name = name
	}
	//地址
	if addr, ok := tag.Attr("data-address"); ok {
		homeDetail.Address = addr
	}
	//户型
	doc.Find("div.houselist").Find("p.p1").Each(func(i int, s *goquery.Selection) {
		if h, err := s.Find("span").First().Html(); err == nil {
			homeDetail.HouseStyle = append(homeDetail.HouseStyle, h)
		}
	})
	//评分
	if s, err := doc.Find("div.totalscore").Find("span.score").Html(); err == nil {
		homeDetail.Score = s
	}
	//楼盘详情
	//查看索引代码注释
	//tag = doc.Find("div.box-loupan").Find("span.label")
	//tag.Each(func(i int, s *goquery.Selection) {
	//	fmt.Println(i, s.Text())
	//})
	tag = doc.Find("div.box-loupan").Find("span.label-val")
	homeDetail.Dev = tag.Eq(2).Text()
	homeDetail.CarRadio = tag.Eq(10).Text() + "/" + strings.Trim(tag.Eq(12).Text(), " ")
	homeDetail.Property = tag.Eq(11).Text()
	homeDetail.Heating = tag.Eq(13).Text()
	homeDetail.Water = tag.Eq(14).Text()
	homeDetail.Electricity = tag.Eq(15).Text()

	//客服电话
	if tel, ok := doc.Find("div.panel-tab").Find("span").Attr("data-all"); ok {
		homeDetail.Tel = tel
	}

	return *homeDetail
}

func GetDetail(url string) HomeDetail {
	doc := LoadUrl(url)
	d := GetHomeDetail(doc)
	return d
}

func ExcelFile() *excel.File {
	if xls, err := excel.OpenFile(ytInfo.file); err != nil {
		return xls
	}
	return nil
}

func writeFile(file *excel.File, index int, data HomeDetail) {
	file, _ = excel.OpenFile(ytInfo.file)
	//i := "5"
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
		fmt.Println(err)
	}
}

func main() {
	list := GetHomeList(ytInfo.url)
	f := ExcelFile()
	for i, v := range list.HomeLink {
		info := GetDetail(v.Link)
		info.Link = v.Link
		fmt.Printf("%+v", info)
		writeFile(f, i, info)
		//GetDetail("https://xa.fang.lianjia.com/loupan/p_tlyclmaaswu/#house-details")
	}
}
