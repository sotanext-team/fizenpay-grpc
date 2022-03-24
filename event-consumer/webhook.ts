/* eslint-disable */
export const protobufPackage = "event_consumer";

export interface WebhookById {
  id: number;
}

export interface WebhookResponse {
  id: number;
  name: string;
}

export interface Webhook {
  findOne(request: WebhookById): Promise<WebhookResponse>;
}
