LIB_DIR=$(DEPS_DIR)/linux

linux:
	if [ -d "$(LIB_DIR)" ]; then\
		rm -rf $(LIB_DIR);\
	fi
	mkdir $(LIB_DIR)

	cp $(BUILD_DIR)/kakasi/include/* $(DEPS_DIR)/include/
	cp $(BUILD_DIR)/kakasi/lib/libkakasi.a $(LIB_DIR)/
	cp $(BUILD_DIR)/kakasi/share/kakasi/* $(DEPS_DIR)/share/
	echo "package linux" > $(LIB_DIR)/vendor.go
