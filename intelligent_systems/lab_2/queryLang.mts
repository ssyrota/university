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
}
export class Rule {
  parsed: {
    l: Predicate[];
    r: Predicate;
  };
  constructor(readonly text: string) {
    this.parsed = {
      l: text
        .split("=>")[0]
        .split("&")
        .map((e) => e.trim())
        .map((e) => new Predicate(e)),
      r: new Predicate(text.split("=>")[1].trim()),
    };
  }
}
