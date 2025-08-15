# InlineResponse20112

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**[]InlineResponse20112Clients**](InlineResponse20112Clients.md) | The list of clients to provision | [optional] 
**DevicePolicy** | Pointer to **string** | The name of the client&#39;s policy | [optional] 
**GroupPolicyId** | Pointer to **string** | The group policy identifier of the client | [optional] 

## Methods

### NewInlineResponse20112

`func NewInlineResponse20112() *InlineResponse20112`

NewInlineResponse20112 instantiates a new InlineResponse20112 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20112WithDefaults

`func NewInlineResponse20112WithDefaults() *InlineResponse20112`

NewInlineResponse20112WithDefaults instantiates a new InlineResponse20112 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *InlineResponse20112) GetClients() []InlineResponse20112Clients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *InlineResponse20112) GetClientsOk() (*[]InlineResponse20112Clients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *InlineResponse20112) SetClients(v []InlineResponse20112Clients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *InlineResponse20112) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetDevicePolicy

`func (o *InlineResponse20112) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *InlineResponse20112) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *InlineResponse20112) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.

### HasDevicePolicy

`func (o *InlineResponse20112) HasDevicePolicy() bool`

HasDevicePolicy returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *InlineResponse20112) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *InlineResponse20112) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *InlineResponse20112) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *InlineResponse20112) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


