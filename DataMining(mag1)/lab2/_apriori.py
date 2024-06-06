import pandas as pd
from itertools import combinations

def apriori(data, min_support, min_confidence):
    def generate_candidates(itemset, length):
        return set([a.union(b) for a in itemset for b in itemset if len(a.union(b)) == length])

    def support(item):
        return float(sum([1 for transaction in data if item.issubset(transaction)])) / len(data)

    def confidence(item, consequent):
        return float(support(item)) / support(item.difference(consequent))

    data = list(map(set, data))
    itemset = set(frozenset([item]) for transaction in data for item in transaction)
    L = [set(item for item in itemset if support(item) >= min_support)]

    k = 2
    while L[k-2]:
        candidate_itemset = generate_candidates(L[k-2], k)
        L.append(set(item for item in candidate_itemset if support(item) >= min_support))
        k += 1

    rules = []
    for itemset in L[1:]:
        for item in itemset:
            for i in range(1, len(item)):
                for subset in combinations(item, i):
                    antecedent = frozenset(subset)
                    consequent = item.difference(antecedent)
                    if confidence(item, consequent) >= min_confidence:
                        rules.append((antecedent, consequent, support(item), confidence(item, consequent)))
    return rules

data = [['milk', 'bread', 'butter'],
        ['beer', 'bread'],
        ['milk', 'bread', 'beer', 'butter'],
        ['bread', 'butter'],
        ['milk', 'bread', 'butter']]

min_support = 0.5
min_confidence = 0.7

rules = apriori(data, min_support, min_confidence)
for rule in rules:
    antecedent, consequent, rule_support, rule_confidence = rule
    print(f"Rule: {set(antecedent)} -> {set(consequent)}, support: {rule_support}, confidence: {rule_confidence}")
