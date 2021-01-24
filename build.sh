#!/bin/sh

KAKASI_VERSION="2.3.6"
BUILD_DIR="${PWD}/kakasi"

if [ ! -d "${BUILD_DIR}" ]; then
  mkdir ${BUILD_DIR}
fi

cd ${BUILD_DIR}
curl -o kakasi-${KAKASI_VERSION}.tar.gz http://kakasi.namazu.org/stable/kakasi-${KAKASI_VERSION}.tar.gz
tar xvfz kakasi-${KAKASI_VERSION}.tar.gz
patch -u kakasi-${KAKASI_VERSION}/src/kanjiio.c < ../patches/kakasi.patch
patch -u kakasi-${KAKASI_VERSION}/configure < ../patches/configure.patch

cd ${BUILD_DIR}/kakasi-${KAKASI_VERSION}
if [ -f "Makefile"]; then
  make clean
fi
./configure --prefix=${BUILD_DIR}/kakasi
make && make install
