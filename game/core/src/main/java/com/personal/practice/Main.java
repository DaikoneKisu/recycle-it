package com.personal.practice;

import java.util.ArrayList;
import java.util.Random;

import com.badlogic.gdx.ApplicationAdapter;
import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.graphics.Color;
import com.badlogic.gdx.graphics.GL20;
import com.badlogic.gdx.graphics.Texture;
import com.badlogic.gdx.graphics.g2d.SpriteBatch;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;
import com.badlogic.gdx.utils.ScreenUtils;

/**
 * {@link com.badlogic.gdx.ApplicationListener} implementation shared by all
 * platforms.
 */
public class Main extends ApplicationAdapter {
	ShapeRenderer shapeRenderer;
	ArrayList<Ball> balls = new ArrayList<>();
	Random randomGenerator = new Random();
	Paddle paddle;

	@Override
	public void create() {
		shapeRenderer = new ShapeRenderer();

		for (int i = 0; i < 1; i++) {
			balls.add(new Ball(Gdx.graphics.getWidth() / 2,
					Gdx.graphics.getHeight() / 2,
					randomGenerator.nextInt(15) + 10, 3, 3));
		}

		var initialPaddleX = Gdx.graphics.getWidth() / 2;
		var initialPaddleY = 50;

		paddle = new Paddle(100, 10, initialPaddleX, initialPaddleY);
	}

	@Override
	public void render() {
		ScreenUtils.clear(Color.BLACK);

		shapeRenderer.begin(ShapeRenderer.ShapeType.Filled);

		paddle.update();
		paddle.draw(shapeRenderer);

		for (Ball ball : balls) {
			ball.update();
			ball.draw(shapeRenderer, paddle);
		}

		shapeRenderer.end();
	}

	@Override
	public void dispose() {
		shapeRenderer.dispose();
	}
}
