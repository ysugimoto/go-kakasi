#include <stdio.h>

#include <libkakasi.h>
#include "kakasi_wrapper.h"

char* Kakasi(int argc, char** argv, char* word)
{
  if (kakasi_getopt_argv(argc, argv)) {
    printf("kakasi initialization failed\n");
    return "";
  }

  char *result = kakasi_do(word);
  kakasi_close_kanwadict();
  return result;
}
