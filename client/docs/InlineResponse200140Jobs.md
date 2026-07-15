# InlineResponse200140Jobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Auto locate job ID | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**FloorPlanId** | Pointer to **string** | Floor plan ID | [optional] 
**Status** | Pointer to **string** | Auto locate job status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;canceling&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;published&#39;, &#39;canceled&#39; | [optional] 
**ScheduledAt** | Pointer to **time.Time** | Scheduled start time for auto locate job | [optional] 
**Completed** | Pointer to [**InlineResponse200140Completed**](InlineResponse200140Completed.md) |  | [optional] 
**Ranging** | Pointer to [**InlineResponse200140Ranging**](InlineResponse200140Ranging.md) |  | [optional] 
**Gnss** | Pointer to [**InlineResponse200140Gnss**](InlineResponse200140Gnss.md) |  | [optional] 
**Errors** | Pointer to [**[]InlineResponse200140Errors**](InlineResponse200140Errors.md) | List of errors that occurred during a failed run of auto locate | [optional] 

## Methods

### NewInlineResponse200140Jobs

`func NewInlineResponse200140Jobs() *InlineResponse200140Jobs`

NewInlineResponse200140Jobs instantiates a new InlineResponse200140Jobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200140JobsWithDefaults

`func NewInlineResponse200140JobsWithDefaults() *InlineResponse200140Jobs`

NewInlineResponse200140JobsWithDefaults instantiates a new InlineResponse200140Jobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200140Jobs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200140Jobs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200140Jobs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200140Jobs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200140Jobs) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200140Jobs) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200140Jobs) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200140Jobs) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetFloorPlanId

`func (o *InlineResponse200140Jobs) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *InlineResponse200140Jobs) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *InlineResponse200140Jobs) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *InlineResponse200140Jobs) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200140Jobs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200140Jobs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200140Jobs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200140Jobs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScheduledAt

`func (o *InlineResponse200140Jobs) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *InlineResponse200140Jobs) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *InlineResponse200140Jobs) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *InlineResponse200140Jobs) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse200140Jobs) GetCompleted() InlineResponse200140Completed`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse200140Jobs) GetCompletedOk() (*InlineResponse200140Completed, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse200140Jobs) SetCompleted(v InlineResponse200140Completed)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse200140Jobs) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetRanging

`func (o *InlineResponse200140Jobs) GetRanging() InlineResponse200140Ranging`

GetRanging returns the Ranging field if non-nil, zero value otherwise.

### GetRangingOk

`func (o *InlineResponse200140Jobs) GetRangingOk() (*InlineResponse200140Ranging, bool)`

GetRangingOk returns a tuple with the Ranging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanging

`func (o *InlineResponse200140Jobs) SetRanging(v InlineResponse200140Ranging)`

SetRanging sets Ranging field to given value.

### HasRanging

`func (o *InlineResponse200140Jobs) HasRanging() bool`

HasRanging returns a boolean if a field has been set.

### GetGnss

`func (o *InlineResponse200140Jobs) GetGnss() InlineResponse200140Gnss`

GetGnss returns the Gnss field if non-nil, zero value otherwise.

### GetGnssOk

`func (o *InlineResponse200140Jobs) GetGnssOk() (*InlineResponse200140Gnss, bool)`

GetGnssOk returns a tuple with the Gnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnss

`func (o *InlineResponse200140Jobs) SetGnss(v InlineResponse200140Gnss)`

SetGnss sets Gnss field to given value.

### HasGnss

`func (o *InlineResponse200140Jobs) HasGnss() bool`

HasGnss returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200140Jobs) GetErrors() []InlineResponse200140Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200140Jobs) GetErrorsOk() (*[]InlineResponse200140Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200140Jobs) SetErrors(v []InlineResponse200140Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200140Jobs) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


