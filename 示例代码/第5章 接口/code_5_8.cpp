//第5章/code_5_8.cpp
Type* pts[2] = {new A, new B};
for(int i = 0; i < 2; ++i) {
	cout << pts[i]->Name() << "," << pts[i]->Size() << endl;
}
