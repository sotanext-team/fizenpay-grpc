/**
* This file is auto-generated by nestjs-proto-gen-ts
*/

import { Observable } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';

export namespace event_consumer {
    export interface WebhookService {
        subscribeWebhook(data: SubscribeWebhookRequest, metadata?: Metadata): Observable<SubscribeWebhookResponse>;
        findAllTypes(data: Empty, metadata?: Metadata): Observable<WebhooksResponse>;
        findAll(data: FindAllRequest, metadata?: Metadata): Observable<WebhooksResponse>;
        unsubscribeWebhook(data: UnsubscribeWebhookRequest, metadata?: Metadata): Observable<UnsubscribeWebhookResponse>;
    }
    export interface FindAllRequest {
        userId?: string;
    }
    export interface SubscribeWebhookRequest {
        userId?: string;
        webhookId?: string;
    }
    export interface SubscribeWebhookResponse {
        userId?: string;
        webhooks?: event_consumer.WebhookResponse[];
    }
    export interface UnsubscribeWebhookRequest {
        userId?: string;
        webhookId?: string;
    }
    export interface WebhookResponse {
        id?: string;
        topic?: string;
    }
    export interface WebhooksResponse {
        webhooks?: event_consumer.WebhookResponse[];
    }
    // tslint:disable-next-line:no-empty-interface
    export interface UnsubscribeWebhookResponse {
    }
    // tslint:disable-next-line:no-empty-interface
    export interface Empty {
    }
}
export namespace google {
    export namespace protobuf {
        export interface Timestamp {
            seconds?: number;
            nanos?: number;
        }
    }
}

