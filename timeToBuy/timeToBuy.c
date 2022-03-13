#include <stdio.h>

// Polynomial time (n squared)
int maxProfitPoly(int* prices, int pricesSize) {
    int buyIndex;
    int sellIndex;
    int currentMaxProfit = 0;
    int gain;

    for(buyIndex=0; buyIndex<pricesSize; buyIndex++) {
        for(sellIndex=buyIndex+1; sellIndex<pricesSize; sellIndex++) {
            gain = prices[sellIndex] - prices[buyIndex];
            if(gain > currentMaxProfit) {
                currentMaxProfit = gain;
            }
        }
    }
    return currentMaxProfit;
}

// Linear time O(n)
int maxProfit(int* prices, int pricesSize) {
    int index;
    int maxGain = 0;
    int bestSell = prices[0];
    int bestBuy = prices[0];

    for(index=0; index<pricesSize; index++) {
        if(prices[index] < bestBuy) {
            // Must reset the best sell once a new best buy is established
            // because otherwise we would be trying to sell a stock in the past
            // (sell date would be before the buy date)
            // 
            // If the previous max gain (profit) before this context switch
            // is still the largest that is retained
            //
            bestSell = prices[index];
            // New best buy price
            bestBuy = prices[index];
        }

        if(prices[index] > bestSell) {
            // New best sell price
            bestSell = prices[index];
        }
        
        if((bestSell - bestBuy) > maxGain) {
            printf("New best\n");
            printf("Best sell %d\n", bestSell);
            printf("Best buy %d\n", bestBuy);
            maxGain = (bestSell - bestBuy);
        }
    }
    return maxGain;
}

int main(int argc, char **argv) {
    int arr[] = {7,1,5,3,6,4};
    int length = sizeof(arr) / sizeof(arr[0]);
    printf("%d\n", maxProfit(arr, length));
}