# lccli

  lccli is a simple command-line tool that serve as a alternative to the official Lunar Client launcher.

### Table of contents
  * [🎉 Explanation](#explanation)
  * [🤔 Usage and Examples](#usage-and-examples)
  * [🏗️  Building](#building)

## Explanation 
  The first question you may ask is: "*why should I use this tool instead of the official launcher?*", and the
  answer for that question is really simple: the official launcher is Electron-based and is really resource-hungry
  in CPU, so it's not suitable for lower end devices or for people who want the maximum performance.

## Usage and Examples
  This tool is intended to be simple, so it just take one argument, the rest of the arguments are ignored and a few
  settings are used in a `lccli.toml` file, which should be located in the same path as the executable. For example,
  if I want to run the `1.7` version, I first need to download the version using the official launcher (yes, I know
  that it is a tricky way to do it, but until I don't find a better way this will be the only option). Then I can just
  run `lccli 1.7` in the terminal, simple as that and the game will be launched just fine. **This tool is untested on 
  Windows, so if it does not work, please make sure to open a issue.**

## Building 
  The building process for this tool is really simple and take just a few seconds or minutes, since it's a small Go
  project. You just need to install the `go` package with the package manager of your choice, and then run the following
  command: `go build -ldflags "-s -w"`.
