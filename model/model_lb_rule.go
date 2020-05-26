package nsxt

// Load balancer rules allow customization of load balancing behavior using match/action rules. Currently, load balancer rules are supported for only layer 7 virtual servers with application profile LbHttpProfile. Each application rule consists of one or more match conditions and one or more actions. Load balancer rules could be used by different load balancer services. 
type LbRule struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Each load balancer rule is used at a specific phase of load balancer processing. Currently five phases are supported, HTTP_REQUEST_REWRITE, HTTP_FORWARDING, HTTP_RESPONSE_REWRITE, HTTP_ACCESS and TRANSPORT. When an HTTP request message is received by load balancer, all HTTP_REQUEST_REWRITE rules, if present are executed in the order they are applied to virtual server. And then if HTTP_FORWARDING rules present, only first matching rule's action is executed, remaining rules are not checked. HTTP_FORWARDING rules can have only one action. If the request is forwarded to a backend server and the response goes back to load balancer, all HTTP_RESPONSE_REWRITE rules, if present, are executed in the order they are applied to the virtual server. In HTTP_ACCESS phase, user can define action to control access using JWT authentication. In TRANSPORT phase, user can define the condition to match SNI in TLS client hello and define the action to do SSL end-to-end, SSL offloading or SSL passthrough using a specific load balancer server pool. 
	Phase string `json:"phase"`
	// A list of match conditions used to match application traffic. Multiple match conditions can be specified in one load balancer rule, each match condition defines a criterion to match application traffic. If no match conditions are specified, then the load balancer rule will always match and it is used typically to define default rules. If more than one match condition is specified, then match strategy determines if all conditions should match or any one condition should match for the load balancer rule to considered a match. 
	MatchConditions []LbRuleCondition `json:"match_conditions,omitempty"`
	// A list of actions to be executed at specified phase when load balancer rule matches. The actions are used to manipulate application traffic, such as rewrite URI of HTTP messages, redirect HTTP messages, etc. 
	Actions []LbRuleAction `json:"actions"`
	// Strategy to define how load balancer rule is considered a match when multiple match conditions are specified in one rule. If match_stragety is set to ALL, then load balancer rule is considered a match only if all the conditions match. If match_strategy is set to ANY, then load balancer rule is considered a match if any one of the conditions match. 
	MatchStrategy string `json:"match_strategy"`
}