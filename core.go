// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openal

/*
#include <stdlib.h>

// It's sad but the OpenAL C API uses lots and lots of typedefs
// that require wrapper function to make things work with cgo.

#include <AL/al.h>

//AL_API void AL_APIENTRY alGenSources( ALsizei n, ALuint* sources );
//AL_API void AL_APIENTRY alDeleteSources( ALsizei n, const ALuint* sources );
//AL_API void AL_APIENTRY alGenBuffers( ALsizei n, ALuint* buffers );
//AL_API void AL_APIENTRY alDeleteBuffers( ALsizei n, const ALuint* buffers );
//AL_API void AL_APIENTRY alBufferData( ALuint bid, ALenum format, const ALvoid* data, ALsizei size, ALsizei freq );
//AL_API void AL_APIENTRY alSourcePlay( ALuint sid );
//AL_API void AL_APIENTRY alSourceStop( ALuint sid );
//AL_API void AL_APIENTRY alSourceRewind( ALuint sid );
//AL_API void AL_APIENTRY alSourcePause( ALuint sid );
//AL_API void AL_APIENTRY alSourceQueueBuffers( ALuint sid, ALsizei numEntries, const ALuint *bids );
//AL_API void AL_APIENTRY alSourceUnqueueBuffers( ALuint sid, ALsizei numEntries, ALuint *bids );
//AL_API ALenum AL_APIENTRY alGetError( void );
//AL_API void AL_APIENTRY alSourcei( ALuint sid, ALenum param, ALint value );
//AL_API void AL_APIENTRY alGetSourcei( ALuint sid,  ALenum param, ALint* value );

#include <AL/alc.h>

ALCdevice *walcOpenDevice(const char *devicename) {
	return alcOpenDevice(devicename);
}

int walcCloseDevice(ALCdevice *device) {
	return alcCloseDevice(device);
}

int walcGetError(ALCdevice *device) {
	return alcGetError(device);
}

//ALC_API ALCcontext *    ALC_APIENTRY alcCreateContext( ALCdevice *device, const ALCint* attrlist );
//ALC_API ALCboolean      ALC_APIENTRY alcMakeContextCurrent( ALCcontext *context );
//ALC_API void            ALC_APIENTRY alcDestroyContext( ALCcontext *context );

ALCdevice *walcCaptureOpenDevice(const char *devicename, int frequency, int format, int buffersize) {
	return alcCaptureOpenDevice(devicename, frequency, format, buffersize);
}

int walcCaptureCloseDevice(ALCdevice *device) {
	return alcCaptureCloseDevice(device);
}

// Silly! You ask for the number of samples, but depending on the format
// you're recording in there can be 1-4 bytes per sample!

void walcCaptureSamples(ALCdevice *device, void *buffer, int samples) {
	alcCaptureSamples(device, buffer, samples);
}

int walcGetInteger(ALCdevice *device, int param) {
	int result;
	alcGetIntegerv(device, param, sizeof(result), &result);
	return result;
}

//#include <AL/alext.h>
//#include <AL/alut.h>
*/
import "C"
import "unsafe"

// All of the following are eventually going to be
// private to the Go OpenAL binding. For now I am
// just playing around, so they are public. Not for
// long I hope. :-D

// what Device.GetError returns
const (
	AlcNoError = 0;
	AlcInvalidDevice =0xA001;
	AlcInvalidContext = 0xA002;
	AlcInvalidEnum = 0xA003;
	AlcInvalidValue = 0xA004;
	AlcOutOfMemory = 0xA005;
)

// format for CaptureOpenDevice
const (
	AlFormatMono8 = 0x1100;
	AlFormatMono16 = 0x1101;
	AlFormatStereo8 = 0x1102;
	AlFormatStereo16 = 0x1103;
)

const (
	AlcCaptureSamples = 0x312;
)

type Device struct {
	handle *C.ALCdevice;
}

func OpenDevice(name string) (device *Device) {
	p := C.CString(name);
	h := C.walcOpenDevice(p);
	C.free(unsafe.Pointer(p));

	if h == nil {
		return;
	}

	device = new(Device);
	device.handle = h;
	return;
}

func (self *Device) CloseDevice() bool {
	return C.walcCloseDevice(self.handle) != 0;
}

func (self *Device) GetError() int {
	return int(C.walcGetError(self.handle));
}

type CaptureDevice struct {
	handle *C.ALCdevice;
}

func CaptureOpenDevice(name string, freq int, format int, size int) (device *CaptureDevice) {
	p := C.CString(name);
	h := C.walcCaptureOpenDevice(p, C.int(freq), C.int(format), C.int(size));
	C.free(unsafe.Pointer(p));

	if h == nil {
		return;
	}

	device = new(CaptureDevice);
	device.handle = h;
	return;
}

func (self *CaptureDevice) CaptureCloseDevice() bool {
	return C.walcCaptureCloseDevice(self.handle) != 0;
}

func (self *CaptureDevice) GetError() int {
	return int(C.walcGetError(self.handle));
}

func (self *CaptureDevice) CaptureStart() {
	C.alcCaptureStart(self.handle);
}

func (self *CaptureDevice) CaptureStop() {
	C.alcCaptureStop(self.handle);
}

func (self *CaptureDevice) GetInteger(param int) int {
	return int(C.walcGetInteger(self.handle, C.int(param)));
}

func (self *CaptureDevice) CaptureSamples(size int) []byte {
	var buffer [16*1024]byte;
	C.walcCaptureSamples(self.handle, unsafe.Pointer(&buffer), C.int(size));
	return buffer[0:];
}
