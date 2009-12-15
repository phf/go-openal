// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// C-level binding for OpenAL's "al" API.
//
// Please consider using the Go-level binding instead.
//
// OpenAL types are (in principle) mapped to Go types as
// follows:
//
//	ALboolean	bool	(al.h says char, but Go's bool should be compatible)
//	ALchar		uint8	(although al.h suggests int8, Go's uint8 (aka byte) seems better)
//	ALbyte		int8	(al.h says char, implying that char is signed)
//	ALubyte		uint8	(al.h says unsigned char)
//	ALshort		int16
//	ALushort	uint16
//	ALint		int32
//	ALuint		uint32
//	ALsizei		uint32	(although al.h suggests int32)
//	ALenum		uint32	(although al.h suggests int32)
//	ALfloat		float32
//	ALdouble	float64
//	ALvoid		-
//
// We also stick to these (not mentioned explicitly in
// OpenAL):
//
//	ALvoid*		unsafe.Pointer
//	ALchar*		string
//
// Overall, the correspondence of types hopefully feels
// natural enough.
//
// XXX: conflicts between constants and functions, for
// example DistanceModel and DistanceModel() and how
// they are resolved.
//
// TODO: write wrappers with better names around GetInteger()
// and friends? for example GetDistanceModel()?
package al

/*
#include <stdlib.h>

// It's sad but the OpenAL C API uses lots and lots of typedefs
// that require wrapper functions (using basic C types) for cgo
// to grok them. So there's a lot more C code here than I would
// like...

#include <AL/al.h>

// I keep all the al.h prototypes here for now, for reference.
// They'll go away eventually. Those commented out are already
// accessible from Go.

// void alEnable( ALenum capability );
// void alDisable( ALenum capability ); 
// ALboolean alIsEnabled( ALenum capability ); 
// const ALchar* alGetString( ALenum param );
const char *walGetString(ALenum param) {
	return alGetString(param);
}
void alGetBooleanv( ALenum param, ALboolean* data );
void alGetIntegerv( ALenum param, ALint* data );
void alGetFloatv( ALenum param, ALfloat* data );
void alGetDoublev( ALenum param, ALdouble* data );
ALboolean alGetBoolean( ALenum param );
ALint alGetInteger( ALenum param );
ALfloat alGetFloat( ALenum param );
ALdouble alGetDouble( ALenum param );
// ALenum alGetError( void );
ALboolean alIsExtensionPresent( const ALchar* extname );
void* alGetProcAddress( const ALchar* fname );
ALenum alGetEnumValue( const ALchar* ename );
// void alListenerf( ALenum param, ALfloat value );
// void alListener3f( ALenum param, ALfloat value1, ALfloat value2, ALfloat value3 );
// void alListenerfv( ALenum param, const ALfloat* values ); 
void walListenerfv(ALenum param, const void* values) {
	alListenerfv(param, values);
}
// void alListeneri( ALenum param, ALint value );
// void alListener3i( ALenum param, ALint value1, ALint value2, ALint value3 );
// void alListeneriv( ALenum param, const ALint* values );
void walListeneriv(ALenum param, const void* values) {
	alListeneriv(param, values);
}
// void alGetListenerf( ALenum param, ALfloat* value );
ALfloat walGetListenerf(ALenum param) {
	ALfloat result;
	alGetListenerf(param, &result);
	return result;
}
// void alGetListener3f( ALenum param, ALfloat *value1, ALfloat *value2, ALfloat *value3 );
void walGetListener3f(ALenum param, void *value1, void *value2, void *value3) {
	alGetListener3f(param, value1, value2, value3);
}
// void alGetListenerfv( ALenum param, ALfloat* values );
void walGetListenerfv(ALenum param, void* values) {
	alGetListenerfv(param, values);
}
// void alGetListeneri( ALenum param, ALint* value );
ALint walGetListeneri(ALenum param) {
	ALint result;
	alGetListeneri(param, &result);
	return result;
}
// void alGetListener3i( ALenum param, ALint *value1, ALint *value2, ALint *value3 );
void walGetListener3i(ALenum param, void *value1, void *value2, void *value3) {
	alGetListener3i(param, value1, value2, value3);
}
// void alGetListeneriv( ALenum param, ALint* values );
void walGetListeneriv(ALenum param, void* values) {
	alGetListeneriv(param, values);
}
//void alGenSources( ALsizei n, ALuint* sources ); 
void walGenSources(ALsizei n, void *sources) {
	alGenSources(n, sources);
}
//void alDeleteSources( ALsizei n, const ALuint* sources );
void walDeleteSources(ALsizei n, const void *sources) {
	alDeleteSources(n, sources);
}
// ALboolean alIsSource( ALuint sid ); 
void alSourcef( ALuint sid, ALenum param, ALfloat value ); 
void alSource3f( ALuint sid, ALenum param, ALfloat value1, ALfloat value2, ALfloat value3 );
void alSourcefv( ALuint sid, ALenum param, const ALfloat* values ); 
void alSourcei( ALuint sid, ALenum param, ALint value ); 
void alSource3i( ALuint sid, ALenum param, ALint value1, ALint value2, ALint value3 );
void alSourceiv( ALuint sid, ALenum param, const ALint* values );
void alGetSourcef( ALuint sid, ALenum param, ALfloat* value );
void alGetSource3f( ALuint sid, ALenum param, ALfloat* value1, ALfloat* value2, ALfloat* value3);
void alGetSourcefv( ALuint sid, ALenum param, ALfloat* values );
void alGetSourcei( ALuint sid,  ALenum param, ALint* value );
void alGetSource3i( ALuint sid, ALenum param, ALint* value1, ALint* value2, ALint* value3);
void alGetSourceiv( ALuint sid,  ALenum param, ALint* values );
void alSourcePlayv( ALsizei ns, const ALuint *sids );
void alSourceStopv( ALsizei ns, const ALuint *sids );
void alSourceRewindv( ALsizei ns, const ALuint *sids );
void alSourcePausev( ALsizei ns, const ALuint *sids );
//void alSourcePlay( ALuint sid );
//void alSourceStop( ALuint sid );
//void alSourceRewind( ALuint sid );
//void alSourcePause( ALuint sid );
void alSourceQueueBuffers( ALuint sid, ALsizei numEntries, const ALuint *bids );
void alSourceUnqueueBuffers( ALuint sid, ALsizei numEntries, ALuint *bids );
//void alGenBuffers( ALsizei n, ALuint* buffers );
void walGenBuffers(ALsizei n, void *buffers) {
	alGenBuffers(n, buffers);
}
//void alDeleteBuffers( ALsizei n, const ALuint* buffers );
void walDeleteBuffers(ALsizei n, const void *buffers) {
	alDeleteBuffers(n, buffers);
}
// ALboolean alIsBuffer( ALuint bid );
// void alBufferData( ALuint bid, ALenum format, const ALvoid* data, ALsizei size, ALsizei freq );
void alBufferf( ALuint bid, ALenum param, ALfloat value );
void alBuffer3f( ALuint bid, ALenum param, ALfloat value1, ALfloat value2, ALfloat value3 );
void alBufferfv( ALuint bid, ALenum param, const ALfloat* values );
void alBufferi( ALuint bid, ALenum param, ALint value );
void alBuffer3i( ALuint bid, ALenum param, ALint value1, ALint value2, ALint value3 );
void alBufferiv( ALuint bid, ALenum param, const ALint* values );
void alGetBufferf( ALuint bid, ALenum param, ALfloat* value );
void alGetBuffer3f( ALuint bid, ALenum param, ALfloat* value1, ALfloat* value2, ALfloat* value3);
void alGetBufferfv( ALuint bid, ALenum param, ALfloat* values );
void alGetBufferi( ALuint bid, ALenum param, ALint* value );
void alGetBuffer3i( ALuint bid, ALenum param, ALint* value1, ALint* value2, ALint* value3);
void alGetBufferiv( ALuint bid, ALenum param, ALint* values );
// void alDopplerFactor( ALfloat value );
// void alDopplerVelocity( ALfloat value );
// void alSpeedOfSound( ALfloat value );
// void alDistanceModel( ALenum distanceModel );

// For convenience we offer "singular" versions of the following
// calls as well, which require different wrappers if we want to
// be efficient. The main reason for "singular" versions is that
// Go doesn't allow us to treat a variable as an array.

ALuint walGenSource(void) {
	ALuint source;
	alGenSources(1, &source);
	return source;
}
void walDeleteSource(ALuint source) {
	alDeleteSources(1, &source);
}
ALuint walGenBuffer(void) {
	ALuint buffer;
	alGenBuffers(1, &buffer);
	return buffer;
}
void walDeleteBuffer(ALuint buffer) {
	alDeleteBuffers(1, &buffer);
}
*/
import "C"
import "unsafe"

