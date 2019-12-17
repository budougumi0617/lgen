package usecase

type {{ .Model | title}}{{ .Action | title }}Input struct{}

type {{ .Model | title}}{{ .Action | title }}Result struct{}

type {{ .Model | title}}{{ .Action | title }}Usecase interface {
  Run({{ .Model | title}}{{ .Action | title }}Input) ({{ .Model | title}}{{ .Action | title }}Result, error)
}

func New{{ .Model | title}}{{ .Action | title }}Usecase() {{ .Model | title}}{{ .Action | title }}Usecase {
  return &{{ .Model }}{{ .Action | title }}Usecase{
  }
}

type {{ .Model }}{{ .Action | title }}Usecase struct {}

func (u *{{ .Model }}{{ .Action | title }}Usecase) Run(
    in {{ .Model | title}}{{ .Action | title }}Input,
  ) ({{ .Model | title}}{{ .Action | title }}Result, error){
  // Need to implement usercase logic
  return {{ .Model | title}}{{ .Action | title }}Result{
    // Need to build result
  }
}
