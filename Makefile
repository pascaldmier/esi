GO ?= go

ESI_SWAGGER_SPEC ?= https://esi.tech.ccp.is/latest/swagger.json?datasource=tranquility

.PHONY: client
client:
	swagger generate client -f swagger/swagger.flattened.json -A esi

.PHONY: server
server:
	swagger generate server -f swagger/swagger.flattened.json -A esi

.PHONY: clean
clean:
	rm -rf client restapi models
