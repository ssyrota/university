import { createClient } from "@clickhouse/client";

const makeClient = () => {
  return createClient({
    url: "http://localhost:8123",
    username: "default",
    password: "",
    request_timeout: 300_000,
    keep_alive: {
      enabled: false,
      retry_on_expired_socket: true,
    },
  });
};

export class ClickhouseIngester {
  client = makeClient();

  constructor() {}
  async init() {
    const query = `
    CREATE TABLE IF NOT EXISTS sells (
      id UUID,
      time Date32,
      product_id UUID,
      product_name String,
      product_type String,
      store_id UUID,
      store_name String,
      user_id UUID,
      quantity UInt8,
      PRIMARY KEY (id)
    ) ENGINE = MergeTree();
    `;
    await this.client.query({
      query: query,
      format: "json",
    });
  }

  /**
   * @param {Array<import('../core/sell').Sell>} data
   */
  async ingest(data) {
    try {
      await this._ingest(data);
    } catch (e) {
      console.error("error ingesting into clickhouse", e);
    }
  }
  /**
   * @param {Array<import('../core/sell').Sell>} data
   */
  async _ingest(data) {
    const time = new Date();
    console.log("ingesting", data.length, "rows into clickhouse");
    await this.client.insert({
      table: "sells",
      values: data,
      format: "JSONEachRow",
    });
    console.log(
      "ingested",
      data.length,
      "rows into clickhouse in ",
      new Date() - time,
      "ms"
    );
  }

  async backfillForgottenData() {
    const query = `
    
    `;
    const client = await this.pool.connect();
    await client.query(query);
    console.log("psql backfilled");
  }
}
