oui
====
search vender information from OUI(Organizationally Unique Identifier)

## Usage

    $ oui <MAC Address>

e.g.

    $ oui 000000 # 00-00-00 or 00:00:00
    XEROX CORPORATION

Use standard input with `input` flag

    $ echo "00:00:00:00:00:00" | oui -i
    XEROX CORPORATION

`verbose` flag is show detailed information

    $ oui -v 00:00:00:00:00:00
    OUI/MA-L :      00-00-00
    Organization :  XEROX CORPORATION
    Address :       M/S 105-50C WEBSTER NY US 14580

## Installation
Use `go get` command:

    $ go get github.com/ayatk/oui

or Download binary


## License
### OUI(Organizationally Unique Identifier) Data
Copyright 2016 [IEEE][ieee] - All rights reserved.

### Source Code
Released under the [MIT License][license] /
Copyright (c) 2016 Aya Tokikaze.

[ieee]: http://www.ieee.org/
[license]: https://github.com/ayatk/oui/blob/master/LICENSE
