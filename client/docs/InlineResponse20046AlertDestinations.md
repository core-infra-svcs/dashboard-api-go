# InlineResponse20046AlertDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** | A list of emails that will receive information about the alert | [optional] 
**SmsNumbers** | Pointer to **[]string** | A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts. | [optional] 
**AllAdmins** | Pointer to **bool** | If true, then all network admins will receive emails for this alert | [optional] 
**Snmp** | Pointer to **bool** | If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network | [optional] 
**HttpServerIds** | Pointer to **[]string** | A list of HTTP server IDs to send a Webhook to for this alert | [optional] 

## Methods

### NewInlineResponse20046AlertDestinations

`func NewInlineResponse20046AlertDestinations() *InlineResponse20046AlertDestinations`

NewInlineResponse20046AlertDestinations instantiates a new InlineResponse20046AlertDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20046AlertDestinationsWithDefaults

`func NewInlineResponse20046AlertDestinationsWithDefaults() *InlineResponse20046AlertDestinations`

NewInlineResponse20046AlertDestinationsWithDefaults instantiates a new InlineResponse20046AlertDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *InlineResponse20046AlertDestinations) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *InlineResponse20046AlertDestinations) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *InlineResponse20046AlertDestinations) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *InlineResponse20046AlertDestinations) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetSmsNumbers

`func (o *InlineResponse20046AlertDestinations) GetSmsNumbers() []string`

GetSmsNumbers returns the SmsNumbers field if non-nil, zero value otherwise.

### GetSmsNumbersOk

`func (o *InlineResponse20046AlertDestinations) GetSmsNumbersOk() (*[]string, bool)`

GetSmsNumbersOk returns a tuple with the SmsNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumbers

`func (o *InlineResponse20046AlertDestinations) SetSmsNumbers(v []string)`

SetSmsNumbers sets SmsNumbers field to given value.

### HasSmsNumbers

`func (o *InlineResponse20046AlertDestinations) HasSmsNumbers() bool`

HasSmsNumbers returns a boolean if a field has been set.

### GetAllAdmins

`func (o *InlineResponse20046AlertDestinations) GetAllAdmins() bool`

GetAllAdmins returns the AllAdmins field if non-nil, zero value otherwise.

### GetAllAdminsOk

`func (o *InlineResponse20046AlertDestinations) GetAllAdminsOk() (*bool, bool)`

GetAllAdminsOk returns a tuple with the AllAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAdmins

`func (o *InlineResponse20046AlertDestinations) SetAllAdmins(v bool)`

SetAllAdmins sets AllAdmins field to given value.

### HasAllAdmins

`func (o *InlineResponse20046AlertDestinations) HasAllAdmins() bool`

HasAllAdmins returns a boolean if a field has been set.

### GetSnmp

`func (o *InlineResponse20046AlertDestinations) GetSnmp() bool`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *InlineResponse20046AlertDestinations) GetSnmpOk() (*bool, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *InlineResponse20046AlertDestinations) SetSnmp(v bool)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *InlineResponse20046AlertDestinations) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetHttpServerIds

`func (o *InlineResponse20046AlertDestinations) GetHttpServerIds() []string`

GetHttpServerIds returns the HttpServerIds field if non-nil, zero value otherwise.

### GetHttpServerIdsOk

`func (o *InlineResponse20046AlertDestinations) GetHttpServerIdsOk() (*[]string, bool)`

GetHttpServerIdsOk returns a tuple with the HttpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerIds

`func (o *InlineResponse20046AlertDestinations) SetHttpServerIds(v []string)`

SetHttpServerIds sets HttpServerIds field to given value.

### HasHttpServerIds

`func (o *InlineResponse20046AlertDestinations) HasHttpServerIds() bool`

HasHttpServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


