#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void merge(void *lineptr[], int left, int right, int (*comp)(void *, void *)){
	int mid = (left+right)/2;
	mid++;
	for(int i = mid; i<right;i++){
		for(int j = 0 ;j<=mid;j++){
			if(0>(*comp)(lineptr[i],lineptr[j])){
				void* tmp = lineptr[i];
				lineptr[i]=lineptr[j];
				lineptr[j]=tmp;
			}
		}
	}
	
}
void mergesort(void *lineptr[], int left, int right,int (*comp)(void *, void *)) {
	if(right > left){
		int mid = (left+right)/2;
		mergesort(lineptr,left,mid,(*comp));
		mergesort(lineptr,mid+1,right,(*comp));
		merge(lineptr,left,right,(*comp));
	}
}
