package com.recycleIt.game;

import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.graphics.Color;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;

public class Ball {
  int x;
  int y;
  int radius;
  int ySpeed;
  int xSpeed;
  Color color = Color.WHITE;

  public Ball(int x, int y, int radius, int xSpeed, int ySpeed) {
    this.x = x;
    this.y = y;
    this.radius = radius;
    this.xSpeed = xSpeed;
    this.ySpeed = ySpeed;
  }

  public void update() {
    x += xSpeed;
    y += ySpeed;
    if (x <= radius || x >= Gdx.graphics.getWidth() - radius) {
      xSpeed = -xSpeed;
    }
    if (y <= radius || y >= Gdx.graphics.getHeight() - radius) {
      ySpeed = -ySpeed;
    }
  }

  public void draw(ShapeRenderer shape, Paddle paddle) {
    this.checkCollision(paddle);
    shape.setColor(color);
    shape.circle(x, y, radius);
  }

  public void checkCollision(Paddle paddle) {
    if (collidesWith(paddle)) {
      color = Color.GREEN;
      ySpeed = -ySpeed;
    } else {
      color = Color.WHITE;
    }
  }

  private boolean collidesWith(Paddle paddle) {
    return (paddle.x + paddle.width >= this.x - this.radius &&
        paddle.x <= this.x + this.radius &&
        paddle.y + paddle.height >= this.y - this.radius &&
        paddle.y <= this.y + this.radius);
  }
}