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
int walGenSource(void) {
	ALuint source;
	alGenSources(1, &source);
	return source;
}
//AL_API void AL_APIENTRY alDeleteSources( ALsizei n, const ALuint* sources );
//AL_API void AL_APIENTRY alGenBuffers( ALsizei n, ALuint* buffers );
//AL_API void AL_APIENTRY alDeleteBuffers( ALsizei n, const ALuint* buffers );
//AL_API void AL_APIENTRY alBufferData( ALuint bid, ALenum format, const ALvoid* data, ALsizei size, ALsizei freq );
void walSourcePlay(int sid) {
	alSourcePlay(sid);
}
void walSourceStop(int sid) {
	alSourceStop(sid);
}
void walSourceRewind(int sid) {
	alSourceRewind(sid);
}
void walSourcePause(int sid) {
	alSourceRewind(sid);
}
//AL_API void AL_APIENTRY alSourceQueueBuffers( ALuint sid, ALsizei numEntries, const ALuint *bids );
//AL_API void AL_APIENTRY alSourceUnqueueBuffers( ALuint sid, ALsizei numEntries, ALuint *bids );

int walGetError(void) {
	return alGetError();
}
void walSourcei(int sid, int param, int value) {
	alSourcei(sid, param, value);
}
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

// Example for attrlist setup from, ALC_INVALID terminates
// http://www.fifi.org/cgi-bin/info2www?%28openal%29alc 
// int attrlist[] = { ALC_SYNC, AL_TRUE,
// ALC_SOURCES, 100,
// ALC_FREQUENCY, 44100,
// ALC_INVALID };
// ALCdevice *dev = alcOpenDevice( NULL );
// void *context = alcCreateContext( dev, attrlist );
ALCcontext* walcCreateContext(ALCdevice *device, const int* attrlist) {
	return alcCreateContext(device, attrlist);
}

int walcMakeContextCurrent(ALCcontext *context) {
	return alcMakeContextCurrent(context);
}

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
#include <AL/alut.h>

// TODO: alutInit(int *argcp, char **argv)
int walutInit(void) {
	return alutInit(NULL, NULL);
}

int walutExit(void) {
	return alutExit();
}

int walutGetError(void) {
	return alutGetError();
}

const char *walutGetErrorString(int error) {
	return alutGetErrorString(error);
}

//ALUT_API ALuint ALUT_APIENTRY alutCreateBufferFromFile (const char *fileName);
//ALUT_API ALuint ALUT_APIENTRY alutCreateBufferFromFileImage (const ALvoid *data, ALsizei length);
//ALUT_API ALuint ALUT_APIENTRY alutCreateBufferHelloWorld (void);
int walutCreateBufferHelloWorld(void) {
	return alutCreateBufferHelloWorld();
}
//ALUT_API ALuint ALUT_APIENTRY alutCreateBufferWaveform (ALenum waveshape, ALfloat frequency, ALfloat phase, ALfloat duration);

//ALUT_API ALvoid *ALUT_APIENTRY alutLoadMemoryFromFile (const char *fileName, ALenum *format, ALsizei *size, ALfloat *frequency);
//ALUT_API ALvoid *ALUT_APIENTRY alutLoadMemoryFromFileImage (const ALvoid *data, ALsizei length, ALenum *format, ALsizei *size, ALfloat *frequency);
//ALUT_API ALvoid *ALUT_APIENTRY alutLoadMemoryHelloWorld (ALenum *format, ALsizei *size, ALfloat *frequency);
//ALUT_API ALvoid *ALUT_APIENTRY alutLoadMemoryWaveform (ALenum waveshape, ALfloat frequency, ALfloat phase, ALfloat duration, ALenum *format, ALsizei *size, ALfloat *freq);

//ALUT_API const char *ALUT_APIENTRY alutGetMIMETypes (ALenum loader);

//ALUT_API ALint ALUT_APIENTRY alutGetMajorVersion (void);
//ALUT_API ALint ALUT_APIENTRY alutGetMinorVersion (void);

//ALUT_API ALboolean ALUT_APIENTRY alutSleep (ALfloat duration);
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
	AlBuffer = 0x1009;
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

func (self *Device) CreateContext() (context *Context) {
	// TODO: attrlist support
	h := C.walcCreateContext(self.handle, nil);

	if h == nil {
		return;
	}

	context = new(Context);
	context.handle = h;
	return;
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
	// TODO: iffy iffy iffy :-D
	var buffer [16*1024]byte;
	C.walcCaptureSamples(self.handle, unsafe.Pointer(&buffer), C.int(size));
	return buffer[0:];
}


type Context struct {
	handle *C.ALCcontext;
}

func (self *Context) MakeContextCurrent() bool {
	return C.walcMakeContextCurrent(self.handle) != 0;
}

func (self *Context) DestroyContext() {
	C.alcDestroyContext(self.handle);
	self.handle = nil;
	// XXX: there used to be a alcDestroyContext() that
	// returned something, but our alc.h doesn't list
	// that one... Hmmm...
}

func Init() {
	C.walutInit();
}

func Exit() {
	C.walutExit();
}


type Buffer struct {
	handle C.int;
}

func CreateBufferHelloWorld() (buffer *Buffer) {
	buffer = new(Buffer);
	buffer.handle = C.walutCreateBufferHelloWorld();
	return;
}


type Source struct {
	handle C.int;
}

func GenSource() (source *Source) {
	source = new(Source);
	source.handle = C.walGenSource();
	return source;
}

// TODO: can't pass buffer really...
func (self *Source) SetAttr(param int, value *Buffer) {
	C.walSourcei(self.handle, C.int(param), value.handle);
}

func (self *Source) Play() {
	C.walSourcePlay(self.handle);
}

