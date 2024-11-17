#include "stdio.h"
#include "stdlib.h"
#include "linear_system.h"
#include <stdbool.h>

// VECTOR

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


// MATRIX

Matrix* make_Matrix(Vector **row_vectors, int rows, int cols) {
  Matrix* matrix = (Matrix*)malloc(sizeof(Matrix));
  matrix->row_vectors = row_vectors;
  matrix->rows = rows;
  matrix->cols = cols;
  return matrix;
}

Matrix* make_EmptyMatrix(int rows, int cols) {
  Matrix* matrix = (Matrix*)malloc(sizeof(Matrix));
  matrix->row_vectors = (Vector **)malloc(rows * sizeof(Vector *));
  for (int i = 0; i < rows; i++) {
    matrix->row_vectors[i] = make_EmptyVector(cols);
  }
  matrix->rows = rows;
  matrix->cols = cols;
  return matrix;
}

number Matrix_determinant(Matrix *matrix) {
    if (matrix->rows != matrix->cols) {
      return 0;
    }

    int n = matrix->rows;
    if (n == 1) {
      return matrix->row_vectors[0]->a[0];
    }
    if (n == 2) {
      return matrix->row_vectors[0]->a[0] * matrix->row_vectors[1]->a[1] -
             matrix->row_vectors[0]->a[1] * matrix->row_vectors[1]->a[0];
    }

    number det = 0;
    for (int p = 0; p < n; p++) {
      Vector **minor_data = (Vector **)malloc((n - 1) * sizeof(Vector *));
      for (int i = 1; i < n; i++) {
        number *row = (number *)malloc((n - 1) * sizeof(number));
        int col = 0;
        for (int j = 0; j < n; j++) {
          if (j == p) continue;
          row[col++] = matrix->row_vectors[i]->a[j];
        }
        minor_data[i - 1] = make_Vector(row, n - 1);
      }
      Matrix *minor = make_Matrix(minor_data, n - 1, n - 1);

      int sign = (p % 2 == 0) ? 1 : -1;
      det += sign * matrix->row_vectors[0]->a[p] * Matrix_determinant(minor);

      for (int i = 0; i < n - 1; i++) {
        free(minor_data[i]->a);
        free(minor_data[i]);
      }
      free(minor_data);
      free(minor);
    }

    return det;
}

Matrix* Matrix_transpose(Matrix *matrix) {
  Matrix *transposed = make_EmptyMatrix(matrix->cols, matrix->rows);
  for (int row = 0; row < transposed->rows; row++) {
    for (int col = 0; col < transposed->cols; col++) {
      transposed->row_vectors[row]->a[col] = matrix->row_vectors[col]->a[row];
    }
  }
  return transposed;
}

Matrix* Matrix_multiply_float(Matrix *matrix, float scalar) {
  Matrix *result = make_EmptyMatrix(matrix->rows, matrix->cols);
  for (int row = 0; row < result->rows; row++) {
    for (int col = 0; col < result->cols; col++) {
      result->row_vectors[row]->a[col] = matrix->row_vectors[row]->a[col] * scalar;
    }
  }
  return result;
}

Matrix* Matrix_minor(Matrix *matrix, int skip_row, int skip_col) {
  Matrix *minor = make_EmptyMatrix(matrix->rows - 1, matrix->cols - 1);
  for (int row = 0; row < minor->rows; row++) {
    minor->row_vectors[row] = Vector_rm_idx(matrix->row_vectors[row < skip_row ? row : row + 1], skip_col);
  }
  return minor;
}

Matrix* Matrix_cofactor(Matrix *matrix) {
  Matrix *cofactor = make_EmptyMatrix(matrix->rows, matrix->cols);
  for (int row = 0; row < cofactor->rows; row++) {
    for (int col = 0; col < cofactor->cols; col++) {
      Matrix *minor = Matrix_minor(matrix, row, col);
      number det = Matrix_determinant(minor);
      cofactor->row_vectors[row]->a[col] = det * ((row + col) % 2 == 0 ? 1 : -1);
      free(minor);
    }
  }
  return cofactor;
}

Matrix* Matrix_adjugate(Matrix *matrix) {
  Matrix *cofactor = Matrix_cofactor(matrix);
  Matrix *adjugate = Matrix_transpose(cofactor);
  return adjugate;
}


bool Matrix_is_singular(Matrix *matrix) {
  bool square = matrix->rows == matrix->cols;
  number det = Matrix_determinant(matrix);
  return square && det == 0;
}

Matrix* Matrix_inverse(Matrix *matrix) {
  printf("[matrix_inverse] matrix:\n");
  print_matrix(matrix);
  Matrix *adjugate = Matrix_adjugate(matrix);
  printf("[matrix_inverse] adjugate:\n");
  print_matrix(adjugate);
  number det = Matrix_determinant(matrix);
  printf("[matrix_inverse] det: %f\n", det);
  Matrix *result = Matrix_multiply_float(adjugate, 1 / det);
  printf("[matrix_inverse] result:\n");
  print_matrix(result);
  return result;
}

Vector* Matrix_multiply_vector(Matrix *matrix, Vector *vector) {
  Vector *result = make_EmptyVector(matrix->rows);
  for (int row = 0; row < matrix->rows; row++) {
    number sum = 0;
    for (int col = 0; col < matrix->cols; col++) {
      sum += matrix->row_vectors[row]->a[col] * vector->a[col];
    }
    result->a[row] = sum;
  }
  return result;
}

void print_matrix(Matrix *matrix) {
  for (int i = 0; i < matrix->rows; i++) {
    for (int j = 0; j < matrix->cols; j++) {
      printf("%f ", matrix->row_vectors[i]->a[j]);
    }
    printf("\n");
  }
}


// LINEAR EQUATION SYSTEM


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