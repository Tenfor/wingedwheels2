local M = {
	speedLevel = 1,
	enemyLevel = 1,
	isOver = false,
	magnetActive = false,
	isStarted = false,
	currentScores = 0,
	isPaused = false,
	newHighscore = false
}

function M.setNewHighscore(val)
	M.newHighscore = val
end

function M.setIsStarted(val)
	M.isStarted = val
end

function M.increaseSpeedLevel()
	M.speedLevel = M.speedLevel + 1
end

function M.setSpeedLevel(val)
	M.speedLevel = val
end

function M.increaseEnemyLevel()
	M.enemyLevel = M.enemyLevel + 1
end

function M.setEnemyLevel(val)
	M.enemyLevel = val
end

function M.setIsOver(val)
	M.isOver = val
end

function M.setMagnetActive(val)
	M.magnetActive = val
end

function M.setCurrentScores(val)
	M.currentScores = val
end

function M.reset()
	M.speedLevel = 1;
	M.enemyLevel = 1;
	M.isOver = false;
	M.magnetActive = false;
	M.isStarted = false;
	M.currentScores = 0;
	M.isPaused = false;
	M.newHighscore = false;
end

function M.togglePause()
	M.isPaused = not M.isPaused
end

return M