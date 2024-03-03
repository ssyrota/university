import { ClickhouseIngester } from "./db/clickhouse.js";
import { PsqlIngester } from "./db/psql.js";
export class DataIngester {
  constructor() {
    this.clickhouseIngester = new ClickhouseIngester();
    this.psqlIngester = new PsqlIngester();
  }
  async init() {
    await this.clickhouseIngester.init();
    await this.psqlIngester.init();
  }

  /**
   * @param {Array<import("./core/sell.js").Sell>} data
   */
  async ingest(data) {
    await Promise.all([
      this.psqlIngester.ingest(data),
      this.clickhouseIngester.ingest(data),
    ]);
  }

  async backfillForgottenData() {
    await Promise.all([
      this.psqlIngester.backfillForgottenData(),
      this.clickhouseIngester.backfillForgottenData(),
    ]);
  }
}
