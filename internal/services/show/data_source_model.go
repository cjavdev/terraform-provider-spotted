// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package show

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type ShowDataSourceModel struct {
	ID                 types.String                                                `tfsdk:"id" path:"id,required"`
	Market             types.String                                                `tfsdk:"market" query:"market,optional"`
	Description        types.String                                                `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                  `tfsdk:"explicit" json:"explicit,computed"`
	Href               types.String                                                `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                `tfsdk:"html_description" json:"html_description,computed"`
	IsExternallyHosted types.Bool                                                  `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	MediaType          types.String                                                `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                 `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                `tfsdk:"uri" json:"uri,computed"`
	AvailableMarkets   customfield.List[types.String]                              `tfsdk:"available_markets" json:"available_markets,computed"`
	Languages          customfield.List[types.String]                              `tfsdk:"languages" json:"languages,computed"`
	Copyrights         customfield.NestedObjectList[ShowCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Episodes           customfield.NestedObject[ShowEpisodesDataSourceModel]       `tfsdk:"episodes" json:"episodes,computed"`
	ExternalURLs       customfield.NestedObject[ShowExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Images             customfield.NestedObjectList[ShowImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
}

func (m *ShowDataSourceModel) toReadParams(_ context.Context) (params spotted.ShowGetParams, diags diag.Diagnostics) {
	params = spotted.ShowGetParams{}

	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}

	return
}

type ShowCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type ShowEpisodesDataSourceModel struct {
	Href     types.String                                                   `tfsdk:"href" json:"href,computed"`
	Limit    types.Int64                                                    `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                   `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                    `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                   `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                    `tfsdk:"total" json:"total,computed"`
	Items    customfield.NestedObjectList[ShowEpisodesItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
}

type ShowEpisodesItemsDataSourceModel struct {
	ID                   types.String                                                           `tfsdk:"id" json:"id,computed"`
	AudioPreviewURL      types.String                                                           `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                           `tfsdk:"description" json:"description,computed"`
	DurationMs           types.Int64                                                            `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                             `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs         customfield.NestedObject[ShowEpisodesItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                           `tfsdk:"href" json:"href,computed"`
	HTMLDescription      types.String                                                           `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[ShowEpisodesItemsImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                             `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	IsPlayable           types.Bool                                                             `tfsdk:"is_playable" json:"is_playable,computed"`
	Languages            customfield.List[types.String]                                         `tfsdk:"languages" json:"languages,computed"`
	Name                 types.String                                                           `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                           `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                           `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Type                 types.String                                                           `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                           `tfsdk:"uri" json:"uri,computed"`
	Language             types.String                                                           `tfsdk:"language" json:"language,computed"`
	Restrictions         customfield.NestedObject[ShowEpisodesItemsRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	ResumePoint          customfield.NestedObject[ShowEpisodesItemsResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type ShowEpisodesItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type ShowEpisodesItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type ShowEpisodesItemsRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type ShowEpisodesItemsResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}

type ShowExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type ShowImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}
