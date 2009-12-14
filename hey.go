package main

import "openal"
import "time"

func main() {
	openal.Init();
	helloBuffer := openal.CreateBufferHelloWorld();
	helloSource := openal.GenSource();
	helloSource.SetAttr(openal.AlBuffer, helloBuffer);
	helloSource.Play();
	time.Sleep(1*1000*1000*1000);
	openal.Exit();
}
