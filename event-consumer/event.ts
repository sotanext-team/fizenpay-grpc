/* eslint-disable */
export const protobufPackage = "event_consumer";

export interface Actor {
  id: string;
  name: string;
  ipAddress: string;
  type: string;
}

export interface Target {
  id: string;
  type: string;
}

export interface DataChanged {
  before: string;
  after: string;
}

export interface CreateEventRequest {
  eventType: string;
  actor: Actor | undefined;
  target: Target | undefined;
  dataChanged: DataChanged | undefined;
}

export interface CreateEventResponse {
  id: string;
  eventType: string;
  actor: Actor | undefined;
  target: Target | undefined;
  dataChanged: DataChanged | undefined;
  createdAt: Date | undefined;
  updatedAt: Date | undefined;
  eventAt: Date | undefined;
}

export interface Event {
  CreateEvent(request: CreateEventRequest): Promise<CreateEventResponse>;
}
