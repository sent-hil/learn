#include <stdio.h>
#include <sys/types.h>

enum {
  TOK_NONE, TOK_OPEN, TOK_CLOSE, TOK_DOT, TOK_QUOTE, TOK_SYM, TOK_NUM
};

enum {
    // special forms
    F_QUOTE=0, F_COND, F_IF, F_AND, F_OR, F_WHILE, F_LAMBDA, F_MACRO, F_LABEL,
    F_PROGN,
    // functions
    F_EQ, F_ATOM, F_CONS, F_CAR, F_CDR, F_READ, F_EVAL, F_PRINT, F_SET, F_NOT,
    F_LOAD, F_SYMBOLP, F_NUMBERP, F_ADD, F_SUB, F_MUL, F_DIV, F_LT, F_PROG1,
    F_APPLY, F_RPLACA, F_RPLACD, F_BOUNDP, N_BUILTINS
};
#define intval(x)  (((int)(x))>>2)

#define isspecial(v) (intval(v) <= (int)F_PROGN)

typedef u_int64_t value_t;

static u_int32_t heapsize = 64*1024;//bytes
static u_int32_t toktype = TOK_NONE;
static u_int32_t SP = 0;

int main () {
  char a[3];
  int i, tr;

  a[0] = F_QUOTE;
  a[1] = F_COND;
  a[2] = TOK_CLOSE;

  for (i = 0; i < 20; i++){
    printf("%d:%d \n", i, intval(i));
  }

  return 0;
}
