// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_show

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type MeShowsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MeShowsItemsDataSourceModel] `json:"items,computed"`
}

type MeShowsDataSourceModel struct {
	Limit    types.Int64                                               `tfsdk:"limit" query:"limit,computed_optional"`
	Offset   types.Int64                                               `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems types.Int64                                               `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[MeShowsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *MeShowsDataSourceModel) toListParams(_ context.Context) (params spotted.MeShowListParams, diags diag.Diagnostics) {
	params = spotted.MeShowListParams{}

	if !m.Limit.IsNull() {
		params.Limit = param.NewOpt(m.Limit.ValueInt64())
	}
	if !m.Offset.IsNull() {
		params.Offset = param.NewOpt(m.Offset.ValueInt64())
	}

	return
}

type MeShowsItemsDataSourceModel struct {
	AddedAt timetypes.RFC3339                                    `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	Show    customfield.NestedObject[MeShowsShowDataSourceModel] `tfsdk:"show" json:"show,computed"`
}

type MeShowsShowDataSourceModel struct {
	ID                 types.String                                                       `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                     `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[MeShowsShowCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                       `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                         `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[MeShowsShowExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                       `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                       `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[MeShowsShowImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                         `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                     `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                       `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                       `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                       `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                        `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                       `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                       `tfsdk:"uri" json:"uri,computed"`
}

type MeShowsShowCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type MeShowsShowExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeShowsShowImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}
