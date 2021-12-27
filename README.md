vim-airline-system
==================

[vim-airline](https://github.com/vim-airline/vim-airline) system information extension.

**note**: _replaces_ the vim-airline "z" section.

![screenshot](https://github.com/ratmav/vim-airline-system/blob/master/screenshot.png?raw=true)

## supported platforms

* 64-bit
  * linux
  * macos
  * windows

## installation

use git or your plugin manager of choice to install vim-airline-system.

## development

### build

**note**: go v1.17

#### *nix host

```bash
cd go

set GOOS=<target os/> # linux, darwin, or windows.
set GOARCH=<target arch/> # amd64.

# if target is *nix:
go build -o ../autoload/airline/extensions/bin/$GOOS --mod=vendor

# if target is windows:
go build -o ../autoload/airline/extensions/bin/$GOOS.exe --mod=vendor
```

#### windows host

```powershell
cd go

$env:GOOS="<target os/>" # linux, darwin, or windows.
$env:GOARCH="<target arch/>" # amd64.

# if target is *nix:
go build -o ..\autoload\airline\extensions\bin\${env:GOOS} --mod=vendor

# if target is windows:
go build -o ..\autoload\airline\extensions\bin\${env:GOOS}.exe --mod=vendor
```

**note**: the `.exe` extension is necessary, because apparently [magic bytes](https://en.wikipedia.org/wiki/List_of_file_signatures) is only a thing on *nix platforms.

## acknowledgements

* [gopsutil](https://github.com/shirou/gopsutil/) made this light work indeed.
* [vim-airline](https://github.com/vim-airline/vim-airline): hard to extend something that doesn't exist.
* [vim-airline-clock](https://github.com/enricobacis/vim-airline-clock) is a great example of a vim-airline extension.
* [vim-ctrlspace](https://github.com/vim-ctrlspace/vim-ctrlspace) is a great example of using golang to extend vim plugins.
* [steve losh's](https://stevelosh.com/) book, [learn vimscript the hard way](https://learnvimscriptthehardway.stevelosh.com/), was a great and useful read.
