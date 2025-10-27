// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist_track

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type PlaylistTracksItemsListDataSourceEnvelope struct {
	Items customfield.NestedObjectList[PlaylistTracksItemsDataSourceModel] `json:"items,computed"`
}

type PlaylistTracksDataSourceModel struct {
	PlaylistID      types.String                                                     `tfsdk:"playlist_id" path:"playlist_id,required"`
	AdditionalTypes types.String                                                     `tfsdk:"additional_types" query:"additional_types,optional"`
	Fields          types.String                                                     `tfsdk:"fields" query:"fields,optional"`
	Market          types.String                                                     `tfsdk:"market" query:"market,optional"`
	Limit           types.Int64                                                      `tfsdk:"limit" query:"limit,computed_optional"`
	Offset          types.Int64                                                      `tfsdk:"offset" query:"offset,computed_optional"`
	MaxItems        types.Int64                                                      `tfsdk:"max_items"`
	Items           customfield.NestedObjectList[PlaylistTracksItemsDataSourceModel] `tfsdk:"items"`
}

func (m *PlaylistTracksDataSourceModel) toListParams(_ context.Context) (params spotted.PlaylistTrackListParams, diags diag.Diagnostics) {
	params = spotted.PlaylistTrackListParams{}

	if !m.AdditionalTypes.IsNull() {
		params.AdditionalTypes = param.NewOpt(m.AdditionalTypes.ValueString())
	}
	if !m.Fields.IsNull() {
		params.Fields = param.NewOpt(m.Fields.ValueString())
	}
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

type PlaylistTracksItemsDataSourceModel struct {
	AddedAt timetypes.RFC3339                                              `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	AddedBy customfield.NestedObject[PlaylistTracksAddedByDataSourceModel] `tfsdk:"added_by" json:"added_by,computed"`
	IsLocal types.Bool                                                     `tfsdk:"is_local" json:"is_local,computed"`
	Track   customfield.NestedObject[PlaylistTracksTrackDataSourceModel]   `tfsdk:"track" json:"track,computed"`
}

type PlaylistTracksAddedByDataSourceModel struct {
	ID           types.String                                                               `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksAddedByExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                               `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                               `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                               `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksAddedByExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksTrackDataSourceModel struct {
	ID                   types.String                                                             `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[PlaylistTracksTrackAlbumDataSourceModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[PlaylistTracksTrackArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                           `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                              `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                              `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                               `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[PlaylistTracksTrackExternalIDsDataSourceModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[PlaylistTracksTrackExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                             `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                               `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                               `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[PlaylistTracksTrackLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                             `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                              `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                             `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions         customfield.NestedObject[PlaylistTracksTrackRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                              `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                             `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                             `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                             `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                             `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                             `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[PlaylistTracksTrackImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                               `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                           `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                             `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                             `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[PlaylistTracksTrackShowDataSourceModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                             `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[PlaylistTracksTrackResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type PlaylistTracksTrackAlbumDataSourceModel struct {
	ID                   types.String                                                                  `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                  `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[PlaylistTracksTrackAlbumArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                                `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[PlaylistTracksTrackAlbumExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                  `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[PlaylistTracksTrackAlbumImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                  `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                  `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                  `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                   `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[PlaylistTracksTrackAlbumRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type PlaylistTracksTrackAlbumArtistsDataSourceModel struct {
	ID           types.String                                                                         `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksTrackAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                         `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                         `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                         `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                         `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksTrackAlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksTrackAlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksTrackAlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksTrackAlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type PlaylistTracksTrackArtistsDataSourceModel struct {
	ID           types.String                                                                    `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksTrackArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                    `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                    `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                    `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                    `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksTrackArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksTrackExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type PlaylistTracksTrackExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksTrackLinkedFromDataSourceModel struct {
	ID           types.String                                                                       `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksTrackLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                       `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                       `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                       `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksTrackLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksTrackRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type PlaylistTracksTrackImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksTrackShowDataSourceModel struct {
	ID                 types.String                                                                   `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                                 `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[PlaylistTracksTrackShowCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                   `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                     `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[PlaylistTracksTrackShowExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                   `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                   `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[PlaylistTracksTrackShowImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                     `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                                 `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                   `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                   `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                    `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                   `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                   `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksTrackShowCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type PlaylistTracksTrackShowExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksTrackShowImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksTrackResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
