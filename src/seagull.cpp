#include <iostream>
using namespace std;
int c(){for(int a=1;a<100;a++){
if(a%15==0){cout<<"fizzbuzz\n"; 
}else if(a%5==0){cout<<"fizz\n";
}else if(a%3==0){cout<<"buzz\n";
}else{cout<<a<<endl;}}}
