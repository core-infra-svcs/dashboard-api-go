# InlineResponse200529SsidSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Security mode of the SSID | [optional] 
**Protocol** | Pointer to **string** | Security protocol used by the SSID | [optional] 
**Encryption** | Pointer to [**InlineResponse200529SsidSecurityEncryption**](InlineResponse200529SsidSecurityEncryption.md) |  | [optional] 
**Dhcp** | Pointer to [**InlineResponse200529SsidSecurityDhcp**](InlineResponse200529SsidSecurityDhcp.md) |  | [optional] 

## Methods

### NewInlineResponse200529SsidSecurity

`func NewInlineResponse200529SsidSecurity() *InlineResponse200529SsidSecurity`

NewInlineResponse200529SsidSecurity instantiates a new InlineResponse200529SsidSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200529SsidSecurityWithDefaults

`func NewInlineResponse200529SsidSecurityWithDefaults() *InlineResponse200529SsidSecurity`

NewInlineResponse200529SsidSecurityWithDefaults instantiates a new InlineResponse200529SsidSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse200529SsidSecurity) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200529SsidSecurity) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200529SsidSecurity) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200529SsidSecurity) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse200529SsidSecurity) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200529SsidSecurity) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200529SsidSecurity) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200529SsidSecurity) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetEncryption

`func (o *InlineResponse200529SsidSecurity) GetEncryption() InlineResponse200529SsidSecurityEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *InlineResponse200529SsidSecurity) GetEncryptionOk() (*InlineResponse200529SsidSecurityEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *InlineResponse200529SsidSecurity) SetEncryption(v InlineResponse200529SsidSecurityEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *InlineResponse200529SsidSecurity) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetDhcp

`func (o *InlineResponse200529SsidSecurity) GetDhcp() InlineResponse200529SsidSecurityDhcp`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *InlineResponse200529SsidSecurity) GetDhcpOk() (*InlineResponse200529SsidSecurityDhcp, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *InlineResponse200529SsidSecurity) SetDhcp(v InlineResponse200529SsidSecurityDhcp)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *InlineResponse200529SsidSecurity) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


