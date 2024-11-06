package com.recycleIt.game.controllers;

import java.util.Map;

import com.badlogic.gdx.InputProcessor;
import com.badlogic.gdx.math.Vector2;

public class KeyboardController implements InputProcessor {
  public boolean left, right, up, down;
  public boolean isMouse1Down, isMouse2Down, isMouse3Down;
  public boolean isDragged;
  public Vector2 mousePos;

  public enum Key {
    Left, Right, Up, Down
  }

  private Map<Key, Integer> keyMapping;
  private final Integer LEFT;
  private final Integer RIGHT;
  private final Integer UP;
  private final Integer DOWN;

  public KeyboardController(Map<Key, Integer> keyMapping) {
    this.keyMapping = keyMapping;
    this.LEFT = this.keyMapping.get(Key.Left);
    this.RIGHT = this.keyMapping.get(Key.Right);
    this.UP = this.keyMapping.get(Key.Up);
    this.DOWN = this.keyMapping.get(Key.Down);
    this.mousePos = new Vector2();
  }

  @Override
  public boolean keyDown(int keycode) {
    boolean isAnyMappedKeyDown = false;

    if (keycode == LEFT) {
      left = true;
      isAnyMappedKeyDown = true;
    }

    if (keycode == RIGHT) {
      right = true;
      isAnyMappedKeyDown = true;
    }

    if (keycode == UP) {
      up = true;
      isAnyMappedKeyDown = true;
    }

    if (keycode == DOWN) {
      down = true;
      isAnyMappedKeyDown = true;
    }

    return isAnyMappedKeyDown;
  }

  @Override
  public boolean keyUp(int keycode) {
    boolean isAnyMappedKeyUp = false;

    if (keycode == LEFT) {
      left = false;
      isAnyMappedKeyUp = true;
    }

    if (keycode == RIGHT) {
      right = false;
      isAnyMappedKeyUp = true;
    }

    if (keycode == UP) {
      up = false;
      isAnyMappedKeyUp = true;
    }

    if (keycode == DOWN) {
      down = false;
      isAnyMappedKeyUp = true;
    }

    return isAnyMappedKeyUp;
  }

  @Override
  public boolean keyTyped(char character) {
    return false;
  }

  @Override
  public boolean touchDown(int screenX, int screenY, int pointer, int button) {
    boolean isAnyMappedTouchDown = false;

    if (button == 0) {
      isMouse1Down = true;
      isAnyMappedTouchDown = true;
    }

    if (button == 1) {
      isMouse2Down = true;
      isAnyMappedTouchDown = true;
    }

    if (button == 2) {
      isMouse3Down = true;
      isAnyMappedTouchDown = true;
    }

    mousePos.set(screenX, screenY);

    return isAnyMappedTouchDown;
  }

  @Override
  public boolean touchUp(int screenX, int screenY, int pointer, int button) {
    isDragged = false;

    boolean isAnyMappedTouchUp = false;

    if (button == 0) {
      isMouse1Down = false;
      isAnyMappedTouchUp = true;
    }

    if (button == 1) {
      isMouse2Down = false;
      isAnyMappedTouchUp = true;
    }

    if (button == 2) {
      isMouse3Down = false;
      isAnyMappedTouchUp = true;
    }

    mousePos.set(screenX, screenY);

    return isAnyMappedTouchUp;
  }

  public boolean touchCancelled(int screenX, int screenY, int pointer, int button) {
    return false;
  }

  @Override
  public boolean touchDragged(int screenX, int screenY, int pointer) {
    isDragged = true;

    mousePos.set(screenX, screenY);

    return isDragged;
  }

  @Override
  public boolean mouseMoved(int screenX, int screenY) {
    mousePos.set(screenX, screenY);

    return true;
  }

  @Override
  public boolean scrolled(float amountX, float amountY) {
    return false;
  }
}