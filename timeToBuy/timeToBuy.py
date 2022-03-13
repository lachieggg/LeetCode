#!/usr/bin/env python3

DEBUG = False

class Solution:
    def loadFile(self):
        f = open('dataset', 'r')
        line = f.readlines()
        line = line[0]
        arr = []
        for char in line:
            if(char == ',' or char == '\n'):
                continue
            arr.append(int(char))
        return arr

    def maxProfit(self, prices):
        prices = self.removeAllZeros(prices)
        print(prices)
        currentMaxProfit = 0
        for index, buyPrice in enumerate(prices):
            for sellPrice in prices[index:]:
                delta = sellPrice - buyPrice
                if(delta > currentMaxProfit):
                    if(DEBUG): print("New max profit {}, sell {}, buy {}".format(delta, sellPrice, buyPrice))
                    currentMaxProfit = delta
        return currentMaxProfit

    
    def removeAllZeros(self, prices):
        return list(filter(lambda x: x > 0, prices))

    def allFutureValuesAreZero(self, prices, index):
        """Determines whether all future values are less than or equal to zero"""
        if(sum(prices[index:]) <= 0):
            return True
        return False

    def allFutureValuesAreLower(self, prices, current):
        if(max(prices) < current):
            return True
    


prices = [7,1,5,3,6,4]

s = Solution()
bestProfit = s.maxProfit(prices)
print(bestProfit)

prices = s.loadFile()
bestProfit = s.maxProfit(prices)
print(bestProfit)


