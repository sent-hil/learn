#include <stdio.h>

int power (int n, int base) {
  int p;

  for (p=1; n>0; --n) {
    p = p*base;
  }

  return p;
}

int main () {
  printf("%d", power(2,10));
}
