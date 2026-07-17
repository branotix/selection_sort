#include <iostream>
#include <vector>
#include <fstream>
#include <chrono> // সময় হিসাব করার জন্য

using namespace std;

// Selection Sort Algorithm
void selectionSort(vector<int>& arr) {
    int n = arr.size();
    for (int i = 0; i < n - 1; i++) {
        int minIndex = i;
        for (int j = i + 1; j < n; j++) {
            if (arr[j] < arr[minIndex]) {
                minIndex = j;
            }
        }
        // সোয়াপ (Swap) করা
        swap(arr[i], arr[minIndex]);
    }
}

int main() {
    // ১. ফাইল থেকে ডাটা রিড করা
    ifstream inputFile("note.txt");
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

    cout << "Loaded " << numbers.size() << " numbers. Starting Selection Sort in C++..." << endl;

    // ২. সময় গণনা শুরু
    auto startTime = chrono::high_resolution_clock::now();

    // সিলেকশন সর্ট কল করা
    selectionSort(numbers);

    // ৩. সময় গণনা শেষ
    auto endTime = chrono::high_resolution_clock::now();
    chrono::duration<double> duration = endTime - startTime;

    cout << "Sorting completed successfully!" << endl;
    cout << "Time taken: " << duration.count() << " seconds" << endl;

    // প্রথম ১০টি সর্ট হওয়া সংখ্যা প্রিন্ট করা
    if (numbers.size() > 10) {
        cout << "First 10 sorted numbers: ";
        for (int i = 0; i < 10; i++) {
            cout << numbers[i] << " ";
        }
        cout << endl;
    }

    return 0;
}