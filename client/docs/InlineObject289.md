# InlineObject289

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only) | [optional] 
**Cidr** | Pointer to **string** | CIDR Value of a policy object (e.g. 10.11.12.1/24\&quot;) | [optional] 
**Fqdn** | Pointer to **string** | Fully qualified domain name of policy object (e.g. \&quot;example.com\&quot;) | [optional] 
**Mask** | Pointer to **string** | Mask of a policy object (e.g. \&quot;255.255.0.0\&quot;) | [optional] 
**Ip** | Pointer to **string** | IP Address of a policy object (e.g. \&quot;1.2.3.4\&quot;) | [optional] 
**GroupIds** | Pointer to **[]string** | The IDs of policy object groups the policy object belongs to | [optional] 

## Methods

### NewInlineObject289

`func NewInlineObject289() *InlineObject289`

NewInlineObject289 instantiates a new InlineObject289 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject289WithDefaults

`func NewInlineObject289WithDefaults() *InlineObject289`

NewInlineObject289WithDefaults instantiates a new InlineObject289 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject289) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject289) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject289) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject289) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCidr

`func (o *InlineObject289) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineObject289) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineObject289) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineObject289) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetFqdn

`func (o *InlineObject289) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *InlineObject289) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *InlineObject289) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *InlineObject289) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetMask

`func (o *InlineObject289) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *InlineObject289) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *InlineObject289) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *InlineObject289) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetIp

`func (o *InlineObject289) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineObject289) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineObject289) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineObject289) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetGroupIds

`func (o *InlineObject289) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *InlineObject289) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *InlineObject289) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *InlineObject289) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


