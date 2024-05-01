# bln
A Unix-like ln command with a better interface

I regularly use the Unix `ln` command, mainly to create symbolic links for configuration files that are supposed to live in folders
scattered under the home directory, but that actually live grouped in a "dotfiles" folder that I keep under version control.

**The problem: if I want to create a link A that points to B, I can never remember if I'm supposed to type `ln -s A B` or `ln -s B A`.**

After studying the [*flags over args* discussion](https://clig.dev/#arguments-and-flags), I thought that I could develop a 
project t develop a better interface for `ln`. Enter `bln`:

```shell
bln -s --link-name A --link-to B
```

This makes is clear: file `A` is a link that points to an existing file `B`.

`bln` does not support all of `ln` flags, but it does accept the `-s` (or `--symbolic`) flag to create symbolic links;
if omitted, hard links are created (which, among other things, cannot be applied to directories).

Run `bln --help` for full instructions.

To be clear: this is a toy project mainly to help me improve my Go skills - but it is coded carefully,
and solves a real problem!

## Installation


You need the [Go](https://go.dev/dl/) toolchain installed. 

```shell
go install https://github.com/fabiofortkamp\bln@main
```

This will install the command under `$GOBIN` (which should be on your `$PATH`).
