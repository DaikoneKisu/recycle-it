package com.recycleIt.game.screens;

import java.util.ArrayList;
import java.util.Random;

import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.Screen;
import com.badlogic.gdx.graphics.Color;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;
import com.badlogic.gdx.utils.ScreenUtils;
import com.recycleIt.game.Ball;
import com.recycleIt.game.Paddle;
import com.recycleIt.game.RecycleIt;

public class GameScreen implements Screen {

  private final RecycleIt game;

  public ArrayList<Ball> balls = new ArrayList<>();
  public Random randomGenerator = new Random();
  public Paddle paddle;

  public GameScreen(RecycleIt game) {
    this.game = game;
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

    this.paddle = new Paddle(100, 10, initialPaddleX, initialPaddleY);
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
    var shapeRenderer = this.game.shapeRenderer;

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
    game.viewport.update(width, height, true);
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
