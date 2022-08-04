package database

import "io"

type Client interface {
	io.Closer
}
