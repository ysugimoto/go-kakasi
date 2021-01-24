linux: copy

copy:
	cp $(BUILD_DIR)/kakasi/include/* $(DEPS_DIR)/include/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.a $(DEPS_DIR)/linux/
