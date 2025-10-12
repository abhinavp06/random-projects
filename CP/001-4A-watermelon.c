// https://codeforces.com/problemset/problem/4/A
#include <stdio.h>

int main() {
	int w = 0;
	scanf("%d", &w);

	if((w - 2) > 0 && (w - 2) % 2 == 0) {
		printf("YES");
	} else {
		printf("NO");
	}

	return 0;
}