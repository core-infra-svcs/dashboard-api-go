# DevicesSerialLiveToolsArpTableCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The callback URL for the webhook target. If using this field, please also specify a sharedSecret. | [optional] 
**SharedSecret** | Pointer to **string** | A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url. | [optional] 
**HttpServer** | Pointer to [**DevicesSerialLiveToolsArpTableCallbackHttpServer**](DevicesSerialLiveToolsArpTableCallbackHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**DevicesSerialLiveToolsArpTableCallbackPayloadTemplate**](DevicesSerialLiveToolsArpTableCallbackPayloadTemplate.md) |  | [optional] 

## Methods

### NewDevicesSerialLiveToolsArpTableCallback

`func NewDevicesSerialLiveToolsArpTableCallback() *DevicesSerialLiveToolsArpTableCallback`

NewDevicesSerialLiveToolsArpTableCallback instantiates a new DevicesSerialLiveToolsArpTableCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsArpTableCallbackWithDefaults

`func NewDevicesSerialLiveToolsArpTableCallbackWithDefaults() *DevicesSerialLiveToolsArpTableCallback`

NewDevicesSerialLiveToolsArpTableCallbackWithDefaults instantiates a new DevicesSerialLiveToolsArpTableCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DevicesSerialLiveToolsArpTableCallback) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevicesSerialLiveToolsArpTableCallback) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevicesSerialLiveToolsArpTableCallback) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DevicesSerialLiveToolsArpTableCallback) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSharedSecret

`func (o *DevicesSerialLiveToolsArpTableCallback) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *DevicesSerialLiveToolsArpTableCallback) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *DevicesSerialLiveToolsArpTableCallback) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *DevicesSerialLiveToolsArpTableCallback) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetHttpServer

`func (o *DevicesSerialLiveToolsArpTableCallback) GetHttpServer() DevicesSerialLiveToolsArpTableCallbackHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *DevicesSerialLiveToolsArpTableCallback) GetHttpServerOk() (*DevicesSerialLiveToolsArpTableCallbackHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *DevicesSerialLiveToolsArpTableCallback) SetHttpServer(v DevicesSerialLiveToolsArpTableCallbackHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *DevicesSerialLiveToolsArpTableCallback) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *DevicesSerialLiveToolsArpTableCallback) GetPayloadTemplate() DevicesSerialLiveToolsArpTableCallbackPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *DevicesSerialLiveToolsArpTableCallback) GetPayloadTemplateOk() (*DevicesSerialLiveToolsArpTableCallbackPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *DevicesSerialLiveToolsArpTableCallback) SetPayloadTemplate(v DevicesSerialLiveToolsArpTableCallbackPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *DevicesSerialLiveToolsArpTableCallback) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


