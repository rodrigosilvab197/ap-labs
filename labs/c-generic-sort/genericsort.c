#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int comparar(const char *x, const char *y){
	int x1 = atoi(x);
	int y1 = atoi(y);
	if (x1 > y1)
		return 1;
	else if(x1 < y1)
		return -1;
	else
		return 0;
}
void quicksort(void *lineptr[], int left, int right,
	   int (*comp)(void *, void *));

void mergesort(void *lineptr[], int left, int right,
	   int (*comp)(void *, void *));

int main(int argc, char **argv){
    if (5>argc){
        printf("Pasa los parametros porfavor\n");
        return -1;
    }
	else{
		if (strcmp(argv[1], "-n") != 0){
		printf("el primer parametro debe de ser -n\n");
        return -1;
		}
		if (strcmp(argv[4], "-o") != 0){
		printf("no se encontro -o \n");
		return -1;
		}
		char *fileName = argv[2];
		FILE *archivo = fopen(fileName,"r");
		if(archivo == NULL){
		printf("el nombre del archivo esta mal\n");
        return -1;
		}
		void *value[1024];
		int size = 0;
		char * linea = NULL;
		char buff[255];
		while(fgets(buff,sizeof(buff),archivo)){
			linea = malloc(strlen(buff));
			value[size]=linea;
			strncpy(linea, buff,  strlen(buff)-1);
			size++;
		}
		fclose(archivo);
		char *x = argv[3];
		if (strcmp(x, "-quicksort") == 0){
		quicksort((void** ) value, 0, size - 1, (int (*)(void *, void *)) (1 ? comparar : strcmp));
		} 
		else if (strcmp(x, "-mergesort") == 0){
		mergesort((void** ) value, 0, size - 1, (int (*)(void *, void *)) (1 ? comparar : strcmp));
		} 
		else {
		printf("Solo se puede hace quickSOrt o MergeSort\n");
		return -1;
		}
		char *nombreArchivo = argv[5];
		FILE *archivoSalida = fopen(nombreArchivo, "wb");
		for (int i = 0; i < size; i++){
			fputs(strcat(value[i], "\n"), archivoSalida);
		}
		fclose(archivoSalida);
		return 0;


	}
}

