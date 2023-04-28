import { makeInitialPopulation } from "./initialPopulation.mjs";
import { TimeTable } from "./timetable.mjs";

type Offspring = [TimeTable, TimeTable];

export class Scheduler {
  constructor(
    private readonly populationSize: number,
    private readonly mutationChance: number,
    private readonly terminationF: (f: number) => boolean,
    private readonly populationGenerator: (count: number) => TimeTable[]
  ) {}

  composeSchedule(): TimeTable | null {
    let basePopulation = this.populationGenerator(this.populationSize);
    let step = 0;
    let terminationMatch: TimeTable[] = [];

    let i = 0;
    while (i < 10e6) {
      terminationMatch = basePopulation.filter((e) =>
        this.terminationF(e.calculateFitness())
      );
      if (terminationMatch.length > 0) {
        break;
      }
      const offspring = this.crossover(
        this.performSelection(basePopulation)
      ).map((e) => this.mutate(e));
      const fitnessSortedPopulation = this.makeSortedPopulation(
        basePopulation
      ).slice(0, -2);

      console.log(`generation step:${step}`);
      console.log(fitnessSortedPopulation[0]);

      basePopulation = [
        ...basePopulation.filter((_, i) =>
          fitnessSortedPopulation.map((e) => e.baseIdx).includes(i)
        ),
        ...offspring,
      ];
      step += 1;
    }

    return terminationMatch.at(0) || null;
  }

  private performSelection(timeTables: TimeTable[]): [TimeTable, TimeTable] {
    const topSortedIndividuals = this.makeSortedPopulation(timeTables);
    return [
      timeTables[topSortedIndividuals[0].baseIdx],
      timeTables[topSortedIndividuals[1].baseIdx],
    ];
  }

  private makeSortedPopulation(timeTables: TimeTable[]) {
    return [
      ...timeTables.map((e, i) => ({ rate: e.calculateFitness(), baseIdx: i })),
    ].sort((a, b) => b.rate - a.rate);
  }

  private crossover(parents: [TimeTable, TimeTable]): Offspring {
    const crossoverPoint = Math.floor(
      Math.random() * parents[0].schedule.length
    );
    const offspring1 = new TimeTable([...parents[0].schedule]);
    const offspring2 = new TimeTable([...parents[1].schedule]);
    [...new Array(crossoverPoint)].forEach((_, i) => {
      offspring1.schedule[i] = parents[1].schedule[i];
      offspring2.schedule[i] = parents[0].schedule[i];
    });
    return [offspring1, offspring2];
  }

  private mutate(t: TimeTable): TimeTable {
    const scheduleMutatedFully = makeInitialPopulation(1)[0].schedule;
    if (Math.random() > this.mutationChance) {
      return t;
    }
    const halfMutatedSchedule = t.schedule.map((e, i) => {
      return Math.random() > 0.5 ? e : scheduleMutatedFully[i];
    });
    return new TimeTable(halfMutatedSchedule);
  }
}
