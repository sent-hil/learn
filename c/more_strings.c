#include <stdio.h>

void my_a();
void my_b();

int main() {
  char my_string[] = "Ted";
  char *my_name = "Ted";

  my_a();
  my_b();

  return 0;
}

void my_a() {
  char a[] = "ABCDE";
}

void my_b() {
  char *cp = "FGHIJ";
}
