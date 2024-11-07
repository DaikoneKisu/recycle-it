package com.recycleIt.game.systems;

import com.badlogic.ashley.core.ComponentMapper;
import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.Family;
import com.badlogic.ashley.systems.IteratingSystem;
import com.recycleIt.game.components.B2dBodyComponent;
import com.recycleIt.game.components.PlayerComponent;
import com.recycleIt.game.components.StateComponent;
import com.recycleIt.game.components.TypeComponent;

public class MultiplayerSystem extends IteratingSystem {
  ComponentMapper<PlayerComponent> pm;
  ComponentMapper<B2dBodyComponent> bodm;
  ComponentMapper<StateComponent> sm;
  ComponentMapper<TypeComponent> tm;

  public MultiplayerSystem() {
    super(Family.all(PlayerComponent.class).get());
    pm = ComponentMapper.getFor(PlayerComponent.class);
    bodm = ComponentMapper.getFor(B2dBodyComponent.class);
    sm = ComponentMapper.getFor(StateComponent.class);
    tm = ComponentMapper.getFor(TypeComponent.class);
  }

  @Override
  protected void processEntity(Entity entity, float deltaTime) {
    B2dBodyComponent b2body = bodm.get(entity);
    StateComponent state = sm.get(entity);
    PlayerComponent player = pm.get(entity);
    TypeComponent type = tm.get(entity);

    if (player.isMe) {
      return;
    }

    if (type.type == TypeComponent.Type.Host) {
      // Host logic
    } else if (type.type == TypeComponent.Type.Guest) {
      // Guest logic
    }
  }
}
