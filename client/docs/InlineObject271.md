# InlineObject271

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only) | 
**Category** | Pointer to **string** | Category of a policy object group (one of: NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup) | [optional] 
**ObjectIds** | Pointer to **[]int32** | A list of Policy Object ID&#39;s that this NetworkObjectGroup should be associated to (note: these ID&#39;s will replace the existing associated Policy Objects) | [optional] 

## Methods

### NewInlineObject271

`func NewInlineObject271(name string, ) *InlineObject271`

NewInlineObject271 instantiates a new InlineObject271 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject271WithDefaults

`func NewInlineObject271WithDefaults() *InlineObject271`

NewInlineObject271WithDefaults instantiates a new InlineObject271 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject271) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject271) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject271) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *InlineObject271) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineObject271) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineObject271) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineObject271) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetObjectIds

`func (o *InlineObject271) GetObjectIds() []int32`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *InlineObject271) GetObjectIdsOk() (*[]int32, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *InlineObject271) SetObjectIds(v []int32)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *InlineObject271) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


