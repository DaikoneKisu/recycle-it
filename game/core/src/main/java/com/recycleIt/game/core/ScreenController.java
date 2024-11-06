package com.recycleIt.game.core;

import com.badlogic.gdx.utils.Array;

import com.recycleIt.game.RecycleIt;
import com.recycleIt.game.screens.GameScreen;
import com.recycleIt.game.screens.MainMenuScreen;

public class ScreenController {
  public enum Screen {
    MainMenu(0), Game(1);

    private int value;

    private Screen(int value) {
      this.value = value;
    }

    public int getValue() {
      return this.value;
    }
  }

  private Array<AbstractScreen> screens = new Array<>();
  private static ScreenController instace;
  private final RecycleIt GAME;

  public static ScreenController getInstace(RecycleIt game) {
    if (ScreenController.instace == null) {
      ScreenController.instace = new ScreenController(game);
    }

    return ScreenController.instace;
  }

  public void register(Screen screen) {
    if (this.contains(screen)) {
      return;
    }

    this.screens.add(this.createScreen(screen));
  }

  public void remove(Screen screen) {
    if (!this.contains(screen)) {
      return;
    }

    AbstractScreen screenToRemove = null;

    for (AbstractScreen lscreen : this.screens) {
      if (this.getScreenClass(screen).isInstance(lscreen)) {
        screenToRemove = lscreen;
        break;
      }
    }

    screenToRemove.dispose();
    this.screens.removeValue(screenToRemove, true);
  }

  public void show(Screen screen) {
    this.GAME.setScreen(this.get(screen));
  }

  public AbstractScreen get(Screen screen) {
    for (AbstractScreen lscreen : this.screens) {
      if (this.getScreenClass(screen).isInstance(lscreen)) {
        return lscreen;
      }
    }

    throw new Error("Screen not registered");
  }

  private ScreenController(RecycleIt game) {
    this.GAME = game;
  }

  private boolean contains(Screen screen) {
    for (AbstractScreen lscreen : this.screens) {
      if (this.getScreenClass(screen).isInstance(lscreen)) {
        return true;
      }
    }

    return false;
  }

  private AbstractScreen createScreen(Screen screen) {
    switch (screen) {
      case MainMenu:
        return new MainMenuScreen(GAME);
      case Game:
        return new GameScreen(GAME);
      default:
        throw new Error("Invalid screen enum type");
    }
  }

  private Class<?> getScreenClass(Screen screen) {
    switch (screen) {
      case MainMenu:
        return MainMenuScreen.class;
      case Game:
        return GameScreen.class;
      default:
        throw new Error("Invalid screen enum type");
    }
  }
}
