// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openal

/*
#include <stdlib.h>

// It's sad but the OpenAL C API uses lots and lots of typedefs
// that require wrapper functions (using basic C types) for cgo
// to grok them. So there's a lot more C code here than I would
// like...

#include <AL/al.h>

// For convenience we offer "singular" versions of the following
// calls as well, which require different wrappers if we want to
// be efficient. The main reason for "singular" versions is that
// Go doesn't allow us to treat a variable as an array.

void walGenSources(ALsizei n, void *sources) {
	alGenSources(n, sources);
}
ALuint walGenSource(void) {
	ALuint source;
	alGenSources(1, &source);
	return source;
}

void walDeleteSources(ALsizei n, const void *sources) {
	alDeleteSources(n, sources);
}
void walDeleteSource(ALuint source) {
	alDeleteSources(1, &source);
}

void walGenBuffers(ALsizei n, void *buffers) {
	alGenBuffers(n, buffers);
}
ALuint walGenBuffer(void) {
	ALuint buffer;
	alGenBuffers(1, &buffer);
	return buffer;
}

void walDeleteBuffers(ALsizei n, const void *buffers) {
	alDeleteBuffers(n, buffers);
}
void walDeleteBuffer(ALuint buffer) {
	alDeleteBuffers(1, &buffer);
}

//AL_API void AL_APIENTRY alBufferData( ALuint bid, ALenum format, const ALvoid* data, ALsizei size, ALsizei freq );
//AL_API void AL_APIENTRY alSourceQueueBuffers( ALuint sid, ALsizei numEntries, const ALuint *bids );
//AL_API void AL_APIENTRY alSourceUnqueueBuffers( ALuint sid, ALsizei numEntries, ALuint *bids );

int walGetError(void) {
	return alGetError();
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

//ALUT_API ALuint ALUT_APIENTRY alutCreateBufferFromFileImage (const ALvoid *data, ALsizei length);
//ALUT_API ALuint ALUT_APIENTRY alutCreateBufferWaveform (ALenum waveshape, ALfloat frequency, ALfloat phase, ALfloat duration);

//ALUT_API ALvoid *ALUT_APIENTRY alutLoadMemoryFromFile (const char *fileName, ALenum *format, ALsizei *size, ALfloat *frequency);
//ALUT_API ALvoid *ALUT_APIENTRY alutLoadMemoryFromFileImage (const ALvoid *data, ALsizei length, ALenum *format, ALsizei *size, ALfloat *frequency);
//ALUT_API ALvoid *ALUT_APIENTRY alutLoadMemoryHelloWorld (ALenum *format, ALsizei *size, ALfloat *frequency);
//ALUT_API ALvoid *ALUT_APIENTRY alutLoadMemoryWaveform (ALenum waveshape, ALfloat frequency, ALfloat phase, ALfloat duration, ALenum *format, ALsizei *size, ALfloat *freq);

//ALUT_API const char *ALUT_APIENTRY alutGetMIMETypes (ALenum loader);

//ALUT_API ALint ALUT_APIENTRY alutGetMajorVersion (void);
//ALUT_API ALint ALUT_APIENTRY alutGetMinorVersion (void);
*/
import "C"
import "unsafe"

import "fmt"

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

// what GetError returns
const (
	AlNoError = 0;
	AlInvalidName = 0xA001;
	AlInvalidEnum = 0xA002;
	AlInvalidValue = 0xA003;
	AlInvalidOperation = 0xA004;
)

// what alutGetError returns
const (
	AlutErrorNoError = 0;
	AlutErrorOutOfMemory = 0x200;
	AlutErrorInvalidEnum = 0x201;
//#define ALUT_ERROR_INVALID_VALUE               0x202
//#define ALUT_ERROR_INVALID_OPERATION           0x203
//#define ALUT_ERROR_NO_CURRENT_CONTEXT          0x204
//#define ALUT_ERROR_AL_ERROR_ON_ENTRY           0x205
//#define ALUT_ERROR_ALC_ERROR_ON_ENTRY          0x206
//#define ALUT_ERROR_OPEN_DEVICE                 0x207
//#define ALUT_ERROR_CLOSE_DEVICE                0x208
//#define ALUT_ERROR_CREATE_CONTEXT              0x209
//#define ALUT_ERROR_MAKE_CONTEXT_CURRENT        0x20A
//#define ALUT_ERROR_DESTROY_CONTEXT             0x20B
//#define ALUT_ERROR_GEN_BUFFERS                 0x20C
//#define ALUT_ERROR_BUFFER_DATA                 0x20D
//#define ALUT_ERROR_IO_ERROR                    0x20E
//#define ALUT_ERROR_UNSUPPORTED_FILE_TYPE       0x20F
//#define ALUT_ERROR_UNSUPPORTED_FILE_SUBTYPE    0x210
//#define ALUT_ERROR_CORRUPT_OR_TRUNCATED_DATA   0x211
)

