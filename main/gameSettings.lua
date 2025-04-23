local M = {
	music = true,
	sounds = true,
	showTutorial = true,
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

return M