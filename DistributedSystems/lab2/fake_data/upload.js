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

async function main() {
  // uploadToMongo();
  await uploadToPsql();
}

main().catch(console.error);
