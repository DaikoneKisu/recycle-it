package com.recycleIt.game.components;

import com.badlogic.ashley.core.Component;

public class TypeComponent implements Component {
  public static final int HOST = 0;
  public static final int GUEST = 1;
  public static final int SCENERY = 3;
  public static final int OTHER = 4;

  public int type = OTHER;
}
