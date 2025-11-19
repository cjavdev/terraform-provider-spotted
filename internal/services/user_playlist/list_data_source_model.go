// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_playlist

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserPlaylistsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[UserPlaylistsItemsDataSourceModel] `json:"items,computed"`
}

type UserPlaylistsDataSourceModel struct {
	UserID   types.String                                                    `tfsdk:"user_id" path:"user_id,required"`
	Limit    types.Int64                                                     `tfsdk:"limit" query:"limit,computed_optional"`
	Offset   types.Int64                                                     `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems types.Int64                                                     `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[UserPlaylistsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *UserPlaylistsDataSourceModel) toListParams(_ context.Context) (params spotted.UserPlaylistListParams, diags diag.Diagnostics) {
	params = spotted.UserPlaylistListParams{}

	if !m.Limit.IsNull() {
		params.Limit = param.NewOpt(m.Limit.ValueInt64())
	}
	if !m.Offset.IsNull() {
		params.Offset = param.NewOpt(m.Offset.ValueInt64())
	}

	return
}

type UserPlaylistsItemsDataSourceModel struct {
	ID            types.String                                                       `tfsdk:"id" json:"id,computed"`
	Collaborative types.Bool                                                         `tfsdk:"collaborative" json:"collaborative,computed"`
	Description   types.String                                                       `tfsdk:"description" json:"description,computed"`
	ExternalURLs  customfield.NestedObject[UserPlaylistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href          types.String                                                       `tfsdk:"href" json:"href,computed"`
	Images        customfield.NestedObjectList[UserPlaylistsImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name          types.String                                                       `tfsdk:"name" json:"name,computed"`
	Owner         customfield.NestedObject[UserPlaylistsOwnerDataSourceModel]        `tfsdk:"owner" json:"owner,computed"`
	Public        types.Bool                                                         `tfsdk:"public" json:"public,computed"`
	SnapshotID    types.String                                                       `tfsdk:"snapshot_id" json:"snapshot_id,computed"`
	Tracks        customfield.NestedObject[UserPlaylistsTracksDataSourceModel]       `tfsdk:"tracks" json:"tracks,computed"`
	Type          types.String                                                       `tfsdk:"type" json:"type,computed"`
	Uri           types.String                                                       `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type UserPlaylistsOwnerDataSourceModel struct {
	ID           types.String                                                            `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistsOwnerExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                            `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                            `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                            `tfsdk:"uri" json:"uri,computed"`
	DisplayName  types.String                                                            `tfsdk:"display_name" json:"display_name,computed"`
}

type UserPlaylistsOwnerExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistsTracksDataSourceModel struct {
	Href  types.String `tfsdk:"href" json:"href,computed"`
	Total types.Int64  `tfsdk:"total" json:"total,computed"`
}
