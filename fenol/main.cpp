#include <iostream>
#include <vector>
#include <fstream>
#include <chrono>

using namespace std;

void selectionSort(vector<int>& arr) {
    int n = arr.size();
    for (int i = 0; i < n - 1; i++) {
        int minIndex = i;
        for (int j = i + 1; j < n; j++) {
            if (arr[j] < arr[minIndex]) {
                minIndex = j;
            }
        }
        swap(arr[i], arr[minIndex]);
    }
}

int main() {
    ifstream inputFile("first_50k.txt");
    if (!inputFile.is_open()) {
        cerr << "Error opening file!" << endl;
        return 1;
    }

    vector<int> numbers;
    int num;
    while (inputFile >> num) {
        numbers.push_back(num);
    }
    inputFile.close();

    auto startTime = chrono::high_resolution_clock::now();
    selectionSort(numbers);

    auto endTime = chrono::high_resolution_clock::now();
    chrono::duration<double> duration = endTime - startTime;

    cout << "Sorting completed successfully Time taken: " << duration.count() << " seconds" << endl;

    return 0;
}