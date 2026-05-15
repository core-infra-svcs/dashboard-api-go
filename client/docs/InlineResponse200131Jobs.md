# InlineResponse200131Jobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Auto locate job ID | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**FloorPlanId** | Pointer to **string** | Floor plan ID | [optional] 
**Status** | Pointer to **string** | Auto locate job status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;canceling&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;published&#39;, &#39;canceled&#39; | [optional] 
**ScheduledAt** | Pointer to **time.Time** | Scheduled start time for auto locate job | [optional] 
**Completed** | Pointer to [**InlineResponse200131Completed**](InlineResponse200131Completed.md) |  | [optional] 
**Ranging** | Pointer to [**InlineResponse200131Ranging**](InlineResponse200131Ranging.md) |  | [optional] 
**Gnss** | Pointer to [**InlineResponse200131Gnss**](InlineResponse200131Gnss.md) |  | [optional] 
**Errors** | Pointer to [**[]InlineResponse200131Errors**](InlineResponse200131Errors.md) | List of errors that occurred during a failed run of auto locate | [optional] 

## Methods

### NewInlineResponse200131Jobs

`func NewInlineResponse200131Jobs() *InlineResponse200131Jobs`

NewInlineResponse200131Jobs instantiates a new InlineResponse200131Jobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200131JobsWithDefaults

`func NewInlineResponse200131JobsWithDefaults() *InlineResponse200131Jobs`

NewInlineResponse200131JobsWithDefaults instantiates a new InlineResponse200131Jobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200131Jobs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200131Jobs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200131Jobs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200131Jobs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *InlineResponse200131Jobs) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse200131Jobs) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse200131Jobs) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse200131Jobs) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetFloorPlanId

`func (o *InlineResponse200131Jobs) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *InlineResponse200131Jobs) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *InlineResponse200131Jobs) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *InlineResponse200131Jobs) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200131Jobs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200131Jobs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200131Jobs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200131Jobs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScheduledAt

`func (o *InlineResponse200131Jobs) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *InlineResponse200131Jobs) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *InlineResponse200131Jobs) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *InlineResponse200131Jobs) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetCompleted

`func (o *InlineResponse200131Jobs) GetCompleted() InlineResponse200131Completed`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *InlineResponse200131Jobs) GetCompletedOk() (*InlineResponse200131Completed, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *InlineResponse200131Jobs) SetCompleted(v InlineResponse200131Completed)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *InlineResponse200131Jobs) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetRanging

`func (o *InlineResponse200131Jobs) GetRanging() InlineResponse200131Ranging`

GetRanging returns the Ranging field if non-nil, zero value otherwise.

### GetRangingOk

`func (o *InlineResponse200131Jobs) GetRangingOk() (*InlineResponse200131Ranging, bool)`

GetRangingOk returns a tuple with the Ranging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanging

`func (o *InlineResponse200131Jobs) SetRanging(v InlineResponse200131Ranging)`

SetRanging sets Ranging field to given value.

### HasRanging

`func (o *InlineResponse200131Jobs) HasRanging() bool`

HasRanging returns a boolean if a field has been set.

### GetGnss

`func (o *InlineResponse200131Jobs) GetGnss() InlineResponse200131Gnss`

GetGnss returns the Gnss field if non-nil, zero value otherwise.

### GetGnssOk

`func (o *InlineResponse200131Jobs) GetGnssOk() (*InlineResponse200131Gnss, bool)`

GetGnssOk returns a tuple with the Gnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnss

`func (o *InlineResponse200131Jobs) SetGnss(v InlineResponse200131Gnss)`

SetGnss sets Gnss field to given value.

### HasGnss

`func (o *InlineResponse200131Jobs) HasGnss() bool`

HasGnss returns a boolean if a field has been set.

### GetErrors

`func (o *InlineResponse200131Jobs) GetErrors() []InlineResponse200131Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse200131Jobs) GetErrorsOk() (*[]InlineResponse200131Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse200131Jobs) SetErrors(v []InlineResponse200131Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse200131Jobs) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


