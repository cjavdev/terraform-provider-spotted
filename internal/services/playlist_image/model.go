// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_image

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/apijson"
)

type PlaylistImageModel struct {
	PlaylistID types.String `tfsdk:"playlist_id" path:"playlist_id,required"`
	Body       types.String `tfsdk:"body" json:"body,required"`
}

func (m PlaylistImageModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Body)
}

func (m PlaylistImageModel) MarshalJSONForUpdate(state PlaylistImageModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Body, state.Body)
}
