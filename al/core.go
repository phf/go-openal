// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// C-level binding for OpenAL's "al" API.
//
// Please consider using the Go-level binding instead.
// See http://connect.creativelabs.com/openal/Documentation/OpenAL%201.1%20Specification.htm
// for details about OpenAL not described here.
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
// In Go, you can say sid.Seti(Buffer_, bid) as well, but
// you probably want to say sid.SetBuffer(bid) instead.
//
// TODO: Not all convenience functions/methods that should
// exist are actually there yet. :-/
//
// TODO: After we have all the convenience functions/methods,
// should we get rid of the generic stuff to unclutter the
// API? Would take us further from OpenAL, but...
package al

/*
#include <stdlib.h>
#include <AL/al.h>
#include "wrappers.c"
*/
import "C"
import "unsafe"

// TODO: General constants for various purposes.
const (
	None = 0; // TODO: no distance model, no buffer for source
	alFalse = 0;
	alTrue = 1;
)

// TODO: GetInteger() queries.
const (
	alDistanceModel = 0xD000;
)

// TODO: GetFloat() queries.
const (
	alDopplerFactor = 0xC000;
	alDopplerVelocity = 0xC001;
	alSpeedOfSound = 0xC003;
)

// TODO: GetString() queries.
const (
	alVendor = 0xB001;
	alVersion = 0xB002;
	alRenderer = 0xB003;
	alExtensions = 0xB004;
)

// TODO: Shared Source/Listener properties.
const (
	Pitch = 0x1003; // TODO al.h says shared, docs say only Source?
	Position = 0x1004;
	Velocity = 0x1006;
	Gain = 0x100A;
)

// TODO: Listener properties.
const (
	Orientation = 0x100F;
)

// TODO: Source queries.
const (
	alSourceState = 0x1010;
	alBuffersQueued = 0x1015;
	alBuffersProcessed = 0x1016;
	SourceType = 0x1027; // TODO: not documented as a query?
)

// Results from Source.State() query.
const (
	Initial = 0x1011;
	Playing = 0x1012;
	Paused = 0x1013;
	Stopped = 0x1014;
)

// TODO: Presumably results from source type query?
const (
	Static = 0x1028;
	Streaming = 0x1029;
	Undetermined = 0x1030;
)

// TODO: Source properties.
const (
	SourceRelative = 0x202;
	ConeInnerAngle = 0x1001;
	ConeOuterAngle = 0x1002;
	Direction = 0x1005;
	Looping = 0x1007;
	Buffer_ = 0x1009;
	MinGain = 0x100D;
	MaxGain = 0x100E;
	SecOffset = 0x1024;
	SampleOffset = 0x1025;
	ByteOffset = 0x1026;
	ReferenceDistance = 0x1020;
	RolloffFactor = 0x1021;
	ConeOuterGain = 0x1022;
	MaxDistance = 0x1023;
)

func GetString(param uint32) string {
	return C.GoString(C.walGetString(C.ALenum(param)));
}

