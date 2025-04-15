# InlineObject273

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only) | 
**Category** | Pointer to **string** | Category of a policy object group (one of: NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup) | [optional] 
**ObjectIds** | Pointer to **[]int32** | A list of Policy Object ID&#39;s that this NetworkObjectGroup should be associated to (note: these ID&#39;s will replace the existing associated Policy Objects) | [optional] 

## Methods

### NewInlineObject273

`func NewInlineObject273(name string, ) *InlineObject273`

NewInlineObject273 instantiates a new InlineObject273 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject273WithDefaults

`func NewInlineObject273WithDefaults() *InlineObject273`

NewInlineObject273WithDefaults instantiates a new InlineObject273 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject273) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject273) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject273) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *InlineObject273) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineObject273) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineObject273) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineObject273) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetObjectIds

`func (o *InlineObject273) GetObjectIds() []int32`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *InlineObject273) GetObjectIdsOk() (*[]int32, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *InlineObject273) SetObjectIds(v []int32)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *InlineObject273) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


