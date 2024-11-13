#ifndef LINEAR_SYSTEM_H
#define LINEAR_SYSTEM_H

typedef struct Vector {
  int *a;
  int length;
} Vector;

Vector make_Vector(int *a, int length);

typedef struct Matrix {
  int **a;
  int rows;
  int cols;
} Matrix;

Matrix make_Matrix(int **a, int rows, int cols);

typedef struct AugmentedMatrix {
  Matrix *matrix;
  Vector *b;
} AugmentedMatrix;

AugmentedMatrix make_AugmentedMatrix(Matrix *matrix, Vector *b);


typedef struct LinearEquationSystem {
  AugmentedMatrix *parameters;
} LinearEquationSystem;


LinearEquationSystem make_LinearEquationSystem(AugmentedMatrix *parameters);

AugmentedMatrix* LinearEquationSystem_solve_matrix(LinearEquationSystem system);
AugmentedMatrix* LinearEquationSystem_solve_gauss(LinearEquationSystem system);

void hello();
#endif