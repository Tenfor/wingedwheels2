local bridge = require("bridge.bridge")

local M = {
	rewardingAdReady = true,
	coins = 0,
	movement = {
		level = 0,
		maxLevel = 5,
		prices = {5,10,25,60,90},
		data = {180,200,220,250,280,310}
	},
	powerup = {
		level = 0, 
		maxLevel = 3,
		prices = {10,35,90},
		data = {1,3,5,6}
	},
	lives = {
		level = 0, 
		maxLevel = 4,
		prices = {10,20,70,150},
		data = {1,2,3,4,5}
	},
	shield = {
		level = 0, 
		maxLevel = 3,
		prices = {15,30,80},
		data = {4,6,8,11}
	},
	magnet = {
		level = 0, 
		maxLevel = 3,
		prices = {15,30,80},
		data = {4,6,8,11}
	},
	weapons = {
		level = 0, 
		maxLevel = 3,
		prices = {15,30,80},
		data = {3,4,6,8}
	},
}

function M.setCoins(val) 
	if val >= 0 and val <= 9999 then 
		M.coins = val
	end
end

function M.setMovementLevel(val)
	if val <= M.movement.maxLevel then 
		M.movement.level = val
	end
end

function M.setPowerupLevel(val)
	if val <= M.powerup.maxLevel then 
		M.powerup.level = val
	end
end

function M.setLivesLevel(val)
	if val <= M.lives.maxLevel then 
		M.lives.level = val
	end
end

function M.setShieldLevel(val)
	if val <= M.shield.maxLevel then 
		M.shield.level = val
	end
end

function M.setMagnetLevel(val)
	if val <= M.magnet.maxLevel then 
		M.magnet.level = val
	end
end

function M.setWeaponsLevel(val)
	if val <= M.weapons.maxLevel then 
		M.weapons.level = val
	end
end

function M.setRewardingAdReady(val)
	M.rewardingAdReady = val
end

function M.saveUpgrades()
	bridge.storage.set ({ 
		coins = M.coins,
		movement = M.movement.level,
		powerup = M.powerup.level,
		lives = M.lives.level,
		magnet = M.magnet.level,
		shield = M.shield.level,
		weapons = M.weapons.level
	})
end

function M.isMaxed()
	return M.movement.level == M.movement.maxLevel and M.powerup.level == M.powerup.maxLevel and M.lives.level == M.lives.maxLevel and M.magnet.level == M.magnet.maxLevel and M.shield.level == M.shield.maxLevel and M.weapons.level == M.weapons.maxLevel
end

function M.resetData()
		M.coins = 0
		M.movement.level = 0
		M.powerup.level = 0
		M.lives.level = 0
		M.magnet.level = 0
		M.shield.level = 0
		M.weapons.level = 0
end

return M