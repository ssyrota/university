const { faker } = require("@faker-js/faker");

const companies = Array.from({ length: 20 }, () => ({
  name: faker.company.name(),
  id: faker.string.uuid(),
}));

const cities = Array.from({ length: 20 }, () => ({
  name: faker.address.city(),
  id: faker.string.uuid(),
}));

const hobbies = Array.from({ length: 20 }, () => ({
  name: faker.music.genre(),
  id: faker.string.uuid(),
}));

const makeData = () => ({
  login: faker.internet.userName(),
  password: faker.internet.password(),
  cv: {
    id: faker.string.uuid(),
    hobbies: Array.from(
      { length: Math.floor(Math.random() * 10) },
      () => hobbies[Math.floor(Math.random() * companies.length)]
    ),
    jobs: Array.from({ length: Math.floor(Math.random() * 10) }, () => ({
      id: faker.string.uuid(),
      from: faker.date.past().toISOString(),
      to: faker.date.future().toISOString(),
      company: companies[Math.floor(Math.random() * companies.length)],
      city: cities[Math.floor(Math.random() * companies.length)],
    })),
  },
});

const data = Array.from({ length: 1000 }, makeData);
const fs = require("fs");
fs.writeFileSync("data.json", JSON.stringify(data, null, 2));
