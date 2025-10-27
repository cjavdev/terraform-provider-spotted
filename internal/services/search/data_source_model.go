// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package search

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

type SearchDataSourceModel struct {
	Q               types.String                                              `tfsdk:"q" query:"q,required"`
	Type            *[]types.String                                           `tfsdk:"type" query:"type,required"`
	IncludeExternal types.String                                              `tfsdk:"include_external" query:"include_external,optional"`
	Market          types.String                                              `tfsdk:"market" query:"market,optional"`
	Limit           types.Int64                                               `tfsdk:"limit" query:"limit,computed_optional"`
	Offset          types.Int64                                               `tfsdk:"offset" query:"offset,computed_optional"`
	Albums          customfield.NestedObject[SearchAlbumsDataSourceModel]     `tfsdk:"albums" json:"albums,computed"`
	Artists         customfield.NestedObject[SearchArtistsDataSourceModel]    `tfsdk:"artists" json:"artists,computed"`
	Audiobooks      customfield.NestedObject[SearchAudiobooksDataSourceModel] `tfsdk:"audiobooks" json:"audiobooks,computed"`
	Episodes        customfield.NestedObject[SearchEpisodesDataSourceModel]   `tfsdk:"episodes" json:"episodes,computed"`
	Playlists       customfield.NestedObject[SearchPlaylistsDataSourceModel]  `tfsdk:"playlists" json:"playlists,computed"`
	Shows           customfield.NestedObject[SearchShowsDataSourceModel]      `tfsdk:"shows" json:"shows,computed"`
	Tracks          customfield.NestedObject[SearchTracksDataSourceModel]     `tfsdk:"tracks" json:"tracks,computed"`
}

