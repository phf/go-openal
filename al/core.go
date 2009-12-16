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
//	ALvoid		not applicable (but see below)
//
// We also stick to these (not mentioned explicitly in
// OpenAL):
//
//	ALvoid*		unsafe.Pointer (but never exported)
//	ALchar*		string
//
// Finally, in places where OpenAL expects pointers to
// C-style arrays, we use Go slices if appropriate:
//
//	ALboolean*	[]bool
//	ALvoid*		[]byte (see SetBufferData() for example)
//	ALint*		[]int32
//	ALuint*		[]uint32 []Source []Buffer
//	ALfloat*	[]float32
//	ALdouble*	[]float64
//
// Overall, the correspondence of types hopefully feels
// natural enough. Note that many of these types do not
// actually occur in the API.
//
// The names of OpenAL constants follow the established
// Go conventions: instead of AL_FORMAT_MONO16 we use
// FormatMono16 for example.
//
// Conversion to Go's camel case notation does however
// lead to name clashes between constants and functions.
// For example, AL_DISTANCE_MODEL becomes DistanceModel
// which collides with the OpenAL function of the same
// name used to set the current distance model. We have
// to rename either the constant or the function, and
// since the function name seems to be at fault (it's a
// setter but doesn't make that obvious), we rename the
// function.
//
// In fact, we renamed plenty of functions, not just the
// ones where collisions with constants were the driving
// force. For example, instead of the Sourcef/GetSourcef
// abomination, we use Getf/Setf methods on a Source type.
// Everything should still be easily recognizable for
// OpenAL hackers, but this structure is a lot more
// sensible (and reveals that the OpenAL API is actually
// not such a bad design).
//
// There are a few cases where constants would collide
// with the names of types we introduced here. Since the
// types serve a much more important function, we renamed
// the constants in those cases. For example AL_BUFFER
// would collide with the type Buffer so it's name is now
// Buffer_ instead. Not pretty, but in many cases you
// don't need the constants anyway as the functionality
// they represent is probably available through one of
// the convenience functions we introduced as well. For
// example consider the task of attaching a buffer to a
// source. In C, you'd say alSourcei(sid, AL_BUFFER, bid).
// In Go, you could say sid.Seti(Buffer_, bid) if you
// wish. But you probably want to say sid.SetBuffer(bid).
//
// TODO: write wrappers with better names around GetInteger()
// and friends? for example GetDistanceModel()? could go into
// util.go or something, pure Go code
package al

