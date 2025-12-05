# Unofficial Spotify Provider

The [Unofficial Spotify provider](https://registry.terraform.io/providers/cjavdev/spotted/latest/docs) provides convenient access to
the [Spotted REST API](https://spotted.stldocs.com) from Terraform.

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    spotted = {
      source  = "cjavdev/spotted"
      version = "~> 0.10.0"
    }
  }
}

# Initialize the provider
provider "spotted" {
  client_id = "My Client ID" # or set SPOTIFY_CLIENT_ID env variable
  client_secret = "My Client Secret" # or set SPOTIFY_CLIENT_SECRET env variable
  access_token = "My Access Token" # or set SPOTIFY_ACCESS_TOKEN env variable
}

# Configure a resource
resource "spotted_user_playlist" "example_user_playlist" {
  user_id = "smedjan"
  name = "New Playlist"
  collaborative = true
  description = "New playlist description"
  public = false
}
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/cjavdev/spotted/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property      | Environment variable    | Required | Default value |
| ------------- | ----------------------- | -------- | ------------- |
| client_secret | `SPOTIFY_CLIENT_SECRET` | false    | —             |
| client_id     | `SPOTIFY_CLIENT_ID`     | false    | —             |
| access_token  | `SPOTIFY_ACCESS_TOKEN`  | false    | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/cjavdev/terraform-provider-spotted/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
