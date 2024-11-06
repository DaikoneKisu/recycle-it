package com.recycleIt.game.core;

import com.badlogic.gdx.Screen;
import com.recycleIt.game.RecycleIt;

public abstract class AbstractScreen implements Screen {
  protected final RecycleIt GAME;

  public AbstractScreen(RecycleIt game) {
    this.GAME = game;
  }

  @Override
  public void show() {

  }

  @Override
  public void render(float delta) {

  }

  @Override
  public void resize(int width, int height) {

  }

  @Override
  public void pause() {

  }

  @Override
  public void resume() {

  }

  @Override
  public void hide() {

  }

  @Override
  public void dispose() {

  }
}
