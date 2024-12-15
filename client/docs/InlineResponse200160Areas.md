# InlineResponse200160Areas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaId** | Pointer to **string** | OSPF area ID | [optional] 
**AreaName** | Pointer to **string** | Name of the OSPF area | [optional] 
**AreaType** | Pointer to **string** | Area types in OSPF. Must be one of: [\&quot;normal\&quot;, \&quot;stub\&quot;, \&quot;nssa\&quot;] | [optional] 

## Methods

### NewInlineResponse200160Areas

`func NewInlineResponse200160Areas() *InlineResponse200160Areas`

NewInlineResponse200160Areas instantiates a new InlineResponse200160Areas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200160AreasWithDefaults

`func NewInlineResponse200160AreasWithDefaults() *InlineResponse200160Areas`

NewInlineResponse200160AreasWithDefaults instantiates a new InlineResponse200160Areas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaId

`func (o *InlineResponse200160Areas) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *InlineResponse200160Areas) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *InlineResponse200160Areas) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.

### HasAreaId

`func (o *InlineResponse200160Areas) HasAreaId() bool`

HasAreaId returns a boolean if a field has been set.

### GetAreaName

`func (o *InlineResponse200160Areas) GetAreaName() string`

GetAreaName returns the AreaName field if non-nil, zero value otherwise.

### GetAreaNameOk

`func (o *InlineResponse200160Areas) GetAreaNameOk() (*string, bool)`

GetAreaNameOk returns a tuple with the AreaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaName

`func (o *InlineResponse200160Areas) SetAreaName(v string)`

SetAreaName sets AreaName field to given value.

### HasAreaName

`func (o *InlineResponse200160Areas) HasAreaName() bool`

HasAreaName returns a boolean if a field has been set.

### GetAreaType

`func (o *InlineResponse200160Areas) GetAreaType() string`

GetAreaType returns the AreaType field if non-nil, zero value otherwise.

### GetAreaTypeOk

`func (o *InlineResponse200160Areas) GetAreaTypeOk() (*string, bool)`

GetAreaTypeOk returns a tuple with the AreaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaType

`func (o *InlineResponse200160Areas) SetAreaType(v string)`

SetAreaType sets AreaType field to given value.

### HasAreaType

`func (o *InlineResponse200160Areas) HasAreaType() bool`

HasAreaType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


