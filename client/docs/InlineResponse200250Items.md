# InlineResponse200250Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The device serial | [optional] 
**Target** | Pointer to **string** | The migration target destination | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time at which a migration was created | [optional] 
**MigratedAt** | Pointer to **time.Time** | The time at which the device initiated migration | [optional] 

## Methods

### NewInlineResponse200250Items

`func NewInlineResponse200250Items() *InlineResponse200250Items`

NewInlineResponse200250Items instantiates a new InlineResponse200250Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200250ItemsWithDefaults

`func NewInlineResponse200250ItemsWithDefaults() *InlineResponse200250Items`

NewInlineResponse200250ItemsWithDefaults instantiates a new InlineResponse200250Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200250Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200250Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200250Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200250Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTarget

`func (o *InlineResponse200250Items) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *InlineResponse200250Items) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *InlineResponse200250Items) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *InlineResponse200250Items) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse200250Items) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse200250Items) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse200250Items) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse200250Items) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMigratedAt

`func (o *InlineResponse200250Items) GetMigratedAt() time.Time`

GetMigratedAt returns the MigratedAt field if non-nil, zero value otherwise.

### GetMigratedAtOk

`func (o *InlineResponse200250Items) GetMigratedAtOk() (*time.Time, bool)`

GetMigratedAtOk returns a tuple with the MigratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedAt

`func (o *InlineResponse200250Items) SetMigratedAt(v time.Time)`

SetMigratedAt sets MigratedAt field to given value.

### HasMigratedAt

`func (o *InlineResponse200250Items) HasMigratedAt() bool`

HasMigratedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


