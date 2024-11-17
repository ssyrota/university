#include "stdio.h"
#include "stdlib.h"
#include "linear_system.h"


Vector make_Vector(int *a, int length) {
  Vector vector;
  vector.a = a;
  vector.length = length;
  return vector;
}

Matrix make_Matrix(int **a, int rows, int cols) {
  Matrix matrix;
  matrix.a = a;
  matrix.rows = rows;
  matrix.cols = cols;
  return matrix;
}

AugmentedMatrix make_AugmentedMatrix(Matrix *matrix, Vector *b) {
  AugmentedMatrix augmentedMatrix;
  augmentedMatrix.matrix = matrix;
  augmentedMatrix.b = b;
  return augmentedMatrix;
}

LinearEquationSystem make_LinearEquationSystem(AugmentedMatrix *parameters) {
  LinearEquationSystem system;
  system.parameters = parameters;
  return system;
}

AugmentedMatrix* LinearEquationSystem_solve_matrix(LinearEquationSystem system) {
  return system.parameters;
}

AugmentedMatrix* LinearEquationSystem_solve_gauss(LinearEquationSystem system) {
  return system.parameters;
}

void hello() {
  printf("Hello from C!\n");
}
