// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browse_category

import (
	"context"

	"github.com/cjavdev/terraform-provider-spotted/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*BrowseCategoriesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"locale": schema.StringAttribute{
				Description: "The desired language, consisting of an [ISO 639-1](http://en.wikipedia.org/wiki/ISO_639-1) language code and an [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2), joined by an underscore. For example: `es_MX`, meaning &quot;Spanish (Mexico)&quot;. Provide this parameter if you want the category strings returned in a particular language.<br/> _**Note**: if `locale` is not supplied, or if the specified language is not available, the category strings returned will be in the Spotify default language (American English)._",
				Optional:    true,
			},
			"limit": schema.Int64Attribute{
				Description: "The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 50),
				},
			},
			"offset": schema.Int64Attribute{
				Description: "The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.",
				Computed:    true,
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[BrowseCategoriesItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) of the category.",
							Computed:    true,
						},
						"href": schema.StringAttribute{
							Description: "A link to the Web API endpoint returning full details of the category.",
							Computed:    true,
						},
						"icons": schema.ListNestedAttribute{
							Description: "The category icon, in various sizes.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[BrowseCategoriesIconsDataSourceModel](ctx),
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
									"published": schema.BoolAttribute{
										Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
										Computed:    true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the category.",
							Computed:    true,
						},
						"published": schema.BoolAttribute{
							Description: "The playlist's public/private status (if it should be added to the user's profile or not): `true` the playlist will be public, `false` the playlist will be private, `null` the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *BrowseCategoriesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *BrowseCategoriesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
