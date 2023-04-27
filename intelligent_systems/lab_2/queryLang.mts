export class Value {
  parsed: string;
  constructor(private readonly t: string) {
    this.parsed = t;
  }
}
export class Variable {
  static is(t: string) {
    return t.startsWith("?");
  }
  parsed: string;
  constructor(readonly t: string) {
    this.parsed = t.replace("?", "");
  }
}
export type Term = Variable | Value;
export class Predicate {
  parsed: {
    relation: string;
    terms: Term[];
  };
  constructor(readonly t: string) {
    this.parsed = {
      relation: t.split(" ")[0],
      terms: t
        .split(" ")
        .slice(1)
        .map((e) => (Variable.is(e) ? new Variable(e) : new Value(e))),
    };
  }

  private toString() {
    return (
      this.parsed.relation +
      " " +
      this.parsed.terms.map((e) => e.parsed).join(" ")
    );
  }

  private mapTerms(f: (t: Term) => Term): Predicate {
    const predicate = new Predicate(this.toString());
    predicate.parsed.terms = predicate.parsed.terms.map((t) => f(t));
    return predicate;
  }

  public bindAssignment(assignments: Substitution): Predicate {
    return this.mapTerms((t) => {
      return t instanceof Variable && t.parsed in assignments.data
        ? assignments.data[t.parsed]
        : t;
    });
  }
}

export class Substitution {
  constructor(public readonly data: Record<string, Value>) {}
}

export class Rule {
  predicates: Predicate[];
  res: Predicate;
  constructor(readonly text: string) {
    this.predicates = text
      .split("=>")[0]
      .split("&")
      .map((e) => e.trim())
      .map((e) => new Predicate(e));

    this.res = new Predicate(text.split("=>")[1].trim());
  }
}
