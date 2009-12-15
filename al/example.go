package main

import "openal/al"
import "fmt"

func main() {
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
