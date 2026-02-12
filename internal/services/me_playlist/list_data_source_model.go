// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_playlist

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MePlaylistsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MePlaylistsItemsDataSourceModel] `json:"items,computed"`
}

type MePlaylistsDataSourceModel struct {
	Limit    types.Int64                                                   `tfsdk:"limit" query:"limit,computed_optional"`
	Offset   types.Int64                                                   `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems types.Int64                                                   `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[MePlaylistsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *MePlaylistsDataSourceModel) toListParams(_ context.Context) (params spotted.MePlaylistListParams, diags diag.Diagnostics) {
	params = spotted.MePlaylistListParams{}

	if !m.Limit.IsNull() {
		params.Limit = param.NewOpt(m.Limit.ValueInt64())
	}
	if !m.Offset.IsNull() {
		params.Offset = param.NewOpt(m.Offset.ValueInt64())
	}

	return
}

type MePlaylistsItemsDataSourceModel struct {
	Href      types.String `tfsdk:"href" json:"href,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Total     types.Int64  `tfsdk:"total" json:"total,computed"`
}

type MePlaylistsExternalURLsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlaylistsImagesDataSourceModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type MePlaylistsOwnerDataSourceModel struct {
	ID           types.String                                                          `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MePlaylistsOwnerExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                          `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                            `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                          `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                          `tfsdk:"uri" json:"uri,computed"`
	DisplayName  types.String                                                          `tfsdk:"display_name" json:"display_name,computed"`
}

type MePlaylistsOwnerExternalURLsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlaylistsTracksDataSourceModel struct {
	Href      types.String `tfsdk:"href" json:"href,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Total     types.Int64  `tfsdk:"total" json:"total,computed"`
}
