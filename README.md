oui
====
[![license](https://img.shields.io/github/license/ayatk/oui.svg?maxAge=2592000?style=flat-square)][license]
[![GitHub release](https://img.shields.io/github/release/ayatk/oui.svg?maxAge=2592000?style=flat-square)][download]  
Look up vendor information from OUI (Organizationally Unique Identifier)

## Usage

    $ oui <MAC Address>

or use pipe

    $ echo <MAC Address> | oui

e.g.

    $ oui 0:0:2:3a:3e:1
    XEROX CORPORATION

`verbose` flag is show detailed information

    $ oui -v 0:0:2:3a:3e:1
    OUI/MA-L :      00:00:02
    Organization :  XEROX CORPORATION
    Address :       M/S 105-50C WEBSTER NY US 14580

## Installation
Use `go get` command:

    $ go get github.com/ayatk/oui

or [Download binary][download]


## License
### OUI (Organizationally Unique Identifier) Data
Copyright 2016 [IEEE][ieee] - All rights reserved.

### Source Code
Released under the [MIT License][license] /
Copyright (c) 2016 Aya Tokikaze.

[ieee]: http://www.ieee.org/
[license]: https://github.com/ayatk/oui/blob/master/LICENSE
[download]: https://github.com/ayatk/oui/releases
