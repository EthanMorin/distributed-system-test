codegen:
	oapi-codegen \
	-generate gin,types,strict-server,spec \
	-package handler -o api/handler/api.gen.go api/config/openapi.yaml