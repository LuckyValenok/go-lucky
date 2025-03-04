package golucky

type Node struct {
	Key     string    `json:"key"`
	Value   bool      `json:"value"`
	Type    string    `json:"type"` // Enum: [ permission, regex_permission, inheritance, prefix, suffix, meta, weight, display_name ]
	Expiry  int       `json:"expiry"`
	Context []Context `json:"context"`
}

// NewNode Only Key is required. Note: Value defaults to false
type NewNode struct {
	Key     string    `json:"key"`
	Value   bool      `json:"value"`
	Expiry  int       `json:"expiry"`
	Context []Context `json:"context"`
}

type Metadata struct {
	Prefix       string `json:"prefix"`
	Suffix       string `json:"suffix"`
	PrimaryGroup string `json:"primaryGroup"`
}

type Context struct {
	Key   string `json:"key"`
	Value bool   `json:"value"`
}

type QueryFlag string

const (
	ResolveInheritance                        QueryFlag = "resolve_inheritance"
	IncludeNodesWithoutServerContext          QueryFlag = "include_nodes_without_server_context"
	IncludeNodesWithoutWorldContext           QueryFlag = "include_nodes_without_world_context"
	ApplyInheritanceNodesWithoutServerContext QueryFlag = "apply_inheritance_nodes_without_server_context"
	ApplyInheritanceNodesWithoutWorldContext  QueryFlag = "apply_inheritance_nodes_without_world_context"
)

type QueryMode string

const (
	Contextual    QueryMode = "contextual"
	NonContextual QueryMode = "non_contextual"
)

type QueryOptions struct {
	Mode     QueryMode   `json:"mode,omitempty"`
	Flags    []QueryFlag `json:"flags,omitempty"`
	Contexts []Context   `json:"contexts,omitempty"`
}

type PermissionCheckRequest struct {
	Permission   string       `json:"permission"`
	QueryOptions QueryOptions `json:"queryOptions,omitempty"`
}

type PermissionCheckResult struct {
	Result string `json:"result"` // Tristate: [ true, false, undefined ]
	Node   Node   `json:"node"`
}

type NodeMergeStrategy string

const (
	None                                 NodeMergeStrategy = "none"
	AddNewDurationToExistingMerge        NodeMergeStrategy = "add_new_duration_to_existing"
	ReplaceExistingIfDurationLongerMerge NodeMergeStrategy = "replace_existing_if_duration_longer"
)
