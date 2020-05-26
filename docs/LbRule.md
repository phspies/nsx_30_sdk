# LbRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | Unique identifier of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**Phase** | **string** | Each load balancer rule is used at a specific phase of load balancer processing. Currently five phases are supported, HTTP_REQUEST_REWRITE, HTTP_FORWARDING, HTTP_RESPONSE_REWRITE, HTTP_ACCESS and TRANSPORT. When an HTTP request message is received by load balancer, all HTTP_REQUEST_REWRITE rules, if present are executed in the order they are applied to virtual server. And then if HTTP_FORWARDING rules present, only first matching rule&#x27;s action is executed, remaining rules are not checked. HTTP_FORWARDING rules can have only one action. If the request is forwarded to a backend server and the response goes back to load balancer, all HTTP_RESPONSE_REWRITE rules, if present, are executed in the order they are applied to the virtual server. In HTTP_ACCESS phase, user can define action to control access using JWT authentication. In TRANSPORT phase, user can define the condition to match SNI in TLS client hello and define the action to do SSL end-to-end, SSL offloading or SSL passthrough using a specific load balancer server pool.  | [default to null]
**MatchConditions** | [**[]LbRuleCondition**](LbRuleCondition.md) | A list of match conditions used to match application traffic. Multiple match conditions can be specified in one load balancer rule, each match condition defines a criterion to match application traffic. If no match conditions are specified, then the load balancer rule will always match and it is used typically to define default rules. If more than one match condition is specified, then match strategy determines if all conditions should match or any one condition should match for the load balancer rule to considered a match.  | [optional] [default to null]
**Actions** | [**[]LbRuleAction**](LbRuleAction.md) | A list of actions to be executed at specified phase when load balancer rule matches. The actions are used to manipulate application traffic, such as rewrite URI of HTTP messages, redirect HTTP messages, etc.  | [default to null]
**MatchStrategy** | **string** | Strategy to define how load balancer rule is considered a match when multiple match conditions are specified in one rule. If match_stragety is set to ALL, then load balancer rule is considered a match only if all the conditions match. If match_strategy is set to ANY, then load balancer rule is considered a match if any one of the conditions match.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

