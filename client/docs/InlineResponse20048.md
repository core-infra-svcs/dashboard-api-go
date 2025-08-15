# InlineResponse20048

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnrollmentId** | Pointer to **string** | Id to check the status of your enrollment | [optional] 
**Url** | Pointer to **string** | Url to check the status of your enrollment | [optional] 
**Request** | Pointer to [**InlineResponse2019Request**](InlineResponse2019Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the enrollment | [optional] 
**EnrollmentStartedAt** | Pointer to **string** | Enrollment started at | [optional] 
**DoorLocks** | Pointer to [**[]InlineResponse20048DoorLocks**](InlineResponse20048DoorLocks.md) | Door locks | [optional] 

## Methods

### NewInlineResponse20048

`func NewInlineResponse20048() *InlineResponse20048`

NewInlineResponse20048 instantiates a new InlineResponse20048 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20048WithDefaults

`func NewInlineResponse20048WithDefaults() *InlineResponse20048`

NewInlineResponse20048WithDefaults instantiates a new InlineResponse20048 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnrollmentId

`func (o *InlineResponse20048) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InlineResponse20048) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InlineResponse20048) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InlineResponse20048) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20048) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20048) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20048) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20048) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20048) GetRequest() InlineResponse2019Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20048) GetRequestOk() (*InlineResponse2019Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20048) SetRequest(v InlineResponse2019Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20048) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20048) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20048) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20048) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20048) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnrollmentStartedAt

`func (o *InlineResponse20048) GetEnrollmentStartedAt() string`

GetEnrollmentStartedAt returns the EnrollmentStartedAt field if non-nil, zero value otherwise.

### GetEnrollmentStartedAtOk

`func (o *InlineResponse20048) GetEnrollmentStartedAtOk() (*string, bool)`

GetEnrollmentStartedAtOk returns a tuple with the EnrollmentStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentStartedAt

`func (o *InlineResponse20048) SetEnrollmentStartedAt(v string)`

SetEnrollmentStartedAt sets EnrollmentStartedAt field to given value.

### HasEnrollmentStartedAt

`func (o *InlineResponse20048) HasEnrollmentStartedAt() bool`

HasEnrollmentStartedAt returns a boolean if a field has been set.

### GetDoorLocks

`func (o *InlineResponse20048) GetDoorLocks() []InlineResponse20048DoorLocks`

GetDoorLocks returns the DoorLocks field if non-nil, zero value otherwise.

### GetDoorLocksOk

`func (o *InlineResponse20048) GetDoorLocksOk() (*[]InlineResponse20048DoorLocks, bool)`

GetDoorLocksOk returns a tuple with the DoorLocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorLocks

`func (o *InlineResponse20048) SetDoorLocks(v []InlineResponse20048DoorLocks)`

SetDoorLocks sets DoorLocks field to given value.

### HasDoorLocks

`func (o *InlineResponse20048) HasDoorLocks() bool`

HasDoorLocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


