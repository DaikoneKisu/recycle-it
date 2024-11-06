package com.recycleIt.game.components;

import com.badlogic.ashley.core.Component;

public class StateComponent implements Component {
  public static final int STATE_IDDLE = 0;
  public static final int STATE_MOVING = 1;
  public static final int STATE_HIT = 2;

  private int state = 0;
  public float time = 0.0f;
  public boolean isLooping = false;

  public void set(int newState) {
    state = newState;
    time = 0.0f;
  }

  public int get() {
    return state;
  }
}
