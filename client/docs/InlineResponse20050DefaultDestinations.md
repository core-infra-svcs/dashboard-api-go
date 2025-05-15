# InlineResponse20050DefaultDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** | A list of emails that will receive the alert(s). | [optional] 
**AllAdmins** | Pointer to **bool** | If true, then all network admins will receive emails. | [optional] 
**Snmp** | Pointer to **bool** | If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network. | [optional] 
**HttpServerIds** | Pointer to **[]string** | A list of HTTP server IDs to send a Webhook to | [optional] 

## Methods

### NewInlineResponse20050DefaultDestinations

`func NewInlineResponse20050DefaultDestinations() *InlineResponse20050DefaultDestinations`

NewInlineResponse20050DefaultDestinations instantiates a new InlineResponse20050DefaultDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20050DefaultDestinationsWithDefaults

`func NewInlineResponse20050DefaultDestinationsWithDefaults() *InlineResponse20050DefaultDestinations`

NewInlineResponse20050DefaultDestinationsWithDefaults instantiates a new InlineResponse20050DefaultDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *InlineResponse20050DefaultDestinations) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *InlineResponse20050DefaultDestinations) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *InlineResponse20050DefaultDestinations) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *InlineResponse20050DefaultDestinations) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetAllAdmins

`func (o *InlineResponse20050DefaultDestinations) GetAllAdmins() bool`

GetAllAdmins returns the AllAdmins field if non-nil, zero value otherwise.

### GetAllAdminsOk

`func (o *InlineResponse20050DefaultDestinations) GetAllAdminsOk() (*bool, bool)`

GetAllAdminsOk returns a tuple with the AllAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAdmins

`func (o *InlineResponse20050DefaultDestinations) SetAllAdmins(v bool)`

SetAllAdmins sets AllAdmins field to given value.

### HasAllAdmins

`func (o *InlineResponse20050DefaultDestinations) HasAllAdmins() bool`

HasAllAdmins returns a boolean if a field has been set.

### GetSnmp

`func (o *InlineResponse20050DefaultDestinations) GetSnmp() bool`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *InlineResponse20050DefaultDestinations) GetSnmpOk() (*bool, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *InlineResponse20050DefaultDestinations) SetSnmp(v bool)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *InlineResponse20050DefaultDestinations) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetHttpServerIds

`func (o *InlineResponse20050DefaultDestinations) GetHttpServerIds() []string`

GetHttpServerIds returns the HttpServerIds field if non-nil, zero value otherwise.

### GetHttpServerIdsOk

`func (o *InlineResponse20050DefaultDestinations) GetHttpServerIdsOk() (*[]string, bool)`

GetHttpServerIdsOk returns a tuple with the HttpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerIds

`func (o *InlineResponse20050DefaultDestinations) SetHttpServerIds(v []string)`

SetHttpServerIds sets HttpServerIds field to given value.

### HasHttpServerIds

`func (o *InlineResponse20050DefaultDestinations) HasHttpServerIds() bool`

HasHttpServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


