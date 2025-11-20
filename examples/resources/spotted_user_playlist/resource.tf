resource "spotted_user_playlist" "example_user_playlist" {
  user_id = "smedjan"
  name = "New Playlist"
  paths_request_body_content_application_json_schema_properties_published = true
  collaborative = true
  description = "New playlist description"
}
