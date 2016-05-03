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

	"github.com/codegangsta/cli"
)

const (
	cliName        = "oui"
	cliDescription = "search vender information for OUI(Organizationally Unique Identifier)"
	version        = "v0.2.0-dev"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	app := cli.NewApp()
	app.Name = cliName
	app.Usage = cliDescription
	app.Version = version
	app.Flags = []cli.Flag{
		cli.HelpFlag,
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 && !sc.Scan() {
			cli.ShowAppHelp(c)
		} else {
			data := InitMalData()

			var mac string
			if sc.Text() == "" {
				mac = c.Args()[0]
			} else {
				mac = sc.Text()
			}

			mac = strings.Replace(mac, ":", "", -1)
			mac = strings.Replace(mac, "-", "", -1)
			for i := 0; i < len(data); i++ {
				if data[i].Hex == strings.ToUpper(mac[0:6]) {
					fmt.Println(data[i].OrgName)
					break
				}
			}
		}
		return nil
	}
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} <Address> [options]

VERSION:
   {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
	app.Run(os.Args)
}
