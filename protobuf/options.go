package protobuf

import "github.com/NYTimes/openapi2proto/internal/option"

const (
	optkeyIndent              = "indent"
	optkeyAutogenerateComment = "autogenerate-message"
)

// WithIndent creates a new Option to control the indentation
// for the encoded definition
func WithIndent(s string) Option {
	return option.New(optkeyIndent, s)
}

// WithAutogenerateComment creates a new Option to add 'DO NOT MODIFY' message at the
// head of the generated proto file
func WithAutogeneratedComment(b bool) Option {
	return option.New(optkeyAutogenerateComment, b)
}