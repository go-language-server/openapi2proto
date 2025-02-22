package openapi2proto

import (
	"go.lsp.dev/openapi2proto/compiler"
	"go.lsp.dev/openapi2proto/internal/option"
	"go.lsp.dev/openapi2proto/protobuf"
)

const (
	optkeyEncoderOptions  = "protobuf-encoder-options"
	optkeyCompilerOptions = "protobuf-compiler-options"
)

// Option is used to pass options to several methods
type Option option.Option

// WithEncoderOptions allows you to specify a list of
// options to `Transpile`, which gets passed to the
// protobuf.Encoder object.
func WithEncoderOptions(options ...protobuf.Option) Option {
	return option.New(optkeyEncoderOptions, options)
}

// WithCompilerOptions allows you to specify a list of
// options to `Transpile`, which gets passed to the
// protobuf.Compile method
func WithCompilerOptions(options ...compiler.Option) Option {
	return option.New(optkeyCompilerOptions, options)
}
