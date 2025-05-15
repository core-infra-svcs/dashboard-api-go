# InlineResponse20052

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Appliance service name | [optional] 
**Access** | Pointer to **string** | A string indicating the rule for which IPs are allowed to use the specified service | [optional] 
**AllowedIps** | Pointer to **[]string** | An array of allowed IPs that can access the service | [optional] 

## Methods

### NewInlineResponse20052

`func NewInlineResponse20052() *InlineResponse20052`

NewInlineResponse20052 instantiates a new InlineResponse20052 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20052WithDefaults

`func NewInlineResponse20052WithDefaults() *InlineResponse20052`

NewInlineResponse20052WithDefaults instantiates a new InlineResponse20052 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *InlineResponse20052) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *InlineResponse20052) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *InlineResponse20052) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *InlineResponse20052) HasService() bool`

HasService returns a boolean if a field has been set.

### GetAccess

`func (o *InlineResponse20052) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *InlineResponse20052) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *InlineResponse20052) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *InlineResponse20052) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAllowedIps

`func (o *InlineResponse20052) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *InlineResponse20052) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *InlineResponse20052) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *InlineResponse20052) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


