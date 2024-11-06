package com.recycleIt.game;

import com.badlogic.gdx.Game;
import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.graphics.g2d.BitmapFont;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;
import com.badlogic.gdx.utils.viewport.FitViewport;
import com.recycleIt.game.core.ScreenController;
import com.recycleIt.game.core.ScreenController.Screen;

/**
 * {@link com.badlogic.gdx.ApplicationListener} implementation shared by all
 * platforms.
 */
public class RecycleIt extends Game {

  public ShapeRenderer shapeRenderer;
  public BitmapFont font;
  public FitViewport viewport;
  public ScreenController screenController;

  public final float WORLD_WIDTH = 800;
  public final float WORLD_HEIGHT = 600;

  @Override
  public void create() {
    this.shapeRenderer = new ShapeRenderer();
    this.font = new BitmapFont(); // default libGDX's font (arial)
    this.viewport = new FitViewport(WORLD_WIDTH, WORLD_HEIGHT);
    this.screenController = ScreenController.getInstace(this);

    this.font.setUseIntegerPositions(false);
    this.font.getData().setScale(viewport.getWorldHeight() / Gdx.graphics.getHeight());

    this.screenController.register(Screen.Game);
    this.screenController.register(Screen.MainMenu);

    this.screenController.show(Screen.MainMenu);
  }

  @Override
  public void render() {
    super.render();
  }

  @Override
  public void dispose() {
    this.shapeRenderer.dispose();
    this.font.dispose();
  }
}
