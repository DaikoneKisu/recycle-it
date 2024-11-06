package com.recycleIt.game.systems;

import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.Family;
import com.badlogic.ashley.systems.IteratingSystem;
import com.badlogic.gdx.graphics.OrthographicCamera;
import com.badlogic.gdx.physics.box2d.Box2DDebugRenderer;
import com.badlogic.gdx.physics.box2d.World;
import com.recycleIt.game.components.B2dBodyComponent;

public class PhysicsDebugSystem extends IteratingSystem {
  private Box2DDebugRenderer debugRenderer;
  private OrthographicCamera camera;
  private World world;

  public PhysicsDebugSystem(World world, OrthographicCamera camera) {
    super(Family.all(B2dBodyComponent.class).get());
    this.debugRenderer = new Box2DDebugRenderer();
    this.camera = camera;
    this.world = world;
  }

  @Override
  public void update(float deltaTime) {
    super.update(deltaTime);
    debugRenderer.render(world, camera.combined);
  }

  @Override
  protected void processEntity(Entity entity, float deltaTime) {
  }

}
