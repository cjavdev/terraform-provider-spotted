// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package me

import (
	"context"

	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*MeDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"country": schema.StringAttribute{
				Description: "The country of the user, as set in the user's account profile. An [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _This field is only available when the current user has granted access to the [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes) scope._",
				Computed:    true,
			},
			"display_name": schema.StringAttribute{
				Description: "The name displayed on the user's profile. `null` if not available.",
				Computed:    true,
			},
			"email": schema.StringAttribute{
				Description: "The user's email address, as entered by the user when creating their account. _**Important!** This email address is unverified; there is no proof that it actually belongs to the user._ _This field is only available when the current user has granted access to the [user-read-email](/documentation/web-api/concepts/scopes/#list-of-scopes) scope._",
				Computed:    true,
			},
			"href": schema.StringAttribute{
				Description: "A link to the Web API endpoint for this user.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for the user.",
				Computed:    true,
			},
			"product": schema.StringAttribute{
				Description: `The user's Spotify subscription level: "premium", "free", etc. (The subscription level "open" can be considered the same as "free".) _This field is only available when the current user has granted access to the [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes) scope._`,
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: `The object type: "user"`,
				Computed:    true,
			},
			"uri": schema.StringAttribute{
				Description: "The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the user.",
				Computed:    true,
			},
			"explicit_content": schema.SingleNestedAttribute{
				Description: "The user's explicit content settings. _This field is only available when the current user has granted access to the [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes) scope._",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[MeExplicitContentDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"filter_enabled": schema.BoolAttribute{
						Description: "When `true`, indicates that explicit content should not be played.",
						Computed:    true,
					},
					"filter_locked": schema.BoolAttribute{
						Description: "When `true`, indicates that the explicit content setting is locked and can't be changed by the user.",
						Computed:    true,
					},
				},
			},
			"external_urls": schema.SingleNestedAttribute{
				Description: "Known external URLs for this user.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[MeExternalURLsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"spotify": schema.StringAttribute{
						Description: "The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.",
						Computed:    true,
					},
				},
			},
			"followers": schema.SingleNestedAttribute{
				Description: "Information about the followers of the user.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[MeFollowersDataSourceModel](ctx),
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
			"images": schema.ListNestedAttribute{
				Description: "The user's profile image.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MeImagesDataSourceModel](ctx),
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
		},
	}
}

func (d *MeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MeDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
