// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package al

// TODO: This is only needed because I decided to map
// ALenum to uint32 instead of int32. Hmmm...
func GetEnum(param uint32) uint32 {
	return uint32(GetInteger(param));
}

// Convenience function, see GetInteger().
func GetDistanceModel() uint32 {
	return GetEnum(DistanceModel);
}
