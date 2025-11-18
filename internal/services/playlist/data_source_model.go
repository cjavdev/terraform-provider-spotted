// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist

import (
	"context"

	"github.com/cjavdev/spotted-go"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type PlaylistDataSourceModel struct {
	PlaylistID      types.String                                                  `tfsdk:"playlist_id" path:"playlist_id,required"`
	AdditionalTypes types.String                                                  `tfsdk:"additional_types" query:"additional_types,optional"`
	Fields          types.String                                                  `tfsdk:"fields" query:"fields,optional"`
	Market          types.String                                                  `tfsdk:"market" query:"market,optional"`
	Collaborative   types.Bool                                                    `tfsdk:"collaborative" json:"collaborative,computed"`
	Description     types.String                                                  `tfsdk:"description" json:"description,computed"`
	Href            types.String                                                  `tfsdk:"href" json:"href,computed"`
	ID              types.String                                                  `tfsdk:"id" json:"id,computed"`
	Name            types.String                                                  `tfsdk:"name" json:"name,computed"`
	Public          types.Bool                                                    `tfsdk:"public" json:"public,computed"`
	SnapshotID      types.String                                                  `tfsdk:"snapshot_id" json:"snapshot_id,computed"`
	Type            types.String                                                  `tfsdk:"type" json:"type,computed"`
	Uri             types.String                                                  `tfsdk:"uri" json:"uri,computed"`
	ExternalURLs    customfield.NestedObject[PlaylistExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Followers       customfield.NestedObject[PlaylistFollowersDataSourceModel]    `tfsdk:"followers" json:"followers,computed"`
	Images          customfield.NestedObjectList[PlaylistImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Owner           customfield.NestedObject[PlaylistOwnerDataSourceModel]        `tfsdk:"owner" json:"owner,computed"`
	Tracks          customfield.NestedObject[PlaylistTracksDataSourceModel]       `tfsdk:"tracks" json:"tracks,computed"`
}

func (m *PlaylistDataSourceModel) toReadParams(_ context.Context) (params spotted.PlaylistGetParams, diags diag.Diagnostics) {
	params = spotted.PlaylistGetParams{}

	if !m.AdditionalTypes.IsNull() {
		params.AdditionalTypes = param.NewOpt(m.AdditionalTypes.ValueString())
	}
	if !m.Fields.IsNull() {
		params.Fields = param.NewOpt(m.Fields.ValueString())
	}
	if !m.Market.IsNull() {
		params.Market = param.NewOpt(m.Market.ValueString())
	}

	return
}

type PlaylistExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistFollowersDataSourceModel struct {
	Href  types.String `tfsdk:"href" json:"href,computed"`
	Total types.Int64  `tfsdk:"total" json:"total,computed"`
}

type PlaylistImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistOwnerDataSourceModel struct {
	ID           types.String                                                       `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistOwnerExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                       `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                       `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                       `tfsdk:"uri" json:"uri,computed"`
	DisplayName  types.String                                                       `tfsdk:"display_name" json:"display_name,computed"`
}

type PlaylistOwnerExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksDataSourceModel struct {
	Href     types.String                                                     `tfsdk:"href" json:"href,computed"`
	Limit    types.Int64                                                      `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                     `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                      `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                     `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                      `tfsdk:"total" json:"total,computed"`
	Items    customfield.NestedObjectList[PlaylistTracksItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
}

type PlaylistTracksItemsDataSourceModel struct {
	AddedAt timetypes.RFC3339                                                   `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	AddedBy customfield.NestedObject[PlaylistTracksItemsAddedByDataSourceModel] `tfsdk:"added_by" json:"added_by,computed"`
	IsLocal types.Bool                                                          `tfsdk:"is_local" json:"is_local,computed"`
	Track   customfield.NestedObject[PlaylistTracksItemsTrackDataSourceModel]   `tfsdk:"track" json:"track,computed"`
}

type PlaylistTracksItemsAddedByDataSourceModel struct {
	ID           types.String                                                                    `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksItemsAddedByExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                    `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                    `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                    `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsAddedByExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackDataSourceModel struct {
	ID                   types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[PlaylistTracksItemsTrackAlbumDataSourceModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[PlaylistTracksItemsTrackArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                                `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                                   `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                                   `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                                    `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[PlaylistTracksItemsTrackExternalIDsDataSourceModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[PlaylistTracksItemsTrackExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                  `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                                    `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                                    `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[PlaylistTracksItemsTrackLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                                   `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                                  `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions         customfield.NestedObject[PlaylistTracksItemsTrackRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                                   `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                                  `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                                  `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                                  `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[PlaylistTracksItemsTrackImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                                    `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                                `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                                  `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                  `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[PlaylistTracksItemsTrackShowDataSourceModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                                  `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[PlaylistTracksItemsTrackResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type PlaylistTracksItemsTrackAlbumDataSourceModel struct {
	ID                   types.String                                                                       `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                       `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[PlaylistTracksItemsTrackAlbumArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                                     `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[PlaylistTracksItemsTrackAlbumExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                       `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[PlaylistTracksItemsTrackAlbumImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                       `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                       `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                       `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                        `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                       `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                       `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[PlaylistTracksItemsTrackAlbumRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type PlaylistTracksItemsTrackAlbumArtistsDataSourceModel struct {
	ID           types.String                                                                              `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksItemsTrackAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                              `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                              `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                              `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                              `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsTrackAlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackAlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackAlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksItemsTrackAlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type PlaylistTracksItemsTrackArtistsDataSourceModel struct {
	ID           types.String                                                                         `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksItemsTrackArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                         `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                         `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                         `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                         `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsTrackArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type PlaylistTracksItemsTrackExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackLinkedFromDataSourceModel struct {
	ID           types.String                                                                            `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksItemsTrackLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                            `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                            `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                            `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsTrackLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type PlaylistTracksItemsTrackImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksItemsTrackShowDataSourceModel struct {
	ID                 types.String                                                                        `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                                      `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[PlaylistTracksItemsTrackShowCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                        `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                          `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[PlaylistTracksItemsTrackShowExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                        `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                        `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[PlaylistTracksItemsTrackShowImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                          `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                                      `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                        `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                        `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                        `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                         `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                        `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                        `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsTrackShowCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type PlaylistTracksItemsTrackShowExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackShowImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksItemsTrackResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
