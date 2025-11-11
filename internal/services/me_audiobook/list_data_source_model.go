// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_audiobook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type MeAudiobooksItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MeAudiobooksItemsDataSourceModel] `json:"items,computed"`
}

type MeAudiobooksDataSourceModel struct {
	Limit    types.Int64                                                    `tfsdk:"limit" query:"limit,computed_optional"`
	Offset   types.Int64                                                    `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems types.Int64                                                    `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[MeAudiobooksItemsDataSourceModel] `tfsdk:"items"`
}

func (m *MeAudiobooksDataSourceModel) toListParams(_ context.Context) (params spotted.MeAudiobookListParams, diags diag.Diagnostics) {
	params = spotted.MeAudiobookListParams{}

	if !m.Limit.IsNull() {
		params.Limit = param.NewOpt(m.Limit.ValueInt64())
	}
	if !m.Offset.IsNull() {
		params.Offset = param.NewOpt(m.Offset.ValueInt64())
	}

	return
}

type MeAudiobooksItemsDataSourceModel struct {
	AddedAt   timetypes.RFC3339                                              `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	Audiobook customfield.NestedObject[MeAudiobooksAudiobookDataSourceModel] `tfsdk:"audiobook" json:"audiobook,computed"`
}

type MeAudiobooksAudiobookDataSourceModel struct {
	ID               types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Authors          customfield.NestedObjectList[MeAudiobooksAudiobookAuthorsDataSourceModel]    `tfsdk:"authors" json:"authors,computed"`
	AvailableMarkets customfield.List[types.String]                                               `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights       customfield.NestedObjectList[MeAudiobooksAudiobookCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description      types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Explicit         types.Bool                                                                   `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs     customfield.NestedObject[MeAudiobooksAudiobookExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href             types.String                                                                 `tfsdk:"href" json:"href,computed"`
	HTMLDescription  types.String                                                                 `tfsdk:"html_description" json:"html_description,computed"`
	Images           customfield.NestedObjectList[MeAudiobooksAudiobookImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	Languages        customfield.List[types.String]                                               `tfsdk:"languages" json:"languages,computed"`
	MediaType        types.String                                                                 `tfsdk:"media_type" json:"media_type,computed"`
	Name             types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Narrators        customfield.NestedObjectList[MeAudiobooksAudiobookNarratorsDataSourceModel]  `tfsdk:"narrators" json:"narrators,computed"`
	Publisher        types.String                                                                 `tfsdk:"publisher" json:"publisher,computed"`
	TotalChapters    types.Int64                                                                  `tfsdk:"total_chapters" json:"total_chapters,computed"`
	Type             types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
	Edition          types.String                                                                 `tfsdk:"edition" json:"edition,computed"`
	Chapters         customfield.NestedObject[MeAudiobooksAudiobookChaptersDataSourceModel]       `tfsdk:"chapters" json:"chapters,computed"`
}

type MeAudiobooksAudiobookAuthorsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type MeAudiobooksAudiobookCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type MeAudiobooksAudiobookExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeAudiobooksAudiobookImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MeAudiobooksAudiobookNarratorsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type MeAudiobooksAudiobookChaptersDataSourceModel struct {
	Href     types.String                                                                    `tfsdk:"href" json:"href,computed"`
	Limit    types.Int64                                                                     `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                                    `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                                     `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                                    `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                                     `tfsdk:"total" json:"total,computed"`
	Items    customfield.NestedObjectList[MeAudiobooksAudiobookChaptersItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
}

type MeAudiobooksAudiobookChaptersItemsDataSourceModel struct {
	ID                   types.String                                                                            `tfsdk:"id" json:"id,computed"`
	AudioPreviewURL      types.String                                                                            `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	ChapterNumber        types.Int64                                                                             `tfsdk:"chapter_number" json:"chapter_number,computed"`
	Description          types.String                                                                            `tfsdk:"description" json:"description,computed"`
	DurationMs           types.Int64                                                                             `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                                              `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs         customfield.NestedObject[MeAudiobooksAudiobookChaptersItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                            `tfsdk:"href" json:"href,computed"`
	HTMLDescription      types.String                                                                            `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[MeAudiobooksAudiobookChaptersItemsImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsPlayable           types.Bool                                                                              `tfsdk:"is_playable" json:"is_playable,computed"`
	Languages            customfield.List[types.String]                                                          `tfsdk:"languages" json:"languages,computed"`
	Name                 types.String                                                                            `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                            `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                            `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Type                 types.String                                                                            `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                            `tfsdk:"uri" json:"uri,computed"`
	AvailableMarkets     customfield.List[types.String]                                                          `tfsdk:"available_markets" json:"available_markets,computed"`
	Restrictions         customfield.NestedObject[MeAudiobooksAudiobookChaptersItemsRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	ResumePoint          customfield.NestedObject[MeAudiobooksAudiobookChaptersItemsResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type MeAudiobooksAudiobookChaptersItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeAudiobooksAudiobookChaptersItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MeAudiobooksAudiobookChaptersItemsRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type MeAudiobooksAudiobookChaptersItemsResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
