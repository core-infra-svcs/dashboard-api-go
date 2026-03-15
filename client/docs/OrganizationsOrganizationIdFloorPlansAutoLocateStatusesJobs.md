# OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Auto locate job ID | [optional] 
**Status** | Pointer to **string** | Auto locate job status. Possible values: &#39;scheduled&#39;, &#39;in progress&#39;, &#39;canceling&#39;, &#39;error&#39;, &#39;finished&#39;, &#39;published&#39;, &#39;canceled&#39; | [optional] 
**ScheduledAt** | Pointer to **time.Time** | Scheduled start time for auto locate job | [optional] 
**Completed** | Pointer to [**InlineResponse200114Completed**](InlineResponse200114Completed.md) |  | [optional] 
**Ranging** | Pointer to [**InlineResponse200114Ranging**](InlineResponse200114Ranging.md) |  | [optional] 
**Gnss** | Pointer to [**InlineResponse200114Gnss**](InlineResponse200114Gnss.md) |  | [optional] 
**Errors** | Pointer to [**[]InlineResponse200114Errors**](InlineResponse200114Errors.md) | List of errors that occurred during a failed run of auto locate | [optional] 

## Methods

### NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs

`func NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs() *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs`

NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobsWithDefaults

`func NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobsWithDefaults() *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs`

NewOrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobsWithDefaults instantiates a new OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScheduledAt

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetCompleted

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetCompleted() InlineResponse200114Completed`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetCompletedOk() (*InlineResponse200114Completed, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetCompleted(v InlineResponse200114Completed)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetRanging

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetRanging() InlineResponse200114Ranging`

GetRanging returns the Ranging field if non-nil, zero value otherwise.

### GetRangingOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetRangingOk() (*InlineResponse200114Ranging, bool)`

GetRangingOk returns a tuple with the Ranging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanging

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetRanging(v InlineResponse200114Ranging)`

SetRanging sets Ranging field to given value.

### HasRanging

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasRanging() bool`

HasRanging returns a boolean if a field has been set.

### GetGnss

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetGnss() InlineResponse200114Gnss`

GetGnss returns the Gnss field if non-nil, zero value otherwise.

### GetGnssOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetGnssOk() (*InlineResponse200114Gnss, bool)`

GetGnssOk returns a tuple with the Gnss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnss

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetGnss(v InlineResponse200114Gnss)`

SetGnss sets Gnss field to given value.

### HasGnss

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasGnss() bool`

HasGnss returns a boolean if a field has been set.

### GetErrors

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetErrors() []InlineResponse200114Errors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) GetErrorsOk() (*[]InlineResponse200114Errors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) SetErrors(v []InlineResponse200114Errors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *OrganizationsOrganizationIdFloorPlansAutoLocateStatusesJobs) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


