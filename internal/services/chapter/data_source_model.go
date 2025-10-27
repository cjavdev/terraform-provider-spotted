// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chapter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type ChapterDataSourceModel struct {
	ID                   types.String                                                 `tfsdk:"id" path:"id,required"`
	Market               types.String                                                 `tfsdk:"market" query:"market,optional"`
	AudioPreviewURL      types.String                                                 `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	ChapterNumber        types.Int64                                                  `tfsdk:"chapter_number" json:"chapter_number,computed"`
	Description          types.String                                                 `tfsdk:"description" json:"description,computed"`
	DurationMs           types.Int64                                                  `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                   `tfsdk:"explicit" json:"explicit,computed"`
	Href                 types.String                                                 `tfsdk:"href" json:"href,computed"`
	HTMLDescription      types.String                                                 `tfsdk:"html_description" json:"html_description,computed"`
	IsPlayable           types.Bool                                                   `tfsdk:"is_playable" json:"is_playable,computed"`
	Name                 types.String                                                 `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                 `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                 `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Type                 types.String                                                 `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                 `tfsdk:"uri" json:"uri,computed"`
	AvailableMarkets     customfield.List[types.String]                               `tfsdk:"available_markets" json:"available_markets,computed"`
	Languages            customfield.List[types.String]                               `tfsdk:"languages" json:"languages,computed"`
	Audiobook            customfield.NestedObject[ChapterAudiobookDataSourceModel]    `tfsdk:"audiobook" json:"audiobook,computed"`
	ExternalURLs         customfield.NestedObject[ChapterExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Images               customfield.NestedObjectList[ChapterImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Restrictions         customfield.NestedObject[ChapterRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	ResumePoint          customfield.NestedObject[ChapterResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

func (m *ChapterDataSourceModel) toReadParams(_ context.Context) (params spotted.ChapterGetParams, diags diag.Diagnostics) {
	params = spotted.ChapterGetParams{}

	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}

	return
}

type ChapterAudiobookDataSourceModel struct {
	ID               types.String                                                            `tfsdk:"id" json:"id,computed"`
	Authors          customfield.NestedObjectList[ChapterAudiobookAuthorsDataSourceModel]    `tfsdk:"authors" json:"authors,computed"`
	AvailableMarkets customfield.List[types.String]                                          `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights       customfield.NestedObjectList[ChapterAudiobookCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description      types.String                                                            `tfsdk:"description" json:"description,computed"`
	Explicit         types.Bool                                                              `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs     customfield.NestedObject[ChapterAudiobookExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href             types.String                                                            `tfsdk:"href" json:"href,computed"`
	HTMLDescription  types.String                                                            `tfsdk:"html_description" json:"html_description,computed"`
	Images           customfield.NestedObjectList[ChapterAudiobookImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	Languages        customfield.List[types.String]                                          `tfsdk:"languages" json:"languages,computed"`
	MediaType        types.String                                                            `tfsdk:"media_type" json:"media_type,computed"`
	Name             types.String                                                            `tfsdk:"name" json:"name,computed"`
	Narrators        customfield.NestedObjectList[ChapterAudiobookNarratorsDataSourceModel]  `tfsdk:"narrators" json:"narrators,computed"`
	Publisher        types.String                                                            `tfsdk:"publisher" json:"publisher,computed"`
	TotalChapters    types.Int64                                                             `tfsdk:"total_chapters" json:"total_chapters,computed"`
	Type             types.String                                                            `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                            `tfsdk:"uri" json:"uri,computed"`
	Edition          types.String                                                            `tfsdk:"edition" json:"edition,computed"`
}

type ChapterAudiobookAuthorsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type ChapterAudiobookCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type ChapterAudiobookExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type ChapterAudiobookImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type ChapterAudiobookNarratorsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type ChapterExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type ChapterImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type ChapterRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type ChapterResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
