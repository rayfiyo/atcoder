#include <stdio.h>

int main(void) {
	int A=100, B=100, C=1;

	scanf("%d %d %d\n", &A, &B, &C);

	if (C==1){ // 青木 先攻
		if (A>=B){
			printf("Takahashi\n");
		} else {
			printf("Aoki\n");
		}
	} else {   // 高橋 先攻
		if (A>B){
			printf("Takahashi\n");
		} else {
			printf("Aoki\n");
		}
	}

	return 0;
}
