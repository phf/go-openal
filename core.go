// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package openal

/*
#include <stdlib.h>
#include <AL/al.h>
//AL_API void AL_APIENTRY alGenSources( ALsizei n, ALuint* sources );
//AL_API void AL_APIENTRY alDeleteSources( ALsizei n, const ALuint* sources );
//AL_API void AL_APIENTRY alGenBuffers( ALsizei n, ALuint* buffers );
//AL_API void AL_APIENTRY alDeleteBuffers( ALsizei n, const ALuint* buffers );
//AL_API void AL_APIENTRY alBufferData( ALuint bid, ALenum format, const ALvoid* data, ALsizei size, ALsizei freq );
//AL_API void AL_APIENTRY alSourcePlay( ALuint sid );
//AL_API void AL_APIENTRY alSourceStop( ALuint sid );
//AL_API void AL_APIENTRY alSourceRewind( ALuint sid );
//AL_API void AL_APIENTRY alSourcePause( ALuint sid );
// AL_API void AL_APIENTRY alSourceQueueBuffers( ALuint sid, ALsizei numEntries, const ALuint *bids );
//AL_API void AL_APIENTRY alSourceUnqueueBuffers( ALuint sid, ALsizei numEntries, ALuint *bids );
//AL_API ALenum AL_APIENTRY alGetError( void );
//AL_API void AL_APIENTRY alSourcei( ALuint sid, ALenum param, ALint value );
//AL_API void AL_APIENTRY alGetSourcei( ALuint sid,  ALenum param, ALint* value );
#include <AL/alc.h>
//ALC_API ALCdevice *     ALC_APIENTRY alcOpenDevice( const ALCchar *devicename );
//ALC_API ALCboolean      ALC_APIENTRY alcCloseDevice( ALCdevice *device );
//ALC_API ALCenum         ALC_APIENTRY alcGetError( ALCdevice *device );
//ALC_API ALCcontext *    ALC_APIENTRY alcCreateContext( ALCdevice *device, const ALCint* attrlist );
//ALC_API ALCboolean      ALC_APIENTRY alcMakeContextCurrent( ALCcontext *context );
//ALC_API void            ALC_APIENTRY alcDestroyContext( ALCcontext *context );
//ALC_API ALCdevice*      ALC_APIENTRY alcCaptureOpenDevice( const ALCchar *devicename, ALCuint frequency, ALCenum format, ALCsizei buffersize );
//ALC_API ALCboolean      ALC_APIENTRY alcCaptureCloseDevice( ALCdevice *device );
//ALC_API void            ALC_APIENTRY alcCaptureStart( ALCdevice *device );
//ALC_API void            ALC_APIENTRY alcCaptureStop( ALCdevice *device );
//ALC_API void            ALC_APIENTRY alcCaptureSamples( ALCdevice *device, ALCvoid *buffer, ALCsizei samples );
//ALC_API void            ALC_APIENTRY alcGetIntegerv( ALCdevice *device, ALCenum param, ALCsizei size, ALCint *data );
#include <AL/alext.h>
#include <AL/alut.h>
*/
import "C"

import (
	"fmt";
	"unsafe";
)

func X() unsafe.Pointer {
	fmt.Println("Argh");
	C.free(nil);
	return nil;
}
