package openapi2proto // go.lsp.dev/openapi2proto

import (
	"fmt"
	"io"

	"go.lsp.dev/openapi2proto/compiler"
	"go.lsp.dev/openapi2proto/openapi"
	"go.lsp.dev/openapi2proto/protobuf"
)

// Transpile is a convenience function that takes an OpenAPI
// spec file and transpiles it into a Protocol Buffers v3 declaration,
// which is written to `dst`.
//
// Options to the compiler and encoder can be passed using
// `WithCompilerOptions` and `WithEncoderOptions`, respectively
//
// For more control, use `openapi`, `compiler`, and `protobuf`
// packages directly.
func Transpile(dst io.Writer, srcFn string, options ...Option) error {
	var encoderOptions []protobuf.Option
	var compilerOptions []compiler.Option

	for _, o := range options {
		switch o.Name() {
		case optkeyEncoderOptions:
			encoderOptions = o.Value().([]protobuf.Option)
		case optkeyCompilerOptions:
			compilerOptions = o.Value().([]compiler.Option)
		}
	}

	s, err := openapi.LoadFile(srcFn)
	if err != nil {
		return fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}

	p, err := compiler.Compile(s, compilerOptions...)
	if err != nil {
		return fmt.Errorf("failed to compile OpenAPI spec to Protocol buffers: %w", err)
	}

	if err := protobuf.NewEncoder(dst, encoderOptions...).Encode(p); err != nil {
		return fmt.Errorf("failed to encode protocol buffers to text: %w", err)
	}

	return nil
}
