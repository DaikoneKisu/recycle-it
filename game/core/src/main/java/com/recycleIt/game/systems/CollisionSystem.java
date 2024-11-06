package com.recycleIt.game.systems;

import com.badlogic.ashley.core.ComponentMapper;
import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.Family;
import com.badlogic.ashley.systems.IteratingSystem;

import com.recycleIt.game.components.CollisionComponent;
import com.recycleIt.game.components.PlayerComponent;
import com.recycleIt.game.components.TypeComponent;

public class CollisionSystem extends IteratingSystem {
  ComponentMapper<CollisionComponent> cm;
  ComponentMapper<PlayerComponent> pm;

  public CollisionSystem() {
    // only need to worry about player collisions
    super(Family.all(CollisionComponent.class, PlayerComponent.class).get());

    cm = ComponentMapper.getFor(CollisionComponent.class);
    pm = ComponentMapper.getFor(PlayerComponent.class);
  }

  @Override
  protected void processEntity(Entity entity, float deltaTime) {
    // get player collision component
    CollisionComponent cc = cm.get(entity);

    Entity collidedEntity = cc.collisionEntity;
    if (collidedEntity != null) {
      TypeComponent tc = collidedEntity.getComponent(TypeComponent.class);
      if (tc != null) {
        switch (tc.type) {
          case TypeComponent.GUEST:
            // do host hit enemy thing
            System.out.println("host hit guest");
            break;
          case TypeComponent.SCENERY:
            // do host hit scenery thing
            System.out.println("host hit scenery");
            break;
          case TypeComponent.OTHER:
            // do host hit other thing
            System.out.println("host hit other");
            break; // technically this isn't needed
        }
        cc.collisionEntity = null; // collision handled reset component
      }
    }
  }
}