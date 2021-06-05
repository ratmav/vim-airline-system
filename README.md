vim-airline-system
==================

[vim-airline](https://github.com/vim-airline/vim-airline) system information extension.

**note**: _replaces_ the vim-airline "z" section.

![screenshot](https://github.com/ratmav/vim-airline-system/blob/master/screenshot.png?raw=true)

## supported platforms

* 64-bit macos
* 64-bit linux

## installation

use git or your plugin manager of choice to install vim-airline-system.

## development

### build

**note**: go v1.15

in the below example, `platform_name` would be "darwin" or "linux".

  ```shell
  $ cd go
  $ go build -o ../autoload/airline/extensions/bin/platform_name --mod=vendor
  ```

## acknowledgements

* [gopsutil](https://github.com/shirou/gopsutil/) made this light work indeed.
* [vim-airline](https://github.com/vim-airline/vim-airline): hard to extend something that doesn't exist.
* [vim-airline-clock](https://github.com/enricobacis/vim-airline-clock) is a great example of a vim-airline extension.
* [vim-ctrlspace](https://github.com/vim-ctrlspace/vim-ctrlspace) is a great example of using golang to extend vim plugins.
* [steve losh's](https://stevelosh.com/) book, [learn vimscript the hard way](https://learnvimscriptthehardway.stevelosh.com/), was a great and useful read.
