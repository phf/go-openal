// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Convenience functions in pure Go.
//
// Not all convenience functions are here: those that need
// to call C code have to be in core.go instead due to cgo
// limitations, while those that are methods have to be in
// core.go due to language limitations. They should all be
// here of course, at least conceptually.

package al

// TODO: This is only needed because I decided to map
// ALenum to uint32 instead of int32. Hmmm...
func GetEnum(param uint32) uint32 {
	i := GetInteger(param);
	if i < 0 {
		panic("GetEnum: value from GetInteger() < 0");
	}
	return uint32(i);
}

// Convenience function, see GetInteger().
func GetDistanceModel() uint32 {
	return GetEnum(alDistanceModel);
}

// Convenience function, see GetFloat().
func GetDopplerFactor() float32 {
	return GetFloat(alDopplerFactor);
}

// Convenience function, see GetFloat().
func GetDopplerVelocity() float32 {
	return GetFloat(alDopplerVelocity);
}

// Convenience function, see GetFloat().
func GetSpeedOfSound() float32 {
	return GetFloat(alSpeedOfSound);
}

// Convenience function, see GetString().
func GetVendor() string {
	return GetString(alVendor);
}

// Convenience function, see GetString().
func GetVersion() string {
	return GetString(alVersion);
}

// Convenience function, see GetString().
func GetRenderer() string {
	return GetString(alRenderer);
}

// Convenience function, see GetString().
func GetExtensions() string {
	return GetString(alExtensions);
}
