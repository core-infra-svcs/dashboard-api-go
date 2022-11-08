# InlineResponse20046

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | device ID | [optional] 
**SsidName** | Pointer to **string** | SSID name | [optional] 
**Name** | Pointer to **string** | device name | [optional] 
**Scope** | Pointer to **string** | scope | [optional] 
**Tags** | Pointer to **[]string** | device tags | [optional] 
**TimeboundType** | Pointer to **string** | type of access period, either a static range or a dynamic period | [optional] 
**AccessStartAt** | Pointer to **time.Time** | time that access starts | [optional] 
**AccessEndAt** | Pointer to **time.Time** | time that access ends | [optional] 

## Methods

### NewInlineResponse20046

`func NewInlineResponse20046() *InlineResponse20046`

NewInlineResponse20046 instantiates a new InlineResponse20046 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20046WithDefaults

`func NewInlineResponse20046WithDefaults() *InlineResponse20046`

NewInlineResponse20046WithDefaults instantiates a new InlineResponse20046 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20046) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20046) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20046) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20046) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSsidName

`func (o *InlineResponse20046) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *InlineResponse20046) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *InlineResponse20046) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *InlineResponse20046) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20046) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20046) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20046) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20046) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *InlineResponse20046) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineResponse20046) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineResponse20046) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineResponse20046) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20046) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20046) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20046) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20046) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeboundType

`func (o *InlineResponse20046) GetTimeboundType() string`

GetTimeboundType returns the TimeboundType field if non-nil, zero value otherwise.

### GetTimeboundTypeOk

`func (o *InlineResponse20046) GetTimeboundTypeOk() (*string, bool)`

GetTimeboundTypeOk returns a tuple with the TimeboundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeboundType

`func (o *InlineResponse20046) SetTimeboundType(v string)`

SetTimeboundType sets TimeboundType field to given value.

### HasTimeboundType

`func (o *InlineResponse20046) HasTimeboundType() bool`

HasTimeboundType returns a boolean if a field has been set.

### GetAccessStartAt

`func (o *InlineResponse20046) GetAccessStartAt() time.Time`

GetAccessStartAt returns the AccessStartAt field if non-nil, zero value otherwise.

### GetAccessStartAtOk

`func (o *InlineResponse20046) GetAccessStartAtOk() (*time.Time, bool)`

GetAccessStartAtOk returns a tuple with the AccessStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStartAt

`func (o *InlineResponse20046) SetAccessStartAt(v time.Time)`

SetAccessStartAt sets AccessStartAt field to given value.

### HasAccessStartAt

`func (o *InlineResponse20046) HasAccessStartAt() bool`

HasAccessStartAt returns a boolean if a field has been set.

### GetAccessEndAt

`func (o *InlineResponse20046) GetAccessEndAt() time.Time`

GetAccessEndAt returns the AccessEndAt field if non-nil, zero value otherwise.

### GetAccessEndAtOk

`func (o *InlineResponse20046) GetAccessEndAtOk() (*time.Time, bool)`

GetAccessEndAtOk returns a tuple with the AccessEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessEndAt

`func (o *InlineResponse20046) SetAccessEndAt(v time.Time)`

SetAccessEndAt sets AccessEndAt field to given value.

### HasAccessEndAt

`func (o *InlineResponse20046) HasAccessEndAt() bool`

HasAccessEndAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

