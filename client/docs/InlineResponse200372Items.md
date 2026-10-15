# InlineResponse200372Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The device serial | [optional] 
**Target** | Pointer to **string** | The migration target destination | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which a migration was created | [optional] 
**MigratedAt** | Pointer to **NullableTime** | The time at which the device initiated migration | [optional] 

## Methods

### NewInlineResponse200372Items

`func NewInlineResponse200372Items() *InlineResponse200372Items`

NewInlineResponse200372Items instantiates a new InlineResponse200372Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200372ItemsWithDefaults

`func NewInlineResponse200372ItemsWithDefaults() *InlineResponse200372Items`

NewInlineResponse200372ItemsWithDefaults instantiates a new InlineResponse200372Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200372Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200372Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200372Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200372Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTarget

`func (o *InlineResponse200372Items) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InlineResponse200372Items) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InlineResponse200372Items) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *InlineResponse200372Items) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200372Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200372Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200372Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200372Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMigratedAt

`func (o *InlineResponse200372Items) GetMigratedAt() time.Time`

GetMigratedAt returns the MigratedAt field if non-nil, zero value otherwise.

### GetMigratedAtOk

`func (o *InlineResponse200372Items) GetMigratedAtOk() (*time.Time, bool)`

GetMigratedAtOk returns a tuple with the MigratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedAt

`func (o *InlineResponse200372Items) SetMigratedAt(v time.Time)`

SetMigratedAt sets MigratedAt field to given value.

### HasMigratedAt

`func (o *InlineResponse200372Items) HasMigratedAt() bool`

HasMigratedAt returns a boolean if a field has been set.

### SetMigratedAtNil

`func (o *InlineResponse200372Items) SetMigratedAtNil(b bool)`

 SetMigratedAtNil sets the value for MigratedAt to be an explicit nil

### UnsetMigratedAt
`func (o *InlineResponse200372Items) UnsetMigratedAt()`

UnsetMigratedAt ensures that no value is present for MigratedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


