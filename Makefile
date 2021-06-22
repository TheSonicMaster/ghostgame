# Makefile for ghostgame
# Copyright (c) 2021 The Sonic Master

# Specify where you want ghostgame to be installed to.
PREFIX=/usr/local

all: ghostgame

ghostgame:
	go build -buildmode=exe

install:
	mkdir -pv $(DESTDIR)$(PREFIX)/bin
	install -vm755 ghostgame $(DESTDIR)$(PREFIX)/bin/ghostgame

strip:
	strip --strip-all ghostgame

uninstall:
	rm -fv $(DESTDIR)$(PREFIX)/bin/ghostgame

clean:
	rm -fv ghostgame

distclean: clean
