package main

import (
	"fmt"
	"github.com/speps/go-hashids"
	)


func main() {
	// woYAJb
	hd := hashids.NewData()
	hd.Salt = "1766"
	hd.MinLength = 6
	hd.Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123457890"

	h, _:= hashids.NewWithData(hd)
	numbers, _ := h.DecodeWithError("woYAJb")
	fmt.Print(numbers)
}