// Error codes returned by GetError().
const (
	NoError = 0;
	InvalidName = 0xA001;
	InvalidEnum = 0xA002;
	InvalidValue = 0xA003;
	InvalidOperation = 0xA004;
)

// GetError() returns the most recent error generated
// in the AL state machine.
func GetError() uint32 {
	return uint32(C.alGetError());
}

// Useless since OpenAL does not yet define any capabilities.
func Enable(capability uint32) {
	C.alEnable(C.ALenum(capability));
}

// Useless since OpenAL does not yet define any capabilities.
func Disable(capability uint32) {
	C.alDisable(C.ALenum(capability));
}

// Useless since OpenAL does not yet define any capabilities.
func IsEnabled(capability uint32) bool {
	return C.alIsEnabled(C.ALenum(capability)) != 0;
}

// IsSource() returns true if id refers to a source.
// Not very useful as we provide a distinct Source type.
func IsSource(id uint32) bool {
	return C.alIsSource(C.ALuint(id)) != 0;
}

// IsBuffer() returns true if id refers to a buffer.
// Not very useful as we provide a distinct Buffer type.
func IsBuffer(id uint32) bool {
	return C.alIsBuffer(C.ALuint(id)) != 0;
}

// Sources represent sound emitters in 3d space.
type Source uint32;

// Buffers are storage space for sample data.
type Buffer uint32;

