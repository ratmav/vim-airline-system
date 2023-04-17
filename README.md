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

### dependencies

* [zig](https://ziglang.org/learn/getting-started/#installing-zig)
* [rust and cargo](https://www.rust-lang.org/tools/install)
* [cargo-zigbuild](https://github.com/rust-cross/cargo-zigbuild)

### setup

```shell
# add zig toolchain.
$ cargo install cargo-zigbuild

# add rust compiler targets.
$ rustup target add x86_64-apple-darwin
$ rustup target add x86_64-unknown-linux-gnu
$ rustup target add x86_64-pc-windows-msvc
```

### TODO: update single platform build instructions
```shell
# move into rust source directory.
$ cd rust

# build for each platform
$ cargo zigbuild --release --target x86_64-apple-darwin
$ cargo zigbuild --release --target x86_64-unknown-linux-gnu
$ cargo zigbuild --release --target x86_64-pc-windows-msvc

# update binaries used by airline plugin.
# TODO: $ cp ./target/x86_64-apple-darwin/release/vasr ../autoload/airline/extensions/bin/vasr_darwin
# TODO: $ cp ./target/x86_64-unknown-linux-gnu/release/vasr ../autoload/airline/extensions/bin/vasr_linux
# TODO: $ cp ./target/x86_64-pc-windows-msvc/release/vasr.exe ../autoload/airline/extensions/bin/vasr_windows.exe
```

## acknowledgements

* [sys-info-rs](https://github.com/FillZpp/sys-info-rs) made this light work indeed.
* [cargo-zigbuild](https://github.com/rust-cross/cargo-zigbuild) was very helpful for crosscompilation.
* [vim-airline](https://github.com/vim-airline/vim-airline): hard to extend something that doesn't exist.
* [vim-airline-clock](https://github.com/enricobacis/vim-airline-clock) is a great example of a vim-airline extension.
* [vim-ctrlspace](https://github.com/vim-ctrlspace/vim-ctrlspace) is a great example of using golang to extend vim plugins.
* [steve losh's](https://stevelosh.com/) book, [learn vimscript the hard way](https://learnvimscriptthehardway.stevelosh.com/), was a great and useful read.
* and of course...the [neovim](https://neovim.io/), [rust](https://www.rust-lang.org/), and [zig](https://ziglang.org/) maintainers.
