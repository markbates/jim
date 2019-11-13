package jim

import "io"

// Iny returns an io.Reader that represents the stdin
type Iny interface {
	Stdin() io.Reader
}

// Outy returns an io.Writer that represents the stdout
type Outy interface {
	Stdout() io.Writer
}

// Erry returns an io.Writer that represents the stderr
type Erry interface {
	Stderr() io.Writer
}
