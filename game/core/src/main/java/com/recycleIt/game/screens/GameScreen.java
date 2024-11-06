package com.recycleIt.game.screens;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Random;

import com.badlogic.ashley.core.Engine;
import com.badlogic.ashley.core.Entity;
import com.badlogic.ashley.core.PooledEngine;
import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.Input.Keys;
import com.badlogic.gdx.graphics.Color;
import com.badlogic.gdx.graphics.OrthographicCamera;
import com.badlogic.gdx.graphics.g2d.SpriteBatch;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;
import com.badlogic.gdx.math.Vector2;
import com.badlogic.gdx.physics.box2d.BodyDef.BodyType;
import com.badlogic.gdx.physics.box2d.World;
import com.badlogic.gdx.utils.ScreenUtils;
import com.badlogic.gdx.utils.viewport.ScreenViewport;

import com.recycleIt.game.Ball;
import com.recycleIt.game.Paddle;
import com.recycleIt.game.RecycleIt;
import com.recycleIt.game.components.CollisionComponent;
import com.recycleIt.game.components.PlayerComponent;
import com.recycleIt.game.components.B2dBodyComponent;
import com.recycleIt.game.components.StateComponent;
import com.recycleIt.game.components.TextureComponent;
import com.recycleIt.game.components.TransformComponent;
import com.recycleIt.game.components.TypeComponent;
import com.recycleIt.game.controllers.KeyboardController;
import com.recycleIt.game.core.AbstractScreen;
import com.recycleIt.game.core.B2dContactListener;
import com.recycleIt.game.core.BodyFactory;
import com.recycleIt.game.core.BodyFactory.BodyMaterial;
import com.recycleIt.game.systems.AnimationSystem;
import com.recycleIt.game.systems.CollisionSystem;
import com.recycleIt.game.systems.PhysicsDebugSystem;
import com.recycleIt.game.systems.PhysicsSystem;
import com.recycleIt.game.systems.PlayerControlSystem;
import com.recycleIt.game.systems.RenderingSystem;

public class GameScreen extends AbstractScreen {
  private World world;
  private KeyboardController controller;
  private BodyFactory bodyFactory;
  private SpriteBatch sb;
  private OrthographicCamera camera;
  private Engine engine;

  private ArrayList<Ball> balls = new ArrayList<>();
  private Random randomGenerator = new Random();
  private Paddle paddle;
  private ScreenViewport viewport;

  public GameScreen(RecycleIt game) {
    super(game);
  }

  @Override
  public void show() {
    Map<KeyboardController.Key, Integer> playerKeyboardMapping = new HashMap<KeyboardController.Key, Integer>();

    playerKeyboardMapping.put(KeyboardController.Key.Left, Keys.LEFT);
    playerKeyboardMapping.put(KeyboardController.Key.Right, Keys.RIGHT);
    playerKeyboardMapping.put(KeyboardController.Key.Up, Keys.UP);
    playerKeyboardMapping.put(KeyboardController.Key.Down, Keys.DOWN);

    this.controller = new KeyboardController(playerKeyboardMapping);

    Gdx.input.setInputProcessor(this.controller);

    world = new World(new Vector2(0, -10f), true);
    world.setGravity(new Vector2(0, 0));
    world.setContactListener(new B2dContactListener());
    bodyFactory = BodyFactory.getInstance(world);

    sb = new SpriteBatch();

    RenderingSystem renderingSystem = new RenderingSystem(this.GAME.spriteBatch);
    camera = renderingSystem.getCamera();
    // sb.setProjectionMatrix(camera.combined);

    this.engine = new PooledEngine();

    engine.addSystem(new AnimationSystem());
    engine.addSystem(renderingSystem);
    engine.addSystem(new PhysicsSystem(world));
    engine.addSystem(new PhysicsDebugSystem(world, camera));
    engine.addSystem(new CollisionSystem());
    engine.addSystem(new PlayerControlSystem(controller));

    // create some game objects
    createPlayer();

    createLimits();

    for (int i = 0; i < 1; i++) {
      balls.add(new Ball(Gdx.graphics.getWidth() / 2,
          Gdx.graphics.getHeight() / 2,
          randomGenerator.nextInt(15) + 10, 3, 3));
    }

    var initialPaddleX = Gdx.graphics.getWidth() / 2;
    var initialPaddleY = 50;

    // this.paddle = new Paddle(100, 10, initialPaddleX, initialPaddleY,
    // this.controller);
    // createPlayer();
  }

