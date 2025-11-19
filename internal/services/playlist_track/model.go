// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_track

import (
	"github.com/cjavdev/terraform-provider-spotted/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PlaylistTrackModel struct {
	PlaylistID   types.String    `tfsdk:"playlist_id" path:"playlist_id,required"`
	InsertBefore types.Int64     `tfsdk:"insert_before" json:"insert_before,optional"`
	RangeLength  types.Int64     `tfsdk:"range_length" json:"range_length,optional"`
	RangeStart   types.Int64     `tfsdk:"range_start" json:"range_start,optional"`
	SnapshotID   types.String    `tfsdk:"snapshot_id" json:"snapshot_id,optional"`
	Uris         *[]types.String `tfsdk:"uris" json:"uris,optional"`
}

func (m PlaylistTrackModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PlaylistTrackModel) MarshalJSONForUpdate(state PlaylistTrackModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
