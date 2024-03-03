import pg from "pg";
const { Pool } = pg;
import format from "pg-format";

const makePool = () => {
  return new Pool({
    user: "postgres",
    host: "localhost",
    database: "postgres",
    password: "postgres",
    port: 5432,
  });
};

export class PsqlIngester {
  pool = makePool();

  constructor() {}
  async init() {
    const query = `
    CREATE TABLE IF NOT EXISTS sells (
      id UUID PRIMARY KEY,
      time DATE NOT NULL,
      product_id UUID NOT NULL,
      product_name TEXT NOT NULL,
      product_type TEXT NOT NULL,
      store_id UUID NOT NULL,
      store_name TEXT NOT NULL,
      user_id UUID NOT NULL,
      quantity INT NOT NULL
    );
    CREATE INDEX IF NOT EXISTS sells_time ON sells (time);
    CREATE INDEX IF NOT EXISTS sells_product_id ON sells (product_id);
    CREATE INDEX IF NOT EXISTS sells_store_id ON sells (store_id);
    `;
    const client = await this.pool.connect();
    await client.query(query);
    console.log("psql initialized");
  }

  /**
   * @param {Array<import('../core/sell').Sell>} data
   */
  async ingest(data) {
    try {
      await this._ingest(data);
    } catch (e) {
      console.error("error ingesting into psql", e);
    }
  }

  /**
   * @param {Array<import('../core/sell').Sell>} data
   */
  async _ingest(data) {
    const time = new Date();
    console.log("ingesting", data.length, "rows into psql");
    const client = await this.pool.connect();

    const values = data.map((sell) => [
      sell.id,
      sell.time,
      sell.product_id,
      sell.product_name,
      sell.product_type,
      sell.store_id,
      sell.store_name,
      sell.user_id,
      sell.quantity,
    ]);
    const query = format(
      `INSERT INTO sells (id, time, product_id, product_name, product_type, store_id, store_name, user_id, quantity) VALUES %L ON CONFLICT (id) DO NOTHING;`,
      values
    );
    await client.query({ text: query });
    console.log(
      "ingested",
      data.length,
      "rows into psql in",
      new Date() - time,
      "ms"
    );
    client.release();
  }

  async backfillForgottenData() {
    const query = `
    CREATE TABLE IF NOT EXISTS
      prices 
    AS 
      select 
        distinct product_id, 
        floor(random() * 10 + 1)::int 
    FROM public.sells;
    `;
    const client = await this.pool.connect();
    await client.query(query);
    console.log("psql backfilled");
  }
}
