#include <stdio.h>
#include <string.h>

struct tag {
  char lname[20];
  char fname[20];
  int age;
  float rate;
};

struct tag my_struct;
void show_name(struct tag *p);

int main() {
  struct tag *st_ptr;
  st_ptr = &my_struct;

  strcpy(my_struct.lname, "Jones");
  strcpy(my_struct.fname, "Indiana");

  puts(my_struct.fname);
  puts(my_struct.lname);

  my_struct.age = 63;
  show_name(st_ptr);

  return 0;
}

void show_name(struct tag *p) {
  puts(p->fname);
  puts(p->lname);
  printf("%d\n", p->age);
}
