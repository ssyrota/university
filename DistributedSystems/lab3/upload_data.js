import { SellsGenerator } from "./data/generator.js";
import { DataIngester } from "./ingester.js";

const ingester = new DataIngester();
await ingester.init();
const sells = new SellsGenerator(10_000, 300_000).generate();
let step = 0;
let mbCount = 0;
for await (const sell of sells) {
  step++;
  const size = new TextEncoder().encode(JSON.stringify(sell)).length;
  const kiloBytes = size / 1024;
  const megaBytes = kiloBytes / 1024;
  mbCount += megaBytes;
  console.log(
    `Uploading chunk ${step}. Size: ${megaBytes} MB. Total: ${mbCount} MB.`
  );

  await ingester.ingest(sell);
}
