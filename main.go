// The MIT License (MIT)
//
// Copyright (c) 2016 Aya Tokikaze
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	cliName        = "oui"
	cliDescription = "Look up vendor information from OUI"
	version        = "v0.2.3"
)

var (
	verboseFlag bool
	versionFlag bool
)

// MAC struct is MAC Address date format
type MAC struct {
	Registry   string
	Hex        string
	OrgName    string
	OrgAddress string
}

func main() {
	var input []string
	var oui string

	flag.Usage = func() {
		fmt.Printf("%s : %s\n", cliName, cliDescription)
		flag.PrintDefaults()
	}

	flag.BoolVar(&verboseFlag, "v", false, "Print detailed information")
	flag.BoolVar(&versionFlag, "version", false, "Print oui version")
	flag.Parse()

	if versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}
	// パイプを使用するとき
	if !terminal.IsTerminal(0) {
		body, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		input = strings.Split(string(body), "\n")
	} else if len(flag.Args()) == 0 { // コマンド単体で使用されたとき
		fmt.Printf("%s : %s\n", cliName, cliDescription)
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(input) == 0 {
		input = flag.Args()
	}

	// ベンダの情報をインポート
	data := InitMalData()

	for _, m := range input {
		oui = ""
		mac := parseAddress(m)
		if len(mac) >= 3 {
			oui = strings.ToUpper(mac[0] + mac[1] + mac[2])
		}
		res, isFound := searchOUI(data, oui)
		if verboseFlag && isFound {
			split := []string{oui[0:2], oui[2:4], oui[4:6]}
			fmt.Printf("OUI/%s :      %s\n", res.Registry, strings.Join(split, ":"))
			fmt.Printf("Organization :  %s\n", res.OrgName)
			fmt.Printf("Address :       %s\n", res.OrgAddress)
		} else {
			fmt.Println(res.OrgName)
		}
	}
	// if opt.Org {
	// 	for i := range data {
	// 		if strings.Contains(strings.ToUpper(data[i].OrgName), strings.ToUpper(args[0])) {
	// 			fmt.Printf("%s:%s:%s  %s\n", data[i].Hex[0:2], data[i].Hex[2:4], data[i].Hex[4:6], data[i].OrgName)
	// 		}
	// 	}
	// 	os.Exit(0)
	// }
}
