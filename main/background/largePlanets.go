components {
  id: "planets"
  component: "/main/background/planets.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"largeplanets\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/background/background.atlas\"\n"
  "}\n"
  ""
}
