import { from, lastValueFrom, reduce, takeWhile } from "rxjs";

export interface Entity {
  getProp(name: symbol): string | null;
}
export type Fact = {
  a: Entity;
  relation: symbol;
  b: Entity;
};

export class KnowledgeBase {
  private facts = new Set<Fact>();

  tellNewFact(s: Fact) {
    !this.facts.has(s) && this.facts.add(s);
  }
   hasFact(fact: Fact) {
    const predicate = (f: Fact) =>
      f.a === fact.a && f.relation === fact.relation && f.b === fact.b;
    for (const f of this.facts) {
      if (predicate(f)){
        return true
      }
    }
    return false
  }
}
