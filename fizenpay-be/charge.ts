/**
* This file is auto-generated by nestjs-proto-gen-ts
*/

import { Observable } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';

export namespace fizenpay_be {
    export interface ChargeService {
        getOne(data: GetOneRequest, metadata?: Metadata): Observable<ChargeResponse>;
    }
    export interface GetOneRequest {
        code?: string;
    }
    export interface ChargeResponse {
        id?: string;
        name?: string;
        description?: string;
        pricingType?: string;
        redirectUrl?: string;
        cancelUrl?: string;
        merchantId?: string;
        status?: string;
        code?: string;
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

