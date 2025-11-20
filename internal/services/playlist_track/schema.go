// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_track

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*PlaylistTrackResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"playlist_id": schema.StringAttribute{
				Description:   "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) of the playlist.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"insert_before": schema.Int64Attribute{
				Description: "The position where the items should be inserted.<br/>To reorder the items to the end of the playlist, simply set _insert_before_ to the position after the last item.<br/>Examples:<br/>To reorder the first item to the last position in a playlist with 10 items, set _range_start_ to 0, and _insert_before_ to 10.<br/>To reorder the last item in a playlist with 10 items to the start of the playlist, set _range_start_ to 9, and _insert_before_ to 0.",
				Optional:    true,
			},
			"range_length": schema.Int64Attribute{
				Description: "The amount of items to be reordered. Defaults to 1 if not set.<br/>The range of items to be reordered begins from the _range_start_ position, and includes the _range_length_ subsequent items.<br/>Example:<br/>To move the items at index 9-10 to the start of the playlist, _range_start_ is set to 9, and _range_length_ is set to 2.",
				Optional:    true,
			},
			"range_start": schema.Int64Attribute{
				Description: "The position of the first item to be reordered.",
				Optional:    true,
			},
			"snapshot_id": schema.StringAttribute{
				Description: "The playlist's snapshot ID against which you want to make the changes.",
				Optional:    true,
			},
			"uris": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *PlaylistTrackResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PlaylistTrackResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
