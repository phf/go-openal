// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

/*
#include <stdlib.h>
#include <AL/al.h>
#include "wrapper.h"
*/
import "C"
import "unsafe"

// Results from Source.State() query.
const (
	Initial = 0x1011;
	Playing = 0x1012;
	Paused = 0x1013;
	Stopped = 0x1014;
)

// Results from Source.Type() query.
const (
	Static = 0x1028;
	Streaming = 0x1029;
	Undetermined = 0x1030;
)

// TODO: Source properties.
// Regardless of what your al.h header may claim, Pitch
// only applies to Sources, not to Listeners. And I got
// that from Chris Robinson himself.
const (
	alSourceRelative = 0x202;
	alConeInnerAngle = 0x1001;
	alConeOuterAngle = 0x1002;
	alPitch = 0x1003;
	alDirection = 0x1005;
	alLooping = 0x1007;
	alBuffer = 0x1009;
	alMinGain = 0x100D;
	alMaxGain = 0x100E;
	alReferenceDistance = 0x1020;
	alRolloffFactor = 0x1021;
	alConeOuterGain = 0x1022;
	alMaxDistance = 0x1023;
	alSecOffset = 0x1024;
	alSampleOffset = 0x1025;
	alByteOffset = 0x1026;
)

// Sources represent sound emitters in 3d space.
type Source uint32;

// NewSources() creates n sources.
// Renamed, was GenSources.
func NewSources(n int) (sources []Source) {
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
func (self Source) setf(param int32, value float32) {
	C.alSourcef(C.ALuint(self), C.ALenum(param), C.ALfloat(value));
}

// Renamed, was Source3f.
func (self Source) set3f(param int32, value1, value2, value3 float32) {
	C.alSource3f(C.ALuint(self), C.ALenum(param), C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3));
}

