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

// Listener properties.
const (
	alOrientation = 0x100F;
)

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
func (self Listener) setf(param int32, value float32) {
	C.alListenerf(C.ALenum(param), C.ALfloat(value));
}

// Renamed, was Listener3f.
func (self Listener) set3f(param int32, value1, value2, value3 float32) {
	C.alListener3f(C.ALenum(param), C.ALfloat(value1), C.ALfloat(value2), C.ALfloat(value3));
}

// Renamed, was Listenerfv.
func (self Listener) setfv(param int32, values []float32) {
	C.walListenerfv(C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was Listeneri.
func (self Listener) seti(param int32, value int32) {
	C.alListeneri(C.ALenum(param), C.ALint(value));
}

// Renamed, was Listener3i.
func (self Listener) set3i(param int32, value1, value2, value3 int32) {
	C.alListener3i(C.ALenum(param), C.ALint(value1), C.ALint(value2), C.ALint(value3));
}

// Renamed, was Listeneriv.
func (self Listener) setiv(param int32, values []int32) {
	C.walListeneriv(C.ALenum(param), unsafe.Pointer(&values[0]));
}

// Renamed, was GetListenerf.
func (self Listener) getf(param int32) float32 {
	return float32(C.walGetListenerf(C.ALenum(param)));
}

// Renamed, was GetListener3f.
func (self Listener) get3f(param int32) (value1, value2, value3 float32) {
	var v1, v2, v3 float32;
	C.walGetListener3f(C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetListenerfv.
func (self Listener) getfv(param int32, values []float32) {
	C.walGetListenerfv(C.ALenum(param), unsafe.Pointer(&values[0]));
	return;
}

// Renamed, was GetListeneri.
func (self Listener) geti(param int32) int32 {
	return int32(C.walGetListeneri(C.ALenum(param)));
}

// Renamed, was GetListener3i.
func (self Listener) get3i(param int32) (value1, value2, value3 int32) {
	var v1, v2, v3 int32;
	C.walGetListener3i(C.ALenum(param), unsafe.Pointer(&v1),
		unsafe.Pointer(&v2), unsafe.Pointer(&v3));
	value1, value2, value3 = v1, v2, v3;
	return;
}

// Renamed, was GetListeneriv.
func (self Listener) getiv(param int32, values []int32) {
	C.walGetListeneriv(C.ALenum(param), unsafe.Pointer(&values[0]));
}

///// Convenience ////////////////////////////////////////////////////

// Convenience method, see Listener.Setf().
func (self Listener) SetGain(gain float32) {
	self.setf(alGain, gain);
}

// Convenience method, see Listener.Getf().
func (self Listener) GetGain() (gain float32) {
	return self.getf(alGain);
}

// Convenience method, see Listener.Setfv().
func (self Listener) SetPosition(vector Vector) {
	self.setfv(alPosition, vector[0:]);
}

// Convenience method, see Listener.Getfv().
func (self Listener) GetPosition() Vector {
	v := Vector{};
	self.getfv(alPosition, v[0:]);
	return v;
}

// Convenience method, see Listener.Setfv().
func (self Listener) SetVelocity(vector Vector) {
	self.setfv(alVelocity, vector[0:]);
}

// Convenience method, see Listener.Getfv().
func (self Listener) GetVelocity() Vector {
	v := Vector{};
	self.getfv(alVelocity, v[0:]);
	return v;
}

// TODO: is there a better way to do this?
// Convenience method, see Listener.Setfv().
func (self Listener) SetOrientation(at Vector, up Vector) {
	t := [6]float32{};
	t[0] = at[0]; t[1] = at[1]; t[2] = at[2];
	t[3] = up[0]; t[4] = up[1]; t[5] = up[2];
	self.setfv(alOrientation, t[0:]);
}

// TODO: is there a better way to do this?
// Convenience method, see Listener.Getfv().
func (self Listener) GetOrientation() (at Vector, up Vector) {
	t := [6]float32{};
	self.getfv(alOrientation, t[0:]);
	at[0] = t[0]; at[1] = t[1]; at[2] = t[2];
	up[0] = t[3]; up[1] = t[4]; up[2] = t[5];
	return;
}
