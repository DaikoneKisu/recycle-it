package com.recycleIt.game.systems;

import com.badlogic.ashley.core.ComponentMapper;
import com.badlogic.ashley.core.Entity;
import com.recycleIt.game.components.TransformComponent;

import java.util.Comparator;

public class ZComparator implements Comparator<Entity> {
  private ComponentMapper<TransformComponent> cmTrans;

  public ZComparator() {
    cmTrans = ComponentMapper.getFor(TransformComponent.class);
  }

  @Override
  public int compare(Entity entityA, Entity entityB) {
    float az = cmTrans.get(entityA).position.z;
    float bz = cmTrans.get(entityB).position.z;
    int res = 0;
    if (az > bz) {
      res = 1;
    } else if (az < bz) {
      res = -1;
    }
    return res;
  }
}