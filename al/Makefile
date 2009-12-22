# mostly copied from Eden Li's mysql interface
# "Who is supposed to grok this mess?" --- phf

include $(GOROOT)/src/Make.$(GOARCH)

TARG=openal/al
CGOFILES=core.go
GOFILES=util.go
CGO_LDFLAGS=wrappers.o -lopenal
CLEANFILES+=wrappers.o

include $(GOROOT)/src/Make.pkg

# cute hack to trigger wrappers.o on make install
_cgo_.so: wrappers.o

wrappers.o: wrappers.c
	gcc -fPIC -O2 -o $@ -c $^
