/**
* This file is auto-generated by nestjs-proto-gen-ts
*/

import { Observable } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';

export namespace fizenpay_be {
    export interface UsersService {
        getAllUser(data: GetAllUserRequest, metadata?: Metadata): Observable<GetAllUserResponse>;
        activateAccount(data: ActivateAccountRequest, metadata?: Metadata): Observable<ActivateAccountResponse>;
        deactivateAccount(data: DeactivateAccountRequest, metadata?: Metadata): Observable<DeactivateAccountResponse>;
    }
    // tslint:disable-next-line:no-empty-interface
    export interface Empty {
    }
    export interface FlexiblePayment {
        type?: number;
        underPaymentThresholdValue?: number;
        underPaymentCurrency?: string;
        overPaymentThresholdValue?: number;
        overPaymentCurrency?: string;
    }
    export interface Wallet {
        bip44Path?: string;
        currency?: string;
        isMainnet?: boolean;
        provider?: string;
        wssProvider?: string;
        network?: string;
        fpMasterContract?: string;
        transactionUrl?: string;
        contract?: string;
        address?: string;
    }
    export interface CryptoCurrency {
        name?: string;
        slug?: string;
        network?: string;
        address?: string;
        cmcId?: number;
        decimal?: number;
        type?: string[];
        gasLimit?: number;
        isEnabled?: boolean;
        confirmedBalance?: string;
        unconfirmedBalance?: string;
        fizenFeeBalance?: string;
        status?: number;
    }
    export interface Security {
        is2faEnabled?: boolean;
        isEmailVerifyEnabled?: boolean;
        secret?: string;
        tmpSecret?: string;
        emailCode?: string;
        emailCodeExp?: number;
    }
    export interface User {
        id?: string;
        currentStep?: number;
        isPendingApproved?: boolean;
        name?: string;
        autoWithdrawal?: number;
        flexiblePaymentSettings?: fizenpay_be.FlexiblePayment;
        email?: string;
        supportEmail?: string;
        userType?: number;
        wallets?: fizenpay_be.Wallet[];
        backupCode?: string[];
        cryptoCurrencySettings?: fizenpay_be.CryptoCurrency[];
        provider?: number;
        password?: string;
        status?: number;
        activeCode?: string;
        forgotPassword?: string;
        security?: fizenpay_be.Security;
        localCurrency?: string;
        fizenpayFeePerc?: number;
        sharedSecret?: string;
        createdAt?: google.protobuf.Timestamp;
        requestCloseAccountAt?: google.protobuf.Timestamp;
        closedAccountAt?: google.protobuf.Timestamp;
        updateAt?: google.protobuf.Timestamp;
    }
    export interface GetAllUserResponse {
        items?: fizenpay_be.User[];
        total?: number;
    }
    export interface GetAllUserRequest {
        isPendingApproved?: string;
        search?: string;
        status?: string;
        limit?: number;
        page?: number;
    }
    export interface ActivateAccountRequest {
        userId?: string;
    }
    // tslint:disable-next-line:no-empty-interface
    export interface ActivateAccountResponse {
    }
    export interface DeactivateAccountRequest {
        userId?: string;
    }
    // tslint:disable-next-line:no-empty-interface
    export interface DeactivateAccountResponse {
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

