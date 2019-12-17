lgen
===================

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/budougumi0617/lgen/blob/master/LICENSE

## Description

## Synopsis
```
$ lgen -a get -m user -template /your/templates/directory -dist /your/project/root/directory
```

## Options

```
$ lgen -h
Usage of /var/folders/sy/ls4cfp216x774g54brzl67yw0000gn/T/go-build403725816/b001/exe/lgen:
  -a string
        action name
  -action string
        action name
  -d string
        output directory (default "./")
  -dist string
        output directory
  -m string
        model name
  -model string
        model name
  -t string
        templates directory (default "./templates")
  -template string
        templates directory
  -v    print version information and quit.
  -version
        print version information and quit.
```

## Installation

```
$ go get github.com/budougumi0617/lgen
```

Built binaries are available on gihub releases. https://github.com/budougumi0617/lgen/releases

### MacOS
If you want to install on MacOS, you can use Homebrew.
```
brew install budougumi0617/tap/lgen
```

## Contribution
1. Fork ([https://github.com/budougumi0617/lgen/fork](https://github.com/budougumi0617/lgen/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request


## License

[MIT](https://github.com/budougumi0617/lgen/blob/master/LICENSE)

## Author
[budougumi0617](https://github.com/budougumi0617)

