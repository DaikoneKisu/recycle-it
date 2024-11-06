package com.recycleIt.game.systems;

import com.badlogic.ashley.core.ComponentMapper;
import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.Family;
import com.badlogic.ashley.systems.IteratingSystem;
import com.badlogic.gdx.graphics.OrthographicCamera;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;

import com.recycleIt.game.components.PolygonBodyComponent;

public class PhysicsDebugSystem extends IteratingSystem {
  private ShapeRenderer debugRenderer;
  private OrthographicCamera camera;
  private ComponentMapper<PolygonBodyComponent> rm;

  public PhysicsDebugSystem(OrthographicCamera camera) {
    super(Family.all(PolygonBodyComponent.class).get());
    this.debugRenderer = new ShapeRenderer();
    this.rm = ComponentMapper.getFor(PolygonBodyComponent.class);
    this.camera = camera;
  }

  @Override
  public void update(float deltaTime) {
    super.update(deltaTime);

    debugRenderer.setProjectionMatrix(camera.combined);
  }

  @Override
  protected void processEntity(Entity entity, float deltaTime) {
    PolygonBodyComponent rb = rm.get(entity);

    debugRenderer.begin(ShapeRenderer.ShapeType.Line);

    debugRenderer.setColor(1, 0, 0, 1); // Red color for bodies
    debugRenderer.polygon(rb.body.getTransformedVertices());

    debugRenderer.end();
  }

}
