// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package album

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type AlbumDataSourceModel struct {
	ID                   types.String                                                 `tfsdk:"id" path:"id,required"`
	Market               types.String                                                 `tfsdk:"market" query:"market,optional"`
	AlbumType            types.String                                                 `tfsdk:"album_type" json:"album_type,computed"`
	Href                 types.String                                                 `tfsdk:"href" json:"href,computed"`
	Label                types.String                                                 `tfsdk:"label" json:"label,computed"`
	Name                 types.String                                                 `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                  `tfsdk:"popularity" json:"popularity,computed"`
	ReleaseDate          types.String                                                 `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                 `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                  `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                 `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                 `tfsdk:"uri" json:"uri,computed"`
	AvailableMarkets     customfield.List[types.String]                               `tfsdk:"available_markets" json:"available_markets,computed"`
	Genres               customfield.List[types.String]                               `tfsdk:"genres" json:"genres,computed"`
	Artists              customfield.NestedObjectList[AlbumArtistsDataSourceModel]    `tfsdk:"artists" json:"artists,computed"`
	Copyrights           customfield.NestedObjectList[AlbumCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	ExternalIDs          customfield.NestedObject[AlbumExternalIDsDataSourceModel]    `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[AlbumExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Images               customfield.NestedObjectList[AlbumImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	Restrictions         customfield.NestedObject[AlbumRestrictionsDataSourceModel]   `tfsdk:"restrictions" json:"restrictions,computed"`
	Tracks               customfield.NestedObject[AlbumTracksDataSourceModel]         `tfsdk:"tracks" json:"tracks,computed"`
}

func (m *AlbumDataSourceModel) toReadParams(_ context.Context) (params spotted.AlbumGetParams, diags diag.Diagnostics) {
	params = spotted.AlbumGetParams{}

	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}

	return
}

type AlbumArtistsDataSourceModel struct {
	ID           types.String                                                      `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[AlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                      `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                      `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                      `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                      `tfsdk:"uri" json:"uri,computed"`
}

type AlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type AlbumCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type AlbumExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type AlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type AlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type AlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type AlbumTracksDataSourceModel struct {
	Href     types.String                                                  `tfsdk:"href" json:"href,computed"`
	Limit    types.Int64                                                   `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                  `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                   `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                  `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                   `tfsdk:"total" json:"total,computed"`
	Items    customfield.NestedObjectList[AlbumTracksItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
}

type AlbumTracksItemsDataSourceModel struct {
	ID               types.String                                                          `tfsdk:"id" json:"id,computed"`
	Artists          customfield.NestedObjectList[AlbumTracksItemsArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets customfield.List[types.String]                                        `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber       types.Int64                                                           `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs       types.Int64                                                           `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit         types.Bool                                                            `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs     customfield.NestedObject[AlbumTracksItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href             types.String                                                          `tfsdk:"href" json:"href,computed"`
	IsLocal          types.Bool                                                            `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable       types.Bool                                                            `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom       customfield.NestedObject[AlbumTracksItemsLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name             types.String                                                          `tfsdk:"name" json:"name,computed"`
	PreviewURL       types.String                                                          `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions     customfield.NestedObject[AlbumTracksItemsRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber      types.Int64                                                           `tfsdk:"track_number" json:"track_number,computed"`
	Type             types.String                                                          `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                          `tfsdk:"uri" json:"uri,computed"`
}

type AlbumTracksItemsArtistsDataSourceModel struct {
	ID           types.String                                                                 `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[AlbumTracksItemsArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                 `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
}

type AlbumTracksItemsArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type AlbumTracksItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type AlbumTracksItemsLinkedFromDataSourceModel struct {
	ID           types.String                                                                    `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[AlbumTracksItemsLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                    `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                    `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                    `tfsdk:"uri" json:"uri,computed"`
}

type AlbumTracksItemsLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type AlbumTracksItemsRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}
