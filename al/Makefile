# mostly copied from Eden Li's mysql interface
# "Who is supposed to grok this mess?" --- phf

include $(GOROOT)/src/Make.$(GOARCH)

TARG=openal/al
CGOFILES=core.go
GOFILES=util.go
CGO_LDFLAGS=-lopenal
#CLEANFILES+=example

include $(GOROOT)/src/Make.pkg

#example: install example.go
#	$(GC) example.go
#	$(LD) -o $@ example.$O
