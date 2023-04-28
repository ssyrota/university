import { makeInitialPopulation } from "./initialPopulation.mjs";
import { TimeTable } from "./timetable.mjs";

const POPULATION_SIZE = 100;
const MUTATION_CHANCE_PERCENTAGES = 0.1;
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
const mutation = (t: TimeTable, chance: number): TimeTable => {
  const scheduleMutatedFully = makeInitialPopulation(1)[0].schedule;
  if (Math.random() > chance) {
    return t;
  }
  const halfMutatedSchedule = t.schedule.map((e, i) => {
    return Math.random() > 0.5 ? e : scheduleMutatedFully[i];
  });
  return new TimeTable(halfMutatedSchedule);
};

const basePopulation = makeInitialPopulation(POPULATION_SIZE);
console.log(selection(basePopulation).map((e) => e.calculateFitness()));
console.log(
  crossover(selection(basePopulation)).map((e) => e.calculateFitness())
);
console.log(
  mutation(
    crossover(selection(basePopulation))[0],
    MUTATION_CHANCE_PERCENTAGES * 1000
  ).calculateFitness()
);
