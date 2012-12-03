#include <stdio.h>

int main () {
  int i;
  int *ptr;
  int my_array[] = {1,23,17,4,-5,100}; 

  ptr = my_array;

  for(i = 0; i < 6; i++) {
    printf("my_array[%d] = %d\n", i, *(ptr++));
  }

  return 0;
}
