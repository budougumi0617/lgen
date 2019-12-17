package controller

import (
  "net/http"
)

// New{{ .Action | title }}{{ .Model | title }}Controller create a new instance of {{ .Action | title }}{{ .Model | title }}Controller.
func New{{ .Action | title }}{{ .Model | title }}Controller(
  u usecase.{{ .Action | title }}{{ .Model | title }}UseCase,
) {{ .Action | title }}{{ .Model | title }}Controller {
  return {{ .Action }}{{ .Model | title }}Controller{
    usercase: u,
  }
}

// {{ .Action | title }}{{ .Model | title }}Controller handle the request to {{ .Action | title }} {{ .Model | title }} records.
type {{ .Action | title }}{{ .Model | title }}Controller struct {
  usercase usercase.{{ .Action | title }}{{ .Model | title }}Usecase
}

// Handler is the http handler to handle request to GET adjustment records.
func (c *{{ .Action | title }}{{ .Model | title }}Controller) Handler(w http.ResponseWriter, r *http.Request) {
  // Need to implement handle request manualy
  in := usercase.{{ .Action | title }}{{ .Model | title }}Input{
    // Need to implement handle request manualy
  }
  result, err := c.usecase.Run(in)
  if err != nil{
    http.Error(w, http.StatusText(http.StatusInternalServerError),
        http.StatusInternalServerError)
  }
  // Need to implement send response used by result manualy
}