  @Override
  public void render(float delta) {
    ScreenUtils.clear(Color.BLACK);

    engine.update(delta);
  }

  private void logic() {
    this.paddle.update();

    for (Ball ball : this.balls) {
      ball.update();
    }
  }

  private void draw() {
    var shapeRenderer = this.GAME.shapeRenderer;

    ScreenUtils.clear(Color.BLACK);

    shapeRenderer.begin(ShapeRenderer.ShapeType.Filled);

    paddle.draw(shapeRenderer);

    for (Ball ball : balls) {
      ball.draw(shapeRenderer, paddle);
    }

    shapeRenderer.end();
  }

  private void createPlayer() {
    // Create the Entity and all the components that will go in the entity
    Entity entity = engine.createEntity();
    B2dBodyComponent b2dbody = engine.createComponent(B2dBodyComponent.class);
    TransformComponent position = engine.createComponent(TransformComponent.class);
    TextureComponent texture = engine.createComponent(TextureComponent.class);
    PlayerComponent player = engine.createComponent(PlayerComponent.class);
    CollisionComponent colComp = engine.createComponent(CollisionComponent.class);
    TypeComponent type = engine.createComponent(TypeComponent.class);
    StateComponent stateCom = engine.createComponent(StateComponent.class);

    // create the data for the components and add them to the components
    b2dbody.body = bodyFactory.makeBoxPolyBody(2, 2, 3, 0.2f,
        BodyMaterial.Stone, BodyType.DynamicBody, true);
    // set object position (x,y,z) z used to define draw order 0 first drawn
    // position.position.set(10, 10, 0);
    // texture.region = atlas.findRegion("player");
    type.type = TypeComponent.HOST;
    stateCom.set(StateComponent.STATE_IDLE);
    b2dbody.body.setUserData(entity);

    // add the components to the entity
    entity.add(b2dbody);
    entity.add(position);
    entity.add(texture);
    entity.add(player);
    entity.add(colComp);
    entity.add(type);
    entity.add(stateCom);

    // add the entity to the engine
    engine.addEntity(entity);
  }

  private void createPlatform(float x, float y) {
    Entity entity = engine.createEntity();
    B2dBodyComponent b2dbody = engine.createComponent(B2dBodyComponent.class);
    b2dbody.body = bodyFactory.makeBoxPolyBody(x, y, 3, 0.2f, BodyMaterial.Stone, BodyType.StaticBody);
    TextureComponent texture = engine.createComponent(TextureComponent.class);
    // texture.region = atlas.findRegion("player");
    TypeComponent type = engine.createComponent(TypeComponent.class);
    type.type = TypeComponent.SCENERY;
    b2dbody.body.setUserData(entity);

    entity.add(b2dbody);
    entity.add(texture);
    entity.add(type);

    engine.addEntity(entity);

  }

  private void createFloor() {
    Entity entity = engine.createEntity();
    B2dBodyComponent b2dbody = engine.createComponent(B2dBodyComponent.class);
    b2dbody.body = bodyFactory.makeBoxPolyBody(0, 0, 100, 0.2f, BodyMaterial.Stone, BodyType.StaticBody);
    TextureComponent texture = engine.createComponent(TextureComponent.class);
    // texture.region = atlas.findRegion("player");
    TypeComponent type = engine.createComponent(TypeComponent.class);
    type.type = TypeComponent.SCENERY;

    b2dbody.body.setUserData(entity);

    entity.add(b2dbody);
    entity.add(texture);
    entity.add(type);

    engine.addEntity(entity);
  }

