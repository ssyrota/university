const { MongoClient } = require("mongodb");

const fs = require("fs");

const uploadToMongo = async () => {
  const uri = process.env.MONGO_CONNECTION_STRING;
  const client = new MongoClient(uri);

  try {
    await client.connect();
    console.log("Connected to MongoDB");

    const database = client.db(process.env.MONGO_DB);
    const collection = database.collection(process.env.MONGO_COLLECTION);
    const documents = JSON.parse(
      fs.readFileSync("data.json").toString("utf-8")
    );
    const result = await collection.insertMany(documents);
    console.log(`${result.insertedCount} documents were inserted`);
  } catch (err) {
    console.error("Error occurred while connecting to MongoDB", err);
  } finally {
    await client.close();
  }
};

const uploadToPsql = async () => {
  const { Pool } = require("pg");
  const fs = require("fs");

  const pool = new Pool({
    user: "postgres",
    host: "relational-db",
    database: "postgres",
    password: "not_so_secret",
    port: 5432,
  });

  const fileContents = fs.readFileSync("data.json", "utf8");
  const data = JSON.parse(fileContents);

  for (const user of data) {
    // Insert into users table
    await pool.query(
      "INSERT INTO users (login, password) VALUES ($1, $2) ON CONFLICT (login) DO NOTHING;",
      [user.login, user.password]
    );

    // Insert into cvs table
    await pool.query(
      'INSERT INTO cvs (id, "user") VALUES ($1, $2) ON CONFLICT (id) DO NOTHING;',
      [user.cv.id, user.login]
    );

    // Insert hobbies
    for (const hobby of user.cv.hobbies) {
      await pool.query(
        "INSERT INTO hobbies (id, name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING;",
        [hobby.id, hobby.name]
      );
      await pool.query(
        "INSERT INTO cvs_hobbies (cv_id, hobby_id) VALUES ($1, $2) ON CONFLICT DO NOTHING;",
        [user.cv.id, hobby.id]
      );
    }

    // Insert jobs and cities
    for (const job of user.cv.jobs) {
      await pool.query(
        "INSERT INTO cities (id, name) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING;",
        [job.city.id, job.city.name]
      );
      await pool.query(
        "INSERT INTO jobs (id, cv_id, city_id) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING;",
        [job.id, user.cv.id, job.city.id]
      );
    }
  }
};

const uploadToNeo4j = async () => {
  const neo4j = require("neo4j-driver");
  const fs = require("fs");
  const uri = "bolt://neo4j:7687";
  const driver = neo4j.driver(uri);
  const session = driver.session();

  const fileContents = fs.readFileSync("data.json", "utf8");
  const users = JSON.parse(fileContents);

  for (const user of users) {
    // Create User node
    await session.run(
      `MERGE (u:User {login: $login, password: $password})`,
      user
    );

    // Create CV node and relationship to User
    await session.run(
      `
            MATCH (u:User {login: $login})
            MERGE (c:CV {id: $cvId})
            MERGE (u)-[:HAS_CV]->(c)
        `,
      { login: user.login, cvId: user.cv.id }
    );

    // Create and relate Hobbies
    for (const hobby of user.cv.hobbies) {
      await session.run(
        `
                MERGE (h:Hobby {id: $id, name: $name})
                WITH h
                MATCH (c:CV {id: $cvId})
                MERGE (c)-[:HAS_HOBBY]->(h)
            `,
        { id: hobby.id, name: hobby.name, cvId: user.cv.id }
      );
    }

    // Create and relate Jobs with Cities
    for (const job of user.cv.jobs) {
      await session.run(
        `
                MERGE (city:City {id: $cityId, name: $cityName})
                MERGE (j:Job {id: $jobId})
                ON CREATE SET j.from = $from, j.to = $to, j.company = $company
                WITH j, city
                MATCH (c:CV {id: $cvId})
                MERGE (j)-[:LOCATED_IN]->(city)
                MERGE (c)-[:HAS_JOB]->(j)
            `,
        {
          cityId: job.city.id,
          cityName: job.city.name,
          jobId: job.id,
          from: job.from,
          to: job.to,
          company: job.company,
          cvId: user.cv.id,
        }
      );
    }
  }
};

async function main() {
  // await uploadToMongo();
  // await uploadToPsql();
  await uploadToNeo4j();
}

main().catch(console.error);
