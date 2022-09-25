//第5章/code_5_7.cpp
class Type {
public:
	virtual string Name() = 0;
	virtual size_t Size() = 0;
};

class A : public Type {
public:
	string Name() {
		return "A";
	}
	size_t Size() {
		return sizeof(*this);
	}
};

class B : public Type {
public:
	string Name() {
		return "B";
	}
	size_t Size() {
		return sizeof(*this);
	}
private:
	int somedata;
};
