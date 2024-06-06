import numpy as np
import matplotlib.pyplot as plt
from sklearn.datasets import make_blobs

class KMeans:
    def __init__(self, k=3, max_iters=100, tol=0.0001):
        self.k = k
        self.max_iters = max_iters
        self.tol = tol

    def fit(self, data):
        self.centroids = {}

        # Initialize centroids with first k points in the dataset
        for i in range(self.k):
            self.centroids[i] = data[i]

        for i in range(self.max_iters):
            self.classes = {}
            for i in range(self.k):
                self.classes[i] = []

            # Assign each point to the nearest centroid
            for features in data:
                distances = [np.linalg.norm(features - self.centroids[centroid]) for centroid in self.centroids]
                classification = distances.index(min(distances))
                self.classes[classification].append(features)

            previous_centroids = dict(self.centroids)

            # Recalculate centroids
            for classification in self.classes:
                self.centroids[classification] = np.average(self.classes[classification], axis=0)

            isOptimal = True
            for centroid in self.centroids:
                original_centroid = previous_centroids[centroid]
                current_centroid = self.centroids[centroid]
                if np.sum((current_centroid - original_centroid) / original_centroid * 100.0) > self.tol:
                    isOptimal = False

            if isOptimal:
                break

    def predict(self, data):
        distances = [np.linalg.norm(data - self.centroids[centroid]) for centroid in self.centroids]
        classification = distances.index(min(distances))
        return classification

# Generate sample data
X, y = make_blobs(n_samples=300, centers=4, cluster_std=0.60, random_state=0)

# Apply k-Means
kmeans = KMeans(k=4)
kmeans.fit(X)

# Plot the results
colors = 10 * ['r', 'g', 'b', 'c', 'k', 'y', 'm']

for centroid in kmeans.centroids:
    plt.scatter(*kmeans.centroids[centroid], color='k', marker='x', s=150, linewidths=5)

for classification in kmeans.classes:
    color = colors[classification]
    for features in kmeans.classes[classification]:
        plt.scatter(*features, color=color, s=30)

plt.show()
