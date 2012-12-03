#ifndef _ex19_h
#define _ex19_h

#include "object.h"

typedef struct {
  Object proto;
  int hit_points;
} Monster;

int Monster_attack(void *self, int damange);
int Monster_init(void *self);

typedef struct {
  Object proto;

  Monster *bad_guy;

  struct Room *north;
  struct Room *south;
  struct Room *east;
  struct Room *west;
} Room;

void *Room_move(void *self, Direction direction);
int Room_attack(void *self, int damage);
int Room_init(void *self);

#endif
