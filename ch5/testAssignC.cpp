#include <iostream>
using namespace std;

int main() {
	int a[3]={1};
	int b[4]={5,2};
	cout<<a[0]<<endl;
	cout<<b[0]<<endl;
	//a = b;
	int* c = a;
	cout<<c[0]<<endl;
	c = b;
	cout<<c[0]<<endl;
}
