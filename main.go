package main

import (
	examples "github.com/diskfs/go-diskfs/examples"
	"os"
)

const (
	img = "/tmp/disk.img"
)

func main() {
	os.Remove(img)
	examples.CreateEfi(img)
}
