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

async function main() {
  // uploadToMongo();
}

main().catch(console.error);
