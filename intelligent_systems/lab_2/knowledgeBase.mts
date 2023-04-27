import { Predicate } from "./queryLang.mjs";

export class KnowledgeBase {
  facts = new Set<Predicate>();
  constructor(private readonly inferenceEngine: InferenceEngine) {}

  query(q: string): string[] {
    return [
      ...this.inferenceEngine
        .performQuery(this.facts, new Predicate(q))
        .values(),
    ].map((e) => e.t);
  }

  addFact(f: string) {
    this.facts = this.inferenceEngine.onNewFact(this.facts, new Predicate(f));
  }
}

export interface InferenceEngine {
  performQuery(facts: Set<Predicate>, q: Predicate): Set<Predicate>;
  onNewFact(facts: Set<Predicate>, newFact: Predicate): Set<Predicate>;
}
