# InlineResponse200523

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisenrollmentId** | Pointer to **string** | Id to check the status of your disenrollment | [optional] 
**Url** | Pointer to **string** | Url to check the status of your disenrollment | [optional] 
**Request** | Pointer to [**InlineResponse20129Request**](InlineResponse20129Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the disenrollment | [optional] 
**DoorLocks** | Pointer to [**[]InlineResponse200523DoorLocks**](InlineResponse200523DoorLocks.md) | Door locks | [optional] 

## Methods

### NewInlineResponse200523

`func NewInlineResponse200523() *InlineResponse200523`

NewInlineResponse200523 instantiates a new InlineResponse200523 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200523WithDefaults

`func NewInlineResponse200523WithDefaults() *InlineResponse200523`

NewInlineResponse200523WithDefaults instantiates a new InlineResponse200523 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisenrollmentId

`func (o *InlineResponse200523) GetDisenrollmentId() string`

GetDisenrollmentId returns the DisenrollmentId field if non-nil, zero value otherwise.

### GetDisenrollmentIdOk

`func (o *InlineResponse200523) GetDisenrollmentIdOk() (*string, bool)`

GetDisenrollmentIdOk returns a tuple with the DisenrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisenrollmentId

`func (o *InlineResponse200523) SetDisenrollmentId(v string)`

SetDisenrollmentId sets DisenrollmentId field to given value.

### HasDisenrollmentId

`func (o *InlineResponse200523) HasDisenrollmentId() bool`

HasDisenrollmentId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200523) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200523) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200523) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200523) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse200523) GetRequest() InlineResponse20129Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse200523) GetRequestOk() (*InlineResponse20129Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse200523) SetRequest(v InlineResponse20129Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse200523) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200523) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200523) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200523) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200523) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDoorLocks

`func (o *InlineResponse200523) GetDoorLocks() []InlineResponse200523DoorLocks`

GetDoorLocks returns the DoorLocks field if non-nil, zero value otherwise.

### GetDoorLocksOk

`func (o *InlineResponse200523) GetDoorLocksOk() (*[]InlineResponse200523DoorLocks, bool)`

GetDoorLocksOk returns a tuple with the DoorLocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorLocks

`func (o *InlineResponse200523) SetDoorLocks(v []InlineResponse200523DoorLocks)`

SetDoorLocks sets DoorLocks field to given value.

### HasDoorLocks

`func (o *InlineResponse200523) HasDoorLocks() bool`

HasDoorLocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


