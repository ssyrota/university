#include "stdio.h"
#include "stdlib.h"
#include "linear_system.h"


Vector* make_Vector(int *a, int length) {
  Vector *vector = (Vector*)malloc(sizeof(Vector));
  vector->a = a;
  vector->length = length;
  return vector;
}

Matrix* make_Matrix(Vector **a, int rows, int cols) {
  Matrix* matrix = (Matrix*)malloc(sizeof(Matrix));
  matrix->a = a;
  matrix->rows = rows;
  matrix->cols = cols;
  return matrix;
}

AugmentedMatrix* make_AugmentedMatrix(Matrix *matrix, Vector *b) {
  AugmentedMatrix* augmentedMatrix = (AugmentedMatrix*)malloc(sizeof(AugmentedMatrix));
  augmentedMatrix->matrix = matrix;
  augmentedMatrix->b = b;
  return augmentedMatrix;
}

LinearEquationSystem* make_LinearEquationSystem(AugmentedMatrix *parameters) {
  LinearEquationSystem *system = (LinearEquationSystem*)malloc(sizeof(LinearEquationSystem));
  system->parameters = parameters;
  return system;
}

AugmentedMatrix* LinearEquationSystem_solve_matrix(LinearEquationSystem *system) {
  return system->parameters;
}

AugmentedMatrix* LinearEquationSystem_solve_gauss(LinearEquationSystem *system) {
  return system->parameters;
}

void hello() {
  printf("Hello from C, dude!\n");
}

void print_vector(Vector *vector) {
  for (int i = 0; i < vector->length; i++) {
    printf("%d ", vector->a[i]);
  }
  printf("\n");
}

void print_matrix(Matrix *matrix) {
  for (int i = 0; i < matrix->rows; i++) {
    for (int j = 0; j < matrix->cols; j++) {
      printf("%d ", matrix->a[i]->a[j]);
    }
    printf("\n");
  }
}
