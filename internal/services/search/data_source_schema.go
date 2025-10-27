// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package search

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/spotted-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*SearchDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"q": schema.StringAttribute{
				Description: "Your search query.\n\nYou can narrow down your search using field filters. The available filters are `album`, `artist`, `track`, `year`, `upc`, `tag:hipster`, `tag:new`, `isrc`, and `genre`. Each field filter only applies to certain result types.\n\nThe `artist` and `year` filters can be used while searching albums, artists and tracks. You can filter on a single `year` or a range (e.g. 1955-1960).<br />\nThe `album` filter can be used while searching albums and tracks.<br />\nThe `genre` filter can be used while searching artists and tracks.<br />\nThe `isrc` and `track` filters can be used while searching tracks.<br />\nThe `upc`, `tag:new` and `tag:hipster` filters can only be used while searching albums. The `tag:new` filter will return albums released in the past two weeks and `tag:hipster` can be used to return only albums with the lowest 10% popularity.<br />",
				Required:    true,
			},
			"type": schema.ListAttribute{
				Description: "A comma-separated list of item types to search across. Search results include hits\nfrom all the specified item types. For example: `q=abacab&type=album,track` returns\nboth albums and tracks matching \"abacab\".",
				Required:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"album",
							"artist",
							"playlist",
							"track",
							"show",
							"episode",
							"audiobook",
						),
					),
				},
				ElementType: types.StringType,
			},
			"include_external": schema.StringAttribute{
				Description: "If `include_external=audio` is specified it signals that the client can play externally hosted audio content, and marks\nthe content as playable in the response. By default externally hosted audio content is marked as unplayable in the response.\nAvailable values: \"audio\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("audio"),
				},
			},
			"market": schema.StringAttribute{
				Description: "An [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).\n  If a country code is specified, only content that is available in that market will be returned.<br/>\n  If a valid user access token is specified in the request header, the country associated with\n  the user account will take priority over this parameter.<br/>\n  _**Note**: If neither market or user country are provided, the content is considered unavailable for the client._<br/>\n  Users can view the country that is associated with their account in the [account settings](https://www.spotify.com/account/overview/).",
				Optional:    true,
			},
			"limit": schema.Int64Attribute{
				Description: "The maximum number of results to return in each item type.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 50),
				},
			},
			"offset": schema.Int64Attribute{
				Description: "The index of the first result to return. Use\nwith limit to get the next page of search results.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1000),
				},
			},
			"albums": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[SearchAlbumsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
					},
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[SearchAlbumsItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the album.",
									Computed:    true,
								},
								"album_type": schema.StringAttribute{
									Description: "The type of the album.\nAvailable values: \"album\", \"single\", \"compilation\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"album",
											"single",
											"compilation",
										),
									},
								},
								"artists": schema.ListNestedAttribute{
									Description: "The artists of the album. Each artist object includes a link in `href` to more detailed information about the artist.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchAlbumsItemsArtistsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
												Computed:    true,
											},
											"external_urls": schema.SingleNestedAttribute{
												Description: "Known external URLs for this artist.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[SearchAlbumsItemsArtistsExternalURLsDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"spotify": schema.StringAttribute{
														Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
														Computed:    true,
													},
												},
											},
											"href": schema.StringAttribute{
												Description: "A link to the Web API endpoint providing full details of the artist.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the artist.",
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "The object type.\nAvailable values: \"artist\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("artist"),
												},
											},
											"uri": schema.StringAttribute{
												Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
												Computed:    true,
											},
										},
									},
								},
								"available_markets": schema.ListAttribute{
									Description: "The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known external URLs for this album.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchAlbumsItemsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the album.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "The cover art for the album in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchAlbumsItemsImagesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"height": schema.Int64Attribute{
												Description: "The image height in pixels.",
												Computed:    true,
											},
											"url": schema.StringAttribute{
												Description: "The source URL of the image.",
												Computed:    true,
											},
											"width": schema.Int64Attribute{
												Description: "The image width in pixels.",
												Computed:    true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									Description: "The name of the album. In case of an album takedown, the value may be an empty string.",
									Computed:    true,
								},
								"release_date": schema.StringAttribute{
									Description: "The date the album was first released.",
									Computed:    true,
								},
								"release_date_precision": schema.StringAttribute{
									Description: "The precision with which `release_date` value is known.\nAvailable values: \"year\", \"month\", \"day\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"year",
											"month",
											"day",
										),
									},
								},
								"total_tracks": schema.Int64Attribute{
									Description: "The number of tracks in the album.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"album\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("album"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the album.",
									Computed:    true,
								},
								"restrictions": schema.SingleNestedAttribute{
									Description: "Included in the response when a content restriction is applied.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchAlbumsItemsRestrictionsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"reason": schema.StringAttribute{
											Description: "The reason for the restriction. Albums may be restricted if the content is not available in a given market, to the user's subscription type, or when the user's account is set to not play explicit content.\nAdditional reasons may be added in the future.\nAvailable values: \"market\", \"product\", \"explicit\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"market",
													"product",
													"explicit",
												),
											},
										},
									},
								},
							},
						},
					},
					"limit": schema.Int64Attribute{
						Description: "The maximum number of items in the response (as set in the query or by default).",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "URL to the next page of items. ( `null` if none)",
						Computed:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "The offset of the items returned (as set in the query or by default)",
						Computed:    true,
					},
					"previous": schema.StringAttribute{
						Description: "URL to the previous page of items. ( `null` if none)",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of items available to return.",
						Computed:    true,
					},
				},
			},
			"artists": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[SearchArtistsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
					},
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[SearchArtistsItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known external URLs for this artist.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchArtistsItemsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"followers": schema.SingleNestedAttribute{
									Description: "Information about the followers of the artist.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchArtistsItemsFollowersDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"href": schema.StringAttribute{
											Description: "This will always be set to null, as the Web API does not support it at the moment.",
											Computed:    true,
										},
										"total": schema.Int64Attribute{
											Description: "The total number of followers.",
											Computed:    true,
										},
									},
								},
								"genres": schema.ListAttribute{
									Description: "A list of the genres the artist is associated with. If not yet classified, the array is empty.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the artist.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "Images of the artist in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchArtistsItemsImagesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"height": schema.Int64Attribute{
												Description: "The image height in pixels.",
												Computed:    true,
											},
											"url": schema.StringAttribute{
												Description: "The source URL of the image.",
												Computed:    true,
											},
											"width": schema.Int64Attribute{
												Description: "The image width in pixels.",
												Computed:    true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									Description: "The name of the artist.",
									Computed:    true,
								},
								"popularity": schema.Int64Attribute{
									Description: "The popularity of the artist. The value will be between 0 and 100, with 100 being the most popular. The artist's popularity is calculated from the popularity of all the artist's tracks.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"artist\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("artist"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
									Computed:    true,
								},
							},
						},
					},
					"limit": schema.Int64Attribute{
						Description: "The maximum number of items in the response (as set in the query or by default).",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "URL to the next page of items. ( `null` if none)",
						Computed:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "The offset of the items returned (as set in the query or by default)",
						Computed:    true,
					},
					"previous": schema.StringAttribute{
						Description: "URL to the previous page of items. ( `null` if none)",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of items available to return.",
						Computed:    true,
					},
				},
			},
			"audiobooks": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[SearchAudiobooksDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
					},
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[SearchAudiobooksItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the audiobook.",
									Computed:    true,
								},
								"authors": schema.ListNestedAttribute{
									Description: "The author(s) for the audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchAudiobooksItemsAuthorsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the author.",
												Computed:    true,
											},
										},
									},
								},
								"available_markets": schema.ListAttribute{
									Description: "A list of the countries in which the audiobook can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"copyrights": schema.ListNestedAttribute{
									Description: "The copyright statements of the audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchAudiobooksItemsCopyrightsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"text": schema.StringAttribute{
												Description: "The copyright text for this content.",
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "The type of copyright: `C` = the copyright, `P` = the sound recording (performance) copyright.",
												Computed:    true,
											},
										},
									},
								},
								"description": schema.StringAttribute{
									Description: "A description of the audiobook. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
									Computed:    true,
								},
								"explicit": schema.BoolAttribute{
									Description: "Whether or not the audiobook has explicit content (true = yes it does; false = no it does not OR unknown).",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "External URLs for this audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchAudiobooksItemsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the audiobook.",
									Computed:    true,
								},
								"html_description": schema.StringAttribute{
									Description: "A description of the audiobook. This field may contain HTML tags.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "The cover art for the audiobook in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchAudiobooksItemsImagesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"height": schema.Int64Attribute{
												Description: "The image height in pixels.",
												Computed:    true,
											},
											"url": schema.StringAttribute{
												Description: "The source URL of the image.",
												Computed:    true,
											},
											"width": schema.Int64Attribute{
												Description: "The image width in pixels.",
												Computed:    true,
											},
										},
									},
								},
								"languages": schema.ListAttribute{
									Description: "A list of the languages used in the audiobook, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"media_type": schema.StringAttribute{
									Description: "The media type of the audiobook.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the audiobook.",
									Computed:    true,
								},
								"narrators": schema.ListNestedAttribute{
									Description: "The narrator(s) for the audiobook.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchAudiobooksItemsNarratorsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "The name of the Narrator.",
												Computed:    true,
											},
										},
									},
								},
								"publisher": schema.StringAttribute{
									Description: "The publisher of the audiobook.",
									Computed:    true,
								},
								"total_chapters": schema.Int64Attribute{
									Description: "The number of chapters in this audiobook.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"audiobook\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("audiobook"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the audiobook.",
									Computed:    true,
								},
								"edition": schema.StringAttribute{
									Description: "The edition of the audiobook.",
									Computed:    true,
								},
							},
						},
					},
					"limit": schema.Int64Attribute{
						Description: "The maximum number of items in the response (as set in the query or by default).",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "URL to the next page of items. ( `null` if none)",
						Computed:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "The offset of the items returned (as set in the query or by default)",
						Computed:    true,
					},
					"previous": schema.StringAttribute{
						Description: "URL to the previous page of items. ( `null` if none)",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of items available to return.",
						Computed:    true,
					},
				},
			},
			"episodes": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[SearchEpisodesDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
					},
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[SearchEpisodesItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the episode.",
									Computed:    true,
								},
								"audio_preview_url": schema.StringAttribute{
									Description:        "A URL to a 30 second preview (MP3 format) of the episode. `null` if not available.",
									Computed:           true,
									DeprecationMessage: "This attribute is deprecated.",
								},
								"description": schema.StringAttribute{
									Description: "A description of the episode. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
									Computed:    true,
								},
								"duration_ms": schema.Int64Attribute{
									Description: "The episode length in milliseconds.",
									Computed:    true,
								},
								"explicit": schema.BoolAttribute{
									Description: "Whether or not the episode has explicit content (true = yes it does; false = no it does not OR unknown).",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "External URLs for this episode.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchEpisodesItemsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the episode.",
									Computed:    true,
								},
								"html_description": schema.StringAttribute{
									Description: "A description of the episode. This field may contain HTML tags.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "The cover art for the episode in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchEpisodesItemsImagesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"height": schema.Int64Attribute{
												Description: "The image height in pixels.",
												Computed:    true,
											},
											"url": schema.StringAttribute{
												Description: "The source URL of the image.",
												Computed:    true,
											},
											"width": schema.Int64Attribute{
												Description: "The image width in pixels.",
												Computed:    true,
											},
										},
									},
								},
								"is_externally_hosted": schema.BoolAttribute{
									Description: "True if the episode is hosted outside of Spotify's CDN.",
									Computed:    true,
								},
								"is_playable": schema.BoolAttribute{
									Description: "True if the episode is playable in the given market. Otherwise false.",
									Computed:    true,
								},
								"languages": schema.ListAttribute{
									Description: "A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"name": schema.StringAttribute{
									Description: "The name of the episode.",
									Computed:    true,
								},
								"release_date": schema.StringAttribute{
									Description: "The date the episode was first released, for example `\"1981-12-15\"`. Depending on the precision, it might be shown as `\"1981\"` or `\"1981-12\"`.",
									Computed:    true,
								},
								"release_date_precision": schema.StringAttribute{
									Description: "The precision with which `release_date` value is known.\nAvailable values: \"year\", \"month\", \"day\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"year",
											"month",
											"day",
										),
									},
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"episode\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("episode"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the episode.",
									Computed:    true,
								},
								"language": schema.StringAttribute{
									Description:        "The language used in the episode, identified by a [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated and might be removed in the future. Please use the `languages` field instead.",
									Computed:           true,
									DeprecationMessage: "This attribute is deprecated.",
								},
								"restrictions": schema.SingleNestedAttribute{
									Description: "Included in the response when a content restriction is applied.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchEpisodesItemsRestrictionsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"reason": schema.StringAttribute{
											Description: "The reason for the restriction. Supported values:\n- `market` - The content item is not available in the given market.\n- `product` - The content item is not available for the user's subscription type.\n- `explicit` - The content item is explicit and the user's account is set to not play explicit content.\n\nAdditional reasons may be added in the future.\n**Note**: If you use this field, make sure that your application safely handles unknown values.",
											Computed:    true,
										},
									},
								},
								"resume_point": schema.SingleNestedAttribute{
									Description: "The user's most recent position in the episode. Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchEpisodesItemsResumePointDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"fully_played": schema.BoolAttribute{
											Description: "Whether or not the episode has been fully played by the user.",
											Computed:    true,
										},
										"resume_position_ms": schema.Int64Attribute{
											Description: "The user's most recent position in the episode in milliseconds.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
					"limit": schema.Int64Attribute{
						Description: "The maximum number of items in the response (as set in the query or by default).",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "URL to the next page of items. ( `null` if none)",
						Computed:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "The offset of the items returned (as set in the query or by default)",
						Computed:    true,
					},
					"previous": schema.StringAttribute{
						Description: "URL to the previous page of items. ( `null` if none)",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of items available to return.",
						Computed:    true,
					},
				},
			},
			"playlists": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[SearchPlaylistsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
					},
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[SearchPlaylistsItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the playlist.",
									Computed:    true,
								},
								"collaborative": schema.BoolAttribute{
									Description: "`true` if the owner allows other users to modify the playlist.",
									Computed:    true,
								},
								"description": schema.StringAttribute{
									Description: "The playlist description. _Only returned for modified, verified playlists, otherwise_ `null`.",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known external URLs for this playlist.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchPlaylistsItemsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the playlist.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "Images for the playlist. The array may be empty or contain up to three images. The images are returned by size in descending order. See [Working with Playlists](/documentation/web-api/concepts/playlists). _**Note**: If returned, the source URL for the image (`url`) is temporary and will expire in less than a day._",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchPlaylistsItemsImagesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"height": schema.Int64Attribute{
												Description: "The image height in pixels.",
												Computed:    true,
											},
											"url": schema.StringAttribute{
												Description: "The source URL of the image.",
												Computed:    true,
											},
											"width": schema.Int64Attribute{
												Description: "The image width in pixels.",
												Computed:    true,
											},
										},
									},
								},
								"name": schema.StringAttribute{
									Description: "The name of the playlist.",
									Computed:    true,
								},
								"owner": schema.SingleNestedAttribute{
									Description: "The user who owns the playlist",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchPlaylistsItemsOwnerDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this user.",
											Computed:    true,
										},
										"external_urls": schema.SingleNestedAttribute{
											Description: "Known public external URLs for this user.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[SearchPlaylistsItemsOwnerExternalURLsDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"spotify": schema.StringAttribute{
													Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
													Computed:    true,
												},
											},
										},
										"href": schema.StringAttribute{
											Description: "A link to the Web API endpoint for this user.",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: "The object type.\nAvailable values: \"user\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("user"),
											},
										},
										"uri": schema.StringAttribute{
											Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this user.",
											Computed:    true,
										},
										"display_name": schema.StringAttribute{
											Description: "The name displayed on the user's profile. `null` if not available.",
											Computed:    true,
										},
									},
								},
								"public": schema.BoolAttribute{
									Description: "The playlist's public/private status (if it is added to the user's profile): `true` the playlist is public, `false` the playlist is private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
									Computed:    true,
								},
								"snapshot_id": schema.StringAttribute{
									Description: "The version identifier for the current playlist. Can be supplied in other requests to target a specific playlist version",
									Computed:    true,
								},
								"tracks": schema.SingleNestedAttribute{
									Description: "A collection containing a link ( `href` ) to the Web API endpoint where full details of the playlist's tracks can be retrieved, along with the `total` number of tracks in the playlist. Note, a track object may be `null`. This can happen if a track is no longer available.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchPlaylistsItemsTracksDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"href": schema.StringAttribute{
											Description: "A link to the Web API endpoint where full details of the playlist's tracks can be retrieved.",
											Computed:    true,
										},
										"total": schema.Int64Attribute{
											Description: "Number of tracks in the playlist.",
											Computed:    true,
										},
									},
								},
								"type": schema.StringAttribute{
									Description: `The object type: "playlist"`,
									Computed:    true,
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the playlist.",
									Computed:    true,
								},
							},
						},
					},
					"limit": schema.Int64Attribute{
						Description: "The maximum number of items in the response (as set in the query or by default).",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "URL to the next page of items. ( `null` if none)",
						Computed:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "The offset of the items returned (as set in the query or by default)",
						Computed:    true,
					},
					"previous": schema.StringAttribute{
						Description: "URL to the previous page of items. ( `null` if none)",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of items available to return.",
						Computed:    true,
					},
				},
			},
			"shows": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[SearchShowsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
					},
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[SearchShowsItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the show.",
									Computed:    true,
								},
								"available_markets": schema.ListAttribute{
									Description: "A list of the countries in which the show can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"copyrights": schema.ListNestedAttribute{
									Description: "The copyright statements of the show.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchShowsItemsCopyrightsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"text": schema.StringAttribute{
												Description: "The copyright text for this content.",
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "The type of copyright: `C` = the copyright, `P` = the sound recording (performance) copyright.",
												Computed:    true,
											},
										},
									},
								},
								"description": schema.StringAttribute{
									Description: "A description of the show. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed.",
									Computed:    true,
								},
								"explicit": schema.BoolAttribute{
									Description: "Whether or not the show has explicit content (true = yes it does; false = no it does not OR unknown).",
									Computed:    true,
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "External URLs for this show.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchShowsItemsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the show.",
									Computed:    true,
								},
								"html_description": schema.StringAttribute{
									Description: "A description of the show. This field may contain HTML tags.",
									Computed:    true,
								},
								"images": schema.ListNestedAttribute{
									Description: "The cover art for the show in various sizes, widest first.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchShowsItemsImagesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"height": schema.Int64Attribute{
												Description: "The image height in pixels.",
												Computed:    true,
											},
											"url": schema.StringAttribute{
												Description: "The source URL of the image.",
												Computed:    true,
											},
											"width": schema.Int64Attribute{
												Description: "The image width in pixels.",
												Computed:    true,
											},
										},
									},
								},
								"is_externally_hosted": schema.BoolAttribute{
									Description: "True if all of the shows episodes are hosted outside of Spotify's CDN. This field might be `null` in some cases.",
									Computed:    true,
								},
								"languages": schema.ListAttribute{
									Description: "A list of the languages used in the show, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"media_type": schema.StringAttribute{
									Description: "The media type of the show.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the episode.",
									Computed:    true,
								},
								"publisher": schema.StringAttribute{
									Description: "The publisher of the show.",
									Computed:    true,
								},
								"total_episodes": schema.Int64Attribute{
									Description: "The total number of episodes in the show.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The object type.\nAvailable values: \"show\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("show"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the show.",
									Computed:    true,
								},
							},
						},
					},
					"limit": schema.Int64Attribute{
						Description: "The maximum number of items in the response (as set in the query or by default).",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "URL to the next page of items. ( `null` if none)",
						Computed:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "The offset of the items returned (as set in the query or by default)",
						Computed:    true,
					},
					"previous": schema.StringAttribute{
						Description: "URL to the previous page of items. ( `null` if none)",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of items available to return.",
						Computed:    true,
					},
				},
			},
			"tracks": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[SearchTracksDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"href": schema.StringAttribute{
						Description: "A link to the Web API endpoint returning the full result of the request",
						Computed:    true,
					},
					"items": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[SearchTracksItemsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
									Computed:    true,
								},
								"album": schema.SingleNestedAttribute{
									Description: "The album on which the track appears. The album object includes a link in `href` to full information about the album.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchTracksItemsAlbumDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the album.",
											Computed:    true,
										},
										"album_type": schema.StringAttribute{
											Description: "The type of the album.\nAvailable values: \"album\", \"single\", \"compilation\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"album",
													"single",
													"compilation",
												),
											},
										},
										"artists": schema.ListNestedAttribute{
											Description: "The artists of the album. Each artist object includes a link in `href` to more detailed information about the artist.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[SearchTracksItemsAlbumArtistsDataSourceModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
														Computed:    true,
													},
													"external_urls": schema.SingleNestedAttribute{
														Description: "Known external URLs for this artist.",
														Computed:    true,
														CustomType:  customfield.NewNestedObjectType[SearchTracksItemsAlbumArtistsExternalURLsDataSourceModel](ctx),
														Attributes: map[string]schema.Attribute{
															"spotify": schema.StringAttribute{
																Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
																Computed:    true,
															},
														},
													},
													"href": schema.StringAttribute{
														Description: "A link to the Web API endpoint providing full details of the artist.",
														Computed:    true,
													},
													"name": schema.StringAttribute{
														Description: "The name of the artist.",
														Computed:    true,
													},
													"type": schema.StringAttribute{
														Description: "The object type.\nAvailable values: \"artist\".",
														Computed:    true,
														Validators: []validator.String{
															stringvalidator.OneOfCaseInsensitive("artist"),
														},
													},
													"uri": schema.StringAttribute{
														Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
														Computed:    true,
													},
												},
											},
										},
										"available_markets": schema.ListAttribute{
											Description: "The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._",
											Computed:    true,
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
										"external_urls": schema.SingleNestedAttribute{
											Description: "Known external URLs for this album.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[SearchTracksItemsAlbumExternalURLsDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"spotify": schema.StringAttribute{
													Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
													Computed:    true,
												},
											},
										},
										"href": schema.StringAttribute{
											Description: "A link to the Web API endpoint providing full details of the album.",
											Computed:    true,
										},
										"images": schema.ListNestedAttribute{
											Description: "The cover art for the album in various sizes, widest first.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectListType[SearchTracksItemsAlbumImagesDataSourceModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"height": schema.Int64Attribute{
														Description: "The image height in pixels.",
														Computed:    true,
													},
													"url": schema.StringAttribute{
														Description: "The source URL of the image.",
														Computed:    true,
													},
													"width": schema.Int64Attribute{
														Description: "The image width in pixels.",
														Computed:    true,
													},
												},
											},
										},
										"name": schema.StringAttribute{
											Description: "The name of the album. In case of an album takedown, the value may be an empty string.",
											Computed:    true,
										},
										"release_date": schema.StringAttribute{
											Description: "The date the album was first released.",
											Computed:    true,
										},
										"release_date_precision": schema.StringAttribute{
											Description: "The precision with which `release_date` value is known.\nAvailable values: \"year\", \"month\", \"day\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"year",
													"month",
													"day",
												),
											},
										},
										"total_tracks": schema.Int64Attribute{
											Description: "The number of tracks in the album.",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: "The object type.\nAvailable values: \"album\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("album"),
											},
										},
										"uri": schema.StringAttribute{
											Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the album.",
											Computed:    true,
										},
										"restrictions": schema.SingleNestedAttribute{
											Description: "Included in the response when a content restriction is applied.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[SearchTracksItemsAlbumRestrictionsDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"reason": schema.StringAttribute{
													Description: "The reason for the restriction. Albums may be restricted if the content is not available in a given market, to the user's subscription type, or when the user's account is set to not play explicit content.\nAdditional reasons may be added in the future.\nAvailable values: \"market\", \"product\", \"explicit\".",
													Computed:    true,
													Validators: []validator.String{
														stringvalidator.OneOfCaseInsensitive(
															"market",
															"product",
															"explicit",
														),
													},
												},
											},
										},
									},
								},
								"artists": schema.ListNestedAttribute{
									Description: "The artists who performed the track. Each artist object includes a link in `href` to more detailed information about the artist.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[SearchTracksItemsArtistsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
												Computed:    true,
											},
											"external_urls": schema.SingleNestedAttribute{
												Description: "Known external URLs for this artist.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[SearchTracksItemsArtistsExternalURLsDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"spotify": schema.StringAttribute{
														Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
														Computed:    true,
													},
												},
											},
											"href": schema.StringAttribute{
												Description: "A link to the Web API endpoint providing full details of the artist.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the artist.",
												Computed:    true,
											},
											"type": schema.StringAttribute{
												Description: "The object type.\nAvailable values: \"artist\".",
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("artist"),
												},
											},
											"uri": schema.StringAttribute{
												Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the artist.",
												Computed:    true,
											},
										},
									},
								},
								"available_markets": schema.ListAttribute{
									Description: "A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"disc_number": schema.Int64Attribute{
									Description: "The disc number (usually `1` unless the album consists of more than one disc).",
									Computed:    true,
								},
								"duration_ms": schema.Int64Attribute{
									Description: "The track length in milliseconds.",
									Computed:    true,
								},
								"explicit": schema.BoolAttribute{
									Description: "Whether or not the track has explicit lyrics ( `true` = yes it does; `false` = no it does not OR unknown).",
									Computed:    true,
								},
								"external_ids": schema.SingleNestedAttribute{
									Description: "Known external IDs for the track.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchTracksItemsExternalIDsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"ean": schema.StringAttribute{
											Description: "[International Article Number](http://en.wikipedia.org/wiki/International_Article_Number_%28EAN%29)",
											Computed:    true,
										},
										"isrc": schema.StringAttribute{
											Description: "[International Standard Recording Code](http://en.wikipedia.org/wiki/International_Standard_Recording_Code)",
											Computed:    true,
										},
										"upc": schema.StringAttribute{
											Description: "[Universal Product Code](http://en.wikipedia.org/wiki/Universal_Product_Code)",
											Computed:    true,
										},
									},
								},
								"external_urls": schema.SingleNestedAttribute{
									Description: "Known external URLs for this track.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchTracksItemsExternalURLsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"spotify": schema.StringAttribute{
											Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
											Computed:    true,
										},
									},
								},
								"href": schema.StringAttribute{
									Description: "A link to the Web API endpoint providing full details of the track.",
									Computed:    true,
								},
								"is_local": schema.BoolAttribute{
									Description: "Whether or not the track is from a local file.",
									Computed:    true,
								},
								"is_playable": schema.BoolAttribute{
									Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied. If `true`, the track is playable in the given market. Otherwise `false`.",
									Computed:    true,
								},
								"linked_from": schema.SingleNestedAttribute{
									Description: "Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied, and the requested track has been replaced with different track. The track in the `linked_from` object contains information about the originally requested track.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchTracksItemsLinkedFromDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
											Computed:    true,
										},
										"external_urls": schema.SingleNestedAttribute{
											Description: "Known external URLs for this track.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[SearchTracksItemsLinkedFromExternalURLsDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"spotify": schema.StringAttribute{
													Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
													Computed:    true,
												},
											},
										},
										"href": schema.StringAttribute{
											Description: "A link to the Web API endpoint providing full details of the track.",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: `The object type: "track".`,
											Computed:    true,
										},
										"uri": schema.StringAttribute{
											Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
											Computed:    true,
										},
									},
								},
								"name": schema.StringAttribute{
									Description: "The name of the track.",
									Computed:    true,
								},
								"popularity": schema.Int64Attribute{
									Description: "The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.<br/>The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.<br/>Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._",
									Computed:    true,
								},
								"preview_url": schema.StringAttribute{
									Description:        "A link to a 30 second preview (MP3 format) of the track. Can be `null`",
									Computed:           true,
									DeprecationMessage: "This attribute is deprecated.",
								},
								"restrictions": schema.SingleNestedAttribute{
									Description: "Included in the response when a content restriction is applied.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[SearchTracksItemsRestrictionsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"reason": schema.StringAttribute{
											Description: "The reason for the restriction. Supported values:\n- `market` - The content item is not available in the given market.\n- `product` - The content item is not available for the user's subscription type.\n- `explicit` - The content item is explicit and the user's account is set to not play explicit content.\n\nAdditional reasons may be added in the future.\n**Note**: If you use this field, make sure that your application safely handles unknown values.",
											Computed:    true,
										},
									},
								},
								"track_number": schema.Int64Attribute{
									Description: "The number of the track. If an album has several discs, the track number is the number on the specified disc.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The object type: \"track\".\nAvailable values: \"track\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("track"),
									},
								},
								"uri": schema.StringAttribute{
									Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.",
									Computed:    true,
								},
							},
						},
					},
					"limit": schema.Int64Attribute{
						Description: "The maximum number of items in the response (as set in the query or by default).",
						Computed:    true,
					},
					"next": schema.StringAttribute{
						Description: "URL to the next page of items. ( `null` if none)",
						Computed:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "The offset of the items returned (as set in the query or by default)",
						Computed:    true,
					},
					"previous": schema.StringAttribute{
						Description: "URL to the previous page of items. ( `null` if none)",
						Computed:    true,
					},
					"total": schema.Int64Attribute{
						Description: "The total number of items available to return.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *SearchDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *SearchDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
