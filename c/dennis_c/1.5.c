#include <stdio.h>

int main() {
  int c;

  c = getchar();
  while (c!= EOF) {
    putchar(c);
    c = getchar();
  }

  while((c = getchar()) != EOF)
    putchar(c);
    printf("%d", c != EOF);

  printf("%d", EOF);
}
