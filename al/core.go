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

// Format of sound samples passed to BufferData().
const (
	FormatMono8 = 0x1100;
	FormatMono16 = 0x1101;
	FormatStereo8 = 0x1102;
	FormatStereo16 = 0x1103;
)

// BufferData() specifies the sample data the buffer should use.
// Depending on the format, the data slice has to be a multiple
// of two or four bytes long. The frequency is given in Hz.
func (self Buffer) BufferData(format uint32, data []byte, frequency uint32) {
	C.alBufferData(C.ALuint(self), C.ALenum(format), unsafe.Pointer(&data[0]),
		C.ALsizei(len(data)), C.ALsizei(frequency));
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
}

func (self Listener) Getiv(param uint32) (values []int32) {
	C.walGetListeneriv(C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}


// NOT CLEANED UP YET

const (
	AlBuffer = 0x1009;
)
