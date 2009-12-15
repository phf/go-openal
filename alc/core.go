// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// C-level binding for OpenAL's "alc" API.
//
// Please consider using the Go-level binding instead.
package alc

/*
#include <stdlib.h>

// It's sad but the OpenAL C API uses lots and lots of typedefs
// that require wrapper functions (using basic C types) for cgo
// to grok them. So there's a lot more C code here than I would
// like...

#include <AL/al.h>
#include <AL/alc.h>

// I keep all the alc.h prototypes here for now, for reference.
// They'll go away eventually. Those commented out are already
// accessible from Go.

ALCcontext *alcCreateContext( ALCdevice *device, const ALCint* attrlist );
ALCboolean alcMakeContextCurrent( ALCcontext *context );
void alcProcessContext( ALCcontext *context );
void alcSuspendContext( ALCcontext *context );
void alcDestroyContext( ALCcontext *context );
ALCcontext *alcGetCurrentContext( void );
ALCdevice *alcGetContextsDevice( ALCcontext *context );
// ALCdevice *alcOpenDevice( const ALCchar *devicename );
ALCdevice *walcOpenDevice(const char *devicename) {
	return alcOpenDevice(devicename);
}
// ALCboolean alcCloseDevice( ALCdevice *device );
// ALCenum alcGetError( ALCdevice *device );
ALCboolean alcIsExtensionPresent( ALCdevice *device, const ALCchar *extname );
void *alcGetProcAddress( ALCdevice *device, const ALCchar *funcname );
ALCenum alcGetEnumValue( ALCdevice *device, const ALCchar *enumname );
const ALCchar *alcGetString( ALCdevice *device, ALCenum param );
//void alcGetIntegerv( ALCdevice *device, ALCenum param, ALCsizei size, ALCint *data );
void walcGetIntegerv(ALCdevice *device, ALCenum param, ALCsizei size, void *data) {
	alcGetIntegerv(device, param, size, data);
}
// ALCdevice *alcCaptureOpenDevice( const ALCchar *devicename, ALCuint frequency, ALCenum format, ALCsizei buffersize );
ALCdevice *walcCaptureOpenDevice(const char *devicename, ALCuint frequency, ALCenum format, ALCsizei buffersize) {
	return alcCaptureOpenDevice(devicename, frequency, format, buffersize);
}
ALCboolean alcCaptureCloseDevice( ALCdevice *device );
void alcCaptureStart( ALCdevice *device );
void alcCaptureStop( ALCdevice *device );
void alcCaptureSamples( ALCdevice *device, ALCvoid *buffer, ALCsizei samples );

// For convenience we offer "singular" versions of the following
// calls as well, which require different wrappers if we want to
// be efficient. The main reason for "singular" versions is that
// Go doesn't allow us to treat a variable as an array.

ALCint walcGetInteger(ALCdevice *device, ALCenum param) {
	ALCint result;
	alcGetIntegerv(device, param, 1, &result);
	return result;
}
*/
import "C"
import "unsafe"

import "openal/al"

// Error codes returned by Device.GetError().
const (
	NoError = 0;
	InvalidDevice =0xA001;
	InvalidContext = 0xA002;
	InvalidEnum = 0xA003;
	InvalidValue = 0xA004;
	OutOfMemory = 0xA005;
)

const (
	Frequency = 0x1007; // int Hz
	Refresh = 0x1008; // int Hz
	Sync = 0x1009; // bool
	MonoSources = 0x1010; // int
	StereoSources = 0x1011; // int
)

// The Specifier string for default device?
const (
	DefaultDeviceSpecifier = 0x1004;
	DeviceSpecifier = 0x1005;
	Extensions = 0x1006;
)

// ?
const (
	MajorVersion = 0x1000;
	MinorVersion = 0x1001;
)

// ?
const (
	AttributesSize = 0x1002;
	AllAttributes = 0x1003;
)

// Capture extension
const (
	CaptureDeviceSpecifier = 0x310;
	CaptureDefaultDeviceSpecifier = 0x311;
	CaptureSamples = 0x312;
)


type Device struct {
	handle *C.ALCdevice;
}

// GetError() returns the most recent error generated
// in the AL state machine.
func (self Device) GetError() uint32 {
	return uint32(C.alcGetError(self.handle));
}

func OpenDevice(name string) Device {
	// TODO: turn empty string into nil?
	// TODO: what about an error return?
	p := C.CString(name);
	h := C.walcOpenDevice(p);
	C.free(unsafe.Pointer(p));
	return Device{h};
}

func (self Device) CloseDevice() bool {
	//TODO: really a method? or not?
	return C.alcCloseDevice(self.handle) != 0;
}

func (self Device) CreateContext() Context {
	// TODO: really a method?
	// TODO: attrlist support
	return Context{C.alcCreateContext(self.handle, nil)};
}

func (self Device) GetIntegerv(param uint32, size uint32) (result []int32) {
	result = make([]int32, size);
	C.walcGetIntegerv(self.handle, C.ALCenum(param), C.ALCsizei(size), unsafe.Pointer(&result[0]));
	return;
}

func (self Device) GetInteger(param uint32) int32 {
	return int32(C.walcGetInteger(self.handle, C.ALCenum(param)));
}




type CaptureDevice struct {
	Device;
	sampleSize uint32;
}

func CaptureOpenDevice(name string, freq uint32, format uint32, size uint32) (device CaptureDevice) {
	// TODO: turn empty string into nil?
	// TODO: what about an error return?
	p := C.CString(name);
	h := C.walcCaptureOpenDevice(p, C.ALCuint(freq), C.ALCenum(format), C.ALCsizei(size));
	C.free(unsafe.Pointer(p));
	s := map[uint32]uint32{al.FormatMono8: 1, al.FormatMono16: 2, al.FormatStereo8: 2, al.FormatStereo16: 4}[format];
	return CaptureDevice{Device{h},s};
}

func (self CaptureDevice) CloseDevice() bool {
	return C.alcCaptureCloseDevice(self.handle) != 0;
}

func (self CaptureDevice) CaptureCloseDevice() bool {
	return self.CloseDevice();
}

func (self CaptureDevice) CaptureStart() {
	C.alcCaptureStart(self.handle);
}

func (self CaptureDevice) CaptureStop() {
	C.alcCaptureStop(self.handle);
}

func (self *CaptureDevice) CaptureSamples(size uint32) (data []byte) {
	data = make([]byte, size * self.sampleSize);
	C.alcCaptureSamples(self.handle, unsafe.Pointer(&data[0]), C.ALCsizei(size));
	return;
}

type Context struct {
	handle *C.ALCcontext;
}

func (self Context) MakeContextCurrent() bool {
	return C.alcMakeContextCurrent(self.handle) != 0;
}

func (self Context) DestroyContext() {
	C.alcDestroyContext(self.handle);
	self.handle = nil;
}

