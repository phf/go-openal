package main

import "openal/al"
import "openal/alc"

import "fmt"

func main() {
	out := alc.OpenDevice("");
	fmt.Printf("%x\n", al.GetError());

	con := out.CreateContext();
	fmt.Printf("%x\n", al.GetError());

	con.MakeContextCurrent();
	fmt.Printf("%x\n", al.GetError());

	src := al.GenSource();
	src1 := al.GenSource();
	fmt.Printf("%x\n", al.GetError());
	fmt.Println(src);
	fmt.Println(src1);
	al.DeleteSource(src1);
	al.DeleteSource(src);
	fmt.Printf("%x\n", al.GetError());
}
