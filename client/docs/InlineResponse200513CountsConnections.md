# InlineResponse200513CountsConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Wireless LAN controller associated total access point count | [optional] 
**ByStatus** | Pointer to [**InlineResponse200513CountsConnectionsByStatus**](InlineResponse200513CountsConnectionsByStatus.md) |  | [optional] 

## Methods

### NewInlineResponse200513CountsConnections

`func NewInlineResponse200513CountsConnections() *InlineResponse200513CountsConnections`

NewInlineResponse200513CountsConnections instantiates a new InlineResponse200513CountsConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200513CountsConnectionsWithDefaults

`func NewInlineResponse200513CountsConnectionsWithDefaults() *InlineResponse200513CountsConnections`

NewInlineResponse200513CountsConnectionsWithDefaults instantiates a new InlineResponse200513CountsConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200513CountsConnections) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200513CountsConnections) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200513CountsConnections) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200513CountsConnections) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *InlineResponse200513CountsConnections) GetByStatus() InlineResponse200513CountsConnectionsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse200513CountsConnections) GetByStatusOk() (*InlineResponse200513CountsConnectionsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse200513CountsConnections) SetByStatus(v InlineResponse200513CountsConnectionsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse200513CountsConnections) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


