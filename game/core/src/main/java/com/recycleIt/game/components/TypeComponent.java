package com.recycleIt.game.components;

import com.badlogic.ashley.core.Component;

public class TypeComponent implements Component {
  public enum Type {
    Host, Guest, Scenery, Other, Ball
  }

  public Type type = Type.Other;
}
