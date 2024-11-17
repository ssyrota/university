#ifndef LINEAR_SYSTEM_H
#define LINEAR_SYSTEM_H

typedef float number;

typedef struct Vector {
  number *a;
  int length;
} Vector;

Vector* make_Vector(number *a, int length);

typedef struct Matrix {
  Vector **row_vectors;
  int rows;
  int cols;
} Matrix;

Matrix* make_Matrix(Vector **a, int rows, int cols);

typedef struct AugmentedMatrix {
  Matrix *matrix;
  Vector *b;
} AugmentedMatrix;

AugmentedMatrix* make_AugmentedMatrix(Matrix *matrix, Vector *b);


typedef struct LinearEquationSystem {
  AugmentedMatrix *parameters;
} LinearEquationSystem;


LinearEquationSystem* make_LinearEquationSystem(AugmentedMatrix *parameters);

Vector* LinearEquationSystem_solve_matrix(LinearEquationSystem *system);
AugmentedMatrix* LinearEquationSystem_solve_gauss(LinearEquationSystem *system);

void hello();
void print_vector(Vector *vector);
void print_matrix(Matrix *matrix);
#endif