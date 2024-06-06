from sklearn.preprocessing import MinMaxScaler
from sklearn import datasets
import numpy as np
import pandas as pd
from base_classifier import Classifier
from sklearn.preprocessing import MinMaxScaler
from sklearn.metrics import accuracy_score

class NaiveBayesClassifier(Classifier):
    def fit(self, X: pd.DataFrame, y: pd.DataFrame):
        self.classes = np.unique(y)
        self.mean = X.groupby(y.iloc[:, 0]).mean()
        self.variance = X.groupby(y.iloc[:, 0]).var()
        self.class_prior = y.value_counts(normalize=True)

    def predict(self, X: pd.DataFrame):
        predictions = [self._predict(x) for x in X.to_numpy()]
        return np.array(predictions)

    def _predict(self, x):
        posteriors = []
        for c in self.classes:
            prior = np.log(self.class_prior[c])
            class_conditional = np.sum(np.log(self._pdf(c, x)))
            posterior = prior + class_conditional
            posteriors.append(posterior)
        return self.classes[np.argmax(posteriors)]

    def _pdf(self, class_idx, x):
        mean = self.mean.loc[class_idx].values
        var = self.variance.loc[class_idx].values
        numerator = np.exp(- (x - mean) ** 2 / (2 * var))
        denominator = np.sqrt(2 * np.pi * var)
        return numerator / denominator

    def score(self, X: pd.DataFrame, y: pd.DataFrame):
        y_pred = self.predict(X)
        return accuracy_score(y, y_pred)

iris = datasets.load_iris()
X = pd.DataFrame(iris.data, columns=iris.feature_names)
y = pd.DataFrame(iris.target, columns=["target"])
scaler = MinMaxScaler()
X_scaled = pd.DataFrame(scaler.fit_transform(X), columns=X.columns)
nb = NaiveBayesClassifier()
nb.fit(X_scaled, y)
predictions = nb.predict(X_scaled)
score = nb.score(X_scaled, y)
print("Predictions:", predictions)
print("Accuracy:", score)
