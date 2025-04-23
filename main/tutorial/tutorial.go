components {
  id: "tutorial"
  component: "/main/tutorial/tutorial.script"
}
embedded_components {
  id: "mobile"
  type: "sprite"
  data: "default_animation: \"toLeft\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/tutorial/tutorial.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.4
    y: 0.4
  }
}
embedded_components {
  id: "keyLeft"
  type: "sprite"
  data: "default_animation: \"A_stop\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/tutorial/tutorial.atlas\"\n"
  "}\n"
  ""
  position {
    x: -75.0
    y: 168.0
  }
  scale {
    x: 0.3
    y: 0.3
  }
}
embedded_components {
  id: "keyRight"
  type: "sprite"
  data: "default_animation: \"D_stop\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 106.0\n"
  "  y: 122.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/tutorial/tutorial.atlas\"\n"
  "}\n"
  ""
  position {
    x: 71.0
    y: 168.0
  }
  scale {
    x: 0.3
    y: 0.3
  }
}
