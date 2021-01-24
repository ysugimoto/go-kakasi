CC=gcc
LIB_DIR=$(DEPS_DIR)/darwin
CFLAGS= -I$(DEPS_DIR)/include -I/usr/include -L$(LIB_DIR) -lkakasi -dynamiclib -o $(TARGET)
SOURCES=kakasi_wrapper.c
OBJECTS=$(SOURCES:.cc=.o) $(KAKASI_DYLIB)
TARGET=$(LIB_DIR)/libkakasi_wrapper.dylib

darwin: copy $(TARGET) link

copy:
	if [ -d "$(LIB_DIR)" ]; then\
		rm -rf $(LIB_DIR);\
	fi
	mkdir $(LIB_DIR)

	cp $(BUILD_DIR)/kakasi/include/* $(DEPS_DIR)/include/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.a $(LIB_DIR)/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.2.dylib $(LIB_DIR)/libkakasi.dylib
	cp $(BUILD_DIR)/kakasi/share/kakasi/* $(DEPS_DIR)/share/
	echo "package darwin" > $(LIB_DIR)/vendor.go

$(TARGET): $(OBJECTS)
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm $(TARGET) $(OBJECTS)

link:
	install_name_tool -id \
		"@rpath/deps/darwin/libkakasi.dylib" \
		"$(LIB_DIR)/libkakasi.dylib"

	install_name_tool -id \
		"@rpath/deps/darwin/libkakasi_wrapper.dylib" \
		"$(LIB_DIR)/libkakasi_wrapper.dylib"

	install_name_tool -change \
		"$(BUILD_DIR)/kakasi/lib/libkakasi.2.dylib" \
		"@loader_path/libkakasi.dylib" \
		"$(LIB_DIR)/libkakasi_wrapper.dylib"
