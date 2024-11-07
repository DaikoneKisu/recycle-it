package com.recycleIt.game.systems;

import java.lang.Thread.State;

import com.badlogic.ashley.core.ComponentMapper;
import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.Family;
import com.badlogic.ashley.systems.IteratingSystem;
import com.badlogic.gdx.math.Vector2;
import com.recycleIt.game.components.B2dBodyComponent;
import com.recycleIt.game.components.BallComponent;
import com.recycleIt.game.components.CollisionComponent;
import com.recycleIt.game.components.PlayerComponent;
import com.recycleIt.game.components.StateComponent;
import com.recycleIt.game.components.TypeComponent;

public class CollisionSystem extends IteratingSystem {
  ComponentMapper<CollisionComponent> cm;
  ComponentMapper<PlayerComponent> pm;
  ComponentMapper<BallComponent> bm;

  public CollisionSystem() {
    // only need to worry about players and balls collisions
    super(Family.all(CollisionComponent.class, PlayerComponent.class, BallComponent.class).get());

    cm = ComponentMapper.getFor(CollisionComponent.class);
    pm = ComponentMapper.getFor(PlayerComponent.class);
    bm = ComponentMapper.getFor(BallComponent.class);
  }

  @Override
  protected void processEntity(Entity entity, float deltaTime) {
    // get player/ball collision component
    CollisionComponent cc = cm.get(entity);

    Entity collidedEntity = cc.collisionEntity;
    if (collidedEntity != null) {
      TypeComponent tce = entity.getComponent(TypeComponent.class);

      StringBuilder sb = new StringBuilder();
      switch (tce.type) {
        case Host:
          sb.append("Host ");
          break;
        case Guest:
          sb.append("Guest ");
          break;
        case Ball:
          sb.append("Ball ");
          break;
        default:
          sb.append("Other ");
          break;
      }

      TypeComponent tcce = collidedEntity.getComponent(TypeComponent.class);
      if (tcce != null) {
        switch (tcce.type) {
          case Host:
            // do player hit enemy thing
            sb.append("hit Host");
            break;
          case Guest:
            // do player hit enemy thing
            sb.append("hit Guest");
            break;
          case Scenery:
            // do player hit scenery thing
            sb.append("hit Scenery");
            break;
          case Ball:
            B2dBodyComponent bc = collidedEntity.getComponent(B2dBodyComponent.class);
            StateComponent state = collidedEntity.getComponent(StateComponent.class);

            state.set(StateComponent.STATE_HIT);

            bc.body.setLinearVelocity(new Vector2(-bc.body.getLinearVelocity().x, -bc.body.getLinearVelocity().y));

            state.set(StateComponent.STATE_MOVING);

            sb.append("hit Ball");
            break;
          case Other:
            // do player hit ball thing
            sb.append("hit Other");
            break; // technically this isn't needed
        }
        System.out.println(sb.toString());
        cc.collisionEntity = null; // collision handled reset component
      }
    }
  }
}