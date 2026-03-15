# InlineResponse200202

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assoc** | Pointer to **int32** | The number of failed association attempts | [optional] 
**Auth** | Pointer to **int32** | The number of failed authentication attempts | [optional] 
**Dhcp** | Pointer to **int32** | The number of failed DHCP attempts | [optional] 
**Dns** | Pointer to **int32** | The number of failed DNS attempts | [optional] 
**Success** | Pointer to **int32** | The number of successful connection attempts | [optional] 

## Methods

### NewInlineResponse200202

`func NewInlineResponse200202() *InlineResponse200202`

NewInlineResponse200202 instantiates a new InlineResponse200202 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200202WithDefaults

`func NewInlineResponse200202WithDefaults() *InlineResponse200202`

NewInlineResponse200202WithDefaults instantiates a new InlineResponse200202 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssoc

`func (o *InlineResponse200202) GetAssoc() int32`

GetAssoc returns the Assoc field if non-nil, zero value otherwise.

### GetAssocOk

`func (o *InlineResponse200202) GetAssocOk() (*int32, bool)`

GetAssocOk returns a tuple with the Assoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssoc

`func (o *InlineResponse200202) SetAssoc(v int32)`

SetAssoc sets Assoc field to given value.

### HasAssoc

`func (o *InlineResponse200202) HasAssoc() bool`

HasAssoc returns a boolean if a field has been set.

### GetAuth

`func (o *InlineResponse200202) GetAuth() int32`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *InlineResponse200202) GetAuthOk() (*int32, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *InlineResponse200202) SetAuth(v int32)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *InlineResponse200202) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetDhcp

`func (o *InlineResponse200202) GetDhcp() int32`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *InlineResponse200202) GetDhcpOk() (*int32, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *InlineResponse200202) SetDhcp(v int32)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *InlineResponse200202) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetDns

`func (o *InlineResponse200202) GetDns() int32`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *InlineResponse200202) GetDnsOk() (*int32, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *InlineResponse200202) SetDns(v int32)`

SetDns sets Dns field to given value.

### HasDns

`func (o *InlineResponse200202) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetSuccess

`func (o *InlineResponse200202) GetSuccess() int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InlineResponse200202) GetSuccessOk() (*int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InlineResponse200202) SetSuccess(v int32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *InlineResponse200202) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


