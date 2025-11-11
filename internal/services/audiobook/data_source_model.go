// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audiobook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type AudiobookDataSourceModel struct {
	ID               types.String                                                     `tfsdk:"id" path:"id,required"`
	Market           types.String                                                     `tfsdk:"market" query:"market,optional"`
	Description      types.String                                                     `tfsdk:"description" json:"description,computed"`
	Edition          types.String                                                     `tfsdk:"edition" json:"edition,computed"`
	Explicit         types.Bool                                                       `tfsdk:"explicit" json:"explicit,computed"`
	Href             types.String                                                     `tfsdk:"href" json:"href,computed"`
	HTMLDescription  types.String                                                     `tfsdk:"html_description" json:"html_description,computed"`
	MediaType        types.String                                                     `tfsdk:"media_type" json:"media_type,computed"`
	Name             types.String                                                     `tfsdk:"name" json:"name,computed"`
	Publisher        types.String                                                     `tfsdk:"publisher" json:"publisher,computed"`
	TotalChapters    types.Int64                                                      `tfsdk:"total_chapters" json:"total_chapters,computed"`
	Type             types.String                                                     `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                     `tfsdk:"uri" json:"uri,computed"`
	AvailableMarkets customfield.List[types.String]                                   `tfsdk:"available_markets" json:"available_markets,computed"`
	Languages        customfield.List[types.String]                                   `tfsdk:"languages" json:"languages,computed"`
	Authors          customfield.NestedObjectList[AudiobookAuthorsDataSourceModel]    `tfsdk:"authors" json:"authors,computed"`
	Chapters         customfield.NestedObject[AudiobookChaptersDataSourceModel]       `tfsdk:"chapters" json:"chapters,computed"`
	Copyrights       customfield.NestedObjectList[AudiobookCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	ExternalURLs     customfield.NestedObject[AudiobookExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Images           customfield.NestedObjectList[AudiobookImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	Narrators        customfield.NestedObjectList[AudiobookNarratorsDataSourceModel]  `tfsdk:"narrators" json:"narrators,computed"`
}

func (m *AudiobookDataSourceModel) toReadParams(_ context.Context) (params spotted.AudiobookGetParams, diags diag.Diagnostics) {
	params = spotted.AudiobookGetParams{}

	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}

	return
}

type AudiobookAuthorsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type AudiobookChaptersDataSourceModel struct {
	Href     types.String                                                        `tfsdk:"href" json:"href,computed"`
	Limit    types.Int64                                                         `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                        `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                         `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                        `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                         `tfsdk:"total" json:"total,computed"`
	Items    customfield.NestedObjectList[AudiobookChaptersItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
}

type AudiobookChaptersItemsDataSourceModel struct {
	ID                   types.String                                                                `tfsdk:"id" json:"id,computed"`
	AudioPreviewURL      types.String                                                                `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	ChapterNumber        types.Int64                                                                 `tfsdk:"chapter_number" json:"chapter_number,computed"`
	Description          types.String                                                                `tfsdk:"description" json:"description,computed"`
	DurationMs           types.Int64                                                                 `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                                  `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs         customfield.NestedObject[AudiobookChaptersItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                `tfsdk:"href" json:"href,computed"`
	HTMLDescription      types.String                                                                `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[AudiobookChaptersItemsImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsPlayable           types.Bool                                                                  `tfsdk:"is_playable" json:"is_playable,computed"`
	Languages            customfield.List[types.String]                                              `tfsdk:"languages" json:"languages,computed"`
	Name                 types.String                                                                `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Type                 types.String                                                                `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                `tfsdk:"uri" json:"uri,computed"`
	AvailableMarkets     customfield.List[types.String]                                              `tfsdk:"available_markets" json:"available_markets,computed"`
	Restrictions         customfield.NestedObject[AudiobookChaptersItemsRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	ResumePoint          customfield.NestedObject[AudiobookChaptersItemsResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type AudiobookChaptersItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type AudiobookChaptersItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type AudiobookChaptersItemsRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type AudiobookChaptersItemsResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}

type AudiobookCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type AudiobookExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type AudiobookImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type AudiobookNarratorsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}
