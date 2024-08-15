# InlineResponse20029

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandId** | Pointer to **string** | ID to check the status of the command request | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the command was triggered | [optional] 
**CompletedAt** | Pointer to **time.Time** | Time when the command was completed | [optional] 
**CreatedBy** | Pointer to [**DevicesSerialSensorCommandsCreatedBy**](DevicesSerialSensorCommandsCreatedBy.md) |  | [optional] 
**Operation** | Pointer to **string** | Operation run on the sensor | [optional] 
**Status** | Pointer to **string** | Status of the command request | [optional] 
**Errors** | Pointer to **[]string** | Array of errors if failed | [optional] 

## Methods

### NewInlineResponse20029

`func NewInlineResponse20029() *InlineResponse20029`

NewInlineResponse20029 instantiates a new InlineResponse20029 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029WithDefaults

`func NewInlineResponse20029WithDefaults() *InlineResponse20029`

NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandId

`func (o *InlineResponse20029) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *InlineResponse20029) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *InlineResponse20029) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *InlineResponse20029) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InlineResponse20029) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse20029) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse20029) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InlineResponse20029) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *InlineResponse20029) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *InlineResponse20029) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *InlineResponse20029) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *InlineResponse20029) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InlineResponse20029) GetCreatedBy() DevicesSerialSensorCommandsCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse20029) GetCreatedByOk() (*DevicesSerialSensorCommandsCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse20029) SetCreatedBy(v DevicesSerialSensorCommandsCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse20029) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOperation

`func (o *InlineResponse20029) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *InlineResponse20029) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *InlineResponse20029) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *InlineResponse20029) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20029) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20029) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20029) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20029) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse20029) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse20029) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse20029) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse20029) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


