KAKASI_INC=$(PWD)/kakasi/kakasi/include
KAKASI_LIBDIR=$(PWD)/kakasi/kakasi/lib

CC=gcc
CFLAGS= -I$(KAKASI_INC) -I/usr/include -L$(KAKASI_LIBDIR) -lkakasi -dynamiclib -o $(TARGET)
SOURCES=kakasi_wrap.cc
OBJECTS=$(SOURCES:.cc=.o) $(KAKASI_DYLIB)
TARGET=$(PWD)/kakasi/kakasi/lib/libkakasi_wrap.dylib

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm $(TARGET) $(OBJECTS)
