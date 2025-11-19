// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package episode

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EpisodeDataSourceModel struct {
	ID                   types.String                                                 `tfsdk:"id" path:"id,required"`
	Market               types.String                                                 `tfsdk:"market" query:"market,optional"`
	AudioPreviewURL      types.String                                                 `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                 `tfsdk:"description" json:"description,computed"`
	DurationMs           types.Int64                                                  `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                   `tfsdk:"explicit" json:"explicit,computed"`
	Href                 types.String                                                 `tfsdk:"href" json:"href,computed"`
	HTMLDescription      types.String                                                 `tfsdk:"html_description" json:"html_description,computed"`
	IsExternallyHosted   types.Bool                                                   `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	IsPlayable           types.Bool                                                   `tfsdk:"is_playable" json:"is_playable,computed"`
	Language             types.String                                                 `tfsdk:"language" json:"language,computed"`
	Name                 types.String                                                 `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                 `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                 `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Type                 types.String                                                 `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                 `tfsdk:"uri" json:"uri,computed"`
	Languages            customfield.List[types.String]                               `tfsdk:"languages" json:"languages,computed"`
	ExternalURLs         customfield.NestedObject[EpisodeExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Images               customfield.NestedObjectList[EpisodeImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Restrictions         customfield.NestedObject[EpisodeRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	ResumePoint          customfield.NestedObject[EpisodeResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
	Show                 customfield.NestedObject[EpisodeShowDataSourceModel]         `tfsdk:"show" json:"show,computed"`
}

func (m *EpisodeDataSourceModel) toReadParams(_ context.Context) (params spotted.EpisodeGetParams, diags diag.Diagnostics) {
	params = spotted.EpisodeGetParams{}

	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}

	return
}

type EpisodeExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type EpisodeImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type EpisodeRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type EpisodeResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}

type EpisodeShowDataSourceModel struct {
	ID                 types.String                                                       `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                     `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[EpisodeShowCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                       `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                         `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[EpisodeShowExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                       `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                       `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[EpisodeShowImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                         `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                     `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                       `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                       `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                       `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                        `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                       `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                       `tfsdk:"uri" json:"uri,computed"`
}

type EpisodeShowCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type EpisodeShowExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type EpisodeShowImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}
