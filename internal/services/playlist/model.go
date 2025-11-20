// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package playlist

import (
	"github.com/cjavdev/terraform-provider-spotted/internal/apijson"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PlaylistModel struct {
	PlaylistID                           types.String                                        `tfsdk:"playlist_id" path:"playlist_id,required"`
	Collaborative                        types.Bool                                          `tfsdk:"collaborative" json:"collaborative,optional"`
	Description                          types.String                                        `tfsdk:"description" json:"description,optional"`
	Name                                 types.String                                        `tfsdk:"name" json:"name,optional"`
	Public                               types.Bool                                          `tfsdk:"public" json:"public,optional,no_refresh"`
	ComponentsSchemasPropertiesPublished types.Bool                                          `tfsdk:"components_schemas_properties_published" json:"$.components.schemas.*.properties.published,computed"`
	Href                                 types.String                                        `tfsdk:"href" json:"href,computed"`
	ID                                   types.String                                        `tfsdk:"id" json:"id,computed"`
	SnapshotID                           types.String                                        `tfsdk:"snapshot_id" json:"snapshot_id,computed"`
	Type                                 types.String                                        `tfsdk:"type" json:"type,computed"`
	Uri                                  types.String                                        `tfsdk:"uri" json:"uri,computed"`
	ExternalURLs                         customfield.NestedObject[PlaylistExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Followers                            customfield.NestedObject[PlaylistFollowersModel]    `tfsdk:"followers" json:"followers,computed"`
	Images                               customfield.NestedObjectList[PlaylistImagesModel]   `tfsdk:"images" json:"images,computed"`
	Owner                                customfield.NestedObject[PlaylistOwnerModel]        `tfsdk:"owner" json:"owner,computed"`
	Tracks                               customfield.NestedObject[PlaylistTracksModel]       `tfsdk:"tracks" json:"tracks,computed"`
}

func (m PlaylistModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m PlaylistModel) MarshalJSONForUpdate(state PlaylistModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type PlaylistExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistFollowersModel struct {
	Href  types.String `tfsdk:"href" json:"href,computed"`
	Total types.Int64  `tfsdk:"total" json:"total,computed"`
}

type PlaylistImagesModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistOwnerModel struct {
	ID           types.String                                             `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistOwnerExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                             `tfsdk:"href" json:"href,computed"`
	Type         types.String                                             `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                             `tfsdk:"uri" json:"uri,computed"`
	DisplayName  types.String                                             `tfsdk:"display_name" json:"display_name,computed"`
}

type PlaylistOwnerExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksModel struct {
	Href     types.String                                           `tfsdk:"href" json:"href,computed"`
	Limit    types.Int64                                            `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                           `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                            `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                           `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                            `tfsdk:"total" json:"total,computed"`
	Items    customfield.NestedObjectList[PlaylistTracksItemsModel] `tfsdk:"items" json:"items,computed"`
}

type PlaylistTracksItemsModel struct {
	AddedAt timetypes.RFC3339                                         `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	AddedBy customfield.NestedObject[PlaylistTracksItemsAddedByModel] `tfsdk:"added_by" json:"added_by,computed"`
	IsLocal types.Bool                                                `tfsdk:"is_local" json:"is_local,computed"`
	Track   customfield.NestedObject[PlaylistTracksItemsTrackModel]   `tfsdk:"track" json:"track,computed"`
}

type PlaylistTracksItemsAddedByModel struct {
	ID           types.String                                                          `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksItemsAddedByExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                          `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                          `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                          `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsAddedByExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackModel struct {
	ID                   types.String                                                        `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[PlaylistTracksItemsTrackAlbumModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[PlaylistTracksItemsTrackArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                      `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                         `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                         `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                          `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[PlaylistTracksItemsTrackExternalIDsModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[PlaylistTracksItemsTrackExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                        `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                          `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                          `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[PlaylistTracksItemsTrackLinkedFromModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                        `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                         `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                        `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions         customfield.NestedObject[PlaylistTracksItemsTrackRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                         `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                        `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                        `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                        `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                        `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                        `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[PlaylistTracksItemsTrackImagesModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                          `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                      `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                        `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                        `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[PlaylistTracksItemsTrackShowModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                        `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[PlaylistTracksItemsTrackResumePointModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type PlaylistTracksItemsTrackAlbumModel struct {
	ID                   types.String                                                             `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                             `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[PlaylistTracksItemsTrackAlbumArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                           `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[PlaylistTracksItemsTrackAlbumExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                             `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[PlaylistTracksItemsTrackAlbumImagesModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                             `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                             `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                             `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                              `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                             `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                             `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[PlaylistTracksItemsTrackAlbumRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type PlaylistTracksItemsTrackAlbumArtistsModel struct {
	ID           types.String                                                                    `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksItemsTrackAlbumArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                    `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                    `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                    `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                    `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsTrackAlbumArtistsExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackAlbumExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackAlbumImagesModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksItemsTrackAlbumRestrictionsModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type PlaylistTracksItemsTrackArtistsModel struct {
	ID           types.String                                                               `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksItemsTrackArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                               `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                               `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                               `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                               `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsTrackArtistsExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackExternalIDsModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type PlaylistTracksItemsTrackExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackLinkedFromModel struct {
	ID           types.String                                                                  `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[PlaylistTracksItemsTrackLinkedFromExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                  `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsTrackLinkedFromExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackRestrictionsModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type PlaylistTracksItemsTrackImagesModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksItemsTrackShowModel struct {
	ID                 types.String                                                              `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                            `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[PlaylistTracksItemsTrackShowCopyrightsModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                              `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[PlaylistTracksItemsTrackShowExternalURLsModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                              `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                              `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[PlaylistTracksItemsTrackShowImagesModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                            `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                              `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                              `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                              `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                               `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                              `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                              `tfsdk:"uri" json:"uri,computed"`
}

type PlaylistTracksItemsTrackShowCopyrightsModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type PlaylistTracksItemsTrackShowExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type PlaylistTracksItemsTrackShowImagesModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type PlaylistTracksItemsTrackResumePointModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
