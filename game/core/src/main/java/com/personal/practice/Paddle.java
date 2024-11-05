package com.personal.practice;

import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;

public class Paddle {
  int x;
  int y;
  int initialX;
  int initialY;
  int width;
  int height;

  public Paddle(int width, int height, int initialX, int initialY) {
    this.x = initialX;
    this.y = initialY;
    this.width = width;
    this.height = height;
    this.initialX = initialX;
    this.initialY = initialY;
  }

  public void update() {
    var playerMouseX = Gdx.input.getX();
    if (playerMouseX >= this.width / 2
        && playerMouseX <= Gdx.graphics.getWidth() - this.width / 2) {
      this.x = Gdx.input.getX() - this.width / 2;
    }

    // var playerMouseY = Gdx.input.getY();
    // if (playerMouseY >= this.height / 2
    // && playerMouseY <= Gdx.graphics.getHeight() - this.height / 2) {
    // this.y = -(Gdx.input.getY() - Gdx.graphics.getHeight()) - this.height / 2;
    // }
  }

  public void draw(ShapeRenderer shape) {
    shape.rect(this.x, this.y, this.width, this.height);
  }
}
