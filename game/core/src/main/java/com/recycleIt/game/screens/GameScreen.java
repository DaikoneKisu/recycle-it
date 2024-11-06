package com.recycleIt.game.screens;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Random;

import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.Input.Keys;
import com.badlogic.gdx.graphics.Color;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;
import com.badlogic.gdx.utils.ScreenUtils;
import com.recycleIt.game.Ball;
import com.recycleIt.game.Paddle;
import com.recycleIt.game.RecycleIt;
import com.recycleIt.game.controllers.KeyboardController;
import com.recycleIt.game.core.AbstractScreen;

public class GameScreen extends AbstractScreen {
  public ArrayList<Ball> balls = new ArrayList<>();
  public Random randomGenerator = new Random();
  public Paddle paddle;
  public KeyboardController controller;

  public GameScreen(RecycleIt game) {
    super(game);
  }

  @Override
  public void show() {
    for (int i = 0; i < 1; i++) {
      balls.add(new Ball(Gdx.graphics.getWidth() / 2,
          Gdx.graphics.getHeight() / 2,
          randomGenerator.nextInt(15) + 10, 3, 3));
    }

    var initialPaddleX = Gdx.graphics.getWidth() / 2;
    var initialPaddleY = 50;

    Map<KeyboardController.Key, Integer> playerKeyboardMapping = new HashMap<KeyboardController.Key, Integer>();

    playerKeyboardMapping.put(KeyboardController.Key.Left, Keys.LEFT);
    playerKeyboardMapping.put(KeyboardController.Key.Right, Keys.RIGHT);
    playerKeyboardMapping.put(KeyboardController.Key.Up, Keys.UP);
    playerKeyboardMapping.put(KeyboardController.Key.Down, Keys.DOWN);

    this.controller = new KeyboardController(playerKeyboardMapping);

    Gdx.input.setInputProcessor(this.controller);

    this.paddle = new Paddle(100, 10, initialPaddleX, initialPaddleY, this.controller);

  }

  @Override
  public void render(float delta) {
    logic();
    draw();
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
