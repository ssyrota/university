import * as dgram from "dgram";

export class UdpMessage {
  constructor(
    private readonly host: string,
    private readonly port: number,
    private readonly data: Uint8Array
  ) {}

  public async query(): Promise<Buffer> {
    const client = await this.makeClient();
    await new Promise<void>((res, rej) =>
      client.send([this.data], (err) => {
        if (err) {
          rej(err);
        } else res();
      })
    );
    return new Promise<Buffer>((res) => {
      const cb = (msg: Buffer) => {
        res(msg);
        client.close();
      };
      client.on("message", cb);
    });
  }

  public async command(): Promise<void> {
    const client = await this.makeClient();
    await new Promise<void>((res, rej) =>
      client.send([this.data], (err) => {
        if (err) {
          rej(err);
        } else {
          res();
        }
        client.close();
      })
    );
    return;
  }
  private async makeClient(): Promise<dgram.Socket> {
    const client = dgram.createSocket("udp4");
    await new Promise<void>((res) => client.connect(this.port, this.host, res));
    return client;
  }
}
