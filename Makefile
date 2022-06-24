compile-event-consumer:
	./node_modules/nestjs-proto-gen-ts/bin/cli --path event-consumer

compile-fizenpay-be:
	./node_modules/nestjs-proto-gen-ts/bin/cli --path fizenpay-be

compile-fizenpay-accounts-be:
	./node_modules/nestjs-proto-gen-ts/bin/cli --path fizenpay-accounts-be

compile-fizenpay-blockchain-service:
	./node_modules/nestjs-proto-gen-ts/bin/cli --path fizenpay-blockchain-service
