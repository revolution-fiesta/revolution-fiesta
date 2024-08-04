/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "api.v1";

export interface RegisterRequest {
  name: string;
  id: string;
  passwd: string;
  email: string;
  phone: string;
}

export interface RegisterResponse {
}

export interface LoginRequest {
  name: string;
  passwd: string;
}

export interface LoginResponse {
  token: string;
  sessionId: string;
}

export interface LogoutRequest {
  token: string;
  id: string;
  name: string;
}

export interface LogoutResponse {
}

function createBaseRegisterRequest(): RegisterRequest {
  return { name: "", id: "", passwd: "", email: "", phone: "" };
}

export const RegisterRequest = {
  encode(message: RegisterRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.passwd !== "") {
      writer.uint32(26).string(message.passwd);
    }
    if (message.email !== "") {
      writer.uint32(34).string(message.email);
    }
    if (message.phone !== "") {
      writer.uint32(42).string(message.phone);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.id = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.passwd = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.email = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.phone = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RegisterRequest {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      passwd: isSet(object.passwd) ? globalThis.String(object.passwd) : "",
      email: isSet(object.email) ? globalThis.String(object.email) : "",
      phone: isSet(object.phone) ? globalThis.String(object.phone) : "",
    };
  },

  toJSON(message: RegisterRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.passwd !== "") {
      obj.passwd = message.passwd;
    }
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.phone !== "") {
      obj.phone = message.phone;
    }
    return obj;
  },

  create(base?: DeepPartial<RegisterRequest>): RegisterRequest {
    return RegisterRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<RegisterRequest>): RegisterRequest {
    const message = createBaseRegisterRequest();
    message.name = object.name ?? "";
    message.id = object.id ?? "";
    message.passwd = object.passwd ?? "";
    message.email = object.email ?? "";
    message.phone = object.phone ?? "";
    return message;
  },
};

function createBaseRegisterResponse(): RegisterResponse {
  return {};
}

export const RegisterResponse = {
  encode(_: RegisterResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterResponse();
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

  fromJSON(_: any): RegisterResponse {
    return {};
  },

  toJSON(_: RegisterResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<RegisterResponse>): RegisterResponse {
    return RegisterResponse.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<RegisterResponse>): RegisterResponse {
    const message = createBaseRegisterResponse();
    return message;
  },
};

function createBaseLoginRequest(): LoginRequest {
  return { name: "", passwd: "" };
}

export const LoginRequest = {
  encode(message: LoginRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.passwd !== "") {
      writer.uint32(18).string(message.passwd);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.passwd = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): LoginRequest {
    return {
      name: isSet(object.name) ? globalThis.String(object.name) : "",
      passwd: isSet(object.passwd) ? globalThis.String(object.passwd) : "",
    };
  },

  toJSON(message: LoginRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.passwd !== "") {
      obj.passwd = message.passwd;
    }
    return obj;
  },

  create(base?: DeepPartial<LoginRequest>): LoginRequest {
    return LoginRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<LoginRequest>): LoginRequest {
    const message = createBaseLoginRequest();
    message.name = object.name ?? "";
    message.passwd = object.passwd ?? "";
    return message;
  },
};

function createBaseLoginResponse(): LoginResponse {
  return { token: "", sessionId: "" };
}

export const LoginResponse = {
  encode(message: LoginResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.sessionId !== "") {
      writer.uint32(18).string(message.sessionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.sessionId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): LoginResponse {
    return {
      token: isSet(object.token) ? globalThis.String(object.token) : "",
      sessionId: isSet(object.sessionId) ? globalThis.String(object.sessionId) : "",
    };
  },

  toJSON(message: LoginResponse): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    if (message.sessionId !== "") {
      obj.sessionId = message.sessionId;
    }
    return obj;
  },

  create(base?: DeepPartial<LoginResponse>): LoginResponse {
    return LoginResponse.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<LoginResponse>): LoginResponse {
    const message = createBaseLoginResponse();
    message.token = object.token ?? "";
    message.sessionId = object.sessionId ?? "";
    return message;
  },
};

function createBaseLogoutRequest(): LogoutRequest {
  return { token: "", id: "", name: "" };
}

export const LogoutRequest = {
  encode(message: LogoutRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LogoutRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLogoutRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.id = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): LogoutRequest {
    return {
      token: isSet(object.token) ? globalThis.String(object.token) : "",
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      name: isSet(object.name) ? globalThis.String(object.name) : "",
    };
  },

  toJSON(message: LogoutRequest): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create(base?: DeepPartial<LogoutRequest>): LogoutRequest {
    return LogoutRequest.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<LogoutRequest>): LogoutRequest {
    const message = createBaseLogoutRequest();
    message.token = object.token ?? "";
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseLogoutResponse(): LogoutResponse {
  return {};
}

export const LogoutResponse = {
  encode(_: LogoutResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LogoutResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLogoutResponse();
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

  fromJSON(_: any): LogoutResponse {
    return {};
  },

  toJSON(_: LogoutResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<LogoutResponse>): LogoutResponse {
    return LogoutResponse.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<LogoutResponse>): LogoutResponse {
    const message = createBaseLogoutResponse();
    return message;
  },
};

export type AuthServiceDefinition = typeof AuthServiceDefinition;
export const AuthServiceDefinition = {
  name: "AuthService",
  fullName: "api.v1.AuthService",
  methods: {
    login: {
      name: "Login",
      requestType: LoginRequest,
      requestStream: false,
      responseType: LoginResponse,
      responseStream: false,
      options: {},
    },
    register: {
      name: "Register",
      requestType: RegisterRequest,
      requestStream: false,
      responseType: RegisterResponse,
      responseStream: false,
      options: {},
    },
    logout: {
      name: "Logout",
      requestType: LogoutRequest,
      requestStream: false,
      responseType: LogoutResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
