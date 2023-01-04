/**
* This file is auto-generated by nestjs-proto-gen-ts
*/

import { Observable } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';

export namespace fizen_auth {
    export interface AdminTokenService {
        verifyAdminToken(data: VerifyAdminTokenRequest, metadata?: Metadata): Observable<VerifyAdminTokenResponse>;
        generateToken(data: GenerateTokenRequest, metadata?: Metadata): Observable<GenerateTokenResponse>;
    }
    export interface VerifyAdminTokenRequest {
        token?: string;
    }
    export interface VerifyAdminTokenResponse {
        id?: string;
        email?: string;
        firstName?: string;
        lastName?: string;
        createdAt?: string;
        updatedAt?: string;
        lastLoginAt?: string;
    }
    export interface GenerateTokenRequest {
        email?: string;
        firstName?: string;
        lastName?: string;
        picture?: string;
        accessToken?: string;
    }
    export interface GenerateTokenResponse {
        token?: string;
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