func GetBoolean(param uint32) bool {
	return C.alGetBoolean(C.ALenum(param)) != alFalse;
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

// Renamed, was GetBooleanv.
func GetBooleans(param uint32, data []bool) {
	C.walGetBooleanv(C.ALenum(param), unsafe.Pointer(&data[0]));
}

// Renamed, was GetIntegerv.
func GetIntegers(param uint32, data []int32) {
	C.walGetIntegerv(C.ALenum(param), unsafe.Pointer(&data[0]));
}

// Renamed, was GetFloatv.
func GetFloats(param uint32, data []float32) {
	C.walGetFloatv(C.ALenum(param), unsafe.Pointer(&data[0]));
}

// Renamed, was GetDoublev.
func GetDoubles(param uint32, data []float64) {
	C.walGetDoublev(C.ALenum(param), unsafe.Pointer(&data[0]));
}

// Error codes from GetError()/for GetString().
const (
	NoError = alFalse;
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
	return C.alIsSource(C.ALuint(id)) != alFalse;
}

// IsBuffer() returns true if id refers to a buffer.
// Not very useful as we provide a distinct Buffer type.
func IsBuffer(id uint32) bool {
	return C.alIsBuffer(C.ALuint(id)) != alFalse;
}

// Renamed, was DopplerFactor.
func SetDopplerFactor (value float32) {
	C.alDopplerFactor(C.ALfloat(value));
}

// Renamed, was DopplerVelocity.
func SetDopplerVelocity (value float32) {
	C.alDopplerVelocity(C.ALfloat(value));
}

// Renamed, was SpeedOfSound.
func SetSpeedOfSound (value float32) {
	C.alSpeedOfSound(C.ALfloat(value));
}

// Distance models for SetDistanceModel() and GetDistanceModel().
const (
	InverseDistance = 0xD001;
	InverseDistanceClamped = 0xD002;
	LinearDistance = 0xD003;
	LinearDistanceClamped = 0xD004;
	ExponentDistance = 0xD005;
	ExponentDistanceClamped = 0xD006;
)

// SetDistanceModel() changes the current distance model.
// Pass "None" to disable distance attenuation.
// Renamed, was DistanceModel.
func SetDistanceModel(model uint32) {
	C.alDistanceModel(C.ALenum(model));
}

///// Source /////////////////////////////////////////////////////////

// Sources represent sound emitters in 3d space.
type Source uint32;

// TODO: Attributes that can be set/queried with Buffer.*().
// (Please use convenience methods instead.)

// GenSources() creates n sources.
func GenSources(n int) (sources []Source) {
	sources = make([]Source, n);
	C.walGenSources(C.ALsizei(n), unsafe.Pointer(&sources[0]));
	return;
}

// DeleteSources() deletes the given sources.
func DeleteSources(sources []Source) {
	n := len(sources);
	C.walDeleteSources(C.ALsizei(n), unsafe.Pointer(&sources[0]));
}

// Renamed, was SourcePlayv.
func PlaySources(sources []Source) {
	C.walSourcePlayv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]));
}

// Renamed, was SourceStopv.
func StopSources(sources []Source) {
	C.walSourceStopv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]));
}

// Renamed, was SourceRewindv.
func RewindSources(sources []Source) {
	C.walSourceRewindv(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]));
}

// Renamed, was SourcePausev.
func PauseSources(sources []Source) {
	C.walSourcePausev(C.ALsizei(len(sources)), unsafe.Pointer(&sources[0]));
}

// Renamed, was Sourcef.
func (self Source) Setf(param uint32, value float32) {
	C.alSourcef(C.ALuint(self), C.ALenum(param), C.ALfloat(value));
}

// Renamed, was Source3f.
func (self Source) Set3f(param uint32, value1, value2, value3 float32) {
	C.alSource3f(C.ALuint(self), C.ALenum(param), C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3));
}

