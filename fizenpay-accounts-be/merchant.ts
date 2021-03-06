/**
 * This file is auto-generated by nestjs-proto-gen-ts
 */

import { Observable } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';

export namespace fizenpay_accounts {
	export interface MerchantService {
		getOne(
			data: GetOneRequest,
			metadata?: Metadata,
		): Observable<GetOneResponse>;
	}
	export interface GetOneRequest {
		email?: string;
	}
	export interface GetOneResponse {
		id?: string;
		email?: string;
		firstName?: string;
		lastName?: string;
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
