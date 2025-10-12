// https://codeforces.com/problemset/problem/71/A
 
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
 
int main() {
	const int THRESHOLD = 10;
	int n = 0;
	scanf("%d", &n);
 
	char *p[n];
 
	for(int i=0; i<n; i++) {
		p[i] = malloc(1024);
		scanf("%s", p[i]);
	}
 
	for(int i=0; i<n; i++) {
		int str_length = strlen(p[i]);
 
		if(str_length > THRESHOLD) {
			char str_length_str[12];
            sprintf(str_length_str, "%d", str_length - 2);

            sprintf(p[i], "%c%s%c", p[i][0], str_length_str, p[i][str_length - 1]);
		}
 
		printf("%s\n", p[i]);
		free(p[i]);
	}
 
	return 0;
}