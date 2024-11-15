# NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloorPlanId** | **string** | The ID of the floor plan to run auto locate for | 
**Refresh** | Pointer to **[]string** | The types of location data that should be refreshed for this job. The list must either contain both &#39;gnss&#39; and &#39;ranging&#39; or be empty, as we currently only support refreshing both &#39;gnss&#39; and &#39;ranging&#39;, or neither. | [optional] 
**ScheduledAt** | Pointer to **time.Time** | Timestamp in ISO8601 format which indicates when the auto locate job should be run. If omitted, the auto locate job will start immediately. | [optional] 

## Methods

### NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs

`func NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs(floorPlanId string, ) *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs`

NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobsWithDefaults

`func NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobsWithDefaults() *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs`

NewNetworksNetworkIdFloorPlansAutoLocateJobsBatchJobsWithDefaults instantiates a new NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloorPlanId

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.


### GetRefresh

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetRefresh() []string`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetRefreshOk() (*[]string, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) SetRefresh(v []string)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetScheduledAt

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *NetworksNetworkIdFloorPlansAutoLocateJobsBatchJobs) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


