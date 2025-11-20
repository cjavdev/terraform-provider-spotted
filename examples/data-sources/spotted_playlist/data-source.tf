data "spotted_playlist" "example_playlist" {
  playlist_id = "3cEYpjA9oz9GiPac4AsH4n"
  additional_types = "additional_types"
  fields = "items(added_by.id,track(name,href,album(name,href)))"
  market = "ES"
}
