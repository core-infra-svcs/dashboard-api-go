# InlineResponse200108Jobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Auto locate job ID | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**FloorPlanId** | Pointer to **string** | Floor plan ID | [optional] 
**Status** | Pointer to **string** | Auto locate job status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;canceling&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;published&#39;, &#39;canceled&#39; | [optional] 
**ScheduledAt** | Pointer to **time.Time** | Scheduled start time for auto locate job | [optional] 
**Completed** | Pointer to [**InlineResponse200108Completed**](InlineResponse200108Completed.md) |  | [optional] 
**Ranging** | Pointer to [**InlineResponse200108Ranging**](InlineResponse200108Ranging.md) |  | [optional] 
**Gnss** | Pointer to [**InlineResponse200108Gnss**](InlineResponse200108Gnss.md) |  | [optional] 
**Errors** | Pointer to [**[]InlineResponse200108Errors**](InlineResponse200108Errors.md) | List of errors that occurred during a failed run of auto locate | [optional] 

## Methods

### NewInlineResponse200108Jobs

`func NewInlineResponse200108Jobs() *InlineResponse200108Jobs`

NewInlineResponse200108Jobs instantiates a new InlineResponse200108Jobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200108JobsWithDefaults

`func NewInlineResponse200108JobsWithDefaults() *InlineResponse200108Jobs`

NewInlineResponse200108JobsWithDefaults instantiates a new InlineResponse200108Jobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200108Jobs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200108Jobs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200108Jobs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200108Jobs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200108Jobs) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200108Jobs) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200108Jobs) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200108Jobs) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetFloorPlanId

`func (o *InlineResponse200108Jobs) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *InlineResponse200108Jobs) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *InlineResponse200108Jobs) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *InlineResponse200108Jobs) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200108Jobs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200108Jobs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200108Jobs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200108Jobs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScheduledAt

`func (o *InlineResponse200108Jobs) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *InlineResponse200108Jobs) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *InlineResponse200108Jobs) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *InlineResponse200108Jobs) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse200108Jobs) GetCompleted() InlineResponse200108Completed`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse200108Jobs) GetCompletedOk() (*InlineResponse200108Completed, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse200108Jobs) SetCompleted(v InlineResponse200108Completed)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse200108Jobs) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetRanging

`func (o *InlineResponse200108Jobs) GetRanging() InlineResponse200108Ranging`

GetRanging returns the Ranging field if non-nil, zero value otherwise.

### GetRangingOk

`func (o *InlineResponse200108Jobs) GetRangingOk() (*InlineResponse200108Ranging, bool)`

GetRangingOk returns a tuple with the Ranging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanging

`func (o *InlineResponse200108Jobs) SetRanging(v InlineResponse200108Ranging)`

SetRanging sets Ranging field to given value.

### HasRanging

`func (o *InlineResponse200108Jobs) HasRanging() bool`

HasRanging returns a boolean if a field has been set.

### GetGnss

`func (o *InlineResponse200108Jobs) GetGnss() InlineResponse200108Gnss`

GetGnss returns the Gnss field if non-nil, zero value otherwise.

### GetGnssOk

`func (o *InlineResponse200108Jobs) GetGnssOk() (*InlineResponse200108Gnss, bool)`

GetGnssOk returns a tuple with the Gnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnss

`func (o *InlineResponse200108Jobs) SetGnss(v InlineResponse200108Gnss)`

SetGnss sets Gnss field to given value.

### HasGnss

`func (o *InlineResponse200108Jobs) HasGnss() bool`

HasGnss returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200108Jobs) GetErrors() []InlineResponse200108Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200108Jobs) GetErrorsOk() (*[]InlineResponse200108Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200108Jobs) SetErrors(v []InlineResponse200108Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200108Jobs) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


