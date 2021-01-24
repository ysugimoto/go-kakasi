CC=gcc
CFLAGS= -I$(DEPS_DIR)/include -I/usr/include -L$(DEPS_DIR)/darwin -lkakasi -dynamiclib -o $(TARGET)
SOURCES=kakasi_wrapper.c
OBJECTS=$(SOURCES:.cc=.o) $(KAKASI_DYLIB)
TARGET=$(DEPS_DIR)/darwin/libkakasi_wrapper.dylib

darwin: copy $(TARGET) link

copy:
	if [ -d "$(DEPS_DIR)/darwin" ]; then\
		rm -rf $(DEPS_DIR)/darwin;\
	fi
	mkdir $(DEPS_DIR)/darwin

	cp $(BUILD_DIR)/kakasi/include/* $(DEPS_DIR)/include/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.a $(DEPS_DIR)/darwin/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.2.dylib $(DEPS_DIR)/darwin/libkakasi.dylib
	cp $(BUILD_DIR)/kakasi/share/kakasi/* $(DEPS_DIR)/share/
	echo "package darwin" > $(DEPS_DIR)/darwin/vendor.go

$(TARGET): $(OBJECTS)
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm $(TARGET) $(OBJECTS)

link:
	install_name_tool -id \
		"@rpath/deps/darwin/libkakasi.dylib" \
		"$(DEPS_DIR)/darwin/libkakasi.dylib"

	install_name_tool -id \
		"@rpath/deps/darwin/libkakasi_wrapper.dylib" \
		"$(DEPS_DIR)/darwin/libkakasi_wrapper.dylib"

	install_name_tool -change \
		"$(BUILD_DIR)/kakasi/lib/libkakasi.2.dylib" \
		"@loader_path/libkakasi.dylib" \
		"$(DEPS_DIR)/darwin/libkakasi_wrapper.dylib"
