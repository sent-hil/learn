#include <stdio.h>

char *my_strcopy(char dest[], char source[]) {
  int i = 0;

  while(source[i] != '\0') {
    dest[i] = source[i];
    i++;
  }

  dest[i] = '\0';
  
  return dest;
}

int main() {
  char stA[] = "Hello";
  char stB[10];

  my_strcopy(stB, stA);

  return 0;
}
