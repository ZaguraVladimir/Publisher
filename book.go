package main

import "fmt"

type book struct {
	name  string
	parts map[int]fmt.Stringer
}
