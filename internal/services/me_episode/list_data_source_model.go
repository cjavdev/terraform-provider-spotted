// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_episode

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type MeEpisodesItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MeEpisodesItemsDataSourceModel] `json:"items,computed"`
}

type MeEpisodesDataSourceModel struct {
	Market   types.String                                                 `tfsdk:"market" query:"market,optional"`
	Limit    types.Int64                                                  `tfsdk:"limit" query:"limit,computed_optional"`
	Offset   types.Int64                                                  `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems types.Int64                                                  `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[MeEpisodesItemsDataSourceModel] `tfsdk:"items"`
}

func (m *MeEpisodesDataSourceModel) toListParams(_ context.Context) (params spotted.MeEpisodeListParams, diags diag.Diagnostics) {
	params = spotted.MeEpisodeListParams{}

	if !m.Limit.IsNull() {
		params.Limit = param.NewOpt(m.Limit.ValueInt64())
	}
	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}
	if !m.Offset.IsNull() {
		params.Offset = param.NewOpt(m.Offset.ValueInt64())
	}

	return
}

type MeEpisodesItemsDataSourceModel struct {
	AddedAt timetypes.RFC3339                                          `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	Episode customfield.NestedObject[MeEpisodesEpisodeDataSourceModel] `tfsdk:"episode" json:"episode,computed"`
}

type MeEpisodesEpisodeDataSourceModel struct {
	ID                   types.String                                                           `tfsdk:"id" json:"id,computed"`
	AudioPreviewURL      types.String                                                           `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                           `tfsdk:"description" json:"description,computed"`
	DurationMs           types.Int64                                                            `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                             `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs         customfield.NestedObject[MeEpisodesEpisodeExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                           `tfsdk:"href" json:"href,computed"`
	HTMLDescription      types.String                                                           `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[MeEpisodesEpisodeImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                             `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	IsPlayable           types.Bool                                                             `tfsdk:"is_playable" json:"is_playable,computed"`
	Languages            customfield.List[types.String]                                         `tfsdk:"languages" json:"languages,computed"`
	Name                 types.String                                                           `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                           `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                           `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[MeEpisodesEpisodeShowDataSourceModel]         `tfsdk:"show" json:"show,computed"`
	Type                 types.String                                                           `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                           `tfsdk:"uri" json:"uri,computed"`
	Language             types.String                                                           `tfsdk:"language" json:"language,computed"`
	Restrictions         customfield.NestedObject[MeEpisodesEpisodeRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	ResumePoint          customfield.NestedObject[MeEpisodesEpisodeResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type MeEpisodesEpisodeExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeEpisodesEpisodeImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MeEpisodesEpisodeShowDataSourceModel struct {
	ID                 types.String                                                                 `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                               `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[MeEpisodesEpisodeShowCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                   `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[MeEpisodesEpisodeShowExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                 `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                 `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[MeEpisodesEpisodeShowImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                   `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                               `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                 `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                 `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                  `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
}

type MeEpisodesEpisodeShowCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type MeEpisodesEpisodeShowExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeEpisodesEpisodeShowImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MeEpisodesEpisodeRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type MeEpisodesEpisodeResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
