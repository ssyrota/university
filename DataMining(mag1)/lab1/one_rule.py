from sklearn import datasets
import pandas as pd
from base_classifier import Classifier
from sklearn.preprocessing import MinMaxScaler


class OneRule(Classifier):
    def __init__(self, feature_range):
        self.feature_range = feature_range
        super()

    def fit(self, X: pd.DataFrame, y: pd.DataFrame):
        unique_classes = y[0].unique()
        top_err_rate = 1
        rule = None
        top_feature = None
        for feature in X:
            data = X[feature]
            tagged = y.join(data).rename(
                columns={feature: "feature", 0: "class"})
            counted = tagged.groupby(["feature", "class"]).size(
            ).reset_index(name="count")

            empty = self._make_empty_table(unique_classes, counted)
            final_table = empty.merge(
                counted, on=["feature", "class"], how="left").fillna(0)
            final_table['count'] = final_table.apply(
                lambda row: row['count_y'] if row['count_y'] != 0 else row['count_x'], axis=1)
            final_table = final_table.drop(
                columns=['count_x', 'count_y'])

            err_rate = self._err_rate(final_table)
            print("pretender: ", err_rate, feature)
            if err_rate < top_err_rate:
                top_err_rate = err_rate
                rule = final_table
                top_feature = feature

        self._rule = rule
        self._top_err_rate = top_err_rate
        self._feature = top_feature
        return (top_err_rate, top_feature)

    def _make_empty_table(self, unique_classes, counted: pd.DataFrame):
        features = range(
            self.feature_range[0], self.feature_range[1])
        present_features = counted.groupby(
            "feature", as_index=False)['count'].sum()
        present_feature_classes = counted.groupby(
            "feature")['class'].nunique().reset_index()

        def fill_empty_count(feature):
            if feature in present_features['feature'] and len(unique_classes)-present_feature_classes['class'][feature] == 0:
                return 0
            return (len(y)-present_features['count'][feature])/(len(unique_classes)-present_feature_classes['class'][feature]) if feature in present_features['feature'] else len(y)/len(unique_classes)

        empty = pd.DataFrame(data=[{
            'feature': feature,
            'class': cls,
            'count': fill_empty_count(feature)}
            for feature in features for cls in unique_classes])
        return empty

    def _err_rate(self, table: pd.DataFrame):
        feature_classes = table.groupby("feature")['count'].max().reset_index()[
            'count'].sum()
        total_rows = table.groupby("feature")['count'].sum(
        ).reset_index()['count'].sum()
        return 1-feature_classes/total_rows


scaling_range = (0, 4)
one_rule_classifier = OneRule(scaling_range)
iris = datasets.load_iris()
X = pd.DataFrame(iris.data)
y = pd.DataFrame(iris.target)
scaler = MinMaxScaler(feature_range=scaling_range)
X = pd.DataFrame(scaler.fit_transform(
    X).astype(int), columns=["scaled-{}".format(original) for original in iris.feature_names])

print(X)
print(y)
print("winner: ", one_rule_classifier.fit(X, y))
