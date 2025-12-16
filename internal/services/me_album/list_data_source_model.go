// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_album

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MeAlbumsItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[MeAlbumsItemsDataSourceModel] `json:"items,computed"`
}

type MeAlbumsDataSourceModel struct {
	Market   types.String                                               `tfsdk:"market" query:"market,optional"`
	Limit    types.Int64                                                `tfsdk:"limit" query:"limit,computed_optional"`
	Offset   types.Int64                                                `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[MeAlbumsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *MeAlbumsDataSourceModel) toListParams(_ context.Context) (params spotted.MeAlbumListParams, diags diag.Diagnostics) {
	params = spotted.MeAlbumListParams{}

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

type MeAlbumsItemsDataSourceModel struct {
	AddedAt   timetypes.RFC3339                                      `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	Album     customfield.NestedObject[MeAlbumsAlbumDataSourceModel] `tfsdk:"album" json:"album,computed"`
	Published types.Bool                                             `tfsdk:"published" json:"published,computed"`
}

type MeAlbumsAlbumDataSourceModel struct {
	ID                   types.String                                                         `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                         `tfsdk:"album_type" json:"album_type,computed"`
	AvailableMarkets     customfield.List[types.String]                                       `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[MeAlbumsAlbumExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                         `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[MeAlbumsAlbumImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                         `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                         `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                         `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                          `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                         `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                         `tfsdk:"uri" json:"uri,computed"`
	Artists              customfield.NestedObjectList[MeAlbumsAlbumArtistsDataSourceModel]    `tfsdk:"artists" json:"artists,computed"`
	Copyrights           customfield.NestedObjectList[MeAlbumsAlbumCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	ExternalIDs          customfield.NestedObject[MeAlbumsAlbumExternalIDsDataSourceModel]    `tfsdk:"external_ids" json:"external_ids,computed"`
	Genres               customfield.List[types.String]                                       `tfsdk:"genres" json:"genres,computed"`
	Label                types.String                                                         `tfsdk:"label" json:"label,computed"`
	Popularity           types.Int64                                                          `tfsdk:"popularity" json:"popularity,computed"`
	Published            types.Bool                                                           `tfsdk:"published" json:"published,computed"`
	Restrictions         customfield.NestedObject[MeAlbumsAlbumRestrictionsDataSourceModel]   `tfsdk:"restrictions" json:"restrictions,computed"`
	Tracks               customfield.NestedObject[MeAlbumsAlbumTracksDataSourceModel]         `tfsdk:"tracks" json:"tracks,computed"`
}

type MeAlbumsAlbumExternalURLsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeAlbumsAlbumImagesDataSourceModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type MeAlbumsAlbumArtistsDataSourceModel struct {
	ID           types.String                                                              `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MeAlbumsAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                              `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                              `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                              `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                              `tfsdk:"uri" json:"uri,computed"`
}

type MeAlbumsAlbumArtistsExternalURLsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeAlbumsAlbumCopyrightsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Text      types.String `tfsdk:"text" json:"text,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}

type MeAlbumsAlbumExternalIDsDataSourceModel struct {
	Ean       types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc      types.String `tfsdk:"isrc" json:"isrc,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Upc       types.String `tfsdk:"upc" json:"upc,computed"`
}

type MeAlbumsAlbumRestrictionsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type MeAlbumsAlbumTracksDataSourceModel struct {
	Href      types.String                                                          `tfsdk:"href" json:"href,computed"`
	Limit     types.Int64                                                           `tfsdk:"limit" json:"limit,computed"`
	Next      types.String                                                          `tfsdk:"next" json:"next,computed"`
	Offset    types.Int64                                                           `tfsdk:"offset" json:"offset,computed"`
	Previous  types.String                                                          `tfsdk:"previous" json:"previous,computed"`
	Total     types.Int64                                                           `tfsdk:"total" json:"total,computed"`
	Items     customfield.NestedObjectList[MeAlbumsAlbumTracksItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Published types.Bool                                                            `tfsdk:"published" json:"published,computed"`
}

type MeAlbumsAlbumTracksItemsDataSourceModel struct {
	ID               types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Artists          customfield.NestedObjectList[MeAlbumsAlbumTracksItemsArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets customfield.List[types.String]                                                `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber       types.Int64                                                                   `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs       types.Int64                                                                   `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit         types.Bool                                                                    `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs     customfield.NestedObject[MeAlbumsAlbumTracksItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href             types.String                                                                  `tfsdk:"href" json:"href,computed"`
	IsLocal          types.Bool                                                                    `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable       types.Bool                                                                    `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom       customfield.NestedObject[MeAlbumsAlbumTracksItemsLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name             types.String                                                                  `tfsdk:"name" json:"name,computed"`
	PreviewURL       types.String                                                                  `tfsdk:"preview_url" json:"preview_url,computed"`
	Published        types.Bool                                                                    `tfsdk:"published" json:"published,computed"`
	Restrictions     customfield.NestedObject[MeAlbumsAlbumTracksItemsRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber      types.Int64                                                                   `tfsdk:"track_number" json:"track_number,computed"`
	Type             types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
}

type MeAlbumsAlbumTracksItemsArtistsDataSourceModel struct {
	ID           types.String                                                                         `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MeAlbumsAlbumTracksItemsArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                         `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                         `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                           `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                         `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                         `tfsdk:"uri" json:"uri,computed"`
}

type MeAlbumsAlbumTracksItemsArtistsExternalURLsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeAlbumsAlbumTracksItemsExternalURLsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeAlbumsAlbumTracksItemsLinkedFromDataSourceModel struct {
	ID           types.String                                                                            `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MeAlbumsAlbumTracksItemsLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                            `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                                              `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                            `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                            `tfsdk:"uri" json:"uri,computed"`
}

type MeAlbumsAlbumTracksItemsLinkedFromExternalURLsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MeAlbumsAlbumTracksItemsRestrictionsDataSourceModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}
