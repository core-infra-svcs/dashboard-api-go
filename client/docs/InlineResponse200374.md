# InlineResponse200374

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisenrollmentId** | Pointer to **string** | Id to check the status of your disenrollment | [optional] 
**Url** | Pointer to **string** | Url to check the status of your disenrollment | [optional] 
**Request** | Pointer to [**InlineResponse20123Request**](InlineResponse20123Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the disenrollment | [optional] 
**DoorLocks** | Pointer to [**[]InlineResponse200374DoorLocks**](InlineResponse200374DoorLocks.md) | Door locks | [optional] 

## Methods

### NewInlineResponse200374

`func NewInlineResponse200374() *InlineResponse200374`

NewInlineResponse200374 instantiates a new InlineResponse200374 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200374WithDefaults

`func NewInlineResponse200374WithDefaults() *InlineResponse200374`

NewInlineResponse200374WithDefaults instantiates a new InlineResponse200374 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisenrollmentId

`func (o *InlineResponse200374) GetDisenrollmentId() string`

GetDisenrollmentId returns the DisenrollmentId field if non-nil, zero value otherwise.

### GetDisenrollmentIdOk

`func (o *InlineResponse200374) GetDisenrollmentIdOk() (*string, bool)`

GetDisenrollmentIdOk returns a tuple with the DisenrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisenrollmentId

`func (o *InlineResponse200374) SetDisenrollmentId(v string)`

SetDisenrollmentId sets DisenrollmentId field to given value.

### HasDisenrollmentId

`func (o *InlineResponse200374) HasDisenrollmentId() bool`

HasDisenrollmentId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200374) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200374) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200374) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200374) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse200374) GetRequest() InlineResponse20123Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse200374) GetRequestOk() (*InlineResponse20123Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse200374) SetRequest(v InlineResponse20123Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse200374) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200374) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200374) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200374) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200374) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDoorLocks

`func (o *InlineResponse200374) GetDoorLocks() []InlineResponse200374DoorLocks`

GetDoorLocks returns the DoorLocks field if non-nil, zero value otherwise.

### GetDoorLocksOk

`func (o *InlineResponse200374) GetDoorLocksOk() (*[]InlineResponse200374DoorLocks, bool)`

GetDoorLocksOk returns a tuple with the DoorLocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorLocks

`func (o *InlineResponse200374) SetDoorLocks(v []InlineResponse200374DoorLocks)`

SetDoorLocks sets DoorLocks field to given value.

### HasDoorLocks

`func (o *InlineResponse200374) HasDoorLocks() bool`

HasDoorLocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


