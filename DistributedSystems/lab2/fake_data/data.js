const { faker } = require("@faker-js/faker");

const makeData = () => ({
  login: faker.internet.userName(),
  password: faker.internet.password(),
  cv: {
    id: faker.string.uuid(),
    hobbies: [
      {
        id: faker.string.uuid(),
        name: faker.music.genre(),
      },
      {
        id: faker.string.uuid(),
        name: faker.music.genre(),
      },
    ],
    jobs: Array.from({ length: Math.floor(Math.random() * 10) }, () => ({
      id: faker.string.uuid(),
      from: faker.date.past().toISOString(),
      to: faker.date.future().toISOString(),
      company: {
        name: faker.company.name(),
        id: faker.string.uuid(),
      },
      city: {
        id: faker.string.uuid(),
        name: faker.location.city(),
      },
    })),
  },
});

const data = Array.from({ length: 1000 }, makeData);
const fs = require("fs");
fs.writeFileSync("data.json", JSON.stringify(data, null, 2));
