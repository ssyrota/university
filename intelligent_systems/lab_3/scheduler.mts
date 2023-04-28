import { makeInitialPopulation } from "./initialPopulation.mjs";
import { TimeTable } from "./timetable.mjs";

const POPULATION_SIZE = 100;
export const selection = (timeTables: TimeTable[]): [TimeTable, TimeTable] => {
  const topIndividuals = [
    ...timeTables.map((e, i) => ({ rate: e.calculateFitness(), baseIdx: i })),
  ].sort((a, b) => b.rate - a.rate);
  return [
    timeTables[topIndividuals[0].baseIdx],
    timeTables[topIndividuals[1].baseIdx],
  ];
};

type Offspring = [TimeTable, TimeTable];
const crossover = (parents: [TimeTable, TimeTable]): Offspring => {
  const crossoverPoint = Math.floor(Math.random() * parents[0].schedule.length);
  const offspring1 = new TimeTable([...parents[0].schedule]);
  const offspring2 = new TimeTable([...parents[1].schedule]);
  [...new Array(crossoverPoint)].forEach((_, i) => {
    offspring1.schedule[i] = parents[1].schedule[i];
    offspring2.schedule[i] = parents[0].schedule[i];
  });
  return [offspring1, offspring2];
};

const basePopulation = makeInitialPopulation(POPULATION_SIZE);
console.log(selection(basePopulation).map((e) => e.calculateFitness()));
console.log(
  crossover(selection(basePopulation)).map((e) => e.calculateFitness())
);
