package com.recycleIt.game.systems;

import com.badlogic.ashley.core.ComponentMapper;
import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.Family;
import com.badlogic.ashley.systems.IteratingSystem;
import com.badlogic.gdx.math.Polygon;
import com.badlogic.gdx.utils.Array;

import com.recycleIt.game.components.PolygonBodyComponent;
import com.recycleIt.game.components.TransformComponent;

public class PhysicsSystem extends IteratingSystem {
  private static final float MAX_STEP_TIME = 1 / 45f;
  private static float accumulator = 0f;

  private Array<Entity> bodiesQueue;

  private ComponentMapper<PolygonBodyComponent> bm = ComponentMapper.getFor(PolygonBodyComponent.class);
  private ComponentMapper<TransformComponent> tm = ComponentMapper.getFor(TransformComponent.class);

  public PhysicsSystem() {
    super(Family.all(PolygonBodyComponent.class, TransformComponent.class).get());
    this.bodiesQueue = new Array<Entity>();
  }

  @Override
  public void update(float deltaTime) {
    super.update(deltaTime);
    float frameTime = Math.min(deltaTime, 0.25f);
    accumulator += frameTime;

    if (accumulator >= MAX_STEP_TIME) {
      accumulator -= MAX_STEP_TIME;

      // Entity Queue
      for (Entity entity : bodiesQueue) {
        TransformComponent tfm = tm.get(entity);
        PolygonBodyComponent bodyComp = bm.get(entity);
        Polygon body = bodyComp.body;
        tfm.position.x = body.getX();
        tfm.position.y = body.getY();
        tfm.rotation = body.getRotation();
      }
    }
    bodiesQueue.clear();
  }

  @Override
  protected void processEntity(Entity entity, float deltaTime) {
    bodiesQueue.add(entity);
  }
}
