from sklearn.preprocessing import MinMaxScaler
from sklearn import datasets
import numpy as np
import pandas as pd
from base_classifier import Classifier
from sklearn.metrics import accuracy_score

class KNNClassifier(Classifier):
    def __init__(self, n_neighbors=5):
        self.n_neighbors = n_neighbors

    def fit(self, X: pd.DataFrame, y: pd.DataFrame):
        self.X_train = X
        self.y_train = y

    def predict(self, X: pd.DataFrame):
        predictions = [self._predict(x) for x in X.to_numpy()]
        return np.array(predictions)

    def _predict(self, x):
        distances = [self._euclidean_distance(x, x_train) for x_train in self.X_train.to_numpy()]
        k_indices = np.argsort(distances)[:self.n_neighbors]
        k_nearest_labels = [self.y_train.iloc[i].values[0] for i in k_indices]
        most_common = np.bincount(k_nearest_labels).argmax()
        return most_common

    def _euclidean_distance(self, x1, x2):
        return np.sqrt(np.sum((x1 - x2)**2))

    def score(self, X: pd.DataFrame, y: pd.DataFrame):
        y_pred = self.predict(X)
        return accuracy_score(y, y_pred)

iris = datasets.load_iris()
X = pd.DataFrame(iris.data, columns=iris.feature_names)
y = pd.DataFrame(iris.target, columns=["target"])
scaler = MinMaxScaler()
X_scaled = pd.DataFrame(scaler.fit_transform(X), columns=X.columns)
knn = KNNClassifier(n_neighbors=3)
knn.fit(X_scaled, y)
predictions = knn.predict(X_scaled)
score = knn.score(X_scaled, y)
print("Predictions:", predictions)
print("Accuracy:", score)
