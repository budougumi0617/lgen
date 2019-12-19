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
