# InlineResponse200302

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Policy object ID | [optional] 
**Name** | Pointer to **string** | Name of policy object (alphanumeric, space, dash, or underscore characters only). | [optional] 
**Category** | Pointer to **string** | Category of a policy object (one of: adaptivePolicy, network) | [optional] 
**Type** | Pointer to **string** | Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask) | [optional] 
**Cidr** | Pointer to **string** | CIDR Value of a policy object | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time Stamp of policy object creation. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time Stamp of policy object updation. | [optional] 
**GroupIds** | Pointer to **[]string** | The IDs of policy object groups the policy object belongs to. | [optional] 
**NetworkIds** | Pointer to **[]string** | The IDs of the networks that use the policy object. | [optional] 

## Methods

### NewInlineResponse200302

`func NewInlineResponse200302() *InlineResponse200302`

NewInlineResponse200302 instantiates a new InlineResponse200302 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200302WithDefaults

`func NewInlineResponse200302WithDefaults() *InlineResponse200302`

NewInlineResponse200302WithDefaults instantiates a new InlineResponse200302 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200302) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200302) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200302) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200302) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200302) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200302) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200302) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200302) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse200302) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse200302) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse200302) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse200302) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200302) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200302) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200302) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200302) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse200302) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse200302) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse200302) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse200302) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200302) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200302) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200302) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200302) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200302) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200302) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200302) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200302) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetGroupIds

`func (o *InlineResponse200302) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *InlineResponse200302) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *InlineResponse200302) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *InlineResponse200302) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetNetworkIds

`func (o *InlineResponse200302) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *InlineResponse200302) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *InlineResponse200302) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *InlineResponse200302) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


