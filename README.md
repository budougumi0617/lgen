lgen
===================

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/budougumi0617/lgen/blob/master/LICENSE

## Description

Generate boilerplates for layered architecture by your templates.

At first, we prepare layered structure and templates.

```bash
$ tree templates
templates
├── repositories
│   └── repository.go
├── controllers
│   └── controller.go
└── usercases
    └── usecase.go
```

For instance, the usecase template is below. templates are written with `text/template`.

```go
// templates/usercase/usecase.go
package usecase

type {{ .Action | title}}{{ .Model | title }}Input struct{}

type {{ .Action | title}}{{ .Model | title }}Result struct{}

type {{ .Action | title}}{{ .Model | title }}Usecase interface {
  Run({{ .Action | title}}{{ .Model | title }}Input) ({{ .Action | title}}{{ .Model | title }}Result, error)
}

func New{{ .Action | title}}{{ .Model | title }}Usecase() {{ .Action | title}}{{ .Model | title }}Usecase {
  return &{{ .Action }}{{ .Model | title }}Usecase{
  }
}

type {{ .Action }}{{ .Model | title }}Usecase struct {}

func (u *{{ .Action }}{{ .Model | title }}Usecase) Run(
    in {{ .Action | title}}{{ .Model | title }}Input,
  ) ({{ .Action | title}}{{ .Model | title }}Result, error){
  // Need to implement usercase logic
  return {{ .Action | title}}{{ .Model | title }}Result{
    // Need to build result
  }
}
```

Execute `lgen` with `Action` and `Model` strings.

```bash
$ lgen -action Get -model User -template ./testdata -dist myproduct
```

The generated directories and files are below.


```bash
$ tree myproduct
myproduct
├── controllers
│   └── get_user_controller.go
└── usercases
    └── get_user_usecase.go

2 directories, 2 files
```

The get_user_usercase.go is below. We are enabled to write miltiple files by a command.

```go
package usecase

type GetUserInput struct{}

type GetUserResult struct{}

type GetUserUsecase interface {
        Run(GetUserInput) (GetUserResult, error)
}

func NewGetUserUsecase() GetUserUsecase {
        return &getUserUsecase{}
}

type getUserUsecase struct{}

func (u *getUserUsecase) Run(
        in GetUserInput,
) (GetUserResult, error) {
        // Need to implement usercase logic
        return GetUserResult{
                // Need to build result
        }
}
```


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

