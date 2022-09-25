//第3章/code_3_17.c
int helper(int (*fn)(int, int), int a, int b) {
	return fn(a, b);
}

int main() {
	return helper(0, 0, 0);
}
