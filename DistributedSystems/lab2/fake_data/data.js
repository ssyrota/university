const { faker } = require("@faker-js/faker");

const makeData = () => ({
  login: faker.internet.userName(),
  password: faker.internet.password(),
  cv: {
    Id: faker.string.uuid(),
    Hobbies: [
      {
        Id: faker.string.uuid(),
        Name: faker.music.genre(),
      },
      {
        Id: faker.string.uuid(),
        Name: faker.music.genre(),
      },
    ],
    JobHistory: Array.from({ length: Math.floor(Math.random() * 10) }, () => ({
      Id: faker.string.uuid(),
      From: faker.date.past().toISOString(),
      To: faker.date.future().toISOString(),
      Company: faker.company.name(),
      City: {
        Id: faker.string.uuid(),
        Name: faker.location.city(),
      },
    })),
  },
});

const data = Array.from({ length: 1000 }, makeData);
const fs = require("fs");
fs.writeFileSync("data.json", JSON.stringify(data, null, 2));
