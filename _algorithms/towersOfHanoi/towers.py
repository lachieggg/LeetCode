#!/usr/bin/env python3

NUM_DISCS = 4
NUM_TOWERS = 3
MAX_DIAMETER = 4

class Disc:
	def __init__(self, diameter):
		self.diameter = diameter
	
	def __str__(self):
		return '_' * self.diameter

class Tower:
	def __init__(self, character):
		self.character = character
		self.discs = {}

	def addDisc(self, level, disc):
		self.discs[level] = disc

	def getHeight(self):
		return len(self.discs.keys())
	
	def removeTopDisc(self):
		height = self.getHeight()
		if(height == 0):
			return
		del self.discs[height-1]

	def getTopDisc(self):
		height = self.getHeight()
		if(height == 0):
			return
		return self.discs.get(height-1)

	def setTopDisc(self, disc):
		height = self.getHeight()
		self.discs[height] = disc

class Game:
	def __init__(self):
		self.discs = []
		self.createTowers()
		self.createDiscs()

	def moveDisc(self, fromTower, toTower):
		"""
		Move a disc from one tower to another
		Ensures that this is a legal move

		@param fromTower char
		@param toTower char
		"""
		fromTower = self.towers.get(fromTower)
		toTower = self.towers.get(toTower)

		# Top discs from both towers
		fromTowerDisc = fromTower.getTopDisc()
		toTowerDisc = toTower.getTopDisc()

		# Check the 'from' tower has a disc
		if(not fromTowerDisc):
			print("Attempted to move disc from tower with no discs")
			return
		
		# Check the diameter on the 'to' tower is less than the diameter on the 'from'
		if(toTowerDisc):
			if(toTowerDisc.diameter < fromTowerDisc.diameter):
				print("Attempted to place disc of larger diameter on top of a smaller one")
				return

		# Remove the top disc on the 'from' tower and add it on the 'to' tower
		fromTower.removeTopDisc()
		toTower.setTopDisc(fromTowerDisc)

	def createTowers(self):
		self.towers = {}
		self.towers['A'] = Tower('A')
		self.towers['B'] = Tower('B')
		self.towers['C'] = Tower('C')

	def createDiscs(self):
		diameter = MAX_DIAMETER
		for i in range(NUM_DISCS):
			disc = Disc(diameter)
			self.discs.append(disc)
			self.towers.get('A').addDisc(i, disc)
			diameter -= 1
	
	def represent(self):
		for tower in self.towers.values():
			discs = tower.discs
			if(not discs):
				print("{}: empty.".format(tower.character))
				continue

			print("Tower {}: ".format(tower.character))

			for level in list(tower.discs.keys())[::-1]:
				disc = discs.get(level)
				diameter = str(disc)
				print("{}, Diameter {}".format(tower.character, diameter))
		
		print("----------------------------------------------------------------")
	

	def hanoi(self, fromTower, viaTower, toTower, n):
		"""
		Move the top n discs on 'fromTower' to 'toTower' using the 'viaTower'
		
		For instance, move n discs on 'A' to 'C' using 'B'
		"""
		if(n == 0):
			return
		
		# Move the top n-1 discs on 'A' to 'B' using 'C'
		# 
		self.hanoi(fromTower, toTower, viaTower, n-1)

		# Move one disc on the 'fromTower' to the 'toTower'
		self.moveDisc(fromTower, toTower)

		# Move the top n-1 discs on 'B' to 'C' using 'A'
		#
		self.hanoi(viaTower, fromTower, toTower, n-1)

		self.represent()

game = Game()
game.represent()
game.hanoi('A', 'B', 'C', NUM_DISCS)