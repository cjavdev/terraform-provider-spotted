// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_image

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*PlaylistImageResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"playlist_id": schema.StringAttribute{
				Description:   "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"body": schema.StringAttribute{
				Description: "Base64 encoded JPEG image data, maximum payload size is 256 KB.",
				Required:    true,
			},
		},
	}
}

func (r *PlaylistImageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PlaylistImageResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
