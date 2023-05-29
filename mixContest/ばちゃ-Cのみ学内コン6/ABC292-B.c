#include <stdio.h>

int main(void) {
	int n=100, q=100;
	scanf("%d %d\n", &n, &q);

	int i=0;
	int C=0, playerNum=0;
	int player[100] = {};
	for (;i < q; i++){
		scanf("%d %d\n", &C, &playerNum);
		switch (C){
			case 1:
				player[playerNum-1]++;
				break;
			case 2:
				player[playerNum-1] = 2;
				break;
			case 3:
				if (player[playerNum-1] < 2){
					printf("No\n");
				} else {
					printf("Yes\n");
				}
				break;
		}
	}

	return 0;
}
