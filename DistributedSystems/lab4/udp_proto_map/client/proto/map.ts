/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "udp_proto_map";

export interface MapGetRequest {
  keys: string[];
}

export interface MapGetResponse {
  values: MapValue[];
}

export interface MapSetRequest {
  values: MapValue[];
}

export interface MapSetResponse {
}

export interface MapValue {
  key: string;
  value?: string | undefined;
}

function createBaseMapGetRequest(): MapGetRequest {
  return { keys: [] };
}

export const MapGetRequest = {
  encode(message: MapGetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.keys) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MapGetRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMapGetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.keys.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MapGetRequest {
    return { keys: globalThis.Array.isArray(object?.keys) ? object.keys.map((e: any) => globalThis.String(e)) : [] };
  },

  toJSON(message: MapGetRequest): unknown {
    const obj: any = {};
    if (message.keys?.length) {
      obj.keys = message.keys;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MapGetRequest>, I>>(base?: I): MapGetRequest {
    return MapGetRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MapGetRequest>, I>>(object: I): MapGetRequest {
    const message = createBaseMapGetRequest();
    message.keys = object.keys?.map((e) => e) || [];
    return message;
  },
};

function createBaseMapGetResponse(): MapGetResponse {
  return { values: [] };
}

export const MapGetResponse = {
  encode(message: MapGetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.values) {
      MapValue.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MapGetResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMapGetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.values.push(MapValue.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MapGetResponse {
    return {
      values: globalThis.Array.isArray(object?.values) ? object.values.map((e: any) => MapValue.fromJSON(e)) : [],
    };
  },

  toJSON(message: MapGetResponse): unknown {
    const obj: any = {};
    if (message.values?.length) {
      obj.values = message.values.map((e) => MapValue.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MapGetResponse>, I>>(base?: I): MapGetResponse {
    return MapGetResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MapGetResponse>, I>>(object: I): MapGetResponse {
    const message = createBaseMapGetResponse();
    message.values = object.values?.map((e) => MapValue.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMapSetRequest(): MapSetRequest {
  return { values: [] };
}

export const MapSetRequest = {
  encode(message: MapSetRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.values) {
      MapValue.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MapSetRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMapSetRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.values.push(MapValue.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MapSetRequest {
    return {
      values: globalThis.Array.isArray(object?.values) ? object.values.map((e: any) => MapValue.fromJSON(e)) : [],
    };
  },

  toJSON(message: MapSetRequest): unknown {
    const obj: any = {};
    if (message.values?.length) {
      obj.values = message.values.map((e) => MapValue.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MapSetRequest>, I>>(base?: I): MapSetRequest {
    return MapSetRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MapSetRequest>, I>>(object: I): MapSetRequest {
    const message = createBaseMapSetRequest();
    message.values = object.values?.map((e) => MapValue.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMapSetResponse(): MapSetResponse {
  return {};
}

export const MapSetResponse = {
  encode(_: MapSetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MapSetResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMapSetResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MapSetResponse {
    return {};
  },

  toJSON(_: MapSetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MapSetResponse>, I>>(base?: I): MapSetResponse {
    return MapSetResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MapSetResponse>, I>>(_: I): MapSetResponse {
    const message = createBaseMapSetResponse();
    return message;
  },
};

function createBaseMapValue(): MapValue {
  return { key: "", value: undefined };
}

export const MapValue = {
  encode(message: MapValue, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MapValue {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMapValue();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.value = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MapValue {
    return {
      key: isSet(object.key) ? globalThis.String(object.key) : "",
      value: isSet(object.value) ? globalThis.String(object.value) : undefined,
    };
  },

  toJSON(message: MapValue): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== undefined) {
      obj.value = message.value;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MapValue>, I>>(base?: I): MapValue {
    return MapValue.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MapValue>, I>>(object: I): MapValue {
    const message = createBaseMapValue();
    message.key = object.key ?? "";
    message.value = object.value ?? undefined;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
