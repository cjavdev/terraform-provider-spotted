// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_track

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type MeTracksItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MeTracksItemsDataSourceModel] `json:"items,computed"`
}

type MeTracksDataSourceModel struct {
	Market   types.String                                               `tfsdk:"market" query:"market,optional"`
	Limit    types.Int64                                                `tfsdk:"limit" query:"limit,computed_optional"`
	Offset   types.Int64                                                `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[MeTracksItemsDataSourceModel] `tfsdk:"items"`
}

func (m *MeTracksDataSourceModel) toListParams(_ context.Context) (params spotted.MeTrackListParams, diags diag.Diagnostics) {
	params = spotted.MeTrackListParams{}

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

type MeTracksItemsDataSourceModel struct {
	AddedAt timetypes.RFC3339                                      `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	Track   customfield.NestedObject[MeTracksTrackDataSourceModel] `tfsdk:"track" json:"track,computed"`
}

type MeTracksTrackDataSourceModel struct {
	ID               types.String                                                       `tfsdk:"id" json:"id,computed"`
	Album            customfield.NestedObject[MeTracksTrackAlbumDataSourceModel]        `tfsdk:"album" json:"album,computed"`
	Artists          customfield.NestedObjectList[MeTracksTrackArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets customfield.List[types.String]                                     `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber       types.Int64                                                        `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs       types.Int64                                                        `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit         types.Bool                                                         `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs      customfield.NestedObject[MeTracksTrackExternalIDsDataSourceModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs     customfield.NestedObject[MeTracksTrackExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href             types.String                                                       `tfsdk:"href" json:"href,computed"`
	IsLocal          types.Bool                                                         `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable       types.Bool                                                         `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom       customfield.NestedObject[MeTracksTrackLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name             types.String                                                       `tfsdk:"name" json:"name,computed"`
	Popularity       types.Int64                                                        `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL       types.String                                                       `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions     customfield.NestedObject[MeTracksTrackRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber      types.Int64                                                        `tfsdk:"track_number" json:"track_number,computed"`
	Type             types.String                                                       `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                       `tfsdk:"uri" json:"uri,computed"`
}

type MeTracksTrackAlbumDataSourceModel struct {
	ID                   types.String                                                            `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                            `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[MeTracksTrackAlbumArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                          `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[MeTracksTrackAlbumExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                            `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[MeTracksTrackAlbumImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                            `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                            `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                            `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                             `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                            `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                            `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[MeTracksTrackAlbumRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type MeTracksTrackAlbumArtistsDataSourceModel struct {
	ID           types.String                                                                   `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MeTracksTrackAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                   `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                   `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                   `tfsdk:"uri" json:"uri,computed"`
}

type MeTracksTrackAlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeTracksTrackAlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeTracksTrackAlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MeTracksTrackAlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type MeTracksTrackArtistsDataSourceModel struct {
	ID           types.String                                                              `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MeTracksTrackArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                              `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                              `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                              `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                              `tfsdk:"uri" json:"uri,computed"`
}

type MeTracksTrackArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeTracksTrackExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type MeTracksTrackExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeTracksTrackLinkedFromDataSourceModel struct {
	ID           types.String                                                                 `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MeTracksTrackLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                 `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
}

type MeTracksTrackLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeTracksTrackRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}
