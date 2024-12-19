components {
  id: "enemy"
  component: "/main/scripts/enemy_death.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game_images/enemy_1_death.tilesource\"\n"
  "}\n"
  ""
  position {
    y: -4.0
  }
}
