import { makeInitialPopulation } from "./initialPopulation.mjs";
import type { TimeTable } from "./timetable.mjs";

export const selection = (timeTables: TimeTable[]): [TimeTable, TimeTable] => {
  const topIndividuals = [
    ...timeTables.map((e, i) => ({ rate: e.calculateFitness(), baseIdx: i })),
  ].sort((a, b) => b.rate - a.rate);
  return [
    timeTables[topIndividuals[0].baseIdx],
    timeTables[topIndividuals[1].baseIdx],
  ];
};

const basePopulation = makeInitialPopulation();
console.log(basePopulation.map((e) => e.calculateFitness()));
console.log(selection(basePopulation).map((e) => e.calculateFitness()));
