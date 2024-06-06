from sklearn.preprocessing import MinMaxScaler
from sklearn import datasets
import numpy as np
import pandas as pd
from base_classifier import Classifier
from collections import Counter
from sklearn.metrics import accuracy_score

class Node:
    def __init__(self, feature=None, threshold=None, left=None, right=None, value=None):
        self.feature = feature
        self.threshold = threshold
        self.left = left
        self.right = right
        self.value = value

class DecisionTreeClassifier(Classifier):
    def __init__(self, max_depth=10, min_samples_split=2):
        self.max_depth = max_depth
        self.min_samples_split = min_samples_split
        self.root = None

    def fit(self, X: pd.DataFrame, y: pd.DataFrame):
        self.root = self._grow_tree(X.to_numpy(), y.to_numpy().flatten())

    def predict(self, X: pd.DataFrame):
        return np.array([self._traverse_tree(x, self.root) for x in X.to_numpy()])

    def _grow_tree(self, X, y, depth=0):
        num_samples, num_features = X.shape
        num_labels = len(np.unique(y))

        if (depth >= self.max_depth or num_labels == 1 or num_samples < self.min_samples_split):
            leaf_value = self._most_common_label(y)
            return Node(value=leaf_value)

        rnd_feats = np.random.choice(num_features, num_features, replace=False)

        best_feat, best_thresh = self._best_criteria(X, y, rnd_feats)
        left_idxs, right_idxs = self._split(X[:, best_feat], best_thresh)

        left = self._grow_tree(X[left_idxs, :], y[left_idxs], depth + 1)
        right = self._grow_tree(X[right_idxs, :], y[right_idxs], depth + 1)
        return Node(best_feat, best_thresh, left, right)

    def _best_criteria(self, X, y, feat_idxs):
        best_gain = -1
        split_idx, split_thresh = None, None
        for feat_idx in feat_idxs:
            X_column = X[:, feat_idx]
            thresholds = np.unique(X_column)
            for threshold in thresholds:
                gain = self._information_gain(y, X_column, threshold)

                if gain > best_gain:
                    best_gain = gain
                    split_idx = feat_idx
                    split_thresh = threshold

        return split_idx, split_thresh

    def _information_gain(self, y, X_column, split_thresh):
        parent_entropy = self._entropy(y)

        left_idxs, right_idxs = self._split(X_column, split_thresh)

        if len(left_idxs) == 0 or len(right_idxs) == 0:
            return 0

        num_left, num_right = len(left_idxs), len(right_idxs)
        num_total = len(y)
        left_entropy = self._entropy(y[left_idxs])
        right_entropy = self._entropy(y[right_idxs])
        child_entropy = (num_left / num_total) * left_entropy + (num_right / num_total) * right_entropy

        ig = parent_entropy - child_entropy
        return ig

    def _split(self, X_column, split_thresh):
        left_idxs = np.argwhere(X_column <= split_thresh).flatten()
        right_idxs = np.argwhere(X_column > split_thresh).flatten()
        return left_idxs, right_idxs

    def _entropy(self, y):
        hist = np.bincount(y)
        ps = hist / len(y)
        return -np.sum([p * np.log2(p) for p in ps if p > 0])

    def _most_common_label(self, y):
        counter = Counter(y)
        most_common = counter.most_common(1)[0][0]
        return most_common

    def score(self, X: pd.DataFrame, y: pd.DataFrame):
        y_pred = self.predict(X)
        return accuracy_score(y, y_pred)

    def _traverse_tree(self, x, node):
        if node.value is not None:
            return node.value

        if x[node.feature] <= node.threshold:
            return self._traverse_tree(x, node.left)
        return self._traverse_tree(x, node.right)

iris = datasets.load_iris()
X = pd.DataFrame(iris.data, columns=iris.feature_names)
y = pd.DataFrame(iris.target, columns=["target"])
tree = DecisionTreeClassifier(max_depth=10)
tree.fit(X, y)
predictions = tree.predict(X)
score = tree.score(X, y)
print("Predictions:", predictions)
print("Accuracy:", score)
