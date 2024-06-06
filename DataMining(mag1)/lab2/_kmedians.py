import numpy as np
import matplotlib.pyplot as plt
from sklearn.datasets import make_blobs

class KMedians:
    def __init__(self, k=3, max_iters=100, tol=0.0001):
        self.k = k
        self.max_iters = max_iters
        self.tol = tol

    def fit(self, data):
        self.medians = {}
        for i in range(self.k):
            self.medians[i] = data[i]

        for i in range(self.max_iters):
            self.classes = {}
            for i in range(self.k):
                self.classes[i] = []
            for features in data:
                distances = [np.linalg.norm(features - self.medians[median], ord=1) for median in self.medians]
                classification = distances.index(min(distances))
                self.classes[classification].append(features)

            previous_medians = dict(self.medians)
            for classification in self.classes:
                self.medians[classification] = np.median(self.classes[classification], axis=0)

            isOptimal = True
            for median in self.medians:
                original_median = previous_medians[median]
                current_median = self.medians[median]
                if np.sum((current_median - original_median) / original_median * 100.0) > self.tol:
                    isOptimal = False

            if isOptimal:
                break

    def predict(self, data):
        distances = [np.linalg.norm(data - self.medians[median], ord=1) for median in self.medians]
        classification = distances.index(min(distances))
        return classification

# Generate sample data
X, y = make_blobs(n_samples=300, centers=4, cluster_std=0.60, random_state=0)

# Apply k-Medians
kmedians = KMedians(k=4)
kmedians.fit(X)

# Plot the results
colors = 10 * ['r', 'g', 'b', 'c', 'k', 'y', 'm']

for median in kmedians.medians:
    plt.scatter(*kmedians.medians[median], color='k', marker='x', s=150, linewidths=5)

for classification in kmedians.classes:
    color = colors[classification]
    for features in kmedians.classes[classification]:
        plt.scatter(*features, color=color, s=30)

plt.show()
