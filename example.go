package main

import "openal"
import "fmt"
import "time"

func main() {
	dev := openal.OpenDevice("");
	fmt.Printf("dev: %s\n", dev);

	err := dev.GetError();
	fmt.Printf("err: %s\n", err);

	ok := dev.CloseDevice();
	fmt.Printf("ok: %s\n", ok);


	mic := openal.CaptureOpenDevice("", 8000, openal.AlFormatMono8, 16000);
	fmt.Printf("mic: %s\n", mic);

	err = mic.GetError();
	fmt.Printf("err: %s\n", err);

	mic.CaptureStart();
	fmt.Println("capture started!");

	err = mic.GetError();
	fmt.Printf("err: %s\n", err);

	fmt.Println(time.Sleep(1*1000*1000*1000));

	smp := mic.GetInteger(openal.AlcCaptureSamples);
	fmt.Printf("smp: %s\n", smp);

	mic.CaptureStop();
	fmt.Println("capture stopped!");

	err = mic.GetError();
	fmt.Printf("err: %s\n", err);

	ok = mic.CaptureCloseDevice();
	fmt.Printf("ok: %s\n", ok);
}
