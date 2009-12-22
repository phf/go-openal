# mostly copied from Eden Li's mysql interface
# "Who is supposed to grok this mess?" --- phf

include $(GOROOT)/src/Make.$(GOARCH)

TARG=openal
CGOFILES=core.go
#GOFILES=error.go util.go
CGO_LDFLAGS=-lopenal -lalut
CLEANFILES+=example hello hey al-example

include $(GOROOT)/src/Make.pkg

example: install example.go
	$(GC) example.go
	$(LD) -o $@ example.$O

al-example: install al-example.go
	$(GC) al-example.go
	$(LD) -o $@ al-example.$O

hey: install hey.go
	$(GC) hey.go
	$(LD) -o $@ hey.$O

hello: hello.c
	gcc hello.c -c -o hello.o
	gcc hello.o -lopenal -lalut -o hello

#test.db:
#	sqlite3 test.db <create_db.sql
