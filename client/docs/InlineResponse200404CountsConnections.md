# InlineResponse200404CountsConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Wireless LAN controller associated total access point count | [optional] 
**ByStatus** | Pointer to [**InlineResponse200404CountsConnectionsByStatus**](InlineResponse200404CountsConnectionsByStatus.md) |  | [optional] 

## Methods

### NewInlineResponse200404CountsConnections

`func NewInlineResponse200404CountsConnections() *InlineResponse200404CountsConnections`

NewInlineResponse200404CountsConnections instantiates a new InlineResponse200404CountsConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200404CountsConnectionsWithDefaults

`func NewInlineResponse200404CountsConnectionsWithDefaults() *InlineResponse200404CountsConnections`

NewInlineResponse200404CountsConnectionsWithDefaults instantiates a new InlineResponse200404CountsConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InlineResponse200404CountsConnections) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200404CountsConnections) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200404CountsConnections) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200404CountsConnections) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByStatus

`func (o *InlineResponse200404CountsConnections) GetByStatus() InlineResponse200404CountsConnectionsByStatus`

GetByStatus returns the ByStatus field if non-nil, zero value otherwise.

### GetByStatusOk

`func (o *InlineResponse200404CountsConnections) GetByStatusOk() (*InlineResponse200404CountsConnectionsByStatus, bool)`

GetByStatusOk returns a tuple with the ByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByStatus

`func (o *InlineResponse200404CountsConnections) SetByStatus(v InlineResponse200404CountsConnectionsByStatus)`

SetByStatus sets ByStatus field to given value.

### HasByStatus

`func (o *InlineResponse200404CountsConnections) HasByStatus() bool`

HasByStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


