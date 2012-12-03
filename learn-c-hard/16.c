#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <string.h>

typedef struct Person {
  char *name;
  int age;
  int height;
  int weight;
} Person;

Person *Person_create(char *name, int age, int height, int weight) {
  Person *who = malloc(sizeof(Person));
  assert(who != NULL);

  who->name = strdup(name);
  who->age = age;
  who->height = height;
  who->weight = weight;

  return who;
}

void Person_destroy(Person *who) {
  assert(who != NULL);

  free(who->name);
  free(who);
}

void Person_print(Person *who) {
  printf("Name: %s\n", who->name);
  printf("\tAge: %d\n", who->age);
  printf("\tHeight: %d\n", who->height);
  printf("\tWeight: %d\n", who->weight);
}

int main(int argc, char *argv[]) {
  Person *joe = Person_create(
      "Joe Alex", 32, 64, 140);
  Person_print(joe);

  joe->age += 20;
  Person_print(joe);

  Person_destroy(joe);

  return 0;
}
