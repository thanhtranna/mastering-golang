package main

import (
	v1 "./myModule"
	v2 "./myModule/v2"
)

func main() {
	v1.Version()
	v2.Version()
}
