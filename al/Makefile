# mostly copied from Eden Li's mysql interface
# "Who is supposed to grok this mess?" --- phf

include $(GOROOT)/src/Make.$(GOARCH)

TARG=openal/al
CGOFILES=core.go buffer.go
GOFILES=util.go
CGO_LDFLAGS=wrapper.o -lopenal
CLEANFILES+=wrapper.o

include $(GOROOT)/src/Make.pkg

# cute hack to trigger wrapper.o on make install
_cgo_.so: wrapper.o

wrapper.o: wrapper.c
	gcc $(_CGO_CFLAGS_$(GOARCH)) -fPIC -O2 -o $@ -c $^
