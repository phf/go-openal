package main

import "openal/al"
import "openal/alc"

import "time"
import "fmt"

func main() {
	out := alc.OpenDevice("")
	fmt.Printf("%x\n", out.GetError())
	con := out.CreateContext()
	fmt.Printf("%x\n", out.GetError())
	con.Activate()
	fmt.Printf("%x\n", out.GetError())

	in := alc.CaptureOpenDevice("", 8000, al.FormatMono16, 16000)
	fmt.Printf("%x\n", in.GetError())
	in.CaptureStart();
	fmt.Printf("%x\n", in.GetError())

	time.Sleep(1*1000*1000*1000)

	in.CaptureStop()
	fmt.Printf("%x\n", in.GetError())

	n := in.GetInteger(alc.CaptureSamples)
	fmt.Printf("n: %s\n", n)
	fmt.Printf("%x\n", in.GetError())

	raw := in.CaptureSamples(uint32(n)) // TODO get rid of cast
	fmt.Printf("raw: %v\n", raw)
	fmt.Printf("%x\n", in.GetError())

	buf := al.NewBuffer()
	fmt.Printf("%x\n", al.GetError())

	buf.SetData(al.FormatMono16, raw, 8000)
	fmt.Printf("%x\n", al.GetError())

	src := al.NewSource()
	fmt.Printf("%x\n", al.GetError())

	src.SetBuffer(buf)
	fmt.Printf("%x\n", al.GetError())

	src.Play()
	fmt.Printf("%x\n", al.GetError())

	time.Sleep(1*1000*1000*1000)
}
