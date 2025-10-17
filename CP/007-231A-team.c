// https://codeforces.com/problemset/problem/231/A
 
#include <stdio.h>
 
int main() {
    int n;
    scanf("%d", &n);
    int result = 0;
 
    for(int i = 0; i < n; i++) {
        int petya, vasya, tonya;
        scanf("%d %d %d", &petya, &vasya, &tonya);
        int count = petya + vasya + tonya;
        if(count >= 2) {
            result++;
        }
    }
    
    printf("%d\n", result);
    return 0;
}