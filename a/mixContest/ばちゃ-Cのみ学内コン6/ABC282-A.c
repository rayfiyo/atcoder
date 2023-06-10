#include <stdio.h>

int main(void) {
	int K=26, i=1;
	scanf("%d", &K);
	for (; i<K+1; i++){
		printf("%c", 64+i);
	}
	return 0;
}
