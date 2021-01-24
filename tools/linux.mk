linux: copy

copy:
	cp $(BUILD_DIR)/kakasi/include/* $(DEPS_DIR)/include/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.a $(DEPS_DIR)/linux/
	cp $(BUILD_DIR)/kakasi/share/kakasi/* $(DEPS_DIR)/share/
	echo "package linux" > $(DEPS_DIR)/linux/vendor.go
