// Copyright 2018 TED@Sogou, Inc. All rights reserved.
//
// @Author: wupengfei@sogou-inc.com
// @Date: 2018-03-13 16:18

package communityspider

import (
	"flag"
	"os"
)

var (
	o         bool
	xlsxFile  string
	fileSheet string
)

func usage() {
	flag.BoolVar(&o, "o", false, "do you need to write xls file?")
	flag.StringVar(&xlsxFile, "file", "tmp.xlsx", "need excel file location")
	flag.StringVar(&fileSheet, "sheet", "Sheet1", "xls file work sheet name")
	flag.Parse()
}

func init() {
	usage()
	if len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(0)
	}
}