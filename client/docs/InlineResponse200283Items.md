# InlineResponse200283Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The device serial | [optional] 
**Target** | Pointer to **string** | The migration target destination | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which a migration was created | [optional] 
**MigratedAt** | Pointer to **NullableTime** | The time at which the device initiated migration | [optional] 

## Methods

### NewInlineResponse200283Items

`func NewInlineResponse200283Items() *InlineResponse200283Items`

NewInlineResponse200283Items instantiates a new InlineResponse200283Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200283ItemsWithDefaults

`func NewInlineResponse200283ItemsWithDefaults() *InlineResponse200283Items`

NewInlineResponse200283ItemsWithDefaults instantiates a new InlineResponse200283Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200283Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200283Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200283Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200283Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTarget

`func (o *InlineResponse200283Items) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InlineResponse200283Items) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InlineResponse200283Items) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *InlineResponse200283Items) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200283Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200283Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200283Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200283Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMigratedAt

`func (o *InlineResponse200283Items) GetMigratedAt() time.Time`

GetMigratedAt returns the MigratedAt field if non-nil, zero value otherwise.

### GetMigratedAtOk

`func (o *InlineResponse200283Items) GetMigratedAtOk() (*time.Time, bool)`

GetMigratedAtOk returns a tuple with the MigratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedAt

`func (o *InlineResponse200283Items) SetMigratedAt(v time.Time)`

SetMigratedAt sets MigratedAt field to given value.

### HasMigratedAt

`func (o *InlineResponse200283Items) HasMigratedAt() bool`

HasMigratedAt returns a boolean if a field has been set.

### SetMigratedAtNil

`func (o *InlineResponse200283Items) SetMigratedAtNil(b bool)`

 SetMigratedAtNil sets the value for MigratedAt to be an explicit nil

### UnsetMigratedAt
`func (o *InlineResponse200283Items) UnsetMigratedAt()`

UnsetMigratedAt ensures that no value is present for MigratedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


