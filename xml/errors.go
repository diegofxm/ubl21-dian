package xml

import "errors"

var (
	ErrTemplateNotFound = errors.New("template not found")
	ErrInvalidTemplate  = errors.New("invalid template")
	ErrRenderFailed     = errors.New("render failed")
)
