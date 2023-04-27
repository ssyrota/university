import { ForwardChainingEngine } from "./forwardChainingEngine.mjs";
import { KnowledgeBase } from "./knowledgeBase.mjs";

const inferenceEngine = new ForwardChainingEngine();
[
  "Child ?child ?parent => Parent ?parent ?child",
  "Child ?child ?parent & Child ?parent ?grandparent => Grandparent ?child ?grandparent",
].map((e) => inferenceEngine.addRule(e));

const knowledgeBase = new KnowledgeBase(inferenceEngine);
["Child son mom", "Child mom grandpa"].map((e) => knowledgeBase.addFact(e));

console.log(knowledgeBase.query("Child mom son"));
// console.log(knowledgeBase.query("Grandparent son ?X"));
// console.log(knowledgeBase.query("Parent mom ?X"));
