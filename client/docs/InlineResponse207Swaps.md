# InlineResponse207Swaps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Swap Request ID | 
**Devices** | [**InlineResponse207Devices**](InlineResponse207Devices.md) |  | 
**Status** | **string** | The current status of the swap job. | 
**AfterAction** | **string** | An action to perform on the devices.old object after swap is complete. | 
**CreatedAt** | **string** | An iso8601 timestamp for the creation of the swap request. | 
**CompletedAt** | Pointer to **string** | An iso8601 timestamp for when the swap completed or failed. | [optional] 
**Errors** | Pointer to **[]string** | A list of error messages for why a swap failed. | [optional] 

## Methods

### NewInlineResponse207Swaps

`func NewInlineResponse207Swaps(id string, devices InlineResponse207Devices, status string, afterAction string, createdAt string, ) *InlineResponse207Swaps`

NewInlineResponse207Swaps instantiates a new InlineResponse207Swaps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse207SwapsWithDefaults

`func NewInlineResponse207SwapsWithDefaults() *InlineResponse207Swaps`

NewInlineResponse207SwapsWithDefaults instantiates a new InlineResponse207Swaps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse207Swaps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse207Swaps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse207Swaps) SetId(v string)`

SetId sets Id field to given value.


### GetDevices

`func (o *InlineResponse207Swaps) GetDevices() InlineResponse207Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse207Swaps) GetDevicesOk() (*InlineResponse207Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse207Swaps) SetDevices(v InlineResponse207Devices)`

SetDevices sets Devices field to given value.


### GetStatus

`func (o *InlineResponse207Swaps) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse207Swaps) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse207Swaps) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAfterAction

`func (o *InlineResponse207Swaps) GetAfterAction() string`

GetAfterAction returns the AfterAction field if non-nil, zero value otherwise.

### GetAfterActionOk

`func (o *InlineResponse207Swaps) GetAfterActionOk() (*string, bool)`

GetAfterActionOk returns a tuple with the AfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterAction

`func (o *InlineResponse207Swaps) SetAfterAction(v string)`

SetAfterAction sets AfterAction field to given value.


### GetCreatedAt

`func (o *InlineResponse207Swaps) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InlineResponse207Swaps) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InlineResponse207Swaps) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *InlineResponse207Swaps) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *InlineResponse207Swaps) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *InlineResponse207Swaps) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *InlineResponse207Swaps) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse207Swaps) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse207Swaps) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse207Swaps) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse207Swaps) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


