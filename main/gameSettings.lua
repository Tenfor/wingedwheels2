local bridge = require("bridge.bridge")

local M = {
	music = true,
	sounds = true,
	showTutorial = true,
	highscore = 0,
	playerName = "",
	lang = "eng",
	leaderBoardAvaible = true
}

function M.setMusic(val)
	M.music = val
end

function M.setSounds(val)
	M.sounds = val
end	

function M.setShowTutorial(val)
	M.showTutorial = val
end	

function M.setHighScore(val)
	M.highscore = val
end

function M.setPlayerName(val)
	M.playerName = val
end

function M.setLang(val)
	M.lang = val
end

return M