// https://codeforces.com/contest/71/problem/B

#include <stdio.h>

int main() {
	int n, k, t;
	scanf("%d %d %d", &n, &k, &t);

	int total_volume = n*k;
	int filled_volume = total_volume*t/100;
	int filled_blocks = filled_volume/k;
	int remaining = filled_volume%k;
	
	for(int i = 0; i < filled_blocks; i++) printf("%d ", k);
    if(remaining > 0) printf("%d ", remaining);
    for(int i = filled_blocks + (remaining > 0); i < n; i++) printf("0 ");
    return 0;

}