package main

import "openal"
import "fmt"

func main() {
	dev := openal.OpenDevice("");
	fmt.Printf("dev: %s\n", dev);
	err := dev.GetError();
	fmt.Printf("err: %s\n", err);
	ok := dev.CloseDevice();
	fmt.Printf("ok: %s\n", ok);

	mic := openal.CaptureOpenDevice("", 16000, openal.AlFormatMono8, 16000);
	fmt.Printf("mic: %s\n", mic);
	err = mic.GetError();
	fmt.Printf("err: %s\n", err);
	ok = mic.CaptureCloseDevice();
	fmt.Printf("ok: %s\n", ok);
}
