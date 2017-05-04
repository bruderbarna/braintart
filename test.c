#include<stdio.h>
#include<stdlib.h>
int main() { char *p = calloc(30000, 1); char *psave = p; if (!p) return 1; putchar(*p); p++; while (*p) { putchar(*p); p++;  } free(psave); }