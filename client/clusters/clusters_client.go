package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new clusters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for clusters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateCluster Create a new cluster
*/
func (a *Client) CreateCluster(params *CreateClusterParams) (*CreateClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCluster",
		Method:             "POST",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterCreated), nil
}

/*
DeleteSingleCluster Delete the specified cluster
*/
func (a *Client) DeleteSingleCluster(params *DeleteSingleClusterParams) (*DeleteSingleClusterNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSingleClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSingleCluster",
		Method:             "DELETE",
		PathPattern:        "/clusters/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSingleClusterReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSingleClusterNoContent), nil
}

/*
FindClusters Returns all clusters that the user is able to access
*/
func (a *Client) FindClusters(params *FindClustersParams) (*FindClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findClusters",
		Method:             "GET",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FindClustersReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindClustersOK), nil
}

/*
FindSingleCluster Return detailed information about a single cluster
*/
func (a *Client) FindSingleCluster(params *FindSingleClusterParams) (*FindSingleClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSingleClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findSingleCluster",
		Method:             "GET",
		PathPattern:        "/clusters/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FindSingleClusterReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindSingleClusterOK), nil
}

/*
UpdateSingleCluster Update the specified cluster
*/
func (a *Client) UpdateSingleCluster(params *UpdateSingleClusterParams) (*UpdateSingleClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSingleClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSingleCluster",
		Method:             "PUT",
		PathPattern:        "/clusters/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateSingleClusterReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSingleClusterAccepted), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
