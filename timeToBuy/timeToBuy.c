#include <stdio.h>

int maxProfit(int* prices, int pricesSize) {
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

int main(int argc, char **argv) {
    int arr[] = {7,1,5,3,6,4};
    int length = sizeof(arr) / sizeof(arr[0]);
    printf("%d\n", maxProfit(arr, length));
}