// Renamed, was Sourcefv.
func (self Source) Setfv(param uint32, values []float32) {
	C.walSourcefv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was Sourcei.
func (self Source) Seti(param uint32, value int32) {
	C.alSourcei(C.ALuint(self), C.ALenum(param), C.ALint(value));
}

// Renamed, was Source3i.
func (self Source) Set3i(param uint32, value1, value2, value3 int32) {
	C.alSource3i(C.ALuint(self), C.ALenum(param), C.ALint(value1), C.ALint(value2), C.ALint(value3));
}

// Renamed, was Sourceiv.
func (self Source) Setiv(param uint32, values []int32) {
	C.walSourceiv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was GetSourcef.
func (self Source) Getf(param uint32) float32 {
	return float32(C.walGetSourcef(C.ALuint(self), C.ALenum(param)));
}

// Renamed, was GetSource3f.
func (self Source) Get3f(param uint32) (value1, value2, value3 float32) {
	var v1, v2, v3 float32;
	C.walGetSource3f(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetSourcefv.
func (self Source) Getfv(param uint32) (values []float32) {
	C.walGetSourcefv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

// Renamed, was GetSourcei.
func (self Source) Geti(param uint32) int32 {
	return int32(C.walGetSourcei(C.ALuint(self), C.ALenum(param)));
}

// Renamed, was GetSource3i.
func (self Source) Get3i(param uint32) (value1, value2, value3 int32) {
	var v1, v2, v3 int32;
	C.walGetSource3i(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetSourceiv.
func (self Source) Getiv(param uint32) (values []int32) {
	C.walGetSourceiv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

// Renamed, was SourcePlay.
func (self Source) Play() {
	C.alSourcePlay(C.ALuint(self));
}

// Renamed, was SourceStop.
func (self Source) Stop() {
	C.alSourceStop(C.ALuint(self));
}

// Renamed, was SourceRewind.
func (self Source) Rewind() {
	C.alSourceRewind(C.ALuint(self));
}

// Renamed, was SourcePause.
func (self Source) Pause() {
	C.alSourcePause(C.ALuint(self));
}

// Renamed, was SourceQueueBuffers.
func (self Source) QueueBuffers(buffers []Buffer) {
	C.walSourceQueueBuffers(C.ALuint(self), C.ALsizei(len(buffers)), unsafe.Pointer(&buffers[0]));
}

// Renamed, was SourceUnqueueBuffers.
func (self Source) UnqueueBuffers(buffers []Buffer) {
	C.walSourceUnqueueBuffers(C.ALuint(self), C.ALsizei(len(buffers)), unsafe.Pointer(&buffers[0]));
}

///// Buffer /////////////////////////////////////////////////////////

// Buffers are storage space for sample data.
type Buffer uint32;

// Attributes that can be queried with Buffer.Geti().
// (Please use convenience methods instead.)
const (
	alFrequency = 0x2001;
	alBits = 0x2002;
	alChannels = 0x2003;
	alSize = 0x2004;
)

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

// Format of sound samples passed to Buffer.SetData().
const (
	FormatMono8 = 0x1100;
	FormatMono16 = 0x1101;
	FormatStereo8 = 0x1102;
	FormatStereo16 = 0x1103;
)

// SetData() specifies the sample data the buffer should use.
// For FormatMono16 and FormatStereo8 the data slice must be a
// multiple of two bytes long; for FormatStereo16 the data slice
// must be a multiple of four bytes long. The frequency is given
// in Hz.
// Renamed, was BufferData.
func (self Buffer) SetData(format uint32, data []byte, frequency uint32) {
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

// General

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

// Buffer

// Convenience method, see Buffer.Geti().
func (self Buffer) GetFrequency() uint32 {
	return uint32(self.Geti(alFrequency));
}

// Convenience method, see Buffer.Geti().
func (self Buffer) GetBits() uint32 {
	return uint32(self.Geti(alBits));
}

// Convenience method, see Buffer.Geti().
func (self Buffer) GetChannels() uint32 {
	return uint32(self.Geti(alChannels));
}

// Convenience method, see Buffer.Geti().
func (self Buffer) GetSize() uint32 {
	return uint32(self.Geti(alSize));
}

// Source

// Convenience method, see Source.QueueBuffers().
func (self Source) QueueBuffer(buffer Buffer) {
	C.walSourceQueueBuffer(C.ALuint(self), C.ALuint(buffer));
}

// Convenience method, see Source.QueueBuffers().
func (self Source) UnqueueBuffer() Buffer {
	return Buffer(C.walSourceUnqueueBuffer(C.ALuint(self)));
}

// Convenience method, see Source.Geti().
func (self Source) BuffersQueued() int32 {
	return self.Geti(alBuffersQueued);
}

// Convenience method, see Source.Geti().
func (self Source) BuffersProcessed() int32 {
	return self.Geti(alBuffersProcessed);
}

// Convenience method, see Source.Geti().
func (self Source) State() int32 {
	return self.Geti(alSourceState);
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
//	return C.alIsEnabled(C.ALenum(capability)) != alFalse;
//}

// These constants are documented as "not yet exposed". We
// keep them here in case they ever become valid. They are
// buffer states.
//
//const (
//	Unused = 0x2010;
//	Pending = 0x2011;
//	Processed = 0x2012;
//)
