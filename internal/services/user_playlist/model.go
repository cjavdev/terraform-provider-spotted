// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_playlist

import (
	"github.com/cjavdev/terraform-provider-spotted/internal/apijson"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserPlaylistModel struct {
	ID                                   types.String                                            `tfsdk:"id" json:"id,computed"`
	UserID                               types.String                                            `tfsdk:"user_id" path:"user_id,required"`
	Name                                 types.String                                            `tfsdk:"name" json:"name,required"`
	Collaborative                        types.Bool                                              `tfsdk:"collaborative" json:"collaborative,optional"`
	ComponentsSchemasPropertiesPublished types.Bool                                              `tfsdk:"components_schemas_properties_published" json:"$.components.schemas.*.properties.published,optional"`
	Description                          types.String                                            `tfsdk:"description" json:"description,optional"`
	Href                                 types.String                                            `tfsdk:"href" json:"href,computed"`
	SnapshotID                           types.String                                            `tfsdk:"snapshot_id" json:"snapshot_id,computed"`
	Type                                 types.String                                            `tfsdk:"type" json:"type,computed"`
	Uri                                  types.String                                            `tfsdk:"uri" json:"uri,computed"`
	ExternalURLs                         customfield.NestedObject[UserPlaylistExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Followers                            customfield.NestedObject[UserPlaylistFollowersModel]    `tfsdk:"followers" json:"followers,computed"`
	Images                               customfield.NestedObjectList[UserPlaylistImagesModel]   `tfsdk:"images" json:"images,computed"`
	Owner                                customfield.NestedObject[UserPlaylistOwnerModel]        `tfsdk:"owner" json:"owner,computed"`
	Tracks                               customfield.NestedObject[UserPlaylistTracksModel]       `tfsdk:"tracks" json:"tracks,computed"`
}

func (m UserPlaylistModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserPlaylistModel) MarshalJSONForUpdate(state UserPlaylistModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type UserPlaylistExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistFollowersModel struct {
	Href  types.String `tfsdk:"href" json:"href,computed"`
	Total types.Int64  `tfsdk:"total" json:"total,computed"`
}

type UserPlaylistImagesModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type UserPlaylistOwnerModel struct {
	ID           types.String                                                 `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistOwnerExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                 `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                 `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                 `tfsdk:"uri" json:"uri,computed"`
	DisplayName  types.String                                                 `tfsdk:"display_name" json:"display_name,computed"`
}

type UserPlaylistOwnerExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksModel struct {
	Href     types.String                                               `tfsdk:"href" json:"href,computed"`
	Limit    types.Int64                                                `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                               `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                               `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                `tfsdk:"total" json:"total,computed"`
	Items    customfield.NestedObjectList[UserPlaylistTracksItemsModel] `tfsdk:"items" json:"items,computed"`
}

type UserPlaylistTracksItemsModel struct {
	AddedAt timetypes.RFC3339                                             `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	AddedBy customfield.NestedObject[UserPlaylistTracksItemsAddedByModel] `tfsdk:"added_by" json:"added_by,computed"`
	IsLocal types.Bool                                                    `tfsdk:"is_local" json:"is_local,computed"`
	Track   customfield.NestedObject[UserPlaylistTracksItemsTrackModel]   `tfsdk:"track" json:"track,computed"`
}

type UserPlaylistTracksItemsAddedByModel struct {
	ID           types.String                                                              `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsAddedByExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                              `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                              `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                              `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsAddedByExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackModel struct {
	ID                   types.String                                                            `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[UserPlaylistTracksItemsTrackAlbumModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[UserPlaylistTracksItemsTrackArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                          `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                             `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                             `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                              `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[UserPlaylistTracksItemsTrackExternalIDsModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[UserPlaylistTracksItemsTrackExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                            `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                              `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                              `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[UserPlaylistTracksItemsTrackLinkedFromModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                            `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                             `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                            `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistTracksItemsTrackRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                             `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                            `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                            `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                            `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                            `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                            `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[UserPlaylistTracksItemsTrackImagesModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                              `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                          `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                            `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                            `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[UserPlaylistTracksItemsTrackShowModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                            `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[UserPlaylistTracksItemsTrackResumePointModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type UserPlaylistTracksItemsTrackAlbumModel struct {
	ID                   types.String                                                                 `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                 `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[UserPlaylistTracksItemsTrackAlbumArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                               `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[UserPlaylistTracksItemsTrackAlbumExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                 `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[UserPlaylistTracksItemsTrackAlbumImagesModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                 `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                 `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                 `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                  `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistTracksItemsTrackAlbumRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type UserPlaylistTracksItemsTrackAlbumArtistsModel struct {
	ID           types.String                                                                        `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsTrackAlbumArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                        `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                        `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                        `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                        `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsTrackAlbumArtistsExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackAlbumExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackAlbumImagesModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type UserPlaylistTracksItemsTrackAlbumRestrictionsModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistTracksItemsTrackArtistsModel struct {
	ID           types.String                                                                   `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsTrackArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                   `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                   `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                   `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsTrackArtistsExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackExternalIDsModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type UserPlaylistTracksItemsTrackExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackLinkedFromModel struct {
	ID           types.String                                                                      `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsTrackLinkedFromExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                      `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                      `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                      `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsTrackLinkedFromExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackRestrictionsModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistTracksItemsTrackImagesModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type UserPlaylistTracksItemsTrackShowModel struct {
	ID                 types.String                                                                  `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                                `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[UserPlaylistTracksItemsTrackShowCopyrightsModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                  `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                    `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[UserPlaylistTracksItemsTrackShowExternalURLsModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                  `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                  `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[UserPlaylistTracksItemsTrackShowImagesModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                    `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                                `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                  `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                  `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                   `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsTrackShowCopyrightsModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type UserPlaylistTracksItemsTrackShowExternalURLsModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackShowImagesModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type UserPlaylistTracksItemsTrackResumePointModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
