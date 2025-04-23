components {
  id: "soundController"
  component: "/main/game/music/soundController.script"
}
embedded_components {
  id: "bgMusic"
  type: "sound"
  data: "sound: \"/assets/sounds/gameplay.ogg\"\n"
  "looping: 1\n"
  "group: \"music\"\n"
  "gain: 0.8\n"
  ""
}
embedded_components {
  id: "miss"
  type: "sound"
  data: "sound: \"/assets/sounds/coinhit.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 1.5\n"
  ""
}
embedded_components {
  id: "coinCollect"
  type: "sound"
  data: "sound: \"/assets/sounds/coin.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 0.5\n"
  ""
}
embedded_components {
  id: "crash"
  type: "sound"
  data: "sound: \"/assets/sounds/hit.wav\"\n"
  "group: \"sfx\"\n"
  ""
}
embedded_components {
  id: "shieldCrash"
  type: "sound"
  data: "sound: \"/assets/sounds/laser2.wav\"\n"
  "group: \"sfx\"\n"
  ""
}
embedded_components {
  id: "laser"
  type: "sound"
  data: "sound: \"/assets/sounds/laser.wav\"\n"
  "group: \"sfx\"\n"
  ""
}
embedded_components {
  id: "magnet_activate"
  type: "sound"
  data: "sound: \"/assets/sounds/powerup/electric_sparks.wav\"\n"
  "group: \"sfx\"\n"
  ""
}
embedded_components {
  id: "magnet_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/powerup_magnetic_field.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "powerupCollect"
  type: "sound"
  data: "sound: \"/assets/sounds/powerup/pickuppowerup.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "extralife_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/powerup_extra_life.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "rocket_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/powerup_rocket.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "shield_activate"
  type: "sound"
  data: "sound: \"/assets/sounds/powerup/shield.wav\"\n"
  "group: \"sfx\"\n"
  ""
}
embedded_components {
  id: "shield_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/powerup_shields.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "coin_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/event_coin_shower.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "invasion_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/event_invasion.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "meteor_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/event_meteor_shower.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "peaceful_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/peaceful_sector.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 3.0\n"
  ""
}
embedded_components {
  id: "rocket"
  type: "sound"
  data: "sound: \"/assets/sounds/rocket.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 0.6\n"
  ""
}
embedded_components {
  id: "countDown_vo"
  type: "sound"
  data: "sound: \"/assets/sounds/VO/countdown.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 4.0\n"
  ""
}
embedded_components {
  id: "menuMusic"
  type: "sound"
  data: "sound: \"/assets/sounds/menu.ogg\"\n"
  "looping: 1\n"
  "group: \"music\"\n"
  "gain: 1.3\n"
  ""
}
embedded_components {
  id: "buy"
  type: "sound"
  data: "sound: \"/assets/sounds/buy.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 1.4\n"
  ""
}
embedded_components {
  id: "lock"
  type: "sound"
  data: "sound: \"/assets/sounds/lock.wav\"\n"
  "group: \"sfx\"\n"
  ""
}
embedded_components {
  id: "tick"
  type: "sound"
  data: "sound: \"/assets/sounds/tick.wav\"\n"
  "group: \"sfx\"\n"
  "gain: 2.4\n"
  ""
}
