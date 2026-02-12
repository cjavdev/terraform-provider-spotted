// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_playlist

import (
	"github.com/cjavdev/terraform-provider-spotted/internal/apijson"
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserPlaylistModel struct {
	ID            types.String                                            `tfsdk:"id" json:"id,computed"`
	UserID        types.String                                            `tfsdk:"user_id" path:"user_id,required"`
	Name          types.String                                            `tfsdk:"name" json:"name,required"`
	Collaborative types.Bool                                              `tfsdk:"collaborative" json:"collaborative,optional"`
	Description   types.String                                            `tfsdk:"description" json:"description,optional"`
	Published     types.Bool                                              `tfsdk:"published" json:"published,optional"`
	Href          types.String                                            `tfsdk:"href" json:"href,computed"`
	SnapshotID    types.String                                            `tfsdk:"snapshot_id" json:"snapshot_id,computed"`
	Type          types.String                                            `tfsdk:"type" json:"type,computed"`
	Uri           types.String                                            `tfsdk:"uri" json:"uri,computed"`
	ExternalURLs  customfield.NestedObject[UserPlaylistExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Followers     customfield.NestedObject[UserPlaylistFollowersModel]    `tfsdk:"followers" json:"followers,computed"`
	Images        customfield.NestedObjectList[UserPlaylistImagesModel]   `tfsdk:"images" json:"images,computed"`
	Items         customfield.NestedObject[UserPlaylistItemsModel]        `tfsdk:"items" json:"items,computed"`
	Owner         customfield.NestedObject[UserPlaylistOwnerModel]        `tfsdk:"owner" json:"owner,computed"`
	Tracks        customfield.NestedObject[UserPlaylistTracksModel]       `tfsdk:"tracks" json:"tracks,computed"`
}

func (m UserPlaylistModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserPlaylistModel) MarshalJSONForUpdate(state UserPlaylistModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type UserPlaylistExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistFollowersModel struct {
	Href      types.String `tfsdk:"href" json:"href,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Total     types.Int64  `tfsdk:"total" json:"total,computed"`
}

type UserPlaylistImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsModel struct {
	Href      types.String                                              `tfsdk:"href" json:"href,computed"`
	Limit     types.Int64                                               `tfsdk:"limit" json:"limit,computed"`
	Next      types.String                                              `tfsdk:"next" json:"next,computed"`
	Offset    types.Int64                                               `tfsdk:"offset" json:"offset,computed"`
	Previous  types.String                                              `tfsdk:"previous" json:"previous,computed"`
	Total     types.Int64                                               `tfsdk:"total" json:"total,computed"`
	Items     customfield.NestedObjectList[UserPlaylistItemsItemsModel] `tfsdk:"items" json:"items,computed"`
	Published types.Bool                                                `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsModel struct {
	AddedAt   timetypes.RFC3339                                            `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	AddedBy   customfield.NestedObject[UserPlaylistItemsItemsAddedByModel] `tfsdk:"added_by" json:"added_by,computed"`
	IsLocal   types.Bool                                                   `tfsdk:"is_local" json:"is_local,computed"`
	Item      customfield.NestedObject[UserPlaylistItemsItemsItemModel]    `tfsdk:"item" json:"item,computed"`
	Published types.Bool                                                   `tfsdk:"published" json:"published,computed"`
	Track     customfield.NestedObject[UserPlaylistItemsItemsTrackModel]   `tfsdk:"track" json:"track,computed"`
}

type UserPlaylistItemsItemsAddedByModel struct {
	ID           types.String                                                             `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistItemsItemsAddedByExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                             `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                               `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                             `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                             `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistItemsItemsAddedByExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsItemModel struct {
	ID                   types.String                                                          `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[UserPlaylistItemsItemsItemAlbumModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[UserPlaylistItemsItemsItemArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                        `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                           `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                           `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                            `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[UserPlaylistItemsItemsItemExternalIDsModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[UserPlaylistItemsItemsItemExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                          `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                            `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                            `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[UserPlaylistItemsItemsItemLinkedFromModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                          `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                           `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                          `tfsdk:"preview_url" json:"preview_url,computed"`
	Published            types.Bool                                                            `tfsdk:"published" json:"published,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistItemsItemsItemRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                           `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                          `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                          `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                          `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                          `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                          `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[UserPlaylistItemsItemsItemImagesModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                            `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                        `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                          `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                          `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[UserPlaylistItemsItemsItemShowModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                          `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[UserPlaylistItemsItemsItemResumePointModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type UserPlaylistItemsItemsItemAlbumModel struct {
	ID                   types.String                                                               `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                               `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[UserPlaylistItemsItemsItemAlbumArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                             `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[UserPlaylistItemsItemsItemAlbumExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                               `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[UserPlaylistItemsItemsItemAlbumImagesModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                               `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                               `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                               `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                               `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                               `tfsdk:"uri" json:"uri,computed"`
	Published            types.Bool                                                                 `tfsdk:"published" json:"published,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistItemsItemsItemAlbumRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type UserPlaylistItemsItemsItemAlbumArtistsModel struct {
	ID           types.String                                                                      `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistItemsItemsItemAlbumArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                      `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                      `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                        `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                      `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                      `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistItemsItemsItemAlbumArtistsExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsItemAlbumExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsItemAlbumImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsItemAlbumRestrictionsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistItemsItemsItemArtistsModel struct {
	ID           types.String                                                                 `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistItemsItemsItemArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                 `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                   `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistItemsItemsItemArtistsExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsItemExternalIDsModel struct {
	Ean       types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc      types.String `tfsdk:"isrc" json:"isrc,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Upc       types.String `tfsdk:"upc" json:"upc,computed"`
}

type UserPlaylistItemsItemsItemExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsItemLinkedFromModel struct {
	ID           types.String                                                                    `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistItemsItemsItemLinkedFromExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                    `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                                      `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                    `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                    `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistItemsItemsItemLinkedFromExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsItemRestrictionsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistItemsItemsItemImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsItemShowModel struct {
	ID                 types.String                                                                `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                              `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[UserPlaylistItemsItemsItemShowCopyrightsModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                  `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[UserPlaylistItemsItemsItemShowExternalURLsModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[UserPlaylistItemsItemsItemShowImagesModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                  `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                              `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                 `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                `tfsdk:"uri" json:"uri,computed"`
	Published          types.Bool                                                                  `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsItemShowCopyrightsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Text      types.String `tfsdk:"text" json:"text,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}

type UserPlaylistItemsItemsItemShowExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsItemShowImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsItemResumePointModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	Published        types.Bool  `tfsdk:"published" json:"published,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}

type UserPlaylistItemsItemsTrackModel struct {
	ID                   types.String                                                           `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[UserPlaylistItemsItemsTrackAlbumModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[UserPlaylistItemsItemsTrackArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                         `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                            `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                            `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                             `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[UserPlaylistItemsItemsTrackExternalIDsModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[UserPlaylistItemsItemsTrackExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                           `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                             `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                             `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[UserPlaylistItemsItemsTrackLinkedFromModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                           `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                            `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                           `tfsdk:"preview_url" json:"preview_url,computed"`
	Published            types.Bool                                                             `tfsdk:"published" json:"published,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistItemsItemsTrackRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                            `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                           `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                           `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                           `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                           `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                           `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[UserPlaylistItemsItemsTrackImagesModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                             `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                         `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                           `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                           `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[UserPlaylistItemsItemsTrackShowModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                           `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[UserPlaylistItemsItemsTrackResumePointModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type UserPlaylistItemsItemsTrackAlbumModel struct {
	ID                   types.String                                                                `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[UserPlaylistItemsItemsTrackAlbumArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                              `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[UserPlaylistItemsItemsTrackAlbumExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[UserPlaylistItemsItemsTrackAlbumImagesModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                 `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                `tfsdk:"uri" json:"uri,computed"`
	Published            types.Bool                                                                  `tfsdk:"published" json:"published,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistItemsItemsTrackAlbumRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type UserPlaylistItemsItemsTrackAlbumArtistsModel struct {
	ID           types.String                                                                       `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistItemsItemsTrackAlbumArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                       `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                       `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                         `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                       `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                       `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistItemsItemsTrackAlbumArtistsExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsTrackAlbumExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsTrackAlbumImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsTrackAlbumRestrictionsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistItemsItemsTrackArtistsModel struct {
	ID           types.String                                                                  `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistItemsItemsTrackArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                  `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                    `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistItemsItemsTrackArtistsExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsTrackExternalIDsModel struct {
	Ean       types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc      types.String `tfsdk:"isrc" json:"isrc,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Upc       types.String `tfsdk:"upc" json:"upc,computed"`
}

type UserPlaylistItemsItemsTrackExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsTrackLinkedFromModel struct {
	ID           types.String                                                                     `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistItemsItemsTrackLinkedFromExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                     `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                                       `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                     `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                     `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistItemsItemsTrackLinkedFromExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsTrackRestrictionsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistItemsItemsTrackImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsTrackShowModel struct {
	ID                 types.String                                                                 `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                               `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[UserPlaylistItemsItemsTrackShowCopyrightsModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                   `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[UserPlaylistItemsItemsTrackShowExternalURLsModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                 `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                 `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[UserPlaylistItemsItemsTrackShowImagesModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                   `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                               `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                 `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                 `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                  `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
	Published          types.Bool                                                                   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsTrackShowCopyrightsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Text      types.String `tfsdk:"text" json:"text,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}

type UserPlaylistItemsItemsTrackShowExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistItemsItemsTrackShowImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistItemsItemsTrackResumePointModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	Published        types.Bool  `tfsdk:"published" json:"published,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}

type UserPlaylistOwnerModel struct {
	ID           types.String                                                 `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistOwnerExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                 `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                   `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                 `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                 `tfsdk:"uri" json:"uri,computed"`
	DisplayName  types.String                                                 `tfsdk:"display_name" json:"display_name,computed"`
}

type UserPlaylistOwnerExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksModel struct {
	Href      types.String                                               `tfsdk:"href" json:"href,computed"`
	Limit     types.Int64                                                `tfsdk:"limit" json:"limit,computed"`
	Next      types.String                                               `tfsdk:"next" json:"next,computed"`
	Offset    types.Int64                                                `tfsdk:"offset" json:"offset,computed"`
	Previous  types.String                                               `tfsdk:"previous" json:"previous,computed"`
	Total     types.Int64                                                `tfsdk:"total" json:"total,computed"`
	Items     customfield.NestedObjectList[UserPlaylistTracksItemsModel] `tfsdk:"items" json:"items,computed"`
	Published types.Bool                                                 `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistTracksItemsModel struct {
	AddedAt   timetypes.RFC3339                                             `tfsdk:"added_at" json:"added_at,computed" format:"date-time"`
	AddedBy   customfield.NestedObject[UserPlaylistTracksItemsAddedByModel] `tfsdk:"added_by" json:"added_by,computed"`
	IsLocal   types.Bool                                                    `tfsdk:"is_local" json:"is_local,computed"`
	Item      customfield.NestedObject[UserPlaylistTracksItemsItemModel]    `tfsdk:"item" json:"item,computed"`
	Published types.Bool                                                    `tfsdk:"published" json:"published,computed"`
	Track     customfield.NestedObject[UserPlaylistTracksItemsTrackModel]   `tfsdk:"track" json:"track,computed"`
}

type UserPlaylistTracksItemsAddedByModel struct {
	ID           types.String                                                              `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsAddedByExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                              `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                                `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                              `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                              `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsAddedByExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsItemModel struct {
	ID                   types.String                                                           `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[UserPlaylistTracksItemsItemAlbumModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[UserPlaylistTracksItemsItemArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                         `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                            `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                            `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                             `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[UserPlaylistTracksItemsItemExternalIDsModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[UserPlaylistTracksItemsItemExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                           `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                             `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                             `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[UserPlaylistTracksItemsItemLinkedFromModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                           `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                            `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                           `tfsdk:"preview_url" json:"preview_url,computed"`
	Published            types.Bool                                                             `tfsdk:"published" json:"published,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistTracksItemsItemRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                            `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                           `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                           `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                           `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                           `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                           `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[UserPlaylistTracksItemsItemImagesModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                             `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                         `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                           `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                           `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[UserPlaylistTracksItemsItemShowModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                           `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[UserPlaylistTracksItemsItemResumePointModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type UserPlaylistTracksItemsItemAlbumModel struct {
	ID                   types.String                                                                `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[UserPlaylistTracksItemsItemAlbumArtistsModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                              `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[UserPlaylistTracksItemsItemAlbumExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[UserPlaylistTracksItemsItemAlbumImagesModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                 `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                `tfsdk:"uri" json:"uri,computed"`
	Published            types.Bool                                                                  `tfsdk:"published" json:"published,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistTracksItemsItemAlbumRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type UserPlaylistTracksItemsItemAlbumArtistsModel struct {
	ID           types.String                                                                       `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsItemAlbumArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                       `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                       `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                         `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                       `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                       `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsItemAlbumArtistsExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsItemAlbumExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsItemAlbumImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistTracksItemsItemAlbumRestrictionsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistTracksItemsItemArtistsModel struct {
	ID           types.String                                                                  `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsItemArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                  `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                    `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsItemArtistsExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsItemExternalIDsModel struct {
	Ean       types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc      types.String `tfsdk:"isrc" json:"isrc,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Upc       types.String `tfsdk:"upc" json:"upc,computed"`
}

type UserPlaylistTracksItemsItemExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsItemLinkedFromModel struct {
	ID           types.String                                                                     `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsItemLinkedFromExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                     `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                                       `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                     `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                     `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsItemLinkedFromExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsItemRestrictionsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistTracksItemsItemImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistTracksItemsItemShowModel struct {
	ID                 types.String                                                                 `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                               `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[UserPlaylistTracksItemsItemShowCopyrightsModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                   `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[UserPlaylistTracksItemsItemShowExternalURLsModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                 `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                 `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[UserPlaylistTracksItemsItemShowImagesModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                   `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                               `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                 `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                 `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                  `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
	Published          types.Bool                                                                   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistTracksItemsItemShowCopyrightsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Text      types.String `tfsdk:"text" json:"text,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}

type UserPlaylistTracksItemsItemShowExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsItemShowImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistTracksItemsItemResumePointModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	Published        types.Bool  `tfsdk:"published" json:"published,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
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
	Published            types.Bool                                                              `tfsdk:"published" json:"published,computed"`
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
	Published            types.Bool                                                                   `tfsdk:"published" json:"published,computed"`
	Restrictions         customfield.NestedObject[UserPlaylistTracksItemsTrackAlbumRestrictionsModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type UserPlaylistTracksItemsTrackAlbumArtistsModel struct {
	ID           types.String                                                                        `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsTrackAlbumArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                        `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                        `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                          `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                        `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                        `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsTrackAlbumArtistsExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackAlbumExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackAlbumImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistTracksItemsTrackAlbumRestrictionsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistTracksItemsTrackArtistsModel struct {
	ID           types.String                                                                   `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsTrackArtistsExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                   `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Published    types.Bool                                                                     `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                   `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                   `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsTrackArtistsExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackExternalIDsModel struct {
	Ean       types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc      types.String `tfsdk:"isrc" json:"isrc,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Upc       types.String `tfsdk:"upc" json:"upc,computed"`
}

type UserPlaylistTracksItemsTrackExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackLinkedFromModel struct {
	ID           types.String                                                                      `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[UserPlaylistTracksItemsTrackLinkedFromExternalURLsModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                      `tfsdk:"href" json:"href,computed"`
	Published    types.Bool                                                                        `tfsdk:"published" json:"published,computed"`
	Type         types.String                                                                      `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                      `tfsdk:"uri" json:"uri,computed"`
}

type UserPlaylistTracksItemsTrackLinkedFromExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackRestrictionsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Reason    types.String `tfsdk:"reason" json:"reason,computed"`
}

type UserPlaylistTracksItemsTrackImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
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
	Published          types.Bool                                                                    `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistTracksItemsTrackShowCopyrightsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Text      types.String `tfsdk:"text" json:"text,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}

type UserPlaylistTracksItemsTrackShowExternalURLsModel struct {
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
	Spotify   types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type UserPlaylistTracksItemsTrackShowImagesModel struct {
	Height    types.Int64  `tfsdk:"height" json:"height,computed"`
	URL       types.String `tfsdk:"url" json:"url,computed"`
	Width     types.Int64  `tfsdk:"width" json:"width,computed"`
	Published types.Bool   `tfsdk:"published" json:"published,computed"`
}

type UserPlaylistTracksItemsTrackResumePointModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	Published        types.Bool  `tfsdk:"published" json:"published,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
