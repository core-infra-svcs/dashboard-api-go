# dashboard-api-golang
Dashboard API for Go

The Meraki Dashboard API Golang library provides all current Meraki [dashboard API](https://developer.cisco.com/meraki/api-v1/) calls to interface with the Cisco Meraki cloud-managed platform.

This library is automatically generated using the [OpenAPI generator v5.3](https://github.com/OpenAPITools/openapi-generator) using the [Meraki OpenAPI v2 specification](https://github.com/meraki/openapi/).

## Build Automation

New versions of the API client are automatically generated by a [GitHub action](.github/workflows/main.yml) that triggers on a push to the main branch and will read the targeted version from the file `meraki-openapi-version.txt` in the root of the project. The version string in this file corresponds to the release version on the meraki openapi project page. To initiate a new client build (along with the corresponding tag and release), update the `meraki-openapi-version.txt` to the desired version and push to main.

## Running Manually

[Install the OpenAPI generator jar](https://github.com/OpenAPITools/openapi-generator#1---installation)

```shell
API_VERSION="v1.21.0"
wget https://github.com/meraki/openapi/archive/refs/tags/$API_VERSION.zip && unzip -j $API_VERSION.zip '*/spec2.json'
rm -rf client/
java -jar openapi-generator-cli.jar generate \
  -i spec2.json \
  -g go \
  -o client \
  -p enumClassPrefix=true \
  -p structPrefix=true \
  -p packageVersion=$API_VERSION \
  --package-name client \
  --git-user-id core-infra-svcs \
  --git-repo-id dashboard-api-golang \
  --git-host github.com
rm $API_VERSION.zip; rm spec2.json
```

## Usage

see /client/docs folder for code samples
 