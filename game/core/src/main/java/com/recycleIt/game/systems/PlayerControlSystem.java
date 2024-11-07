package com.recycleIt.game.systems;

import com.badlogic.ashley.core.ComponentMapper;
import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.Family;
import com.badlogic.ashley.systems.IteratingSystem;
import com.badlogic.gdx.math.MathUtils;

import com.recycleIt.game.components.PlayerComponent;
import com.recycleIt.game.components.B2dBodyComponent;
import com.recycleIt.game.components.StateComponent;
import com.recycleIt.game.controllers.KeyboardController;

public class PlayerControlSystem extends IteratingSystem {
  ComponentMapper<PlayerComponent> pm;
  ComponentMapper<B2dBodyComponent> bodm;
  ComponentMapper<StateComponent> sm;
  KeyboardController controller;

  private final float PLAYER_VEL_MAX_SPEED = 7f;
  private final float PLAYER_VEL_MIN_SPEED = 0f;
  private final float PLAYER_VEL_DELTA = 0.5f;

  public PlayerControlSystem(KeyboardController keyCon) {
    super(Family.all(PlayerComponent.class).get());
    controller = keyCon;
    pm = ComponentMapper.getFor(PlayerComponent.class);
    bodm = ComponentMapper.getFor(B2dBodyComponent.class);
    sm = ComponentMapper.getFor(StateComponent.class);
  }

  @Override
  protected void processEntity(Entity entity, float deltaTime) {
    B2dBodyComponent b2body = bodm.get(entity);
    StateComponent state = sm.get(entity);
    PlayerComponent player = pm.get(entity);

    if (!player.isMe) {
      return;
    }

    if (b2body.body.getLinearVelocity().len() > 0) {
      state.set(StateComponent.STATE_MOVING);
    }

    if (b2body.body.getLinearVelocity().len() == 0) {
      state.set(StateComponent.STATE_IDLE);
    }

    if (controller.left) {
      b2body.body.setLinearVelocity(
          MathUtils.lerp(b2body.body.getLinearVelocity().x, -PLAYER_VEL_MAX_SPEED, PLAYER_VEL_DELTA),
          b2body.body.getLinearVelocity().y);
    }
    if (controller.right) {
      b2body.body.setLinearVelocity(
          MathUtils.lerp(b2body.body.getLinearVelocity().x, PLAYER_VEL_MAX_SPEED, PLAYER_VEL_DELTA),
          b2body.body.getLinearVelocity().y);
    }

    if (!controller.left && !controller.right) {
      b2body.body.setLinearVelocity(
          MathUtils.lerp(b2body.body.getLinearVelocity().x, PLAYER_VEL_MIN_SPEED, PLAYER_VEL_DELTA),
          b2body.body.getLinearVelocity().y);
    }

    if (controller.up) {
      b2body.body.setLinearVelocity(
          b2body.body.getLinearVelocity().x,
          MathUtils.lerp(b2body.body.getLinearVelocity().y, PLAYER_VEL_MAX_SPEED, PLAYER_VEL_DELTA));
    }

    if (controller.down) {
      b2body.body.setLinearVelocity(
          b2body.body.getLinearVelocity().x,
          MathUtils.lerp(b2body.body.getLinearVelocity().y, -PLAYER_VEL_MAX_SPEED, PLAYER_VEL_DELTA));
    }

    if (!controller.up && !controller.down) {
      b2body.body.setLinearVelocity(
          b2body.body.getLinearVelocity().x,
          MathUtils.lerp(b2body.body.getLinearVelocity().y, PLAYER_VEL_MIN_SPEED, PLAYER_VEL_DELTA));
    }
  }
}
