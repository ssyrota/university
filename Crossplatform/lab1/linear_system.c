#include "stdio.h"
#include "stdlib.h"
#include "linear_system.h"
#include <stdbool.h>
#include "matrix.c"

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


void hello() {
  printf("Hello from C, dude!\n");
}

Vector* LinearEquationSystem_solve_matrix(LinearEquationSystem *system) {
  if (Matrix_is_singular(system->parameters->matrix)) {
    printf("Matrix is singular\n");
    return NULL;
  }
  printf("Matrix is not singular\n");
  
  Matrix *inverse = Matrix_inverse(system->parameters->matrix);
  Vector *solution = Matrix_multiply_vector(inverse, system->parameters->b);
  return solution;
}

AugmentedMatrix* LinearEquationSystem_solve_gauss(LinearEquationSystem *system) {
  return system->parameters;
}