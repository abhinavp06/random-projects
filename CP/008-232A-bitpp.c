// https://codeforces.com/problemset/problem/282/A
 
#include <stdio.h>
 
int main() {
    int n;
    scanf("%d", &n);
    char str[4] = {0};
    int x = 0;
    for(int i = 0; i < n; i++) {
        scanf("%s", str);
        if(str[0] == '+' || str[2] == '+') {
            x++;
        } else {
            x--;
        }
    }
    printf("%d", x);
    return 0;
}