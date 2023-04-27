import { KnowledgeBase, Entity } from "./base.mjs";

const relations = {
  ID: Symbol.for("id"),
  CHILD: Symbol.for("child"),
  PARENT: Symbol.for("parent"),
  DESCENDANT: Symbol.for("descendant"),
  PREDECESSOR: Symbol.for("predecessor"),
} as const;

const props = {
  PERSON_NAME: Symbol.for("person_name"),
};

class Person implements Entity {
  constructor(private readonly name: string) {}

  getProp(name: symbol): string | null {
    switch (name) {
      case props.PERSON_NAME:
        return this.name;
      default:
        return null;
    }
  }
}

const MomDad = new Person("MomDad");
const MomMother = new Person("MomMother");
const Mom = new Person("Mom");
const Father = new Person("Father");
const Son = new Person("Son");
const Sister = new Person("Sister");

const people = [MomDad, MomMother, Mom, Father, Son, Sister];

describe("fact existence", () => {
  it("should return true if fact exists", () => {
    const base = new KnowledgeBase();
    const uploadPeople = (base: KnowledgeBase) => {
      people.map((e) =>
        base.tellNewFact({
          a: e,
          b: e,
          relation: relations.ID,
        })
      );
    };
    uploadPeople(base);

    const res = base.hasFact({
      a: Mom,
      b: Mom,
      relation: relations.ID,
    });

    expect(res).toBeTruthy();
  });
  it("should return false if fact does not exists", () => {
    const base = new KnowledgeBase();
    const uploadPeople = (base: KnowledgeBase) => {
      people.map((e) =>
        base.tellNewFact({
          a: e,
          b: e,
          relation: relations.ID,
        })
      );
    };
    uploadPeople(base);
    const unknownObj = {
      getProp(_: symbol) {
        return null;
      },
    };
    const res = base.hasFact({
      a: unknownObj,
      b: unknownObj,
      relation: relations.ID,
    });
    expect(res).toBeFalsy();
  });
});
