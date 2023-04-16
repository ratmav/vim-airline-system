vim-airline-system
==================

[vim-airline](https://github.com/vim-airline/vim-airline) system information extension.

**note**: _replaces_ the vim-airline "z" section.

![screenshot](https://github.com/ratmav/vim-airline-system/blob/master/screenshot.png?raw=true)

## supported platforms

* 64-bit linux
* 64-bit macos
* 64-bit windows

## installation

use git or your plugin manager of choice to install vim-airline-system.

## development

### build

**note**: rustc v1.57.0

  ```shell
  $ cd rust
  $ rustup target add x86_64-apple-darwin
  $ rustup target add x84_64-unknown-linux-gnu
  $ rustup target add x86_64-pc-windows-msvc
  $ cargo build --release
  # windows build
  $ cp ./target/relase/vasr.exe ../autoload/airline/extensions/bin/vasr_windows.exe
  ```

## acknowledgements

* [sys-info-rs](https://github.com/FillZpp/sys-info-rs) made this light work indeed.
* [vim-airline](https://github.com/vim-airline/vim-airline): hard to extend something that doesn't exist.
* [vim-airline-clock](https://github.com/enricobacis/vim-airline-clock) is a great example of a vim-airline extension.
* [vim-ctrlspace](https://github.com/vim-ctrlspace/vim-ctrlspace) is a great example of using golang to extend vim plugins.
* [steve losh's](https://stevelosh.com/) book, [learn vimscript the hard way](https://learnvimscriptthehardway.stevelosh.com/), was a great and useful read.
