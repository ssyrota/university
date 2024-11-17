#include "stdio.h"
#include "stdlib.h"
#include "linear_system.h"
#include <stdbool.h>

Vector* make_EmptyVector(int length) {
  Vector *vector = (Vector*)malloc(sizeof(Vector));
  vector->a = (number *)malloc(length * sizeof(number));
  vector->length = length;
  return vector;
}

Vector* make_Vector(number *a, int length) {
  Vector *vector = (Vector*)malloc(sizeof(Vector));
  vector->a = a;
  vector->length = length;
  return vector;
}

Vector* Vector_clone(Vector *vector) {
  Vector *clone = (Vector *)malloc(sizeof(Vector));
  clone->length = vector->length;
  clone->a = (number *)malloc(clone->length * sizeof(number));
  for (int i = 0; i < clone->length; i++) {
    clone->a[i] = vector->a[i];
  }
  return clone;
}

Vector* Vector_rm_idx(Vector *vector, int idx) {
  Vector *result = (Vector *)malloc(sizeof(Vector));
  result->length = vector->length - 1;
  result->a = (number *)malloc(result->length * sizeof(number));
  for (int i = 0; i < result->length; i++) {
    result->a[i] = vector->a[i < idx ? i : i + 1];
  }
  return result;
}

void print_vector(Vector *vector) {
  for (int i = 0; i < vector->length; i++) {
    printf("%f ", vector->a[i]);
  }
  printf("\n");
}
