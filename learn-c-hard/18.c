#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <errno.h>
#include <string.h>

void die(const char *message) {
  if(errno) {
    perror(message);
  } else {
    printf("ERROR:%s\n", message);
  }
}

typedef int (*compare_cb)(int a, int b);

int *bubble_sort(int *numbers, int count, compare_cb cmp) {
  int temp = 0;
  int i = 0;
  int j = 0;
  int *target = malloc(count * sizeof(int));
}
