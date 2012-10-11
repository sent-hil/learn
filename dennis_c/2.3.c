#include <stdio.h>

int strLen(char s[]) {
  int i;

  i = 0;
  while (s[i] != '\0') {
    ++i;
  }

  return i;
}

int main () {
  char c;

  while ((c = getchar()) != EOF) {
    putchar(c);
    printf("%s", c);
  }
}
