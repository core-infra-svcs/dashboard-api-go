# DevicesSerialLiveToolsPingCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The callback URL for the webhook target. If using this field, please also specify a sharedSecret. | [optional] 
**SharedSecret** | Pointer to **string** | A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url. | [optional] 
**HttpServer** | Pointer to [**DevicesSerialLiveToolsPingCallbackHttpServer**](DevicesSerialLiveToolsPingCallbackHttpServer.md) |  | [optional] 
**PayloadTemplate** | Pointer to [**DevicesSerialLiveToolsPingCallbackPayloadTemplate**](DevicesSerialLiveToolsPingCallbackPayloadTemplate.md) |  | [optional] 

## Methods

### NewDevicesSerialLiveToolsPingCallback

`func NewDevicesSerialLiveToolsPingCallback() *DevicesSerialLiveToolsPingCallback`

NewDevicesSerialLiveToolsPingCallback instantiates a new DevicesSerialLiveToolsPingCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsPingCallbackWithDefaults

`func NewDevicesSerialLiveToolsPingCallbackWithDefaults() *DevicesSerialLiveToolsPingCallback`

NewDevicesSerialLiveToolsPingCallbackWithDefaults instantiates a new DevicesSerialLiveToolsPingCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DevicesSerialLiveToolsPingCallback) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevicesSerialLiveToolsPingCallback) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevicesSerialLiveToolsPingCallback) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DevicesSerialLiveToolsPingCallback) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSharedSecret

`func (o *DevicesSerialLiveToolsPingCallback) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *DevicesSerialLiveToolsPingCallback) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *DevicesSerialLiveToolsPingCallback) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *DevicesSerialLiveToolsPingCallback) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetHttpServer

`func (o *DevicesSerialLiveToolsPingCallback) GetHttpServer() DevicesSerialLiveToolsPingCallbackHttpServer`

GetHttpServer returns the HttpServer field if non-nil, zero value otherwise.

### GetHttpServerOk

`func (o *DevicesSerialLiveToolsPingCallback) GetHttpServerOk() (*DevicesSerialLiveToolsPingCallbackHttpServer, bool)`

GetHttpServerOk returns a tuple with the HttpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServer

`func (o *DevicesSerialLiveToolsPingCallback) SetHttpServer(v DevicesSerialLiveToolsPingCallbackHttpServer)`

SetHttpServer sets HttpServer field to given value.

### HasHttpServer

`func (o *DevicesSerialLiveToolsPingCallback) HasHttpServer() bool`

HasHttpServer returns a boolean if a field has been set.

### GetPayloadTemplate

`func (o *DevicesSerialLiveToolsPingCallback) GetPayloadTemplate() DevicesSerialLiveToolsPingCallbackPayloadTemplate`

GetPayloadTemplate returns the PayloadTemplate field if non-nil, zero value otherwise.

### GetPayloadTemplateOk

`func (o *DevicesSerialLiveToolsPingCallback) GetPayloadTemplateOk() (*DevicesSerialLiveToolsPingCallbackPayloadTemplate, bool)`

GetPayloadTemplateOk returns a tuple with the PayloadTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTemplate

`func (o *DevicesSerialLiveToolsPingCallback) SetPayloadTemplate(v DevicesSerialLiveToolsPingCallbackPayloadTemplate)`

SetPayloadTemplate sets PayloadTemplate field to given value.

### HasPayloadTemplate

`func (o *DevicesSerialLiveToolsPingCallback) HasPayloadTemplate() bool`

HasPayloadTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


