#include <iostream>
using namespace std;

void ptr1(int arr[],int size){
	for(int i=0; i<size; i++){
		cout<<arr[i]<<"\t";
	}
	cout<<endl;
}

void ptr2(int arr[4],int size){
	for(int i=0; i<size; i++){
		cout<<arr[i]<<"\t";
	}
	cout<<endl;
}

void ptr3(int* arr,int size){
	for(int i=0; i<size; i++){
		cout<<arr[i]<<"\t";
	}
	cout<<endl;
}

void ref(int (&arr)[4],int size){
	for(int i=0; i<size; i++){
		cout<<arr[i]<<"\t";
	}
	cout<<endl;
}

int main() {
	int a[]={1,2,3,4};
	int* b = &a[0];
	cout<<a[0]<<"\t"<<a[1]<<"\t";
	cout<<b[2]<<"\t"<<b[3]<<endl;
	cout<<a<<endl;
	cout<<b<<endl;
	cout<<sizeof(a)<<endl;
	cout<<sizeof(b)<<endl;

	ptr1(a,4);
	ptr2(a,4);
	ptr3(b,4);
	ref(a,4);
	return 0;
}
