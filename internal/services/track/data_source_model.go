// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package track

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type TrackDataSourceModel struct {
	ID               types.String                                               `tfsdk:"id" path:"id,required"`
	Market           types.String                                               `tfsdk:"market" query:"market,optional"`
	DiscNumber       types.Int64                                                `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs       types.Int64                                                `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit         types.Bool                                                 `tfsdk:"explicit" json:"explicit,computed"`
	Href             types.String                                               `tfsdk:"href" json:"href,computed"`
	IsLocal          types.Bool                                                 `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable       types.Bool                                                 `tfsdk:"is_playable" json:"is_playable,computed"`
	Name             types.String                                               `tfsdk:"name" json:"name,computed"`
	Popularity       types.Int64                                                `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL       types.String                                               `tfsdk:"preview_url" json:"preview_url,computed"`
	TrackNumber      types.Int64                                                `tfsdk:"track_number" json:"track_number,computed"`
	Type             types.String                                               `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                               `tfsdk:"uri" json:"uri,computed"`
	AvailableMarkets customfield.List[types.String]                             `tfsdk:"available_markets" json:"available_markets,computed"`
	Album            customfield.NestedObject[TrackAlbumDataSourceModel]        `tfsdk:"album" json:"album,computed"`
	Artists          customfield.NestedObjectList[TrackArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	ExternalIDs      customfield.NestedObject[TrackExternalIDsDataSourceModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs     customfield.NestedObject[TrackExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	LinkedFrom       customfield.NestedObject[TrackLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Restrictions     customfield.NestedObject[TrackRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

func (m *TrackDataSourceModel) toReadParams(_ context.Context) (params spotted.TrackGetParams, diags diag.Diagnostics) {
	params = spotted.TrackGetParams{}

	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}

	return
}

type TrackAlbumDataSourceModel struct {
	ID                   types.String                                                    `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                    `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[TrackAlbumArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                  `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[TrackAlbumExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                    `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[TrackAlbumImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                    `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                    `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                    `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                     `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                    `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                    `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[TrackAlbumRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type TrackAlbumArtistsDataSourceModel struct {
	ID           types.String                                                           `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[TrackAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                           `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                           `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                           `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                           `tfsdk:"uri" json:"uri,computed"`
}

type TrackAlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type TrackAlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type TrackAlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type TrackAlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type TrackArtistsDataSourceModel struct {
	ID           types.String                                                      `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[TrackArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                      `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                      `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                      `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                      `tfsdk:"uri" json:"uri,computed"`
}

type TrackArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type TrackExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type TrackExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type TrackLinkedFromDataSourceModel struct {
	ID           types.String                                                         `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[TrackLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                         `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                         `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                         `tfsdk:"uri" json:"uri,computed"`
}

type TrackLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type TrackRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}
