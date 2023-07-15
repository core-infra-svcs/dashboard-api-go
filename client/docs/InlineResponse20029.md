# InlineResponse20029

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | A message regarding the events sent. Usually &#39;null&#39; unless there are no events | [optional] 
**PageStartAt** | Pointer to **string** | An UTC ISO8601 string of the earliest occured at time of the listed events of the page. | [optional] 
**PageEndAt** | Pointer to **string** | An UTC ISO8601 string of the latest occured at time of the listed events of the page. | [optional] 
**Events** | Pointer to [**[]InlineResponse20029Events**](InlineResponse20029Events.md) | An array of events that took place in the network. | [optional] 

## Methods

### NewInlineResponse20029

`func NewInlineResponse20029() *InlineResponse20029`

NewInlineResponse20029 instantiates a new InlineResponse20029 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20029WithDefaults

`func NewInlineResponse20029WithDefaults() *InlineResponse20029`

NewInlineResponse20029WithDefaults instantiates a new InlineResponse20029 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InlineResponse20029) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse20029) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse20029) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse20029) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPageStartAt

`func (o *InlineResponse20029) GetPageStartAt() string`

GetPageStartAt returns the PageStartAt field if non-nil, zero value otherwise.

### GetPageStartAtOk

`func (o *InlineResponse20029) GetPageStartAtOk() (*string, bool)`

GetPageStartAtOk returns a tuple with the PageStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageStartAt

`func (o *InlineResponse20029) SetPageStartAt(v string)`

SetPageStartAt sets PageStartAt field to given value.

### HasPageStartAt

`func (o *InlineResponse20029) HasPageStartAt() bool`

HasPageStartAt returns a boolean if a field has been set.

### GetPageEndAt

`func (o *InlineResponse20029) GetPageEndAt() string`

GetPageEndAt returns the PageEndAt field if non-nil, zero value otherwise.

### GetPageEndAtOk

`func (o *InlineResponse20029) GetPageEndAtOk() (*string, bool)`

GetPageEndAtOk returns a tuple with the PageEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageEndAt

`func (o *InlineResponse20029) SetPageEndAt(v string)`

SetPageEndAt sets PageEndAt field to given value.

### HasPageEndAt

`func (o *InlineResponse20029) HasPageEndAt() bool`

HasPageEndAt returns a boolean if a field has been set.

### GetEvents

`func (o *InlineResponse20029) GetEvents() []InlineResponse20029Events`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *InlineResponse20029) GetEventsOk() (*[]InlineResponse20029Events, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *InlineResponse20029) SetEvents(v []InlineResponse20029Events)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *InlineResponse20029) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


