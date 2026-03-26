package resources

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Businnect/letit_go/interfaces"
	"github.com/Businnect/letit_go/schemas"
)

type AdminBlogResource struct {
	client interfaces.ClientInterface
}

func NewAdminBlogResource(c interfaces.ClientInterface) *AdminBlogResource {
	return &AdminBlogResource{client: c}
}

func (r *AdminBlogResource) Get(ctx context.Context) (schemas.AdminBlogArticle, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/client/admin/blog", nil)
	if err != nil {
		return nil, err
	}

	respBody, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	var result schemas.AdminBlogArticle
	if err := json.NewDecoder(respBody).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *AdminBlogResource) List(ctx context.Context) (*schemas.AdminBlogListResponse, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/client/admin/blog/list", nil)
	if err != nil {
		return nil, err
	}

	respBody, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	var result schemas.AdminBlogListResponse
	if err := json.NewDecoder(respBody).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
