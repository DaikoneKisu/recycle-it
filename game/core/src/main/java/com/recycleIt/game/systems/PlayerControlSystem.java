package com.recycleIt.game.systems;

import com.badlogic.ashley.core.ComponentMapper;
import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.Family;
import com.badlogic.ashley.systems.IteratingSystem;
import com.badlogic.gdx.math.MathUtils;
import com.badlogic.gdx.math.Vector2;
import com.recycleIt.game.components.PlayerComponent;
import com.recycleIt.game.components.PolygonBodyComponent;
import com.recycleIt.game.components.StateComponent;
import com.recycleIt.game.components.VelocityComponent;
import com.recycleIt.game.controllers.KeyboardController;

public class PlayerControlSystem extends IteratingSystem {
  ComponentMapper<PlayerComponent> pm;
  ComponentMapper<PolygonBodyComponent> bodm;
  ComponentMapper<VelocityComponent> vm;
  ComponentMapper<StateComponent> sm;
  KeyboardController controller;

  private final float PLAYER_VEL_MAX_SPEED = 5f;
  private final float PLAYER_VEL_MIN_SPEED = 0f;
  private final float PLAYER_VEL_DELTA = 0.5f;

  public PlayerControlSystem(KeyboardController keyCon) {
    super(Family.all(PlayerComponent.class).get());
    controller = keyCon;
    pm = ComponentMapper.getFor(PlayerComponent.class);
    bodm = ComponentMapper.getFor(PolygonBodyComponent.class);
    sm = ComponentMapper.getFor(StateComponent.class);
    vm = ComponentMapper.getFor(VelocityComponent.class);
  }

  @Override
  protected void processEntity(Entity entity, float deltaTime) {
    PolygonBodyComponent pbody = bodm.get(entity);
    VelocityComponent vel = vm.get(entity);
    StateComponent state = sm.get(entity);

    if (vel.velocity.len() > 0) {
      state.set(StateComponent.STATE_MOVING);
    }

    if (vel.velocity.len() == 0) {
      state.set(StateComponent.STATE_IDDLE);
    }

    if (controller.left) {
      vel.velocity
          .set(new Vector2(MathUtils.lerp(vel.velocity.x, -PLAYER_VEL_MAX_SPEED, PLAYER_VEL_DELTA), vel.velocity.y));
    }
    if (controller.right) {
      vel.velocity
          .set(new Vector2(MathUtils.lerp(vel.velocity.x, PLAYER_VEL_MAX_SPEED, PLAYER_VEL_DELTA), vel.velocity.y));
    }

    if (!controller.left && !controller.right && state.get() == StateComponent.STATE_MOVING) {
      vel.velocity
          .set(new Vector2(MathUtils.lerp(vel.velocity.x, PLAYER_VEL_MIN_SPEED, PLAYER_VEL_DELTA), vel.velocity.y));
    }

    if (controller.up) {
      vel.velocity
          .set(new Vector2(vel.velocity.x, MathUtils.lerp(vel.velocity.y, PLAYER_VEL_MAX_SPEED, PLAYER_VEL_DELTA)));
    }

    if (controller.down) {
      vel.velocity
          .set(new Vector2(vel.velocity.x, MathUtils.lerp(vel.velocity.y, -PLAYER_VEL_MAX_SPEED, PLAYER_VEL_DELTA)));
    }

    if (!controller.up && !controller.down && state.get() == StateComponent.STATE_MOVING) {
      vel.velocity
          .set(new Vector2(vel.velocity.x, MathUtils.lerp(vel.velocity.x, PLAYER_VEL_MIN_SPEED, PLAYER_VEL_DELTA)));
    }

    // if (controller.up &&
    // (state.get() == StateComponent.STATE_IDDLE || state.get() ==
    // StateComponent.STATE_MOVING)) {
    // // b2body.body.applyForceToCenter(0, 3000,true);
    // pbody.body.applyLinearImpulse(0, 75f,
    // pbody.body.getWorldCenter().x, pbody.body.getWorldCenter().y, true);
    // state.set(StateComponent.STATE_JUMPING);
    // }
  }
}
