/* eslint-disable */
export const protobufPackage = "webhook";

export interface WebhookById {
  id: number;
}

export interface Webhook {
  id: number;
  name: string;
}

export interface WebhookService {
  findOne(request: WebhookById): Promise<Webhook>;
}