// Listener represents the singleton receiver of
// sound in 3d space.
//
// We "fake" this type so we can provide OpenAL
// listener calls as methods. This is convenient
// and makes all those calls consistent with the
// way they work for Source and Buffer. You can't
// make new listeners, there's only one!
type Listener struct{};

// TODO: Get*() queries.
const (
	DistanceModel = 0xD000;
)

// Distance models passed to SetDistanceModel().
const (
	InverseDistance = 0xD001;
	InverseDistanceClamped = 0xD002;
	LinearDistance = 0xD003;
	LinearDistanceClamped = 0xD004;
	ExponentDistance = 0xD005;
	ExponentDistanceClamped = 0xD006;
)

// SetDistanceModel() changes the current distance model.
// This is just DistanceModel() in OpenAL.
func SetDistanceModel(model uint32) {
	C.alDistanceModel(C.ALenum(model));
}

// NOT DOCUMENTED YET

func GetString(param uint32) string {
	return C.GoString(C.walGetString(C.ALenum(param)));
}


func GenSources(n int) (sources []Source) {
	sources = make([]Source, n);
	C.walGenSources(C.ALsizei(n), unsafe.Pointer(&sources[0]));
	return;
}

func GenSource() Source {
	return Source(C.walGenSource());
}

func DeleteSources(sources []Source) {
	n := len(sources);
	C.walDeleteSources(C.ALsizei(n), unsafe.Pointer(&sources[0]));
}

func DeleteSource(source Source) {
	C.walDeleteSource(C.ALuint(source));
}

func (self Source) Play() {
	C.alSourcePlay(C.ALuint(self));
}

func (self Source) Stop() {
	C.alSourceStop(C.ALuint(self));
}

func (self Source) Rewind() {
	C.alSourceRewind(C.ALuint(self));
}

func (self Source) Pause() {
	C.alSourcePause(C.ALuint(self));
}


func GenBuffers(n int) (buffers []Buffer) {
	buffers = make([]Buffer, n);
	C.walGenBuffers(C.ALsizei(n), unsafe.Pointer(&buffers[0]));
	return;
}

func GenBuffer() Buffer {
	return Buffer(C.walGenBuffer());
}

func DeleteBuffers(buffers []Buffer) {
	n := len(buffers);
	C.walDeleteBuffers(C.ALsizei(n), unsafe.Pointer(&buffers[0]));
}

func DeleteBuffer(buffer Buffer) {
	C.walDeleteSource(C.ALuint(buffer));
}

const (
	FormatMono8 = 0x1100;
	FormatMono16 = 0x1101;
	FormatStereo8 = 0x1102;
	FormatStereo16 = 0x1103;
)

func (self Buffer) BufferData(format uint32, data []byte, size uint32, freq uint32) {
	C.alBufferData(C.ALuint(self), C.ALenum(format), unsafe.Pointer(&data[0]),
		C.ALsizei(size), C.ALsizei(freq));
	// TODO: pass data as array or pointer?
}



func DopplerFactor (value float32) {
	C.alDopplerFactor(C.ALfloat(value));
}

func DopplerVelocity (value float32) {
	C.alDopplerVelocity(C.ALfloat(value));
}

func SpeedOfSound (value float32) {
	C.alSpeedOfSound(C.ALfloat(value));
}




func (self Listener) Setf(param uint32, value float32) {
	C.alListenerf(C.ALenum(param), C.ALfloat(value));
}

func (self Listener) Set3f(param uint32, value1, value2, value3 float32) {
	C.alListener3f(C.ALenum(param), C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3));
}

func (self Listener) Setfv(param uint32, values []float32) {
	C.walListenerfv(C.ALenum(param), unsafe.Pointer(&values[0]));
}

func (self Listener) Seti(param uint32, value int32) {
	C.alListeneri(C.ALenum(param), C.ALint(value));
}

func (self Listener) Set3i(param uint32, value1, value2, value3 int32) {
	C.alListener3i(C.ALenum(param), C.ALint(value1), C.ALint(value2), C.ALint(value3));
}

func (self Listener) Setiv(param uint32, values []int32) {
	C.walListeneriv(C.ALenum(param), unsafe.Pointer(&values[0]));
}

func (self Listener) Getf(param uint32) float32 {
	return float32(C.walGetListenerf(C.ALenum(param)));
}

func (self Listener) Get3f(param uint32) (value1, value2, value3 float32) {
	var v1, v2, v3 float32;
	C.walGetListener3f(C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
	// TODO: this the best way?
}

func (self Listener) Getfv(param uint32) (values []float32) {
	C.walGetListenerfv(C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

func (self Listener) Geti(param uint32) int32 {
	return int32(C.walGetListeneri(C.ALenum(param)));
}

func (self Listener) Get3i(param uint32) (value1, value2, value3 int32) {
	var v1, v2, v3 int32;
	C.walGetListener3i(C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
	// TODO: this the best way?
}

func (self Listener) Getiv(param uint32) (values []int32) {
	C.walGetListeneriv(C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}


// NOT CLEANED UP YET

const (
	AlBuffer = 0x1009;
)