// waveform for alutSomething
const (
	AlutWaveformSine = 0x100;
	AlutWaveformSquare = 0x101;
	AlutWaveformSawtooth = 0x102;
	AlutWaveformWhitenoise = 0x103;
	AlutWaveformImpulse = 0x104;
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



func GetError() int {
	return int(C.walGetError());
}

func AlutGetError() int {
	return int(C.walutGetError());
}


// We maintain these to make results that return lists
// of OpenAL object names intelligible to the Go side.

var bufferRegistry map[C.ALuint]*Buffer = make(map[C.ALuint]*Buffer);
var sourceRegistry map[C.ALuint]*Source = make(map[C.ALuint]*Source);

func rememberBuffer(buffer *Buffer) {
	bufferRegistry[buffer.handle] = buffer;
}
func forgetBuffer(buffer *Buffer) {
	bufferRegistry[buffer.handle] = buffer, false;
}
func rememberSource(source *Source) {
	sourceRegistry[source.handle] = source;
}
func forgetSource(source *Source) {
	sourceRegistry[source.handle] = source, false;
}
func DumpRegistries() {
	fmt.Println("========");
	for _, v := range bufferRegistry {
		fmt.Printf("%s\n", v);
	}
	for _, v := range sourceRegistry {
		fmt.Printf("%s\n", v);
	}
}

// OpenAL Buffers

type Buffer struct {
	handle C.ALuint;
}

func GenBuffers(n int) (buffers []*Buffer) {
	bufferIds := make([]C.ALuint, n);
	C.walGenBuffers(C.ALsizei(n), unsafe.Pointer(&bufferIds[0]));

	if GetError() != AlNoError {
		return;
	}

	buffers = make([]*Buffer, n);
	for i, v := range bufferIds {
		b := new(Buffer);
		b.handle = v;
		rememberBuffer(b);
		buffers[i] = b;
	}
	return;
}

func GenBuffer() (buffer *Buffer) {
	h := C.walGenBuffer();

	if GetError() != AlNoError {
		return;
	}

	buffer = new(Buffer);
	buffer.handle = h;
	rememberBuffer(buffer);
	return;
}

func DeleteBuffers(buffers []*Buffer) {
	n := len(buffers);
	bufferIds := make([]C.ALuint, n);

	for i, v := range buffers {
		forgetBuffer(v);
		bufferIds[i] = v.handle;
	}

	C.walDeleteBuffers(C.ALsizei(n), unsafe.Pointer(&bufferIds[0]));
	return;
}

func DeleteBuffer(buffer *Buffer) {
	forgetBuffer(buffer);
	C.walDeleteBuffer(buffer.handle);
}

func CreateBufferHelloWorld() (buffer *Buffer) {
	h := C.alutCreateBufferHelloWorld();

	if AlutGetError() != AlutErrorNoError {
		return;
	}

	buffer = new(Buffer);
	buffer.handle = h;
	rememberBuffer(buffer);
	return;
}

func CreateBufferFromFile(name string) (buffer *Buffer) {
	p := C.CString(name);
	h := C.alutCreateBufferFromFile(p);
	C.free(unsafe.Pointer(p));

	if AlutGetError() != AlutErrorNoError {
		return;
	}

	buffer = new(Buffer);
	buffer.handle = h;
	rememberBuffer(buffer);
	return;
}

// OpenAL Sources

type Source struct {
	handle C.ALuint;
}

func GenSource() (source *Source) {
	source = new(Source);
	source.handle = C.ALuint(C.walGenSource());
	return source;
}

func GenSources(sources []uint) {
	n := len(sources);
	C.walGenSources(C.ALsizei(n), unsafe.Pointer(&sources[0]));
}

// TODO: can't pass buffer really...
func (self *Source) SetAttr(param int, value *Buffer) {
	C.alSourcei(self.handle, C.ALenum(param), C.ALint(value.handle));
}

func (self *Source) Play() {
	C.alSourcePlay(self.handle);
}

func (self *Source) Stop() {
	C.alSourceStop(self.handle);
}

func (self *Source) Rewind() {
	C.alSourceRewind(self.handle);
}

func (self *Source) Pause() {
	C.alSourcePause(self.handle);
}
