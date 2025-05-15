# InlineResponse200102Jobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Auto locate job ID | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**FloorPlanId** | Pointer to **string** | Floor plan ID | [optional] 
**Status** | Pointer to **string** | Auto locate job status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;canceling&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;published&#39;, &#39;canceled&#39; | [optional] 
**ScheduledAt** | Pointer to **time.Time** | Scheduled start time for auto locate job | [optional] 
**Completed** | Pointer to [**InlineResponse200102Completed**](InlineResponse200102Completed.md) |  | [optional] 
**Ranging** | Pointer to [**InlineResponse200102Ranging**](InlineResponse200102Ranging.md) |  | [optional] 
**Gnss** | Pointer to [**InlineResponse200102Gnss**](InlineResponse200102Gnss.md) |  | [optional] 
**Errors** | Pointer to [**[]InlineResponse200102Errors**](InlineResponse200102Errors.md) | List of errors that occurred during a failed run of auto locate | [optional] 

## Methods

### NewInlineResponse200102Jobs

`func NewInlineResponse200102Jobs() *InlineResponse200102Jobs`

NewInlineResponse200102Jobs instantiates a new InlineResponse200102Jobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200102JobsWithDefaults

`func NewInlineResponse200102JobsWithDefaults() *InlineResponse200102Jobs`

NewInlineResponse200102JobsWithDefaults instantiates a new InlineResponse200102Jobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200102Jobs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200102Jobs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200102Jobs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200102Jobs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200102Jobs) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200102Jobs) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200102Jobs) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200102Jobs) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetFloorPlanId

`func (o *InlineResponse200102Jobs) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *InlineResponse200102Jobs) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *InlineResponse200102Jobs) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *InlineResponse200102Jobs) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200102Jobs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200102Jobs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200102Jobs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200102Jobs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScheduledAt

`func (o *InlineResponse200102Jobs) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *InlineResponse200102Jobs) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *InlineResponse200102Jobs) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *InlineResponse200102Jobs) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse200102Jobs) GetCompleted() InlineResponse200102Completed`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse200102Jobs) GetCompletedOk() (*InlineResponse200102Completed, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse200102Jobs) SetCompleted(v InlineResponse200102Completed)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse200102Jobs) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetRanging

`func (o *InlineResponse200102Jobs) GetRanging() InlineResponse200102Ranging`

GetRanging returns the Ranging field if non-nil, zero value otherwise.

### GetRangingOk

`func (o *InlineResponse200102Jobs) GetRangingOk() (*InlineResponse200102Ranging, bool)`

GetRangingOk returns a tuple with the Ranging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanging

`func (o *InlineResponse200102Jobs) SetRanging(v InlineResponse200102Ranging)`

SetRanging sets Ranging field to given value.

### HasRanging

`func (o *InlineResponse200102Jobs) HasRanging() bool`

HasRanging returns a boolean if a field has been set.

### GetGnss

`func (o *InlineResponse200102Jobs) GetGnss() InlineResponse200102Gnss`

GetGnss returns the Gnss field if non-nil, zero value otherwise.

### GetGnssOk

`func (o *InlineResponse200102Jobs) GetGnssOk() (*InlineResponse200102Gnss, bool)`

GetGnssOk returns a tuple with the Gnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnss

`func (o *InlineResponse200102Jobs) SetGnss(v InlineResponse200102Gnss)`

SetGnss sets Gnss field to given value.

### HasGnss

`func (o *InlineResponse200102Jobs) HasGnss() bool`

HasGnss returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200102Jobs) GetErrors() []InlineResponse200102Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200102Jobs) GetErrorsOk() (*[]InlineResponse200102Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200102Jobs) SetErrors(v []InlineResponse200102Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200102Jobs) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