  private void createLimits() {
    Entity lowerLimit = engine.createEntity();
    B2dBodyComponent llb2dbody = engine.createComponent(B2dBodyComponent.class);
    llb2dbody.body = bodyFactory.makeBoxPolyBody(0, 0, RenderingSystem.FRUSTUM_WIDTH * 2, 0.2f, BodyMaterial.Stone,
        BodyType.StaticBody);
    TextureComponent lltexture = engine.createComponent(TextureComponent.class);
    // texture.region = atlas.findRegion("player");
    TypeComponent lltype = engine.createComponent(TypeComponent.class);
    lltype.type = TypeComponent.SCENERY;

    llb2dbody.body.setUserData(lowerLimit);

    lowerLimit.add(llb2dbody);
    lowerLimit.add(lltexture);
    lowerLimit.add(lltype);

    engine.addEntity(lowerLimit);

    Entity upperLimit = engine.createEntity();
    B2dBodyComponent ulb2dbody = engine.createComponent(B2dBodyComponent.class);
    ulb2dbody.body = bodyFactory.makeBoxPolyBody(0, RenderingSystem.FRUSTUM_HEIGHT,
        RenderingSystem.FRUSTUM_WIDTH * 2, 0.2f,
        BodyMaterial.Stone, BodyType.StaticBody);
    TextureComponent ultexture = engine.createComponent(TextureComponent.class);
    // texture.region = atlas.findRegion("player");
    TypeComponent ultype = engine.createComponent(TypeComponent.class);
    ultype.type = TypeComponent.SCENERY;

    ulb2dbody.body.setUserData(upperLimit);

    upperLimit.add(ulb2dbody);
    upperLimit.add(ultexture);
    upperLimit.add(ultype);

    engine.addEntity(upperLimit);

    Entity lefmostLimit = engine.createEntity();
    B2dBodyComponent lmlb2dbody = engine.createComponent(B2dBodyComponent.class);
    lmlb2dbody.body = bodyFactory.makeBoxPolyBody(0, 0, 0.2f, RenderingSystem.FRUSTUM_HEIGHT * 2, BodyMaterial.Stone,
        BodyType.StaticBody);
    TextureComponent lmltexture = engine.createComponent(TextureComponent.class);
    // texture.region = atlas.findRegion("player");
    TypeComponent lmltype = engine.createComponent(TypeComponent.class);
    lmltype.type = TypeComponent.SCENERY;

    lmlb2dbody.body.setUserData(lefmostLimit);

    lefmostLimit.add(lmlb2dbody);
    lefmostLimit.add(lmltexture);
    lefmostLimit.add(lmltype);

    engine.addEntity(lefmostLimit);

    Entity rightmostLimit = engine.createEntity();
    B2dBodyComponent rmlb2dbody = engine.createComponent(B2dBodyComponent.class);
    rmlb2dbody.body = bodyFactory.makeBoxPolyBody(RenderingSystem.FRUSTUM_WIDTH, 0, 0.2f,
        RenderingSystem.FRUSTUM_HEIGHT * 2, BodyMaterial.Stone,
        BodyType.StaticBody);
    TextureComponent rmltexture = engine.createComponent(TextureComponent.class);
    // texture.region = atlas.findRegion("player");
    TypeComponent rmltype = engine.createComponent(TypeComponent.class);
    rmltype.type = TypeComponent.SCENERY;

    rmlb2dbody.body.setUserData(rightmostLimit);

    rightmostLimit.add(rmlb2dbody);
    rightmostLimit.add(rmltexture);
    rightmostLimit.add(rmltype);

    engine.addEntity(rightmostLimit);
  }

  @Override
  public void resize(int width, int height) {
    this.GAME.viewport.update(width, height, true);
  }

  @Override
  public void hide() {
  }

  @Override
  public void pause() {
  }

  @Override
  public void resume() {
  }

  @Override
  public void dispose() {
    this.balls.clear();
    this.paddle = null;
  }
}
