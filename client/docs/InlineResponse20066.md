# InlineResponse20066

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnrollmentId** | Pointer to **string** | Id to check the status of your enrollment | [optional] 
**Url** | Pointer to **string** | Url to check the status of your enrollment | [optional] 
**Request** | Pointer to [**InlineResponse20115Request**](InlineResponse20115Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the enrollment | [optional] 
**EnrollmentStartedAt** | Pointer to **string** | Enrollment started at | [optional] 
**DoorLocks** | Pointer to [**[]InlineResponse20066DoorLocks**](InlineResponse20066DoorLocks.md) | Door locks | [optional] 

## Methods

### NewInlineResponse20066

`func NewInlineResponse20066() *InlineResponse20066`

NewInlineResponse20066 instantiates a new InlineResponse20066 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20066WithDefaults

`func NewInlineResponse20066WithDefaults() *InlineResponse20066`

NewInlineResponse20066WithDefaults instantiates a new InlineResponse20066 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnrollmentId

`func (o *InlineResponse20066) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InlineResponse20066) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InlineResponse20066) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InlineResponse20066) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20066) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20066) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20066) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20066) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20066) GetRequest() InlineResponse20115Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20066) GetRequestOk() (*InlineResponse20115Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20066) SetRequest(v InlineResponse20115Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20066) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20066) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20066) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20066) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20066) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnrollmentStartedAt

`func (o *InlineResponse20066) GetEnrollmentStartedAt() string`

GetEnrollmentStartedAt returns the EnrollmentStartedAt field if non-nil, zero value otherwise.

### GetEnrollmentStartedAtOk

`func (o *InlineResponse20066) GetEnrollmentStartedAtOk() (*string, bool)`

GetEnrollmentStartedAtOk returns a tuple with the EnrollmentStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentStartedAt

`func (o *InlineResponse20066) SetEnrollmentStartedAt(v string)`

SetEnrollmentStartedAt sets EnrollmentStartedAt field to given value.

### HasEnrollmentStartedAt

`func (o *InlineResponse20066) HasEnrollmentStartedAt() bool`

HasEnrollmentStartedAt returns a boolean if a field has been set.

### GetDoorLocks

`func (o *InlineResponse20066) GetDoorLocks() []InlineResponse20066DoorLocks`

GetDoorLocks returns the DoorLocks field if non-nil, zero value otherwise.

### GetDoorLocksOk

`func (o *InlineResponse20066) GetDoorLocksOk() (*[]InlineResponse20066DoorLocks, bool)`

GetDoorLocksOk returns a tuple with the DoorLocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorLocks

`func (o *InlineResponse20066) SetDoorLocks(v []InlineResponse20066DoorLocks)`

SetDoorLocks sets DoorLocks field to given value.

### HasDoorLocks

`func (o *InlineResponse20066) HasDoorLocks() bool`

HasDoorLocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