// Renamed, was Sourcefv.
func (self Source) setfv(param int32, values []float32) {
	C.walSourcefv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was Sourcei.
func (self Source) seti(param int32, value int32) {
	C.alSourcei(C.ALuint(self), C.ALenum(param), C.ALint(value));
}

// Renamed, was Source3i.
func (self Source) set3i(param int32, value1, value2, value3 int32) {
	C.alSource3i(C.ALuint(self), C.ALenum(param), C.ALint(value1), C.ALint(value2), C.ALint(value3));
}

// Renamed, was Sourceiv.
func (self Source) setiv(param int32, values []int32) {
	C.walSourceiv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was GetSourcef.
func (self Source) getf(param int32) float32 {
	return float32(C.walGetSourcef(C.ALuint(self), C.ALenum(param)));
}

// Renamed, was GetSource3f.
func (self Source) get3f(param int32) (value1, value2, value3 float32) {
	var v1, v2, v3 float32;
	C.walGetSource3f(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetSourcefv.
func (self Source) getfv(param int32, values []float32) {
	C.walGetSourcefv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was GetSourcei.
func (self Source) geti(param int32) int32 {
	return int32(C.walGetSourcei(C.ALuint(self), C.ALenum(param)));
}

// Renamed, was GetSource3i.
func (self Source) get3i(param int32) (value1, value2, value3 int32) {
	var v1, v2, v3 int32;
	C.walGetSource3i(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetSourceiv.
func (self Source) getiv(param int32, values []int32) {
	C.walGetSourceiv(C.ALuint(self), C.ALenum(param), unsafe.Pointer(&values[0]));
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

///// Convenience ////////////////////////////////////////////////////

// NewSource() creates a single source.
// Convenience function, see NewSources().
func NewSource() Source {
	return Source(C.walGenSource());
}

// DeleteSource() deletes a single source.
// Convenience function, see DeleteSources().
func DeleteSource(source Source) {
	C.walDeleteSource(C.ALuint(source));
}

// Convenience method, see Source.QueueBuffers().
func (self Source) QueueBuffer(buffer Buffer) {
	C.walSourceQueueBuffer(C.ALuint(self), C.ALuint(buffer));
}

// Convenience method, see Source.QueueBuffers().
func (self Source) UnqueueBuffer() Buffer {
	return Buffer(C.walSourceUnqueueBuffer(C.ALuint(self)));
}

// Source queries.
// TODO: SourceType isn't documented as a query in the
// al.h header, but it is documented that way in
// the OpenAL 1.1 specification.
const (
	alSourceState = 0x1010;
	alBuffersQueued = 0x1015;
	alBuffersProcessed = 0x1016;
	alSourceType = 0x1027;
)

// Convenience method, see Source.Geti().
func (self Source) BuffersQueued() int32 {
	return self.geti(alBuffersQueued);
}

// Convenience method, see Source.Geti().
func (self Source) BuffersProcessed() int32 {
	return self.geti(alBuffersProcessed);
}

// Convenience method, see Source.Geti().
func (self Source) State() int32 {
	return self.geti(alSourceState);
}

// Convenience method, see Source.Geti().
func (self Source) Type() int32 {
	return self.geti(alSourceType);
}

// Convenience method, see Source.Getf().
func (self Source) GetGain() (gain float32) {
	return self.getf(alGain);
}

// Convenience method, see Source.Setf().
func (self Source) SetGain(gain float32) {
	self.setf(alGain, gain);
}

// Convenience method, see Source.Getf().
func (self Source) GetMinGain() (gain float32) {
	return self.getf(alMinGain);
}

// Convenience method, see Source.Setf().
func (self Source) SetMinGain(gain float32) {
	self.setf(alMinGain, gain);
}

// Convenience method, see Source.Getf().
func (self Source) GetMaxGain() (gain float32) {
	return self.getf(alMaxGain);
}

// Convenience method, see Source.Setf().
func (self Source) SetMaxGain(gain float32) {
	self.setf(alMaxGain, gain);
}

// Convenience method, see Source.Getf().
func (self Source) GetReferenceDistance() (distance float32) {
	return self.getf(alReferenceDistance);
}

// Convenience method, see Source.Setf().
func (self Source) SetReferenceDistance(distance float32) {
	self.setf(alReferenceDistance, distance);
}

// Convenience method, see Source.Getf().
func (self Source) GetMaxDistance() (distance float32) {
	return self.getf(alMaxDistance);
}

// Convenience method, see Source.Setf().
func (self Source) SetMaxDistance(distance float32) {
	self.setf(alMaxDistance, distance);
}

// Convenience method, see Source.Getf().
func (self Source) GetPitch() (gain float32) {
	return self.getf(alPitch);
}

// Convenience method, see Source.Setf().
func (self Source) SetPitch(gain float32) {
	self.setf(alPitch, gain);
}

// Convenience method, see Source.Getf().
func (self Source) GetRolloffFactor() (gain float32) {
	return self.getf(alRolloffFactor);
}

// Convenience method, see Source.Setf().
func (self Source) SetRolloffFactor(gain float32) {
	self.setf(alRolloffFactor, gain);
}

// Convenience method, see Source.Geti().
func (self Source) GetLooping() bool {
	return self.geti(alLooping) != alFalse;
}

var bool2al map[bool]int32 = map[bool]int32{true: alTrue, false: alFalse}

// Convenience method, see Source.Seti().
func (self Source) SetLooping(yes bool) {
	self.seti(alLooping, bool2al[yes]);
}

// Convenience method, see Source.Geti().
func (self Source) GetSourceRelative() bool {
	return self.geti(alSourceRelative) != alFalse;
}

// Convenience method, see Source.Seti().
func (self Source) SetSourceRelative(yes bool) {
	self.seti(alSourceRelative, bool2al[yes]);
}

// Convenience method, see Source.Setfv().
func (self Source) SetPosition(vector Vector) {
	self.setfv(alPosition, vector[0:]);
}

// Convenience method, see Source.Getfv().
func (self Source) GetPosition() Vector {
	v := Vector{};
	self.getfv(alPosition, v[0:]);
	return v;
}

// Convenience method, see Source.Setfv().
func (self Source) SetDirection(vector Vector) {
	self.setfv(alDirection, vector[0:]);
}

// Convenience method, see Source.Getfv().
func (self Source) GetDirection() Vector {
	v := Vector{};
	self.getfv(alDirection, v[0:]);
	return v;
}

// Convenience method, see Source.Setfv().
func (self Source) SetVelocity(vector Vector) {
	self.setfv(alVelocity, vector[0:]);
}

// Convenience method, see Source.Getfv().
func (self Source) GetVelocity() Vector {
	v := Vector{};
	self.getfv(alVelocity, v[0:]);
	return v;
}

// Convenience method, see Source.Getf().
func (self Source) GetOffsetSeconds() float32 {
	return self.getf(alSecOffset);
}

// Convenience method, see Source.Setf().
func (self Source) SetOffsetSeconds(offset float32) {
	self.setf(alSecOffset, offset);
}

// Convenience method, see Source.Geti().
func (self Source) GetOffsetSamples() int32 {
	return self.geti(alSampleOffset);
}

// Convenience method, see Source.Seti().
func (self Source) SetOffsetSamples(offset int32) {
	self.seti(alSampleOffset, offset);
}

// Convenience method, see Source.Geti().
func (self Source) GetOffsetBytes() int32 {
	return self.geti(alByteOffset);
}

// Convenience method, see Source.Seti().
func (self Source) SetOffsetBytes(offset int32) {
	self.seti(alByteOffset, offset);
}

// Convenience method, see Source.Getf().
func (self Source) GetInnerAngle() float32 {
	return self.getf(alConeInnerAngle);
}

// Convenience method, see Source.Setf().
func (self Source) SetInnerAngle(offset float32) {
	self.setf(alConeInnerAngle, offset);
}

// Convenience method, see Source.Getf().
func (self Source) GetOuterAngle() float32 {
	return self.getf(alConeOuterAngle);
}

// Convenience method, see Source.Setf().
func (self Source) SetOuterAngle(offset float32) {
	self.setf(alConeOuterAngle, offset);
}

// Convenience method, see Source.Getf().
func (self Source) GetOuterGain() float32 {
	return self.getf(alConeOuterGain);
}

// Convenience method, see Source.Setf().
func (self Source) SetOuterGain(offset float32) {
	self.setf(alConeOuterGain, offset);
}

// Convenience method, see Source.Geti().
func (self Source) SetBuffer(buffer Buffer) {
	self.seti(alBuffer, int32(buffer));
}

// Convenience method, see Source.Geti().
func (self Source) GetBuffer() (buffer Buffer) {
	return Buffer(self.geti(alBuffer));
}
