# InlineResponse20068

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnrollmentId** | Pointer to **string** | Id to check the status of your enrollment | [optional] 
**Url** | Pointer to **string** | Url to check the status of your enrollment | [optional] 
**Request** | Pointer to [**InlineResponse20116Request**](InlineResponse20116Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the enrollment | [optional] 
**EnrollmentStartedAt** | Pointer to **string** | Enrollment started at | [optional] 
**DoorLocks** | Pointer to [**[]InlineResponse20068DoorLocks**](InlineResponse20068DoorLocks.md) | Door locks | [optional] 

## Methods

### NewInlineResponse20068

`func NewInlineResponse20068() *InlineResponse20068`

NewInlineResponse20068 instantiates a new InlineResponse20068 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20068WithDefaults

`func NewInlineResponse20068WithDefaults() *InlineResponse20068`

NewInlineResponse20068WithDefaults instantiates a new InlineResponse20068 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnrollmentId

`func (o *InlineResponse20068) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InlineResponse20068) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InlineResponse20068) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InlineResponse20068) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20068) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20068) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20068) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20068) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20068) GetRequest() InlineResponse20116Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20068) GetRequestOk() (*InlineResponse20116Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20068) SetRequest(v InlineResponse20116Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20068) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20068) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20068) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20068) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20068) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnrollmentStartedAt

`func (o *InlineResponse20068) GetEnrollmentStartedAt() string`

GetEnrollmentStartedAt returns the EnrollmentStartedAt field if non-nil, zero value otherwise.

### GetEnrollmentStartedAtOk

`func (o *InlineResponse20068) GetEnrollmentStartedAtOk() (*string, bool)`

GetEnrollmentStartedAtOk returns a tuple with the EnrollmentStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentStartedAt

`func (o *InlineResponse20068) SetEnrollmentStartedAt(v string)`

SetEnrollmentStartedAt sets EnrollmentStartedAt field to given value.

### HasEnrollmentStartedAt

`func (o *InlineResponse20068) HasEnrollmentStartedAt() bool`

HasEnrollmentStartedAt returns a boolean if a field has been set.

### GetDoorLocks

`func (o *InlineResponse20068) GetDoorLocks() []InlineResponse20068DoorLocks`

GetDoorLocks returns the DoorLocks field if non-nil, zero value otherwise.

### GetDoorLocksOk

`func (o *InlineResponse20068) GetDoorLocksOk() (*[]InlineResponse20068DoorLocks, bool)`

GetDoorLocksOk returns a tuple with the DoorLocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorLocks

`func (o *InlineResponse20068) SetDoorLocks(v []InlineResponse20068DoorLocks)`

SetDoorLocks sets DoorLocks field to given value.

### HasDoorLocks

`func (o *InlineResponse20068) HasDoorLocks() bool`

HasDoorLocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


