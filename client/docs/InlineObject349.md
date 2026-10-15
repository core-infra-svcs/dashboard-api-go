# InlineObject349

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only) | 
**Category** | **string** | Category of a policy object (one of: adaptivePolicy, network) | 
**Type** | **string** | Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn). DEPRECATED: &#x60;ipAndMask&#x60; is deprecated and will be removed in a future release. Use &#x60;cidr&#x60; instead. | 
**Cidr** | Pointer to **string** | CIDR Value of a policy object (e.g. 10.11.12.1/24\&quot;) | [optional] 
**Fqdn** | Pointer to **string** | Fully qualified domain name of policy object (e.g. \&quot;example.com\&quot;) | [optional] 
**Mask** | Pointer to **string** | Mask of a policy object (e.g. \&quot;255.255.0.0\&quot;). Used only with deprecated &#x60;type&#x3D;ipAndMask&#x60;. | [optional] 
**Ip** | Pointer to **string** | IP Address of a policy object (e.g. \&quot;1.2.3.4\&quot;). Used only with deprecated &#x60;type&#x3D;ipAndMask&#x60;. | [optional] 
**GroupIds** | Pointer to **[]string** | The IDs of policy object groups the policy object belongs to | [optional] 

## Methods

### NewInlineObject349

`func NewInlineObject349(name string, category string, type_ string, ) *InlineObject349`

NewInlineObject349 instantiates a new InlineObject349 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject349WithDefaults

`func NewInlineObject349WithDefaults() *InlineObject349`

NewInlineObject349WithDefaults instantiates a new InlineObject349 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject349) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject349) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject349) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *InlineObject349) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineObject349) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineObject349) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetType

`func (o *InlineObject349) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject349) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject349) SetType(v string)`

SetType sets Type field to given value.


### GetCidr

`func (o *InlineObject349) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineObject349) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineObject349) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineObject349) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetFqdn

`func (o *InlineObject349) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *InlineObject349) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *InlineObject349) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *InlineObject349) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetMask

`func (o *InlineObject349) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *InlineObject349) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *InlineObject349) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *InlineObject349) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetIp

`func (o *InlineObject349) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InlineObject349) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InlineObject349) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InlineObject349) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetGroupIds

`func (o *InlineObject349) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *InlineObject349) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *InlineObject349) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *InlineObject349) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


