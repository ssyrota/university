type People = {
  [name: string]: "male" | "female";
};

const people: People = {
  SomeDad: "male",
  SomeSun: "male",
  SomeGrantDad: "male",
  SomeMama: "female",
  SomeSister1: "female",
  SomeSister2: "female",
  SomeGrantMama: "female",
};

type ParentChild = {
  [parentName: string]: string[];
};

const parentChild: ParentChild = {
  SomeDad: ["SomeSun", "SomeSister1", "SomeSister2"],
  SomeMama: ["SomeSun", "SomeSister1", "SomeSister2"],
  SomeGrantDad: ["SomeDad"],
  SomeGrantMama: ["SomeDad"],
};

function isMother(x: string, y: string): boolean {
  return (
    x in parentChild &&
    parentChild[x].some((e) => e === y) &&
    people[x] === "female"
  );
}

function isFather(x: string, y: string): boolean {
  return (
    x in parentChild &&
    parentChild[x].some((e) => e === y) &&
    people[x] === "male"
  );
}

function isSon(x: string, y: string): boolean {
  return (
    y in parentChild &&
    parentChild[y].some((e) => e === x) &&
    people[x] === "male"
  );
}

function isDaughter(x: string, y: string): boolean {
  return (
    y in parentChild &&
    parentChild[y].some((e) => e === x) &&
    people[x] === "female"
  );
}

function isGrandparent(x: string, y: string): boolean {
  return (
    x in parentChild &&
    parentChild[x].some(
      (t) => t in parentChild && parentChild[t].some((e) => e === y)
    )
  );
}

console.log(isMother("SomeMama", "SomeSun")); // true
console.log(isFather("SomeDad", "SomeSister2")); // true
console.log(isSon("SomeSun", "SomeGrantDad")); // false
console.log(isDaughter("SomeSister1", "SomeDad")); // true
console.log(isGrandparent("SomeGrantMama", "SomeSun")); // true
console.log(isGrandparent("SomeGrantDad", "SomeSister1")); // true
console.log(isGrandparent("SomeGrantDad", "SomeDad")); // false