/*
#include <stdlib.h>
#include <AL/al.h>
#include "wrappers.c"
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


func DopplerFactor (value float32) {
	C.alDopplerFactor(C.ALfloat(value));
}

func DopplerVelocity (value float32) {
	C.alDopplerVelocity(C.ALfloat(value));
}

func SpeedOfSound (value float32) {
	C.alSpeedOfSound(C.ALfloat(value));
}



// TODO: Get*() queries.
const (
	DistanceModel = 0xD000;
)

// TODO: ???
const (
	Buffer_ = 0x1009;
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

func GetString(param uint32) string {
	return C.GoString(C.walGetString(C.ALenum(param)));
}

func GetBoolean(param uint32) bool {
	return C.alGetBoolean(C.ALenum(param)) != 0;
}

func GetInteger(param uint32) int32 {
	return int32(C.alGetInteger(C.ALenum(param)));
}

func GetFloat(param uint32) float32 {
	return float32(C.alGetFloat(C.ALenum(param)));
}

func GetDouble(param uint32) float64 {
	return float64(C.alGetDouble(C.ALenum(param)));
}

func GetBooleanv(param uint32, data []bool) {
	C.walGetBooleanv(C.ALenum(param), unsafe.Pointer(&data[0]));
}

func GetIntegerv(param uint32, data []int32) {
	C.walGetIntegerv(C.ALenum(param), unsafe.Pointer(&data[0]));
}

func GetFloatv(param uint32, data []float32) {
	C.walGetFloatv(C.ALenum(param), unsafe.Pointer(&data[0]));
}

func GetDoublev(param uint32, data []float64) {
	C.walGetDoublev(C.ALenum(param), unsafe.Pointer(&data[0]));
}



///// Source /////////////////////////////////////////////////////////

// Sources represent sound emitters in 3d space.
type Source uint32;

func GenSources(n int) (sources []Source) {
	sources = make([]Source, n);
	C.walGenSources(C.ALsizei(n), unsafe.Pointer(&sources[0]));
	return;
}

func DeleteSources(sources []Source) {
	n := len(sources);
	C.walDeleteSources(C.ALsizei(n), unsafe.Pointer(&sources[0]));
}

func (self Source) Setf(param uint32, value float32) {
	C.alSourcef(C.ALuint(self), C.ALenum(param), C.ALfloat(value));
}

func (self Source) Set3f(param uint32, value1, value2, value3 float32) {
	C.alSource3f(C.ALuint(self), C.ALenum(param), C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3));
}

func (self Source) Setfv(param uint32, values []float32) {
	C.walSourcefv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

func (self Source) Seti(param uint32, value int32) {
	C.alSourcei(C.ALuint(self), C.ALenum(param), C.ALint(value));
}

func (self Source) Set3i(param uint32, value1, value2, value3 int32) {
	C.alSource3i(C.ALuint(self), C.ALenum(param), C.ALint(value1), C.ALint(value2), C.ALint(value3));
}

func (self Source) Setiv(param uint32, values []int32) {
	C.walSourceiv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

func (self Source) Getf(param uint32) float32 {
	return float32(C.walGetSourcef(C.ALuint(self), C.ALenum(param)));
}

func (self Source) Get3f(param uint32) (value1, value2, value3 float32) {
	var v1, v2, v3 float32;
	C.walGetSource3f(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

func (self Source) Getfv(param uint32) (values []float32) {
	C.walGetSourcefv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

func (self Source) Geti(param uint32) int32 {
	return int32(C.walGetSourcei(C.ALuint(self), C.ALenum(param)));
}

func (self Source) Get3i(param uint32) (value1, value2, value3 int32) {
	var v1, v2, v3 int32;
	C.walGetSource3i(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

func (self Source) Getiv(param uint32) (values []int32) {
	C.walGetSourceiv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
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

func (self Source) QueueBuffers(buffers []Buffer) {
	C.walSourceQueueBuffers(C.ALuint(self), C.ALsizei(len(buffers)), unsafe.Pointer(&buffers[0]));
}

func (self Source) UnqueueBuffers(buffers []Buffer) {
	C.walSourceUnqueueBuffers(C.ALuint(self), C.ALsizei(len(buffers)), unsafe.Pointer(&buffers[0]));
}

func SourcePlayv(sources []Source) {
	C.walSourcePlayv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]));
}

func SourceStopv(sources []Source) {
	C.walSourceStopv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]));
}

func SourceRewindv(sources []Source) {
	C.walSourceRewindv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]));
}

func SourcePausev(sources []Source) {
	C.walSourcePausev(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]));
}

///// Buffer /////////////////////////////////////////////////////////

// Buffers are storage space for sample data.
type Buffer uint32;

// GenBuffers() creates n buffers.
func GenBuffers(n int) (buffers []Buffer) {
	buffers = make([]Buffer, n);
	C.walGenBuffers(C.ALsizei(n), unsafe.Pointer(&buffers[0]));
	return;
}

// DeleteBuffers() deletes the given buffers.
func DeleteBuffers(buffers []Buffer) {
	n := len(buffers);
	C.walDeleteBuffers(C.ALsizei(n), unsafe.Pointer(&buffers[0]));
}

// Renamed, was Bufferf.
func (self Buffer) Setf(param uint32, value float32) {
	C.alBufferf(C.ALuint(self), C.ALenum(param), C.ALfloat(value));
}

// Renamed, was Buffer3f.
func (self Buffer) Set3f(param uint32, value1, value2, value3 float32) {
	C.alBuffer3f(C.ALuint(self), C.ALenum(param), C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3));
}

// Renamed, was Bufferfv.
func (self Buffer) Setfv(param uint32, values []float32) {
	C.walBufferfv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was Bufferi.
func (self Buffer) Seti(param uint32, value int32) {
	C.alBufferi(C.ALuint(self), C.ALenum(param), C.ALint(value));
}

// Renamed, was Buffer3i.
func (self Buffer) Set3i(param uint32, value1, value2, value3 int32) {
	C.alBuffer3i(C.ALuint(self), C.ALenum(param), C.ALint(value1), C.ALint(value2), C.ALint(value3));
}

// Renamed, was Bufferiv.
func (self Buffer) Setiv(param uint32, values []int32) {
	C.walBufferiv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was GetBufferf.
func (self Buffer) Getf(param uint32) float32 {
	return float32(C.walGetBufferf(C.ALuint(self), C.ALenum(param)));
}

// Renamed, was GetBuffer3f.
func (self Buffer) Get3f(param uint32) (value1, value2, value3 float32) {
	var v1, v2, v3 float32;
	C.walGetBuffer3f(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetBufferfv.
func (self Buffer) Getfv(param uint32) (values []float32) {
	C.walGetBufferfv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

// Renamed, was GetBufferi.
func (self Buffer) Geti(param uint32) int32 {
	return int32(C.walGetBufferi(C.ALuint(self), C.ALenum(param)));
}

// Renamed, was GetBuffer3i.
func (self Buffer) Get3i(param uint32) (value1, value2, value3 int32) {
	var v1, v2, v3 int32;
	C.walGetBuffer3i(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetBufferiv.
func (self Buffer) Getiv(param uint32) (values []int32) {
	C.walGetBufferiv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

// Format of sound samples passed to SetBufferData().
const (
	FormatMono8 = 0x1100;
	FormatMono16 = 0x1101;
	FormatStereo8 = 0x1102;
	FormatStereo16 = 0x1103;
)

// SetBufferData() specifies the sample data the buffer should use.
// For FormatMono16 and FormatStereo8 the data slice must be a
// multiple of two bytes long; for FormatStereo16 the data slice
// must be a multiple of four bytes long. The frequency is given
// in Hz.
// Renamed, was BufferData.
func (self Buffer) SetBufferData(format uint32, data []byte, frequency uint32) {
	C.alBufferData(C.ALuint(self), C.ALenum(format), unsafe.Pointer(&data[0]),
		C.ALsizei(len(data)), C.ALsizei(frequency));
}

///// Listener ///////////////////////////////////////////////////////

// Listener represents the singleton receiver of
// sound in 3d space.
//
// We "fake" this type so we can provide OpenAL
// listener calls as methods. This is convenient
// and makes all those calls consistent with the
// way they work for Source and Buffer. You can't
// make new listeners, there's only one!
type Listener struct{};

// Renamed, was Listenerf.
func (self Listener) Setf(param uint32, value float32) {
	C.alListenerf(C.ALenum(param), C.ALfloat(value));
}

// Renamed, was Listener3f.
func (self Listener) Set3f(param uint32, value1, value2, value3 float32) {
	C.alListener3f(C.ALenum(param), C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3));
}

// Renamed, was Listenerfv.
func (self Listener) Setfv(param uint32, values []float32) {
	C.walListenerfv(C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was Listeneri.
func (self Listener) Seti(param uint32, value int32) {
	C.alListeneri(C.ALenum(param), C.ALint(value));
}

// Renamed, was Listener3i.
func (self Listener) Set3i(param uint32, value1, value2, value3 int32) {
	C.alListener3i(C.ALenum(param), C.ALint(value1), C.ALint(value2), C.ALint(value3));
}

// Renamed, was Listeneriv.
func (self Listener) Setiv(param uint32, values []int32) {
	C.walListeneriv(C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was GetListenerf.
func (self Listener) Getf(param uint32) float32 {
	return float32(C.walGetListenerf(C.ALenum(param)));
}

// Renamed, was GetListener3f.
func (self Listener) Get3f(param uint32) (value1, value2, value3 float32) {
	var v1, v2, v3 float32;
	C.walGetListener3f(C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetListenerfv.
func (self Listener) Getfv(param uint32) (values []float32) {
	C.walGetListenerfv(C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

// Renamed, was GetListeneri.
func (self Listener) Geti(param uint32) int32 {
	return int32(C.walGetListeneri(C.ALenum(param)));
}

// Renamed, was GetListener3i.
func (self Listener) Get3i(param uint32) (value1, value2, value3 int32) {
	var v1, v2, v3 int32;
	C.walGetListener3i(C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetListeneriv.
func (self Listener) Getiv(param uint32) (values []int32) {
	C.walGetListeneriv(C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

///// Convenience ////////////////////////////////////////////////////

// GenSource() creates a single source.
// Convenience function, see GenSources().
func GenSource() Source {
	return Source(C.walGenSource());
}

// DeleteSource() deletes a single source.
// Convenience function, see DeleteSources().
func DeleteSource(source Source) {
	C.walDeleteSource(C.ALuint(source));
}

// GenBuffer() creates a single buffer.
// Convenience function, see GenBuffers().
func GenBuffer() Buffer {
	return Buffer(C.walGenBuffer());
}

// DeleteBuffer() deletes a single buffer.
// Convenience function, see DeleteBuffers().
func DeleteBuffer(buffer Buffer) {
	C.walDeleteSource(C.ALuint(buffer));
}

///// Crap ///////////////////////////////////////////////////////////

// These functions are wrapped and should work fine, but they
// have no purpose: There are *no* capabilities in OpenAL 1.1
// which is the latest specification. So we removed from from
// the API for now, it's complicated enough without them.
//
//func Enable(capability uint32) {
//	C.alEnable(C.ALenum(capability));
//}
//
//func Disable(capability uint32) {
//	C.alDisable(C.ALenum(capability));
//}
//
//func IsEnabled(capability uint32) bool {
//	return C.alIsEnabled(C.ALenum(capability)) != 0;
//}
