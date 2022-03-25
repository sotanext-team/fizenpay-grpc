/**
* This file is auto-generated by nestjs-proto-gen-ts
*/

import { Observable } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';

export namespace event_consumer {
    export interface EventService {
        createEvent(data: CreateEventRequest, metadata?: Metadata): Observable<CreateEventResponse>;
        findOne(data: CreateEventRequest, metadata?: Metadata): Observable<CreateEventResponse>;
    }
    export interface EventById {
        id?: number;
    }
    export interface EventResponse {
        id?: number;
        name?: string;
    }
    export interface Actor {
        id?: string;
        name?: string;
        ipAddress?: string;
        type?: string;
    }
    export interface Target {
        id?: string;
        type?: string;
    }
    export interface DataChanged {
        before?: string;
        after?: string;
    }
    export interface CreateEventRequest {
        eventType?: string;
        actor?: event_consumer.Actor;
        target?: event_consumer.Target;
        dataChanged?: event_consumer.DataChanged;
    }
    export interface CreateEventResponse {
        id?: string;
        eventType?: string;
        actor?: event_consumer.Actor;
        target?: event_consumer.Target;
        dataChanged?: event_consumer.DataChanged;
        createdAt?: google.protobuf.Timestamp;
        updatedAt?: google.protobuf.Timestamp;
        eventAt?: google.protobuf.Timestamp;
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