func (m *SearchDataSourceModel) toReadParams(_ context.Context) (params spotted.SearchGetParams, diags diag.Diagnostics) {
	mType := []string{}
	if m.Type != nil {
		for _, item := range *m.Type {
			mType = append(mType, string(item.ValueString()))
		}
	}

	params = spotted.SearchGetParams{
		Q:    m.Q.ValueString(),
		Type: mType,
	}

	if !m.IncludeExternal.IsNull() {
		params.IncludeExternal = spotted.SearchGetParamsIncludeExternal(m.IncludeExternal.ValueString())
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

type SearchAlbumsDataSourceModel struct {
	Href     types.String                                                   `tfsdk:"href" json:"href,computed"`
	Items    customfield.NestedObjectList[SearchAlbumsItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Limit    types.Int64                                                    `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                   `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                    `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                   `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                    `tfsdk:"total" json:"total,computed"`
}

type SearchAlbumsItemsDataSourceModel struct {
	ID                   types.String                                                           `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                           `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[SearchAlbumsItemsArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                         `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[SearchAlbumsItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                           `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[SearchAlbumsItemsImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                           `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                           `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                           `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                            `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                           `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                           `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[SearchAlbumsItemsRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type SearchAlbumsItemsArtistsDataSourceModel struct {
	ID           types.String                                                                  `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[SearchAlbumsItemsArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                  `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
}

type SearchAlbumsItemsArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchAlbumsItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchAlbumsItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type SearchAlbumsItemsRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type SearchArtistsDataSourceModel struct {
	Href     types.String                                                    `tfsdk:"href" json:"href,computed"`
	Items    customfield.NestedObjectList[SearchArtistsItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Limit    types.Int64                                                     `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                    `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                     `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                    `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                     `tfsdk:"total" json:"total,computed"`
}

type SearchArtistsItemsDataSourceModel struct {
	ID           types.String                                                            `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[SearchArtistsItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Followers    customfield.NestedObject[SearchArtistsItemsFollowersDataSourceModel]    `tfsdk:"followers" json:"followers,computed"`
	Genres       customfield.List[types.String]                                          `tfsdk:"genres" json:"genres,computed"`
	Href         types.String                                                            `tfsdk:"href" json:"href,computed"`
	Images       customfield.NestedObjectList[SearchArtistsItemsImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name         types.String                                                            `tfsdk:"name" json:"name,computed"`
	Popularity   types.Int64                                                             `tfsdk:"popularity" json:"popularity,computed"`
	Type         types.String                                                            `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                            `tfsdk:"uri" json:"uri,computed"`
}

type SearchArtistsItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchArtistsItemsFollowersDataSourceModel struct {
	Href  types.String `tfsdk:"href" json:"href,computed"`
	Total types.Int64  `tfsdk:"total" json:"total,computed"`
}

type SearchArtistsItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type SearchAudiobooksDataSourceModel struct {
	Href     types.String                                                       `tfsdk:"href" json:"href,computed"`
	Items    customfield.NestedObjectList[SearchAudiobooksItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Limit    types.Int64                                                        `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                       `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                        `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                       `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                        `tfsdk:"total" json:"total,computed"`
}

type SearchAudiobooksItemsDataSourceModel struct {
	ID               types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Authors          customfield.NestedObjectList[SearchAudiobooksItemsAuthorsDataSourceModel]    `tfsdk:"authors" json:"authors,computed"`
	AvailableMarkets customfield.List[types.String]                                               `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights       customfield.NestedObjectList[SearchAudiobooksItemsCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description      types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Explicit         types.Bool                                                                   `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs     customfield.NestedObject[SearchAudiobooksItemsExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href             types.String                                                                 `tfsdk:"href" json:"href,computed"`
	HTMLDescription  types.String                                                                 `tfsdk:"html_description" json:"html_description,computed"`
	Images           customfield.NestedObjectList[SearchAudiobooksItemsImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	Languages        customfield.List[types.String]                                               `tfsdk:"languages" json:"languages,computed"`
	MediaType        types.String                                                                 `tfsdk:"media_type" json:"media_type,computed"`
	Name             types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Narrators        customfield.NestedObjectList[SearchAudiobooksItemsNarratorsDataSourceModel]  `tfsdk:"narrators" json:"narrators,computed"`
	Publisher        types.String                                                                 `tfsdk:"publisher" json:"publisher,computed"`
	TotalChapters    types.Int64                                                                  `tfsdk:"total_chapters" json:"total_chapters,computed"`
	Type             types.String                                                                 `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                                 `tfsdk:"uri" json:"uri,computed"`
	Edition          types.String                                                                 `tfsdk:"edition" json:"edition,computed"`
}

type SearchAudiobooksItemsAuthorsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type SearchAudiobooksItemsCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type SearchAudiobooksItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchAudiobooksItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type SearchAudiobooksItemsNarratorsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type SearchEpisodesDataSourceModel struct {
	Href     types.String                                                     `tfsdk:"href" json:"href,computed"`
	Items    customfield.NestedObjectList[SearchEpisodesItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Limit    types.Int64                                                      `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                     `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                      `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                     `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                      `tfsdk:"total" json:"total,computed"`
}

type SearchEpisodesItemsDataSourceModel struct {
	ID                   types.String                                                             `tfsdk:"id" json:"id,computed"`
	AudioPreviewURL      types.String                                                             `tfsdk:"audio_preview_url" json:"audio_preview_url,computed"`
	Description          types.String                                                             `tfsdk:"description" json:"description,computed"`
	DurationMs           types.Int64                                                              `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit             types.Bool                                                               `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs         customfield.NestedObject[SearchEpisodesItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                             `tfsdk:"href" json:"href,computed"`
	HTMLDescription      types.String                                                             `tfsdk:"html_description" json:"html_description,computed"`
	Images               customfield.NestedObjectList[SearchEpisodesItemsImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted   types.Bool                                                               `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	IsPlayable           types.Bool                                                               `tfsdk:"is_playable" json:"is_playable,computed"`
	Languages            customfield.List[types.String]                                           `tfsdk:"languages" json:"languages,computed"`
	Name                 types.String                                                             `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                             `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                             `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	Type                 types.String                                                             `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                             `tfsdk:"uri" json:"uri,computed"`
	Language             types.String                                                             `tfsdk:"language" json:"language,computed"`
	Restrictions         customfield.NestedObject[SearchEpisodesItemsRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	ResumePoint          customfield.NestedObject[SearchEpisodesItemsResumePointDataSourceModel]  `tfsdk:"resume_point" json:"resume_point,computed"`
}

type SearchEpisodesItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchEpisodesItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type SearchEpisodesItemsRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type SearchEpisodesItemsResumePointDataSourceModel struct {
	FullyPlayed      types.Bool  `tfsdk:"fully_played" json:"fully_played,computed"`
	ResumePositionMs types.Int64 `tfsdk:"resume_position_ms" json:"resume_position_ms,computed"`
}

type SearchPlaylistsDataSourceModel struct {
	Href     types.String                                                      `tfsdk:"href" json:"href,computed"`
	Items    customfield.NestedObjectList[SearchPlaylistsItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Limit    types.Int64                                                       `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                      `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                       `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                      `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                       `tfsdk:"total" json:"total,computed"`
}

type SearchPlaylistsItemsDataSourceModel struct {
	ID            types.String                                                              `tfsdk:"id" json:"id,computed"`
	Collaborative types.Bool                                                                `tfsdk:"collaborative" json:"collaborative,computed"`
	Description   types.String                                                              `tfsdk:"description" json:"description,computed"`
	ExternalURLs  customfield.NestedObject[SearchPlaylistsItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href          types.String                                                              `tfsdk:"href" json:"href,computed"`
	Images        customfield.NestedObjectList[SearchPlaylistsItemsImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name          types.String                                                              `tfsdk:"name" json:"name,computed"`
	Owner         customfield.NestedObject[SearchPlaylistsItemsOwnerDataSourceModel]        `tfsdk:"owner" json:"owner,computed"`
	Public        types.Bool                                                                `tfsdk:"public" json:"public,computed"`
	SnapshotID    types.String                                                              `tfsdk:"snapshot_id" json:"snapshot_id,computed"`
	Tracks        customfield.NestedObject[SearchPlaylistsItemsTracksDataSourceModel]       `tfsdk:"tracks" json:"tracks,computed"`
	Type          types.String                                                              `tfsdk:"type" json:"type,computed"`
	Uri           types.String                                                              `tfsdk:"uri" json:"uri,computed"`
}

type SearchPlaylistsItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchPlaylistsItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type SearchPlaylistsItemsOwnerDataSourceModel struct {
	ID           types.String                                                                   `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[SearchPlaylistsItemsOwnerExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                   `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                   `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                   `tfsdk:"uri" json:"uri,computed"`
	DisplayName  types.String                                                                   `tfsdk:"display_name" json:"display_name,computed"`
}

type SearchPlaylistsItemsOwnerExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchPlaylistsItemsTracksDataSourceModel struct {
	Href  types.String `tfsdk:"href" json:"href,computed"`
	Total types.Int64  `tfsdk:"total" json:"total,computed"`
}

type SearchShowsDataSourceModel struct {
	Href     types.String                                                  `tfsdk:"href" json:"href,computed"`
	Items    customfield.NestedObjectList[SearchShowsItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Limit    types.Int64                                                   `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                  `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                   `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                  `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                   `tfsdk:"total" json:"total,computed"`
}

type SearchShowsItemsDataSourceModel struct {
	ID                 types.String                                                            `tfsdk:"id" json:"id,computed"`
	AvailableMarkets   customfield.List[types.String]                                          `tfsdk:"available_markets" json:"available_markets,computed"`
	Copyrights         customfield.NestedObjectList[SearchShowsItemsCopyrightsDataSourceModel] `tfsdk:"copyrights" json:"copyrights,computed"`
	Description        types.String                                                            `tfsdk:"description" json:"description,computed"`
	Explicit           types.Bool                                                              `tfsdk:"explicit" json:"explicit,computed"`
	ExternalURLs       customfield.NestedObject[SearchShowsItemsExternalURLsDataSourceModel]   `tfsdk:"external_urls" json:"external_urls,computed"`
	Href               types.String                                                            `tfsdk:"href" json:"href,computed"`
	HTMLDescription    types.String                                                            `tfsdk:"html_description" json:"html_description,computed"`
	Images             customfield.NestedObjectList[SearchShowsItemsImagesDataSourceModel]     `tfsdk:"images" json:"images,computed"`
	IsExternallyHosted types.Bool                                                              `tfsdk:"is_externally_hosted" json:"is_externally_hosted,computed"`
	Languages          customfield.List[types.String]                                          `tfsdk:"languages" json:"languages,computed"`
	MediaType          types.String                                                            `tfsdk:"media_type" json:"media_type,computed"`
	Name               types.String                                                            `tfsdk:"name" json:"name,computed"`
	Publisher          types.String                                                            `tfsdk:"publisher" json:"publisher,computed"`
	TotalEpisodes      types.Int64                                                             `tfsdk:"total_episodes" json:"total_episodes,computed"`
	Type               types.String                                                            `tfsdk:"type" json:"type,computed"`
	Uri                types.String                                                            `tfsdk:"uri" json:"uri,computed"`
}

type SearchShowsItemsCopyrightsDataSourceModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
	Type types.String `tfsdk:"type" json:"type,computed"`
}

type SearchShowsItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchShowsItemsImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type SearchTracksDataSourceModel struct {
	Href     types.String                                                   `tfsdk:"href" json:"href,computed"`
	Items    customfield.NestedObjectList[SearchTracksItemsDataSourceModel] `tfsdk:"items" json:"items,computed"`
	Limit    types.Int64                                                    `tfsdk:"limit" json:"limit,computed"`
	Next     types.String                                                   `tfsdk:"next" json:"next,computed"`
	Offset   types.Int64                                                    `tfsdk:"offset" json:"offset,computed"`
	Previous types.String                                                   `tfsdk:"previous" json:"previous,computed"`
	Total    types.Int64                                                    `tfsdk:"total" json:"total,computed"`
}

type SearchTracksItemsDataSourceModel struct {
	ID               types.String                                                           `tfsdk:"id" json:"id,computed"`
	Album            customfield.NestedObject[SearchTracksItemsAlbumDataSourceModel]        `tfsdk:"album" json:"album,computed"`
	Artists          customfield.NestedObjectList[SearchTracksItemsArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets customfield.List[types.String]                                         `tfsdk:"available_markets" json:"available_markets,computed"`
	DiscNumber       types.Int64                                                            `tfsdk:"disc_number" json:"disc_number,computed"`
	DurationMs       types.Int64                                                            `tfsdk:"duration_ms" json:"duration_ms,computed"`
	Explicit         types.Bool                                                             `tfsdk:"explicit" json:"explicit,computed"`
	ExternalIDs      customfield.NestedObject[SearchTracksItemsExternalIDsDataSourceModel]  `tfsdk:"external_ids" json:"external_ids,computed"`
	ExternalURLs     customfield.NestedObject[SearchTracksItemsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href             types.String                                                           `tfsdk:"href" json:"href,computed"`
	IsLocal          types.Bool                                                             `tfsdk:"is_local" json:"is_local,computed"`
	IsPlayable       types.Bool                                                             `tfsdk:"is_playable" json:"is_playable,computed"`
	LinkedFrom       customfield.NestedObject[SearchTracksItemsLinkedFromDataSourceModel]   `tfsdk:"linked_from" json:"linked_from,computed"`
	Name             types.String                                                           `tfsdk:"name" json:"name,computed"`
	Popularity       types.Int64                                                            `tfsdk:"popularity" json:"popularity,computed"`
	PreviewURL       types.String                                                           `tfsdk:"preview_url" json:"preview_url,computed"`
	Restrictions     customfield.NestedObject[SearchTracksItemsRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
	TrackNumber      types.Int64                                                            `tfsdk:"track_number" json:"track_number,computed"`
	Type             types.String                                                           `tfsdk:"type" json:"type,computed"`
	Uri              types.String                                                           `tfsdk:"uri" json:"uri,computed"`
}

type SearchTracksItemsAlbumDataSourceModel struct {
	ID                   types.String                                                                `tfsdk:"id" json:"id,computed"`
	AlbumType            types.String                                                                `tfsdk:"album_type" json:"album_type,computed"`
	Artists              customfield.NestedObjectList[SearchTracksItemsAlbumArtistsDataSourceModel]  `tfsdk:"artists" json:"artists,computed"`
	AvailableMarkets     customfield.List[types.String]                                              `tfsdk:"available_markets" json:"available_markets,computed"`
	ExternalURLs         customfield.NestedObject[SearchTracksItemsAlbumExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href                 types.String                                                                `tfsdk:"href" json:"href,computed"`
	Images               customfield.NestedObjectList[SearchTracksItemsAlbumImagesDataSourceModel]   `tfsdk:"images" json:"images,computed"`
	Name                 types.String                                                                `tfsdk:"name" json:"name,computed"`
	ReleaseDate          types.String                                                                `tfsdk:"release_date" json:"release_date,computed"`
	ReleaseDatePrecision types.String                                                                `tfsdk:"release_date_precision" json:"release_date_precision,computed"`
	TotalTracks          types.Int64                                                                 `tfsdk:"total_tracks" json:"total_tracks,computed"`
	Type                 types.String                                                                `tfsdk:"type" json:"type,computed"`
	Uri                  types.String                                                                `tfsdk:"uri" json:"uri,computed"`
	Restrictions         customfield.NestedObject[SearchTracksItemsAlbumRestrictionsDataSourceModel] `tfsdk:"restrictions" json:"restrictions,computed"`
}

type SearchTracksItemsAlbumArtistsDataSourceModel struct {
	ID           types.String                                                                       `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[SearchTracksItemsAlbumArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                       `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                       `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                       `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                       `tfsdk:"uri" json:"uri,computed"`
}

type SearchTracksItemsAlbumArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchTracksItemsAlbumExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchTracksItemsAlbumImagesDataSourceModel struct {
	Height types.Int64  `tfsdk:"height" json:"height,computed"`
	URL    types.String `tfsdk:"url" json:"url,computed"`
	Width  types.Int64  `tfsdk:"width" json:"width,computed"`
}

type SearchTracksItemsAlbumRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}

type SearchTracksItemsArtistsDataSourceModel struct {
	ID           types.String                                                                  `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[SearchTracksItemsArtistsExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                  `tfsdk:"href" json:"href,computed"`
	Name         types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Type         types.String                                                                  `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                  `tfsdk:"uri" json:"uri,computed"`
}

type SearchTracksItemsArtistsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchTracksItemsExternalIDsDataSourceModel struct {
	Ean  types.String `tfsdk:"ean" json:"ean,computed"`
	Isrc types.String `tfsdk:"isrc" json:"isrc,computed"`
	Upc  types.String `tfsdk:"upc" json:"upc,computed"`
}

type SearchTracksItemsExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchTracksItemsLinkedFromDataSourceModel struct {
	ID           types.String                                                                     `tfsdk:"id" json:"id,computed"`
	ExternalURLs customfield.NestedObject[SearchTracksItemsLinkedFromExternalURLsDataSourceModel] `tfsdk:"external_urls" json:"external_urls,computed"`
	Href         types.String                                                                     `tfsdk:"href" json:"href,computed"`
	Type         types.String                                                                     `tfsdk:"type" json:"type,computed"`
	Uri          types.String                                                                     `tfsdk:"uri" json:"uri,computed"`
}

type SearchTracksItemsLinkedFromExternalURLsDataSourceModel struct {
	Spotify types.String `tfsdk:"spotify" json:"spotify,computed"`
}

type SearchTracksItemsRestrictionsDataSourceModel struct {
	Reason types.String `tfsdk:"reason" json:"reason,computed"`
}
