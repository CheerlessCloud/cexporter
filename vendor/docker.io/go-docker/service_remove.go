package docker // import "docker.io/go-docker"

import "golang.org/x/net/context"

// ServiceRemove kills and removes a service.
func (cli *Client) ServiceRemove(ctx context.Context, serviceID string) error {
	resp, err := cli.delete(ctx, "/services/"+serviceID, nil, nil)
	ensureReaderClosed(resp)
	return wrapResponseError(err, resp, "service", serviceID)
}
