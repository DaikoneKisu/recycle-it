package com.recycleIt.game;

import com.badlogic.gdx.Game;
import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.graphics.g2d.BitmapFont;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;
import com.badlogic.gdx.utils.viewport.FitViewport;
import com.recycleIt.game.screens.GameScreen;

/**
 * {@link com.badlogic.gdx.ApplicationListener} implementation shared by all
 * platforms.
 */
public class RecycleIt extends Game {

  public ShapeRenderer shapeRenderer;
  public BitmapFont font;
  public FitViewport viewport;

  public final float WORLD_WIDTH = 800;
  public final float WORLD_HEIGHT = 600;

  @Override
  public void create() {
    shapeRenderer = new ShapeRenderer();

    font = new BitmapFont(); // default libGDX's font (arial)
    viewport = new FitViewport(WORLD_WIDTH, WORLD_HEIGHT);

    font.setUseIntegerPositions(false);
    font.getData().setScale(viewport.getWorldHeight() / Gdx.graphics.getHeight());

    this.setScreen(new GameScreen(this));
  }

  @Override
  public void render() {
    super.render();
  }

  @Override
  public void dispose() {
    shapeRenderer.dispose();
    font.dispose();
  }
}
