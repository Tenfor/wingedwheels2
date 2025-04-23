components {
  id: "magnet"
  component: "/main/game/enemycoins/powerups/magnet.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"magnet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ADD\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
  scale {
    x: 0.4
    y: 0.4
  }
}
