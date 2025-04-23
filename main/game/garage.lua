local M = {
	enemyDestroyed = 0,
	eventSurvived = 0,
	died = 0,
	powerupCollected = 0,
	selectedCar = 1,
	selectedShip = 1,
}

function M.setEnemyDestroyed(val) 
	M.enemyDestroyed = val
end

function M.setEventSurvived(val) 
	M.eventSurvived = val
end

function M.setDied(val) 
	M.died = val
end

function M.setPowerupCollected(val) 
	M.powerupCollected = val
end

function M.setSelectedCar(val) 
	M.selectedCar = val
end

function M.setSelectedShip(val) 
	M.selectedShip = val
end


function M.saveGarage()
	local filename = sys.get_save_file("sys_save_load", "garage")
	sys.save(filename, { 
		enemyDestroyed = M.enemyDestroyed,
		eventSurvived = M.eventSurvived,
		died = M.died,
		powerupCollected = M.powerupCollected,
		selectedCar = M.selectedCar,
		selectedShip = M.selectedShip,
	})
end

function M.resetData()
	M.enemyDestroyed = 0;
	M.eventSurvived = 0;
	M.died = 0;
	M.powerupCollected = 0;
	M.selectedCar = 1;
	M.selectedShip = 1;
end

return M