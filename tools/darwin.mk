CC=gcc
CFLAGS= -I$(DEPS_DIR)/include -I/usr/include -L$(DEPS_DIR)/darwin -lkakasi -dynamiclib -o $(TARGET)
SOURCES=kakasi_wrapper.c
OBJECTS=$(SOURCES:.cc=.o) $(KAKASI_DYLIB)
TARGET=$(DEPS_DIR)/darwin/libkakasi_wrapper.dylib

darwin: copy $(TARGET)

copy:
	cp $(BUILD_DIR)/kakasi/include/* $(DEPS_DIR)/include/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.a $(DEPS_DIR)/darwin/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.dylib $(DEPS_DIR)/darwin/

$(TARGET): $(OBJECTS)
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm $(TARGET) $(OBJECTS)
