// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is a C-level binding for OpenAL's "alc" API. Please
// consider using the Go-level binding instead.
package alc

/*
#include <stdlib.h>

// It's sad but the OpenAL C API uses lots and lots of typedefs
// that require wrapper functions (using basic C types) for cgo
// to grok them. So there's a lot more C code here than I would
// like...

#include <AL/alc.h>

// I keep all the alc.h prototypes here for now, for reference.
// They'll go away eventually. Those commented out are already
// accessible from Go.

ALCcontext *    alcCreateContext( ALCdevice *device, const ALCint* attrlist );
ALCboolean      alcMakeContextCurrent( ALCcontext *context );
void            alcProcessContext( ALCcontext *context );
void            alcSuspendContext( ALCcontext *context );
void            alcDestroyContext( ALCcontext *context );
ALCcontext *    alcGetCurrentContext( void );
ALCdevice*      alcGetContextsDevice( ALCcontext *context );
ALCdevice *     alcOpenDevice( const ALCchar *devicename );
ALCboolean      alcCloseDevice( ALCdevice *device );
//ALCenum         alcGetError( ALCdevice *device );
ALCboolean      alcIsExtensionPresent( ALCdevice *device, const ALCchar *extname );
void  *         alcGetProcAddress( ALCdevice *device, const ALCchar *funcname );
ALCenum         alcGetEnumValue( ALCdevice *device, const ALCchar *enumname );
const ALCchar * alcGetString( ALCdevice *device, ALCenum param );
void            alcGetIntegerv( ALCdevice *device, ALCenum param, ALCsizei size, ALCint *data );
ALCdevice*      alcCaptureOpenDevice( const ALCchar *devicename, ALCuint frequency, ALCenum format, ALCsizei buffersize );
ALCboolean      alcCaptureCloseDevice( ALCdevice *device );
void            alcCaptureStart( ALCdevice *device );
void            alcCaptureStop( ALCdevice *device );
void            alcCaptureSamples( ALCdevice *device, ALCvoid *buffer, ALCsizei samples );
*/
import "C"
//import "unsafe"

// Error codes returned by Device.GetError().
const (
	NoError = 0;
	InvalidDevice =0xA001;
	InvalidContext = 0xA002;
	InvalidEnum = 0xA003;
	InvalidValue = 0xA004;
	OutOfMemory = 0xA005;
)

type Device struct {
	handle *C.ALCdevice;
}

// GetError() returns the most recent error generated
// in the AL state machine.
func (self Device) GetError() uint32 {
	return uint32(C.alcGetError(self.handle));
}

