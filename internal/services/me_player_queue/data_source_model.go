// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me_player_queue

import (
	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MePlayerQueueDataSourceModel struct {
	CurrentlyPlaying customfield.NestedObject[MePlayerQueueCurrentlyPlayingDataSourceModel] `tfsdk:"currently_playing" json:"currently_playing,computed"`
	Queue            customfield.NestedObjectList[MePlayerQueueQueueDataSourceModel]        `tfsdk:"queue" json:"queue,computed"`
}

type MePlayerQueueCurrentlyPlayingDataSourceModel struct {
	ID                   types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[MePlayerQueueCurrentlyPlayingAlbumDataSourceModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[MePlayerQueueCurrentlyPlayingArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                                     `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                                        `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                                        `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                                         `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[MePlayerQueueCurrentlyPlayingExternalIDsDataSourceModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[MePlayerQueueCurrentlyPlayingExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                       `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                                         `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                                         `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[MePlayerQueueCurrentlyPlayingLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                                       `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                                        `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                                       `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions         customfield.NestedObject[MePlayerQueueCurrentlyPlayingRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                                        `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                                       `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                       `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                                       `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                                       `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                                       `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[MePlayerQueueCurrentlyPlayingImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                                         `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                                     `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                                       `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                       `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[MePlayerQueueCurrentlyPlayingShowDataSourceModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                                       `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[MePlayerQueueCurrentlyPlayingResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type MePlayerQueueCurrentlyPlayingAlbumDataSourceModel struct {
	ID                   types.String                                                                            `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                            `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[MePlayerQueueCurrentlyPlayingAlbumArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                                          `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[MePlayerQueueCurrentlyPlayingAlbumExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                            `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[MePlayerQueueCurrentlyPlayingAlbumImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                            `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                            `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                            `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                             `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                            `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                            `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[MePlayerQueueCurrentlyPlayingAlbumRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type MePlayerQueueCurrentlyPlayingAlbumArtistsDataSourceModel struct {
	ID           types.String                                                                                   `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MePlayerQueueCurrentlyPlayingAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                                   `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                                   `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                                   `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                                   `tfsdk:"uri" json:"uri,computed"`
}

type MePlayerQueueCurrentlyPlayingAlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueCurrentlyPlayingAlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueCurrentlyPlayingAlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MePlayerQueueCurrentlyPlayingAlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type MePlayerQueueCurrentlyPlayingArtistsDataSourceModel struct {
	ID           types.String                                                                              `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MePlayerQueueCurrentlyPlayingArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                              `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                              `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                              `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                              `tfsdk:"uri" json:"uri,computed"`
}

type MePlayerQueueCurrentlyPlayingArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueCurrentlyPlayingExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type MePlayerQueueCurrentlyPlayingExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueCurrentlyPlayingLinkedFromDataSourceModel struct {
	ID           types.String                                                                                 `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MePlayerQueueCurrentlyPlayingLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                                 `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                                 `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                                 `tfsdk:"uri" json:"uri,computed"`
}

type MePlayerQueueCurrentlyPlayingLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueCurrentlyPlayingRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type MePlayerQueueCurrentlyPlayingImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MePlayerQueueCurrentlyPlayingShowDataSourceModel struct {
	ID                 types.String                                                                             `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                                           `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[MePlayerQueueCurrentlyPlayingShowCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                             `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                               `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[MePlayerQueueCurrentlyPlayingShowExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                             `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                             `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[MePlayerQueueCurrentlyPlayingShowImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                               `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                                           `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                             `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                             `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                             `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                              `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                             `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                             `tfsdk:"uri" json:"uri,computed"`
}

type MePlayerQueueCurrentlyPlayingShowCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type MePlayerQueueCurrentlyPlayingShowExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueCurrentlyPlayingShowImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MePlayerQueueCurrentlyPlayingResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}

type MePlayerQueueQueueDataSourceModel struct {
	ID                   types.String                                                            `tfsdk:"id" json:"id,computed"`
	Album                customfield.NestedObject[MePlayerQueueQueueAlbumDataSourceModel]        `tfsdk:"album" json:"album,computed"`
	Artists              customfield.NestedObjectList[MePlayerQueueQueueArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                          `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber           types.Int64                                                             `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs           types.Int64                                                             `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                              `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs          customfield.NestedObject[MePlayerQueueQueueExternalIDsDataSourceModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs         customfield.NestedObject[MePlayerQueueQueueExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                            `tfsdk:"href" json:"href,computed"`
	IsLocal              types.Bool                                                              `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable           types.Bool                                                              `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom           customfield.NestedObject[MePlayerQueueQueueLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name                 types.String                                                            `tfsdk:"name" json:"name,computed"`
	Popularity           types.Int64                                                             `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL           types.String                                                            `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions         customfield.NestedObject[MePlayerQueueQueueRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber          types.Int64                                                             `tfsdk:"track_number" json:"track_number,computed"`
	Type                 types.String                                                            `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                            `tfsdk:"uri" json:"uri,computed"`
	AudioPreviewURL      types.String                                                            `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                            `tfsdk:"description" json:"description,computed"`
	HTMLDescription      types.String                                                            `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[MePlayerQueueQueueImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                              `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages            customfield.List[types.String]                                          `tfsdk:"languages" json:"languages,computed"`
	ReleaseDate          types.String                                                            `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                            `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Show                 customfield.NestedObject[MePlayerQueueQueueShowDataSourceModel]         `tfsdk:"show" json:"show,computed"`
	Language             types.String                                                            `tfsdk:"language" json:"language,computed"`
	ResumePoint          customfield.NestedObject[MePlayerQueueQueueResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type MePlayerQueueQueueAlbumDataSourceModel struct {
	ID                   types.String                                                                 `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                 `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[MePlayerQueueQueueAlbumArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                               `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[MePlayerQueueQueueAlbumExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                 `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[MePlayerQueueQueueAlbumImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                 `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                 `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                 `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                  `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[MePlayerQueueQueueAlbumRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type MePlayerQueueQueueAlbumArtistsDataSourceModel struct {
	ID           types.String                                                                        `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MePlayerQueueQueueAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                        `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                        `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                        `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                        `tfsdk:"uri" json:"uri,computed"`
}

type MePlayerQueueQueueAlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueQueueAlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueQueueAlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MePlayerQueueQueueAlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type MePlayerQueueQueueArtistsDataSourceModel struct {
	ID           types.String                                                                   `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MePlayerQueueQueueArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                   `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                   `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                   `tfsdk:"uri" json:"uri,computed"`
}

type MePlayerQueueQueueArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueQueueExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type MePlayerQueueQueueExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueQueueLinkedFromDataSourceModel struct {
	ID           types.String                                                                      `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[MePlayerQueueQueueLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                      `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                      `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                      `tfsdk:"uri" json:"uri,computed"`
}

type MePlayerQueueQueueLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueQueueRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type MePlayerQueueQueueImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MePlayerQueueQueueShowDataSourceModel struct {
	ID                 types.String                                                                  `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                                `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[MePlayerQueueQueueShowCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                                  `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                                    `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[MePlayerQueueQueueShowExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                                  `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                                  `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[MePlayerQueueQueueShowImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                                    `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                                `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                                  `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                                  `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                                   `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
}

type MePlayerQueueQueueShowCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type MePlayerQueueQueueShowExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type MePlayerQueueQueueShowImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type MePlayerQueueQueueResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}
