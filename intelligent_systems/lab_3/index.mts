import { makeInitialPopulation } from "./initialPopulation.mjs";
import "./scheduler.mjs";
import { Scheduler } from "./scheduler.mjs";

const POPULATION_SIZE = 15;
const MUTATION_RATE = 0.1;

const scheduler = new Scheduler(
  POPULATION_SIZE,
  MUTATION_RATE,
  (f) => f > 0.5,
  makeInitialPopulation
);

console.log(scheduler.composeSchedule());
