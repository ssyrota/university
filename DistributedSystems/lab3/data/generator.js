import { faker } from "@faker-js/faker";
import { Sell } from "../core/sell.js";
const products = Array.from({ length: 50_000 }, () => ({
  type: faker.commerce.product(),
  name: faker.commerce.productName(),
  id: faker.string.uuid(),
}));

const stores = Array.from({ length: 1_000 }, () => ({
  name: faker.company.name(),
  id: faker.string.uuid(),
}));

const users = Array.from({ length: 100_000 }, () => ({
  id: faker.string.uuid(),
}));

const sellsPerMb = 3_200;

export class SellsGenerator {
  constructor(mb, generationStep = 10_000) {
    this.mb = mb;
    this.generationStep = generationStep;
  }

  *generate() {
    const count = this.mb * sellsPerMb;
    const steps = Math.floor(count / this.generationStep) + 1;
    for (let i = 0; i < steps; i++) {
      console.log({ steps, i });
      const product = products[Math.floor(Math.random() * products.length)];
      const store = stores[Math.floor(Math.random() * stores.length)];
      const user = users[Math.floor(Math.random() * users.length)];
      yield Array.from(
        { length: this.generationStep },
        () =>
          new Sell({
            id: faker.string.uuid(),
            time: faker.date
              .past({ years: 1, refDate: "2024-01-01T00:00:00.000Z" })
              .toISOString()
              .split("T")[0],
            product_id: product.id,
            product_name: product.name,
            product_type: product.type,
            store_id: store.id,
            store_name: store.name,
            user_id: user.id,
            quantity: Math.floor(Math.random() * 10),
          })
      );
    }
  }
}
