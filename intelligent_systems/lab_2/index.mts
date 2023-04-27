import { ForwardChainingEngine } from "./forwardChainingEngine.mjs";
import { KnowledgeBase } from "./knowledgeBase.mjs";

const inferenceEngine = new ForwardChainingEngine();

const knowledgeBase = new KnowledgeBase(inferenceEngine);
["Child son mom", "Child mom grandpa"].map((e) => knowledgeBase.addFact(e));

[
  "Child ?child ?parent => Parent ?parent ?child",
  "Parent ?parent ?child => Child ?child ?parent",
  "Child ?child ?parent & Child ?parent ?grandparent => Grandparent ?child ?grandparent",
].map((e) => knowledgeBase.addRule(e));

console.log(knowledgeBase.query("Child ?child ?parent"));
