// Copyright 2009 Peter H. Froehlich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
// void alSourcef( ALuint sid, ALenum param, ALfloat value ); 
// void alSource3f( ALuint sid, ALenum param, ALfloat value1, ALfloat value2, ALfloat value3 );
// void alSourcefv( ALuint sid, ALenum param, const ALfloat* values ); 
void walSourcefv(ALuint sid, ALenum param, const void* values) {
	alSourcefv(sid, param, values);
}
// void alSourcei( ALuint sid, ALenum param, ALint value ); 
// void alSource3i( ALuint sid, ALenum param, ALint value1, ALint value2, ALint value3 );
// void alSourceiv( ALuint sid, ALenum param, const ALint* values );
void walSourceiv(ALuint sid, ALenum param, const void* values) {
	alSourceiv(sid, param, values);
}
// void alGetSourcef( ALuint sid, ALenum param, ALfloat* value );
ALfloat walGetSourcef(ALuint sid, ALenum param) {
	ALfloat result;
	alGetSourcef(sid, param, &result);
	return result;
}
// void alGetSource3f( ALuint sid, ALenum param, ALfloat* value1, ALfloat* value2, ALfloat* value3);
void walGetSource3f(ALuint sid, ALenum param, void *value1, void *value2, void *value3) {
	alGetSource3f(sid, param, value1, value2, value3);
}
// void alGetSourcefv( ALuint sid, ALenum param, ALfloat* values );
void walGetSourcefv(ALuint sid, ALenum param, void* values) {
	alGetSourcefv(sid, param, values);
}
// void alGetSourcei( ALuint sid,  ALenum param, ALint* value );
ALint walGetSourcei(ALuint sid, ALenum param) {
	ALint result;
	alGetSourcei(sid, param, &result);
	return result;
}
// void alGetSource3i( ALuint sid, ALenum param, ALint* value1, ALint* value2, ALint* value3);
void walGetSource3i(ALuint sid, ALenum param, void *value1, void *value2, void *value3) {
	alGetSource3i(sid, param, value1, value2, value3);
}
// void alGetSourceiv( ALuint sid,  ALenum param, ALint* values );
void walGetSourceiv(ALuint sid, ALenum param, void* values) {
	alGetSourceiv(sid, param, values);
}
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
