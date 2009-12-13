package main

import "openal"
import "fmt"

func main() {
	dev := openal.OpenDevice("");
	fmt.Printf("dev: %s\n", dev);
	ok := openal.CloseDevice(dev);
	fmt.Printf("ok: %s\n", ok);
}
