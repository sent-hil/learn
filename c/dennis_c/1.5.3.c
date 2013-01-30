#include <stdio.h>

int main () {
  int c, nl, tabs;

  tabs = 0;
  nl = 0;
  while ((c = getchar()) != EOF) {
    if (c == '\n') {
      ++nl;
    }

    if (c == '\t') {
      ++tabs;
    }
  }

  printf("Newlines:%d \n", nl);
  printf("Tabs:%d \n", tabs);
}
