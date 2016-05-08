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
	"bufio"
	"fmt"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

const (
	cliName        = "oui"
	cliDescription = "search vender information for OUI(Organizationally Unique Identifier)"
	version        = "v0.3.0-dev"
)

// MAC stract is MAC Address date format
type MAC struct {
	Registry   string
	Hex        string
	OrgName    string
	OrgAddress string
}

// list options
type Options struct {
	Verbose bool `short:"v" long:"verbose" description:"print detailed information"`
	Input   bool `short:"i" long:"input" description:"use standard input"`
	Version bool `long:"version" description:"print oui version"`
	Org     bool `short:"o" long:"org" description:"search organization to OUI"`
}

func main() {

	var opt Options
	var mac string
	var res MAC

	sc := bufio.NewScanner(os.Stdin)
	f := flags.NewParser(&opt, flags.Default)
	f.Name = "oui"
	f.Usage = "<ADDRESS> [OPTION]"

	args, _ := f.Parse()

	if opt.Version {
		fmt.Printf("%s version %s\n", cliName, version)
		os.Exit(0)
	}

	if len(args) == 0 {
		if os.Args[1] == "-h" {
			os.Exit(0)
		}
		if !opt.Input {
			f.WriteHelp(os.Stdout)
			os.Exit(1)
		}
	}

	data := InitMalData()

	if opt.Input && sc.Scan() {
		mac = sc.Text()
	} else {
		mac = args[0]
	}

	mac = strings.Replace(mac, ":", "", -1)
	mac = strings.Replace(mac, "-", "", -1)

	if opt.Org {
		for i := range data {
			if strings.Contains(strings.ToUpper(data[i].OrgName), strings.ToUpper(args[0])) {
				fmt.Printf("%s:%s:%s  %s\n", data[i].Hex[0:2], data[i].Hex[2:4], data[i].Hex[4:6], data[i].OrgName)
			}
		}
		os.Exit(0)
	}

	for i := range data {
		if data[i].Hex == strings.ToUpper(mac[0:6]) {
			res = data[i]
			break
		}
	}

	if opt.Verbose {
		split := []string{mac[0:2], mac[2:4], mac[4:6]}
		fmt.Printf("OUI/%s :      %s\n", res.Registry, strings.Join(split, "-"))
		fmt.Printf("Organization :  %s\n", res.OrgName)
		fmt.Printf("Address :       %s\n", res.OrgAddress)
	} else {
		fmt.Println(res.OrgName)
	}
}
