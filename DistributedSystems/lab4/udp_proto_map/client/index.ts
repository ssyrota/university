import * as map from "./proto/map.js";
import { UdpMessage } from "./udp.js";

class Client implements ClientI {
  constructor() {}
  public async get(keys: string[]): Promise<Map<string, string>> {
    const res = await new UdpMessage(
      "host.docker.internal",
      3008,
      map.MapGetRequest.encode({ keys }).finish()
    ).query();
    return map.MapGetResponse.decode(res).values.reduce(
      (acc, curr) => acc.set(curr.key, curr.value ?? null),
      new Map()
    );
  }
  public async set(batch: Map<string, string>): Promise<void> {
    await new UdpMessage(
      "host.docker.internal",
      3008,
      map.MapSetRequest.encode({
        values: [...batch.entries()].map(([key, value]) => ({ key, value })),
      }).finish()
    ).command();
    return;
  }
}

interface ClientI {
  get(keys: string[]): Promise<Map<string, string>>;
  set(batch: Map<string, string>): Promise<void>;
}

const main = async () => {
  const client = new Client();
  const entity1 = { key: crypto.randomUUID(), value: crypto.randomUUID() };
  const entity2 = { key: crypto.randomUUID(), value: crypto.randomUUID() };

  console.log({ initialMap: await client.get([entity1.key, entity2.key]) });

  console.log("setting entity1", new Map([[entity1.key, entity1.value]]));
  await client.set(new Map([[entity1.key, entity1.value]]));

  console.log({ updatedMap: await client.get([entity1.key, entity2.key]) });
};
main();
