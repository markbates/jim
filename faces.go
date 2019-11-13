package jim

import "io"

type Iny interface {
	Stdin() io.Reader
}

type Outy interface {
	Stdout() io.Writer
}

type Erry interface {
	Stderr() io.Writer
}
