// magic function made by stefanpochmann to split up big and small vs even and odd

int getFancyVal(std::vector<int>& nums, int index) {
    int doubleIndexPlusOne = 2 * i + 1;

    /*
    to see how the section works, do this. AS you can see it just rounds it up to the odd
    std::cout << (1|1) << std::endl;
    std::cout << (0|1) << std::endl;
    std::cout << (2|1) << std::endl;
    std::cout << (3|1) << std::endl
        std::cout << (4|1) << std::endl;
    std::cout << (5|1) << std::endl;
    */

    int section = nums.size()|1;
    return nums[doubleIndexPlusOne % section];
}

void wiggleSort(std::vector<int>& nums) {
    int size = nums.size();

    auto midPointer = nums.begin() + size / 2;
    std::nth_element(nums.begin(), midPointer, nums.end());
    int midPointNum = *midPointer;

    int front = 0;
    int back = size - 1;
    int swapPoint = 0;

    while(front <= back) {
        if (getFancyVal(front) > midPointNum)
            std::swap(getFancyVal(swapPoint++), getFancyVal(front++));
        else if (getFancyVal(front) < midPointNum)
            std::swap(getFancyVal(swapPoint++), getFancyVal(back--));
        else
            front++;
    }
}








// problems still to do:
// trapping rainwater
// friendship circles
// wigglesort 1
// all neetcode
// everything from uber
// everything from stripe
// everything from pinterest
// everything from airbnb
// everything from robinhood
// everything from linkedin
// everything from facebook
// everything from google
// everything from twitter
// everything from amazon
// everything from databricks
// everything from snowflake
// everything from quora
// everything from coinbase
// everything from github
// citadel
// 2sigma
// optiver
// nuro
// cruise
// convoy
// instacart
// plaid
// HRT
// DRW
// PDT partners
// square
// chime
// amplitude
// dropbox
// snap
// figma
// scaleai
// airtable
// bytedance
// argo
// asana
// dfinity
// ramp ? (lol)
// brex?



/*
class Solution {
public:
    
    int getFancyVal(std::vector<int>& nums, int index) {
    int doubleIndexPlusOne = 2 * index + 1;
    int section = nums.size()|1;
    return doubleIndexPlusOne % section;
}
    
    void wiggleSort(vector<int>& nums) {
            int size = nums.size();

    auto midPointer = nums.begin() + size / 2;
    std::nth_element(nums.begin(), midPointer, nums.end());
    int midPointNum = *midPointer;

    int front = 0;
    int back = size - 1;
    int swapPoint = 0;

    while(front <= back) {
        if (nums[getFancyVal(nums, front)] > midPointNum)
            std::swap(nums[getFancyVal(nums, swapPoint++)], nums[getFancyVal(nums, front++)]);
        else if (nums[getFancyVal(nums, front)] < midPointNum)
            std::swap(nums[getFancyVal(nums, front)], nums[getFancyVal(nums, back--)]);
        else
            front++;
    }
    }
};
*/




