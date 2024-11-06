package com.recycleIt.game;

import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;
import com.recycleIt.game.controllers.KeyboardController;

public class Paddle {
  int x;
  int y;
  int initialX;
  int initialY;
  int width;
  int height;
  KeyboardController keyboardController;

  public Paddle(int width, int height, int initialX, int initialY, KeyboardController kc) {
    this.x = initialX;
    this.y = initialY;
    this.width = width;
    this.height = height;
    this.initialX = initialX;
    this.initialY = initialY;
    this.keyboardController = kc;
  }

  public void update() {
    if (keyboardController.left && this.x >= 0) {
      this.x -= 5;
    }

    if (keyboardController.right && this.x <= Gdx.graphics.getWidth() - this.width) {
      this.x += 5;
    }

    // var playerMouseX = Gdx.input.getX();
    // if (playerMouseX >= this.width / 2
    // && playerMouseX <= Gdx.graphics.getWidth() - this.width / 2) {
    // this.x = Gdx.input.getX() - this.width / 2;
    // }

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
