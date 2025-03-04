package golucky

import (
	"context"
	"fmt"
)

type userPair struct {
	UniqueId string `json:"uniqueId"`
	Username string `json:"username"`
}

type UserLookupResult userPair

type UserResult struct {
	UniqueId     string   `json:"uniqueId"`
	Username     string   `json:"username"`
	ParentGroups []string `json:"parentGroups"`
	Nodes        []Node   `json:"nodes"`
	Metadata     Metadata `json:"metadata"`
}

type TrackResult struct {
	Success   bool   `json:"success"`
	Status    string `json:"status"` // Enum: [ success, added_to_first_group, malformed_track, end_of_track, ambiguous_call, undefined_failure ]
	GroupFrom string `json:"groupFrom"`
	GroupTo   string `json:"groupTo"`
}

type trackBody struct {
	Track string `json:"track"`
}

type CreateUserResult struct {
	Outcomes         []string `json:"outcomes"`
	PreviousUsername string   `json:"previousUsername"`
	OtherUniqueIds   []string `json:"otherUniqueIds"`
}

func (c *Client) ListUsers(ctx context.Context) (*[]string, error) {
	return getRequest[[]string](ctx, fmt.Sprintf("%s/user", c.config.RestIp), c.config.AuthKey)
}

func (c *Client) CreateUser(ctx context.Context, uuid string, username string) (*CreateUserResult, error) {
	return postRequestBody[CreateUserResult](ctx, fmt.Sprintf("%s/user", c.config.RestIp), userPair{UniqueId: uuid, Username: username}, c.config.AuthKey)
}

func (c *Client) DeleteUser(ctx context.Context, uuid string) error {
	return deleteRequestNoResponse(ctx, fmt.Sprintf("%s/user/%s", c.config.RestIp, uuid), nil, c.config.AuthKey)
}

func (c *Client) LookupUserUUID(ctx context.Context, uuid string) (*UserLookupResult, error) {
	return getRequest[UserLookupResult](ctx, fmt.Sprintf("%s/user/lookup?uniqueId=%s", c.config.RestIp, uuid), c.config.AuthKey)
}

func (c *Client) LookupUserUsername(ctx context.Context, username string) (*UserLookupResult, error) {
	return getRequest[UserLookupResult](ctx, fmt.Sprintf("%s/user/lookup?username=%s", c.config.RestIp, username), c.config.AuthKey)
}

func (c *Client) GetUserData(ctx context.Context, uuid string) (*UserResult, error) {
	return getRequest[UserResult](ctx, fmt.Sprintf("%s/user/%s", c.config.RestIp, uuid), c.config.AuthKey)
}

// GetUserNodes Returns all of a user's Nodes
func (c *Client) GetUserNodes(ctx context.Context, uuid string) (*[]Node, error) {
	return getRequest[[]Node](ctx, fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), c.config.AuthKey)
}

// AddUserNode Adds a Node to a user, then returns the new array of nodes
func (c *Client) AddUserNode(ctx context.Context, uuid string, node NewNode, mergeStrategy NodeMergeStrategy) (*[]Node, error) {
	return postRequestBody[[]Node](ctx, fmt.Sprintf("%s/user/%s/nodes?temporaryNodeMergeStrategy=%s", c.config.RestIp, uuid, mergeStrategy), node, c.config.AuthKey)
}

// AddUserNodes Adds multiple Nodes to a user, then returns the new array of nodes
func (c *Client) AddUserNodes(ctx context.Context, uuid string, nodes []NewNode, mergeStrategy NodeMergeStrategy) (*[]Node, error) {
	return patchRequestBody[[]Node](ctx, fmt.Sprintf("%s/user/%s/nodes?temporaryNodeMergeStrategy=%s", c.config.RestIp, uuid, mergeStrategy), nodes, c.config.AuthKey)
}

// SetUserNodes Replaces all the Nodes of the user with newNodes
func (c *Client) SetUserNodes(ctx context.Context, uuid string, newNodes []NewNode) error {
	return putRequestNoResponse(ctx, fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), newNodes, c.config.AuthKey)
}

// RemoveUserNodes Removes multiple Nodes from a user
func (c *Client) RemoveUserNodes(ctx context.Context, uuid string, nodes []NewNode) error {
	return deleteRequestNoResponse(ctx, fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), nodes, c.config.AuthKey)
}

// ClearUserNodes Removes all Nodes from a user
func (c *Client) ClearUserNodes(ctx context.Context, uuid string) error {
	return deleteRequestNoResponse(ctx, fmt.Sprintf("%s/user/%s/nodes", c.config.RestIp, uuid), nil, c.config.AuthKey)
}

func (c *Client) UserHasPermission(ctx context.Context, uuid string, permission string) (*PermissionCheckResult, error) {
	return getRequest[PermissionCheckResult](ctx, fmt.Sprintf("%s/user/%s/permission-check?permission=%s", c.config.RestIp, uuid, permission), c.config.AuthKey)
}

func (c *Client) UserHasPermissionWithOptions(ctx context.Context, uuid string, request PermissionCheckRequest) (*PermissionCheckResult, error) {
	return postRequestBody[PermissionCheckResult](ctx, fmt.Sprintf("%s/user/%s/permission-check", c.config.RestIp, uuid), request, c.config.AuthKey)
}

func (c *Client) Promote(ctx context.Context, uuid string, track string) (*TrackResult, error) {
	return postRequestBody[TrackResult](ctx, fmt.Sprintf("%s/user/%s/promote", c.config.RestIp, uuid), trackBody{Track: track}, c.config.AuthKey)
}

func (c *Client) Demote(ctx context.Context, uuid string, track string) (*TrackResult, error) {
	return postRequestBody[TrackResult](ctx, fmt.Sprintf("%s/user/%s/demote", c.config.RestIp, uuid), trackBody{Track: track}, c.config.AuthKey)
}
