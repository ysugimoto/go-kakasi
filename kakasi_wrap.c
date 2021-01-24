#include <stdio.h>

#include <libkakasi.h>
#include "kakasi_wrap.h"

char* Kakasi(int argc, char* argv, char* word)
{
  char* result = "";
  int ret;
  printf("%d, %s", argc, argv);
  char *args[] = {"-Ha", "-Ka", "-Ja", "-Ea", "-ka", "-i", "utf-8", "-o", "utf-8", "-u"};
  ret = kakasi_getopt_argv(argc, args);
  if (ret) {
    printf("%d, %s", argc, argv);
    printf("kakasi initialization failed");
    return result;
  }
  printf("kakasi initialization success");
  return result;

  result = kakasi_do(word);
  kakasi_close_kanwadict();
  return result;
}
