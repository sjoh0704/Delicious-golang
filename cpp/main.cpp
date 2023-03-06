#include <stdio.h>

extern "C" {
    #include "libexample.h"
}

int main() {
    int a = 1;
    int b = 2;
    int c = Sum(a, b);
    printf("%d + %d = %d\n", a, b, c);
    return 0;
}