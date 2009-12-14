package main

import "openal"
import "time"
import "fmt"

func main() {
	openal.Init();
//	helloBuffer := openal.CreateBufferHelloWorld();
	helloBuffer := openal.CreateBufferFromFile("welcome.wav");
	helloSource := openal.GenSource();
	helloSource.SetAttr(openal.AlBuffer, helloBuffer);
	helloSource.Play();
	time.Sleep(1*1000*1000*1000);
	// just to test GenSources
	someSources := make([]uint, 10);
	fmt.Println(someSources);
	openal.GenSources(someSources);
	fmt.Println(someSources);
	openal.Exit();
}
