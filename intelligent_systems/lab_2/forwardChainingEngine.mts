import type { InferenceEngine } from "./knowledgeBase.mjs";
import { Rule, Predicate, Variable, Value } from "./queryLang.mjs";

export class ForwardChainingEngine implements InferenceEngine {
  rules = new Set<Rule>();
  addRule(t: string) {
    this.rules = this.rules.add(new Rule(t));
  }
  performQuery(facts: Set<Predicate>, q: Predicate): Set<Predicate> {
    return [...facts].reduce((acc, curr) => {
      const matchRes = this.matchFact(curr, q);
      matchRes && acc.add(matchRes);
      return acc;
    }, new Set<Predicate>());
  }
  private matchFact(stored: Predicate, q: Predicate): Predicate | null {
    const baseCheck =
      stored.parsed.relation === q.parsed.relation &&
      stored.parsed.terms.length === q.parsed.terms.length;
    if (!baseCheck) return null;
    return this.matchTerms({ stored, query: q });
  }
  private matchTerms({
    stored,
    query,
  }: {
    stored: Predicate;
    query: Predicate;
  }): Predicate | null {
    return stored.parsed.terms
      .map((t, i) => {
        return query.parsed.terms[i].parsed === t.parsed;
      })
      .every((e) => e === true)
      ? stored
      : null;
  }

  onNewFact(facts: Set<Predicate>, newFact: Predicate): Set<Predicate> {
    return new Set(facts).add(newFact);
  }
}
