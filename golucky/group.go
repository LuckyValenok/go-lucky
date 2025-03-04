package golucky

import (
	"context"
	"fmt"
)

type Group struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Weight      int      `json:"weight"`
	Nodes       []Node   `json:"nodes"`
	Metadata    Metadata `json:"metadata"`
}

type groupName struct {
	Name string `json:"name"`
}

func (c *Client) ListGroups(ctx context.Context) (*[]string, error) {
	return getRequest[[]string](ctx, fmt.Sprintf("%s/group", c.config.RestIp), c.config.AuthKey)
}

func (c *Client) CreateGroup(ctx context.Context, name string) (*Group, error) {
	return postRequestBody[Group](ctx, fmt.Sprintf("%s/group", c.config.RestIp), groupName{Name: name}, c.config.AuthKey)
}

func (c *Client) DeleteGroup(ctx context.Context, name string) error {
	return deleteRequestNoResponse(ctx, fmt.Sprintf("%s/group/%s", c.config.RestIp, name), nil, c.config.AuthKey)
}

func (c *Client) GetGroup(ctx context.Context, group string) (*Group, error) {
	return getRequest[Group](ctx, fmt.Sprintf("%s/group/%s", c.config.RestIp, group), c.config.AuthKey)
}

// GetGroupNodes Returns all of a group's Nodes
func (c *Client) GetGroupNodes(ctx context.Context, groupName string) (*[]Node, error) {
	return getRequest[[]Node](ctx, fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), c.config.AuthKey)
}

// AddGroupNode Adds a Node to a group, then returns the new array of nodes
func (c *Client) AddGroupNode(ctx context.Context, groupName string, node NewNode, mergeStrategy NodeMergeStrategy) (*[]Node, error) {
	return postRequestBody[[]Node](ctx, fmt.Sprintf("%s/group/%s/nodes?temporaryNodeMergeStrategy=%s", c.config.RestIp, groupName, mergeStrategy), node, c.config.AuthKey)
}

// AddGroupNodes Adds multiple Nodes to a group, then returns the new array of nodes
func (c *Client) AddGroupNodes(ctx context.Context, groupName string, nodes []NewNode, mergeStrategy NodeMergeStrategy) (*[]Node, error) {
	return patchRequestBody[[]Node](ctx, fmt.Sprintf("%s/group/%s/nodes?temporaryNodeMergeStrategy=%s", c.config.RestIp, groupName, mergeStrategy), nodes, c.config.AuthKey)
}

// SetGroupNodes Replaces all the Nodes of the group with newNodes
func (c *Client) SetGroupNodes(ctx context.Context, groupName string, newNodes []NewNode) error {
	return putRequestNoResponse(ctx, fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), newNodes, c.config.AuthKey)
}

// RemoveGroupNodes Removes multiple Nodes from a group
func (c *Client) RemoveGroupNodes(ctx context.Context, groupName string, nodes []NewNode) error {
	return deleteRequestNoResponse(ctx, fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), nodes, c.config.AuthKey)
}

// ClearGroupNodes Removes all Nodes from a group
func (c *Client) ClearGroupNodes(ctx context.Context, groupName string) error {
	return deleteRequestNoResponse(ctx, fmt.Sprintf("%s/group/%s/nodes", c.config.RestIp, groupName), nil, c.config.AuthKey)
}

func (c *Client) GroupHasPermission(ctx context.Context, group string, permission string) (*PermissionCheckResult, error) {
	return getRequest[PermissionCheckResult](ctx, fmt.Sprintf("%s/group/%s/permission-check?permission=%s", c.config.RestIp, group, permission), c.config.AuthKey)
}

func (c *Client) GroupHasPermissionWithOptions(ctx context.Context, group string, request PermissionCheckRequest) (*PermissionCheckResult, error) {
	return postRequestBody[PermissionCheckResult](ctx, fmt.Sprintf("%s/group/%s/permission-check", c.config.RestIp, group), request, c.config.AuthKey)
}
