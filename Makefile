.PHONY: generate deps

VERSION=0.0.0

brew:
	rm -rf client
	rm -rf spec2.json
	wget -O spec2.json https://raw.githubusercontent.com/meraki/openapi/master/openapi/spec2.json
	openapi-generator generate \
                               			-i spec2.json \
                               			-g go \
                               			--package-name client \
                               			-p structPrefix=true \
                               			-p packageVersion=$(VERSION) \
                               			-p enumClassPrefix=true \
                               			--git-repo-id core-infra-svcs/dashboard-api-golang \
                               			--git-host github.com \
                               			-o client
docker:
		rm -rf client
		docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v5.4.0 generate \
			-i https://raw.githubusercontent.com/meraki/openapi/master/openapi/spec2.json \
			-g go \
			--package-name client \
			-p structPrefix=true \
			-p packageVersion=$(VERSION) \
			-p enumClassPrefix=true \
			--git-repo-id core-infra-svcs/dashboard-api-golang \
			--git-host github.com \
			-o client