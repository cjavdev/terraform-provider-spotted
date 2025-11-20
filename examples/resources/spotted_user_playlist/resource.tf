resource "spotted_user_playlist" "example_user_playlist" {
  user_id = "smedjan"
  name = "New Playlist"
  components_schemas_properties_published = true
  collaborative = true
  description = "New playlist description"
}
