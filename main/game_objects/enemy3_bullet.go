components {
  id: "bullet_sound"
  component: "/game resources/audio/sfx/laserSmall_000.ogg"
}
components {
  id: "enemy3_bullet"
  component: "/main/scripts/enemy3_bullet.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game_images/enemy3_bullets.tilesource\"\n"
  "}\n"
  ""
  position {
    x: -1.0
    y: 6.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy_bullets\"\n"
  "mask: \"players\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 0.5\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 2.5\n"
  "  data: 9.5\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
