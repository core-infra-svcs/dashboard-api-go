# InlineResponse200311

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Policy object ID | [optional] 
**Name** | Pointer to **string** | Name of the Policy object group. | [optional] 
**Category** | Pointer to **string** | Type of object groups. (NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time Stamp of policy object creation. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time Stamp of policy object updation. | [optional] 
**ObjectIds** | Pointer to **[]int32** | Policy objects associated with Network Object Group or Port Object Group | [optional] 
**NetworkIds** | Pointer to **[]string** | Network ID&#39;s associated with the policy objects. | [optional] 

## Methods

### NewInlineResponse200311

`func NewInlineResponse200311() *InlineResponse200311`

NewInlineResponse200311 instantiates a new InlineResponse200311 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200311WithDefaults

`func NewInlineResponse200311WithDefaults() *InlineResponse200311`

NewInlineResponse200311WithDefaults instantiates a new InlineResponse200311 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200311) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200311) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200311) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200311) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200311) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200311) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200311) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200311) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse200311) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse200311) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse200311) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse200311) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200311) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200311) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200311) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200311) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InlineResponse200311) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InlineResponse200311) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InlineResponse200311) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InlineResponse200311) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetObjectIds

`func (o *InlineResponse200311) GetObjectIds() []int32`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *InlineResponse200311) GetObjectIdsOk() (*[]int32, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *InlineResponse200311) SetObjectIds(v []int32)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *InlineResponse200311) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### GetNetworkIds

`func (o *InlineResponse200311) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *InlineResponse200311) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *InlineResponse200311) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *InlineResponse200311) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


